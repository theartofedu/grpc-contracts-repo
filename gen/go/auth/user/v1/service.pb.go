// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: auth/user/v1/service.proto

package pb_auth_user

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UserServiceCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Email string   `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Roles []uint64 `protobuf:"varint,3,rep,packed,name=roles,proto3" json:"roles,omitempty"`
}

func (x *UserServiceCreateRequest) Reset() {
	*x = UserServiceCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_user_v1_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserServiceCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserServiceCreateRequest) ProtoMessage() {}

func (x *UserServiceCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_user_v1_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserServiceCreateRequest.ProtoReflect.Descriptor instead.
func (*UserServiceCreateRequest) Descriptor() ([]byte, []int) {
	return file_auth_user_v1_service_proto_rawDescGZIP(), []int{0}
}

func (x *UserServiceCreateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserServiceCreateRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserServiceCreateRequest) GetRoles() []uint64 {
	if x != nil {
		return x.Roles
	}
	return nil
}

type UserServiceCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UserServiceCreateResponse) Reset() {
	*x = UserServiceCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_user_v1_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserServiceCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserServiceCreateResponse) ProtoMessage() {}

func (x *UserServiceCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_auth_user_v1_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserServiceCreateResponse.ProtoReflect.Descriptor instead.
func (*UserServiceCreateResponse) Descriptor() ([]byte, []int) {
	return file_auth_user_v1_service_proto_rawDescGZIP(), []int{1}
}

type UserServiceUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name  string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Roles []uint64 `protobuf:"varint,4,rep,packed,name=roles,proto3" json:"roles,omitempty"`
}

func (x *UserServiceUpdateRequest) Reset() {
	*x = UserServiceUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_user_v1_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserServiceUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserServiceUpdateRequest) ProtoMessage() {}

func (x *UserServiceUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_user_v1_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserServiceUpdateRequest.ProtoReflect.Descriptor instead.
func (*UserServiceUpdateRequest) Descriptor() ([]byte, []int) {
	return file_auth_user_v1_service_proto_rawDescGZIP(), []int{2}
}

func (x *UserServiceUpdateRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UserServiceUpdateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserServiceUpdateRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserServiceUpdateRequest) GetRoles() []uint64 {
	if x != nil {
		return x.Roles
	}
	return nil
}

type UserServiceUpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UserServiceUpdateResponse) Reset() {
	*x = UserServiceUpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_user_v1_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserServiceUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserServiceUpdateResponse) ProtoMessage() {}

func (x *UserServiceUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_auth_user_v1_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserServiceUpdateResponse.ProtoReflect.Descriptor instead.
func (*UserServiceUpdateResponse) Descriptor() ([]byte, []int) {
	return file_auth_user_v1_service_proto_rawDescGZIP(), []int{3}
}

type UserServiceByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UserServiceByIDRequest) Reset() {
	*x = UserServiceByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_user_v1_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserServiceByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserServiceByIDRequest) ProtoMessage() {}

func (x *UserServiceByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_user_v1_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserServiceByIDRequest.ProtoReflect.Descriptor instead.
func (*UserServiceByIDRequest) Descriptor() ([]byte, []int) {
	return file_auth_user_v1_service_proto_rawDescGZIP(), []int{4}
}

func (x *UserServiceByIDRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UserServiceByIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *UserServiceByIDResponse) Reset() {
	*x = UserServiceByIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_user_v1_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserServiceByIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserServiceByIDResponse) ProtoMessage() {}

func (x *UserServiceByIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_auth_user_v1_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserServiceByIDResponse.ProtoReflect.Descriptor instead.
func (*UserServiceByIDResponse) Descriptor() ([]byte, []int) {
	return file_auth_user_v1_service_proto_rawDescGZIP(), []int{5}
}

func (x *UserServiceByIDResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type UserServiceAllRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UserServiceAllRequest) Reset() {
	*x = UserServiceAllRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_user_v1_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserServiceAllRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserServiceAllRequest) ProtoMessage() {}

func (x *UserServiceAllRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_user_v1_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserServiceAllRequest.ProtoReflect.Descriptor instead.
func (*UserServiceAllRequest) Descriptor() ([]byte, []int) {
	return file_auth_user_v1_service_proto_rawDescGZIP(), []int{6}
}

type UserServiceAllResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User []*User `protobuf:"bytes,1,rep,name=user,proto3" json:"user,omitempty"`
}

func (x *UserServiceAllResponse) Reset() {
	*x = UserServiceAllResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_user_v1_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserServiceAllResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserServiceAllResponse) ProtoMessage() {}

func (x *UserServiceAllResponse) ProtoReflect() protoreflect.Message {
	mi := &file_auth_user_v1_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserServiceAllResponse.ProtoReflect.Descriptor instead.
func (*UserServiceAllResponse) Descriptor() ([]byte, []int) {
	return file_auth_user_v1_service_proto_rawDescGZIP(), []int{7}
}

func (x *UserServiceAllResponse) GetUser() []*User {
	if x != nil {
		return x.User
	}
	return nil
}

var File_auth_user_v1_service_proto protoreflect.FileDescriptor

var file_auth_user_v1_service_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x17, 0x61, 0x75, 0x74, 0x68,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x5a, 0x0a, 0x18, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x6f, 0x6c,
	0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x04, 0x52, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x22,
	0x1b, 0x0a, 0x19, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x6a, 0x0a, 0x18,
	0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x04, 0x52, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x22, 0x1b, 0x0a, 0x19, 0x55, 0x73, 0x65, 0x72,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x28, 0x0a, 0x16, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x41, 0x0a, 0x17, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x42, 0x79,
	0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x22, 0x17, 0x0a, 0x15, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x40, 0x0a, 0x16, 0x55,
	0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x32, 0xea, 0x02,
	0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x50, 0x0a,
	0x03, 0x41, 0x6c, 0x6c, 0x12, 0x23, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41,
	0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x53, 0x0a, 0x04, 0x42, 0x79, 0x49, 0x44, 0x12, 0x24, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x59, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x26,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x59, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x26, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x27, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x4d, 0x5a, 0x4b, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x68, 0x65, 0x61, 0x72, 0x74, 0x6f,
	0x66, 0x65, 0x64, 0x75, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x73, 0x2d, 0x72, 0x65, 0x70, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f,
	0x61, 0x75, 0x74, 0x68, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x62, 0x5f,
	0x61, 0x75, 0x74, 0x68, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_auth_user_v1_service_proto_rawDescOnce sync.Once
	file_auth_user_v1_service_proto_rawDescData = file_auth_user_v1_service_proto_rawDesc
)

func file_auth_user_v1_service_proto_rawDescGZIP() []byte {
	file_auth_user_v1_service_proto_rawDescOnce.Do(func() {
		file_auth_user_v1_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_auth_user_v1_service_proto_rawDescData)
	})
	return file_auth_user_v1_service_proto_rawDescData
}

var file_auth_user_v1_service_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_auth_user_v1_service_proto_goTypes = []interface{}{
	(*UserServiceCreateRequest)(nil),  // 0: auth.user.v1.UserServiceCreateRequest
	(*UserServiceCreateResponse)(nil), // 1: auth.user.v1.UserServiceCreateResponse
	(*UserServiceUpdateRequest)(nil),  // 2: auth.user.v1.UserServiceUpdateRequest
	(*UserServiceUpdateResponse)(nil), // 3: auth.user.v1.UserServiceUpdateResponse
	(*UserServiceByIDRequest)(nil),    // 4: auth.user.v1.UserServiceByIDRequest
	(*UserServiceByIDResponse)(nil),   // 5: auth.user.v1.UserServiceByIDResponse
	(*UserServiceAllRequest)(nil),     // 6: auth.user.v1.UserServiceAllRequest
	(*UserServiceAllResponse)(nil),    // 7: auth.user.v1.UserServiceAllResponse
	(*User)(nil),                      // 8: auth.user.v1.User
}
var file_auth_user_v1_service_proto_depIdxs = []int32{
	8, // 0: auth.user.v1.UserServiceByIDResponse.user:type_name -> auth.user.v1.User
	8, // 1: auth.user.v1.UserServiceAllResponse.user:type_name -> auth.user.v1.User
	6, // 2: auth.user.v1.UserService.All:input_type -> auth.user.v1.UserServiceAllRequest
	4, // 3: auth.user.v1.UserService.ByID:input_type -> auth.user.v1.UserServiceByIDRequest
	0, // 4: auth.user.v1.UserService.Create:input_type -> auth.user.v1.UserServiceCreateRequest
	2, // 5: auth.user.v1.UserService.Update:input_type -> auth.user.v1.UserServiceUpdateRequest
	7, // 6: auth.user.v1.UserService.All:output_type -> auth.user.v1.UserServiceAllResponse
	5, // 7: auth.user.v1.UserService.ByID:output_type -> auth.user.v1.UserServiceByIDResponse
	1, // 8: auth.user.v1.UserService.Create:output_type -> auth.user.v1.UserServiceCreateResponse
	3, // 9: auth.user.v1.UserService.Update:output_type -> auth.user.v1.UserServiceUpdateResponse
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_auth_user_v1_service_proto_init() }
func file_auth_user_v1_service_proto_init() {
	if File_auth_user_v1_service_proto != nil {
		return
	}
	file_auth_user_v1_user_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_auth_user_v1_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserServiceCreateRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_auth_user_v1_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserServiceCreateResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_auth_user_v1_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserServiceUpdateRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_auth_user_v1_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserServiceUpdateResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_auth_user_v1_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserServiceByIDRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_auth_user_v1_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserServiceByIDResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_auth_user_v1_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserServiceAllRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_auth_user_v1_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserServiceAllResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_auth_user_v1_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_auth_user_v1_service_proto_goTypes,
		DependencyIndexes: file_auth_user_v1_service_proto_depIdxs,
		MessageInfos:      file_auth_user_v1_service_proto_msgTypes,
	}.Build()
	File_auth_user_v1_service_proto = out.File
	file_auth_user_v1_service_proto_rawDesc = nil
	file_auth_user_v1_service_proto_goTypes = nil
	file_auth_user_v1_service_proto_depIdxs = nil
}
