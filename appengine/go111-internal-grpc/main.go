package main

import (
	"context"
	"contrib.go.opencensus.io/exporter/stackdriver"
	"contrib.go.opencensus.io/exporter/stackdriver/propagation"
	"errors"
	"fmt"
	"github.com/favclip/ucon"
	"github.com/vvakame/til/appengine/go111-internal-grpc/echopb"
	"github.com/vvakame/til/appengine/go111-internal-grpc/log"
	"go.opencensus.io/plugin/ochttp"
	"go.opencensus.io/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	rlog "log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var echoCli echopb.EchoClient
var echoOnce sync.Once

func main() {
	close, err := log.Init()
	if err != nil {
		rlog.Fatalf("Failed to create logger: %v", err)
	}
	defer close()

	exporter, err := stackdriver.NewExporter(stackdriver.Options{
		ProjectID: os.Getenv("GOOGLE_CLOUD_PROJECT"),
	})
	if err != nil {
		rlog.Fatalf("Failed to create stackdriver exporter: %v", err)
	}
	trace.RegisterExporter(exporter)
	defer exporter.Flush()

	ucon.Middleware(func(b *ucon.Bubble) error {
		b.Context = log.WithContext(b.Context, b.R)
		b.R = b.R.WithContext(b.Context)
		return b.Next()
	})
	ucon.Orthodox()

	ucon.HandleFunc("POST", "/echo", echoHandler)
	ucon.HandleFunc("GET", "/", indexHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		rlog.Printf("Defaulting to port %s for HTTP", port)
	}

	rlog.Printf("Listening HTTP on port %s", port)

	server := &http.Server{
		Addr: fmt.Sprintf(":%s", port),
		Handler: &ochttp.Handler{
			Handler:     ucon.DefaultMux,
			Propagation: &propagation.HTTPFormat{},
		},
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			rlog.Fatal(err)
		}
	}()

	port = os.Getenv("GRPC_PORT")
	if port == "" {
		port = "5000"
		rlog.Printf("Defaulting to port %s for gRPC", port)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", port))
	if err != nil {
		rlog.Fatal(err)
	}
	grpcServer := grpc.NewServer()
	echopb.RegisterEchoServer(grpcServer, &echoServiceImpl{})
	reflection.Register(grpcServer)

	rlog.Printf("Listening gRPC on port %s", port)

	go func() {
		if err := grpcServer.Serve(lis); err != nil && err != http.ErrServerClosed {
			rlog.Fatal(err)
		}
	}()

	rlog.Printf("running...")

	// setup graceful shutdown...
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM)
	<-sigCh

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	if err := server.Shutdown(ctx); err != nil {
		rlog.Fatalf("graceful shutdown failure: %s", err)
	}
	grpcServer.GracefulStop()
	rlog.Printf("graceful shutdown successfully")
}

// indexHandler responds to requests with our greeting.
func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	ctx := r.Context()

	ctx, span := trace.StartSpan(ctx, "indexHandler")
	defer span.End()

	log.Debugf(ctx, "Hi, 1")
	log.Infof(ctx, "Hi, 2")

	fmt.Fprint(w, "Hello, World!")
}

func echoHandler(ctx context.Context, req *echopb.SayRequest) (*echopb.SayResponse, error) {
	echoOnce.Do(func() {
		port := os.Getenv("GRPC_PORT")
		if port == "" {
			port = "5000"
		}

		conn, err := grpc.Dial(fmt.Sprintf("localhost:%s", port), grpc.WithInsecure())
		if err != nil {
			return
		}

		echoCli = echopb.NewEchoClient(conn)
	})
	if echoCli == nil {
		return nil, errors.New("echoCli is nil")
	}

	return echoCli.Say(ctx, req)
}
