// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package idp

import (
	"context"

	"github.com/ory/fosite"
	"github.com/vvakame/til/go/oauth2idp-example/domains"
)

// Injectors from wire.go:

func InitializeProvider() (fosite.OAuth2Provider, error) {
	config, err := ProvideConfig()
	if err != nil {
		return nil, err
	}
	client, err := ProvideDatastore()
	if err != nil {
		return nil, err
	}
	storage, err := ProvideStore(client)
	if err != nil {
		return nil, err
	}
	rsaPrivateKey, err := ProvideRSAPrivateKey()
	if err != nil {
		return nil, err
	}
	commonStrategy, err := ProvideStrategy(config, rsaPrivateKey)
	if err != nil {
		return nil, err
	}
	oAuth2Provider, err := ProvideOAuth2Provider(config, storage, commonStrategy)
	if err != nil {
		return nil, err
	}
	return oAuth2Provider, nil
}

func InitializeSession(ctx context.Context, user *domains.User) (Session, error) {
	session, err := ProvideSession(ctx, user)
	if err != nil {
		return nil, err
	}
	return session, nil
}
