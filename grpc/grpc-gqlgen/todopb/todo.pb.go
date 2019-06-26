// Code generated by protoc-gen-go. DO NOT EDIT.
// source: todo.proto

package todopb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/vvakame/til/grpc/grpc-gqlgen/gqlgen-proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ListARequest_DoneFilter int32

const (
	ListARequest_NONE     ListARequest_DoneFilter = 0
	ListARequest_DONE     ListARequest_DoneFilter = 1
	ListARequest_NOT_DONE ListARequest_DoneFilter = 2
)

var ListARequest_DoneFilter_name = map[int32]string{
	0: "NONE",
	1: "DONE",
	2: "NOT_DONE",
}

var ListARequest_DoneFilter_value = map[string]int32{
	"NONE":     0,
	"DONE":     1,
	"NOT_DONE": 2,
}

func (x ListARequest_DoneFilter) String() string {
	return proto.EnumName(ListARequest_DoneFilter_name, int32(x))
}

func (ListARequest_DoneFilter) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0e4b95d0c4e09639, []int{3, 0}
}

type ListBRequest_DoneFilter int32

const (
	ListBRequest_NONE     ListBRequest_DoneFilter = 0
	ListBRequest_DONE     ListBRequest_DoneFilter = 1
	ListBRequest_NOT_DONE ListBRequest_DoneFilter = 2
)

var ListBRequest_DoneFilter_name = map[int32]string{
	0: "NONE",
	1: "DONE",
	2: "NOT_DONE",
}

var ListBRequest_DoneFilter_value = map[string]int32{
	"NONE":     0,
	"DONE":     1,
	"NOT_DONE": 2,
}

func (x ListBRequest_DoneFilter) String() string {
	return proto.EnumName(ListBRequest_DoneFilter_name, int32(x))
}

func (ListBRequest_DoneFilter) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0e4b95d0c4e09639, []int{5, 0}
}

type Todo struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Text                 string               `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	Done                 bool                 `protobuf:"varint,3,opt,name=done,proto3" json:"done,omitempty"`
	DoneAt               *timestamp.Timestamp `protobuf:"bytes,4,opt,name=done_at,json=doneAt,proto3" json:"done_at,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Todo) Reset()         { *m = Todo{} }
func (m *Todo) String() string { return proto.CompactTextString(m) }
func (*Todo) ProtoMessage()    {}
func (*Todo) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e4b95d0c4e09639, []int{0}
}

func (m *Todo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Todo.Unmarshal(m, b)
}
func (m *Todo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Todo.Marshal(b, m, deterministic)
}
func (m *Todo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Todo.Merge(m, src)
}
func (m *Todo) XXX_Size() int {
	return xxx_messageInfo_Todo.Size(m)
}
func (m *Todo) XXX_DiscardUnknown() {
	xxx_messageInfo_Todo.DiscardUnknown(m)
}

var xxx_messageInfo_Todo proto.InternalMessageInfo

func (m *Todo) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Todo) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *Todo) GetDone() bool {
	if m != nil {
		return m.Done
	}
	return false
}

func (m *Todo) GetDoneAt() *timestamp.Timestamp {
	if m != nil {
		return m.DoneAt
	}
	return nil
}

func (m *Todo) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *Todo) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

type CreateRequest struct {
	Text                 string   `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e4b95d0c4e09639, []int{1}
}

func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (m *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(m, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type CreateResponse struct {
	Todo                 *Todo    `protobuf:"bytes,1,opt,name=todo,proto3" json:"todo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e4b95d0c4e09639, []int{2}
}

func (m *CreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResponse.Unmarshal(m, b)
}
func (m *CreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResponse.Marshal(b, m, deterministic)
}
func (m *CreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponse.Merge(m, src)
}
func (m *CreateResponse) XXX_Size() int {
	return xxx_messageInfo_CreateResponse.Size(m)
}
func (m *CreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponse proto.InternalMessageInfo

func (m *CreateResponse) GetTodo() *Todo {
	if m != nil {
		return m.Todo
	}
	return nil
}

type ListARequest struct {
	// TODO uint32 first = 1 [(gqlgen.api.field) = {optional: true}];
	First                int32                   `protobuf:"varint,1,opt,name=first,proto3" json:"first,omitempty"`
	After                string                  `protobuf:"bytes,2,opt,name=after,proto3" json:"after,omitempty"`
	Done                 ListARequest_DoneFilter `protobuf:"varint,3,opt,name=done,proto3,enum=todo.ListARequest_DoneFilter" json:"done,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *ListARequest) Reset()         { *m = ListARequest{} }
func (m *ListARequest) String() string { return proto.CompactTextString(m) }
func (*ListARequest) ProtoMessage()    {}
func (*ListARequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e4b95d0c4e09639, []int{3}
}

func (m *ListARequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListARequest.Unmarshal(m, b)
}
func (m *ListARequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListARequest.Marshal(b, m, deterministic)
}
func (m *ListARequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListARequest.Merge(m, src)
}
func (m *ListARequest) XXX_Size() int {
	return xxx_messageInfo_ListARequest.Size(m)
}
func (m *ListARequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListARequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListARequest proto.InternalMessageInfo

func (m *ListARequest) GetFirst() int32 {
	if m != nil {
		return m.First
	}
	return 0
}

func (m *ListARequest) GetAfter() string {
	if m != nil {
		return m.After
	}
	return ""
}

func (m *ListARequest) GetDone() ListARequest_DoneFilter {
	if m != nil {
		return m.Done
	}
	return ListARequest_NONE
}

type ListAResponse struct {
	Cursor               string   `protobuf:"bytes,1,opt,name=cursor,proto3" json:"cursor,omitempty"`
	Todos                []*Todo  `protobuf:"bytes,2,rep,name=todos,proto3" json:"todos,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListAResponse) Reset()         { *m = ListAResponse{} }
func (m *ListAResponse) String() string { return proto.CompactTextString(m) }
func (*ListAResponse) ProtoMessage()    {}
func (*ListAResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e4b95d0c4e09639, []int{4}
}

func (m *ListAResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListAResponse.Unmarshal(m, b)
}
func (m *ListAResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListAResponse.Marshal(b, m, deterministic)
}
func (m *ListAResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListAResponse.Merge(m, src)
}
func (m *ListAResponse) XXX_Size() int {
	return xxx_messageInfo_ListAResponse.Size(m)
}
func (m *ListAResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListAResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListAResponse proto.InternalMessageInfo

func (m *ListAResponse) GetCursor() string {
	if m != nil {
		return m.Cursor
	}
	return ""
}

func (m *ListAResponse) GetTodos() []*Todo {
	if m != nil {
		return m.Todos
	}
	return nil
}

type ListBRequest struct {
	// TODO uint32 offset = 1 [(gqlgen.api.field) = {optional: true}];
	Offset int32 `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	// TODO uint32 limit = 2 [(gqlgen.api.field) = {optional: true}];
	Limit                int32                   `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Done                 ListBRequest_DoneFilter `protobuf:"varint,3,opt,name=done,proto3,enum=todo.ListBRequest_DoneFilter" json:"done,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *ListBRequest) Reset()         { *m = ListBRequest{} }
func (m *ListBRequest) String() string { return proto.CompactTextString(m) }
func (*ListBRequest) ProtoMessage()    {}
func (*ListBRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e4b95d0c4e09639, []int{5}
}

func (m *ListBRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListBRequest.Unmarshal(m, b)
}
func (m *ListBRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListBRequest.Marshal(b, m, deterministic)
}
func (m *ListBRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListBRequest.Merge(m, src)
}
func (m *ListBRequest) XXX_Size() int {
	return xxx_messageInfo_ListBRequest.Size(m)
}
func (m *ListBRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListBRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListBRequest proto.InternalMessageInfo

func (m *ListBRequest) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *ListBRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ListBRequest) GetDone() ListBRequest_DoneFilter {
	if m != nil {
		return m.Done
	}
	return ListBRequest_NONE
}

type ListBResponse struct {
	Todos                []*Todo  `protobuf:"bytes,2,rep,name=todos,proto3" json:"todos,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListBResponse) Reset()         { *m = ListBResponse{} }
func (m *ListBResponse) String() string { return proto.CompactTextString(m) }
func (*ListBResponse) ProtoMessage()    {}
func (*ListBResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e4b95d0c4e09639, []int{6}
}

func (m *ListBResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListBResponse.Unmarshal(m, b)
}
func (m *ListBResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListBResponse.Marshal(b, m, deterministic)
}
func (m *ListBResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListBResponse.Merge(m, src)
}
func (m *ListBResponse) XXX_Size() int {
	return xxx_messageInfo_ListBResponse.Size(m)
}
func (m *ListBResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListBResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListBResponse proto.InternalMessageInfo

func (m *ListBResponse) GetTodos() []*Todo {
	if m != nil {
		return m.Todos
	}
	return nil
}

type UpdateRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Text                 string   `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	Done                 bool     `protobuf:"varint,3,opt,name=done,proto3" json:"done,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateRequest) Reset()         { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()    {}
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e4b95d0c4e09639, []int{7}
}

func (m *UpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRequest.Unmarshal(m, b)
}
func (m *UpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRequest.Marshal(b, m, deterministic)
}
func (m *UpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRequest.Merge(m, src)
}
func (m *UpdateRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateRequest.Size(m)
}
func (m *UpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRequest proto.InternalMessageInfo

func (m *UpdateRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UpdateRequest) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *UpdateRequest) GetDone() bool {
	if m != nil {
		return m.Done
	}
	return false
}

type UpdateResponse struct {
	Todo                 *Todo    `protobuf:"bytes,1,opt,name=todo,proto3" json:"todo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateResponse) Reset()         { *m = UpdateResponse{} }
func (m *UpdateResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateResponse) ProtoMessage()    {}
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e4b95d0c4e09639, []int{8}
}

func (m *UpdateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateResponse.Unmarshal(m, b)
}
func (m *UpdateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateResponse.Marshal(b, m, deterministic)
}
func (m *UpdateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateResponse.Merge(m, src)
}
func (m *UpdateResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateResponse.Size(m)
}
func (m *UpdateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateResponse proto.InternalMessageInfo

func (m *UpdateResponse) GetTodo() *Todo {
	if m != nil {
		return m.Todo
	}
	return nil
}

func init() {
	proto.RegisterEnum("todo.ListARequest_DoneFilter", ListARequest_DoneFilter_name, ListARequest_DoneFilter_value)
	proto.RegisterEnum("todo.ListBRequest_DoneFilter", ListBRequest_DoneFilter_name, ListBRequest_DoneFilter_value)
	proto.RegisterType((*Todo)(nil), "todo.Todo")
	proto.RegisterType((*CreateRequest)(nil), "todo.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "todo.CreateResponse")
	proto.RegisterType((*ListARequest)(nil), "todo.ListARequest")
	proto.RegisterType((*ListAResponse)(nil), "todo.ListAResponse")
	proto.RegisterType((*ListBRequest)(nil), "todo.ListBRequest")
	proto.RegisterType((*ListBResponse)(nil), "todo.ListBResponse")
	proto.RegisterType((*UpdateRequest)(nil), "todo.UpdateRequest")
	proto.RegisterType((*UpdateResponse)(nil), "todo.UpdateResponse")
}

func init() { proto.RegisterFile("todo.proto", fileDescriptor_0e4b95d0c4e09639) }

var fileDescriptor_0e4b95d0c4e09639 = []byte{
	// 675 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0xcf, 0x6e, 0xd3, 0x4e,
	0x10, 0xfe, 0xd9, 0x4d, 0xfc, 0x4b, 0x27, 0x7f, 0x28, 0x5b, 0x40, 0x2b, 0xab, 0x94, 0x60, 0x7a,
	0x08, 0x48, 0xb5, 0xdb, 0x70, 0xaa, 0x90, 0x90, 0xe2, 0xb6, 0x48, 0x48, 0xa8, 0x05, 0x53, 0x2e,
	0x1c, 0xa8, 0x9c, 0x78, 0x13, 0x56, 0x24, 0x5e, 0xd7, 0x5e, 0x17, 0xb8, 0x72, 0xec, 0x23, 0xc0,
	0x8d, 0x77, 0xc8, 0x5b, 0x20, 0xde, 0x87, 0x1b, 0xda, 0x5d, 0x6f, 0xe2, 0x94, 0x56, 0xe1, 0xc2,
	0x25, 0xd9, 0x9d, 0x99, 0x6f, 0x66, 0xbe, 0xcf, 0x33, 0x0b, 0xc0, 0x59, 0xc4, 0xdc, 0x24, 0x65,
	0x9c, 0xa1, 0x8a, 0x38, 0xdb, 0x1b, 0x23, 0xc6, 0x46, 0x63, 0xe2, 0x85, 0x09, 0xf5, 0xc2, 0x38,
	0x66, 0x3c, 0xe4, 0x94, 0xc5, 0x99, 0x8a, 0xb1, 0xef, 0x15, 0x5e, 0x79, 0xeb, 0xe7, 0x43, 0x8f,
	0xd3, 0x09, 0xc9, 0x78, 0x38, 0x49, 0x8a, 0x00, 0x7b, 0x74, 0x36, 0x1e, 0x91, 0x78, 0x5b, 0xde,
	0x3c, 0x96, 0x94, 0xc0, 0xce, 0x2f, 0x03, 0x2a, 0x27, 0x2c, 0x62, 0xe8, 0x0e, 0x98, 0x34, 0xc2,
	0x46, 0xdb, 0xe8, 0xac, 0xfa, 0xd6, 0xd7, 0x29, 0x36, 0xb1, 0x11, 0x98, 0x34, 0x42, 0x08, 0x2a,
	0x9c, 0x7c, 0xe2, 0xd8, 0x14, 0x9e, 0x40, 0x9e, 0x85, 0x2d, 0x62, 0x31, 0xc1, 0x2b, 0x6d, 0xa3,
	0x53, 0x0b, 0xe4, 0x19, 0x3d, 0x81, 0xff, 0xc5, 0xff, 0x69, 0xc8, 0x71, 0xa5, 0x6d, 0x74, 0xea,
	0x5d, 0xdb, 0x55, 0x7d, 0xb9, 0xba, 0x2f, 0xf7, 0x44, 0xf7, 0xa5, 0x0a, 0xac, 0x19, 0x81, 0x25,
	0x20, 0x3d, 0x8e, 0xf6, 0x00, 0xf2, 0x24, 0x0a, 0x39, 0x89, 0x04, 0xbe, 0xba, 0x0c, 0x1f, 0xac,
	0x16, 0xd1, 0x0a, 0x3a, 0x48, 0x89, 0x86, 0x5a, 0xcb, 0xa1, 0x45, 0x74, 0x8f, 0x3b, 0x0f, 0xa0,
	0xb9, 0x2f, 0x2f, 0x01, 0x39, 0xcb, 0x49, 0xc6, 0x67, 0x5c, 0x8d, 0x39, 0x57, 0x67, 0x07, 0x5a,
	0x3a, 0x28, 0x4b, 0x58, 0x9c, 0x11, 0xb4, 0x09, 0xf2, 0xab, 0xc8, 0xa8, 0x7a, 0x17, 0x5c, 0xf9,
	0xb9, 0x84, 0x86, 0x81, 0xb4, 0x3b, 0x3f, 0x0c, 0x68, 0xbc, 0xa0, 0x19, 0xef, 0xe9, 0xb4, 0x1b,
	0x50, 0x1d, 0xd2, 0x34, 0x53, 0x79, 0xab, 0x33, 0xf2, 0xca, 0x28, 0xbc, 0xe1, 0x90, 0x93, 0x54,
	0x29, 0x3c, 0xf7, 0x4a, 0x23, 0xda, 0x2b, 0x49, 0xdd, 0xea, 0xde, 0x55, 0xc5, 0xca, 0xd9, 0xdd,
	0x03, 0x16, 0x93, 0x67, 0x74, 0xcc, 0x49, 0x3a, 0xc3, 0x4a, 0x88, 0x73, 0x08, 0x30, 0xf7, 0xa1,
	0x1a, 0x54, 0x8e, 0x8e, 0x8f, 0x0e, 0xd7, 0xfe, 0x13, 0xa7, 0x03, 0x71, 0x32, 0x50, 0x03, 0x6a,
	0x47, 0xc7, 0x27, 0xa7, 0xf2, 0x66, 0xda, 0xb7, 0xbf, 0x4f, 0xf1, 0x4d, 0xb8, 0x21, 0x93, 0xcf,
	0x81, 0xce, 0x2b, 0x68, 0x16, 0xf5, 0x66, 0xfc, 0xad, 0x41, 0x9e, 0x66, 0x2c, 0x2d, 0x4f, 0x8b,
	0xf8, 0x98, 0xca, 0x8a, 0xda, 0x50, 0x15, 0x5d, 0x66, 0xd8, 0x6c, 0xaf, 0x5c, 0x12, 0x48, 0x39,
	0x9c, 0x9f, 0x85, 0x42, 0xbe, 0x56, 0x68, 0x13, 0x2c, 0x36, 0x1c, 0x66, 0xe4, 0xb2, 0x44, 0x85,
	0x55, 0x68, 0x34, 0xa6, 0x13, 0xaa, 0xa6, 0xb0, 0xa4, 0xa0, 0x34, 0x5e, 0xaf, 0x91, 0xff, 0x2f,
	0x35, 0xf2, 0x4b, 0x1a, 0xed, 0x2a, 0x8d, 0xfc, 0x99, 0x46, 0xcb, 0x35, 0x38, 0x85, 0xe6, 0x1b,
	0x39, 0xc4, 0x5a, 0x83, 0xeb, 0x16, 0xd0, 0x2e, 0x2f, 0xe0, 0xbc, 0x7d, 0xb9, 0x88, 0x76, 0x79,
	0x11, 0x2f, 0x51, 0xdb, 0x81, 0x96, 0x2e, 0xf0, 0x77, 0x83, 0xdb, 0xfd, 0x66, 0x42, 0x5d, 0x5c,
	0x5f, 0x93, 0xf4, 0x9c, 0x0e, 0x08, 0x3a, 0x04, 0x4b, 0x8d, 0x3e, 0x5a, 0x57, 0xb1, 0x0b, 0xdb,
	0x62, 0xdf, 0x5a, 0x34, 0xaa, 0x22, 0xce, 0xda, 0xc5, 0x14, 0x37, 0x50, 0xb1, 0x93, 0xf2, 0x65,
	0x79, 0x0a, 0x55, 0x39, 0x40, 0x08, 0xfd, 0x39, 0xbd, 0xf6, 0xfa, 0x82, 0xad, 0xc8, 0xd1, 0xb8,
	0x98, 0xe2, 0x1a, 0x58, 0x52, 0xa8, 0x9e, 0xc6, 0xfb, 0x65, 0xbc, 0x7f, 0x05, 0xde, 0xbf, 0x12,
	0xef, 0x0b, 0x1a, 0x4a, 0x08, 0x4d, 0x63, 0x41, 0x77, 0x4d, 0x63, 0x51, 0x2b, 0x4d, 0x43, 0xbd,
	0x33, 0x82, 0x86, 0xff, 0xf1, 0xcb, 0x14, 0xef, 0xc3, 0x26, 0x34, 0xdf, 0x75, 0xdc, 0x47, 0x0f,
	0x0b, 0xf4, 0x16, 0xaa, 0x6f, 0xed, 0x0a, 0xf7, 0xf3, 0x38, 0xc9, 0x39, 0x36, 0xe1, 0x3e, 0xb4,
	0x0a, 0xbf, 0x4a, 0xb4, 0x85, 0x9a, 0x2a, 0xe0, 0x65, 0xf8, 0x79, 0xcc, 0xc2, 0x08, 0x1b, 0x6f,
	0xdd, 0x11, 0xe5, 0xef, 0xf3, 0xbe, 0x3b, 0x60, 0x13, 0xef, 0xfc, 0x3c, 0xfc, 0x10, 0x4e, 0x88,
	0xc7, 0xe9, 0xd8, 0x1b, 0xa5, 0xc9, 0x40, 0xfe, 0x6c, 0xab, 0xf7, 0xda, 0x13, 0x7d, 0x25, 0xfd,
	0xbe, 0x25, 0x5f, 0xb1, 0xc7, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0xe7, 0xf8, 0x4d, 0xfa, 0x18,
	0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TodoServiceClient is the client API for TodoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TodoServiceClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	ListA(ctx context.Context, in *ListARequest, opts ...grpc.CallOption) (*ListAResponse, error)
	ListB(ctx context.Context, in *ListBRequest, opts ...grpc.CallOption) (*ListBResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
}

type todoServiceClient struct {
	cc *grpc.ClientConn
}

func NewTodoServiceClient(cc *grpc.ClientConn) TodoServiceClient {
	return &todoServiceClient{cc}
}

func (c *todoServiceClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/todo.TodoService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) ListA(ctx context.Context, in *ListARequest, opts ...grpc.CallOption) (*ListAResponse, error) {
	out := new(ListAResponse)
	err := c.cc.Invoke(ctx, "/todo.TodoService/ListA", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) ListB(ctx context.Context, in *ListBRequest, opts ...grpc.CallOption) (*ListBResponse, error) {
	out := new(ListBResponse)
	err := c.cc.Invoke(ctx, "/todo.TodoService/ListB", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/todo.TodoService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TodoServiceServer is the server API for TodoService service.
type TodoServiceServer interface {
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	ListA(context.Context, *ListARequest) (*ListAResponse, error)
	ListB(context.Context, *ListBRequest) (*ListBResponse, error)
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
}

func RegisterTodoServiceServer(s *grpc.Server, srv TodoServiceServer) {
	s.RegisterService(&_TodoService_serviceDesc, srv)
}

func _TodoService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.TodoService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_ListA_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListARequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).ListA(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.TodoService/ListA",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).ListA(ctx, req.(*ListARequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_ListB_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).ListB(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.TodoService/ListB",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).ListB(ctx, req.(*ListBRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.TodoService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TodoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "todo.TodoService",
	HandlerType: (*TodoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _TodoService_Create_Handler,
		},
		{
			MethodName: "ListA",
			Handler:    _TodoService_ListA_Handler,
		},
		{
			MethodName: "ListB",
			Handler:    _TodoService_ListB_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _TodoService_Update_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "todo.proto",
}
