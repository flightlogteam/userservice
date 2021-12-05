// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: userservice.proto

package userservice

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

type PrivacyLevel int32

const (
	PrivacyLevel_PUBLIC    PrivacyLevel = 0
	PrivacyLevel_PROTECTED PrivacyLevel = 1
	PrivacyLevel_PRIVATE   PrivacyLevel = 2
)

// Enum value maps for PrivacyLevel.
var (
	PrivacyLevel_name = map[int32]string{
		0: "PUBLIC",
		1: "PROTECTED",
		2: "PRIVATE",
	}
	PrivacyLevel_value = map[string]int32{
		"PUBLIC":    0,
		"PROTECTED": 1,
		"PRIVATE":   2,
	}
)

func (x PrivacyLevel) Enum() *PrivacyLevel {
	p := new(PrivacyLevel)
	*p = x
	return p
}

func (x PrivacyLevel) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PrivacyLevel) Descriptor() protoreflect.EnumDescriptor {
	return file_userservice_proto_enumTypes[0].Descriptor()
}

func (PrivacyLevel) Type() protoreflect.EnumType {
	return &file_userservice_proto_enumTypes[0]
}

func (x PrivacyLevel) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PrivacyLevel.Descriptor instead.
func (PrivacyLevel) EnumDescriptor() ([]byte, []int) {
	return file_userservice_proto_rawDescGZIP(), []int{0}
}

type LoginResponse_LoginStatus int32

const (
	LoginResponse_SUCCESS             LoginResponse_LoginStatus = 0
	LoginResponse_INVALID_CREDENTIALS LoginResponse_LoginStatus = 1
	LoginResponse_NOT_ACTIVATED       LoginResponse_LoginStatus = 2
)

// Enum value maps for LoginResponse_LoginStatus.
var (
	LoginResponse_LoginStatus_name = map[int32]string{
		0: "SUCCESS",
		1: "INVALID_CREDENTIALS",
		2: "NOT_ACTIVATED",
	}
	LoginResponse_LoginStatus_value = map[string]int32{
		"SUCCESS":             0,
		"INVALID_CREDENTIALS": 1,
		"NOT_ACTIVATED":       2,
	}
)

func (x LoginResponse_LoginStatus) Enum() *LoginResponse_LoginStatus {
	p := new(LoginResponse_LoginStatus)
	*p = x
	return p
}

func (x LoginResponse_LoginStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LoginResponse_LoginStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_userservice_proto_enumTypes[1].Descriptor()
}

func (LoginResponse_LoginStatus) Type() protoreflect.EnumType {
	return &file_userservice_proto_enumTypes[1]
}

func (x LoginResponse_LoginStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LoginResponse_LoginStatus.Descriptor instead.
func (LoginResponse_LoginStatus) EnumDescriptor() ([]byte, []int) {
	return file_userservice_proto_rawDescGZIP(), []int{5, 0}
}

type CreateUserRequest_PrivacyLevel int32

const (
	CreateUserRequest_PUBLIC    CreateUserRequest_PrivacyLevel = 0
	CreateUserRequest_PROTECTED CreateUserRequest_PrivacyLevel = 1
	CreateUserRequest_PRIVATE   CreateUserRequest_PrivacyLevel = 2
)

// Enum value maps for CreateUserRequest_PrivacyLevel.
var (
	CreateUserRequest_PrivacyLevel_name = map[int32]string{
		0: "PUBLIC",
		1: "PROTECTED",
		2: "PRIVATE",
	}
	CreateUserRequest_PrivacyLevel_value = map[string]int32{
		"PUBLIC":    0,
		"PROTECTED": 1,
		"PRIVATE":   2,
	}
)

func (x CreateUserRequest_PrivacyLevel) Enum() *CreateUserRequest_PrivacyLevel {
	p := new(CreateUserRequest_PrivacyLevel)
	*p = x
	return p
}

func (x CreateUserRequest_PrivacyLevel) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CreateUserRequest_PrivacyLevel) Descriptor() protoreflect.EnumDescriptor {
	return file_userservice_proto_enumTypes[2].Descriptor()
}

func (CreateUserRequest_PrivacyLevel) Type() protoreflect.EnumType {
	return &file_userservice_proto_enumTypes[2]
}

func (x CreateUserRequest_PrivacyLevel) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CreateUserRequest_PrivacyLevel.Descriptor instead.
func (CreateUserRequest_PrivacyLevel) EnumDescriptor() ([]byte, []int) {
	return file_userservice_proto_rawDescGZIP(), []int{6, 0}
}

type CreateUserResponse_CreationStatus int32

const (
	CreateUserResponse_SUCCESS         CreateUserResponse_CreationStatus = 0
	CreateUserResponse_EMAIL_EXISTS    CreateUserResponse_CreationStatus = 1
	CreateUserResponse_USERNAME_EXISTS CreateUserResponse_CreationStatus = 2
	CreateUserResponse_INVALID_DATA    CreateUserResponse_CreationStatus = 3
)

// Enum value maps for CreateUserResponse_CreationStatus.
var (
	CreateUserResponse_CreationStatus_name = map[int32]string{
		0: "SUCCESS",
		1: "EMAIL_EXISTS",
		2: "USERNAME_EXISTS",
		3: "INVALID_DATA",
	}
	CreateUserResponse_CreationStatus_value = map[string]int32{
		"SUCCESS":         0,
		"EMAIL_EXISTS":    1,
		"USERNAME_EXISTS": 2,
		"INVALID_DATA":    3,
	}
)

func (x CreateUserResponse_CreationStatus) Enum() *CreateUserResponse_CreationStatus {
	p := new(CreateUserResponse_CreationStatus)
	*p = x
	return p
}

func (x CreateUserResponse_CreationStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CreateUserResponse_CreationStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_userservice_proto_enumTypes[3].Descriptor()
}

func (CreateUserResponse_CreationStatus) Type() protoreflect.EnumType {
	return &file_userservice_proto_enumTypes[3]
}

func (x CreateUserResponse_CreationStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CreateUserResponse_CreationStatus.Descriptor instead.
func (CreateUserResponse_CreationStatus) EnumDescriptor() ([]byte, []int) {
	return file_userservice_proto_rawDescGZIP(), []int{7, 0}
}

type ActivateUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *ActivateUserRequest) Reset() {
	*x = ActivateUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userservice_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActivateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivateUserRequest) ProtoMessage() {}

func (x *ActivateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_userservice_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivateUserRequest.ProtoReflect.Descriptor instead.
func (*ActivateUserRequest) Descriptor() ([]byte, []int) {
	return file_userservice_proto_rawDescGZIP(), []int{0}
}

func (x *ActivateUserRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type UserByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *UserByIdRequest) Reset() {
	*x = UserByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userservice_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserByIdRequest) ProtoMessage() {}

func (x *UserByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_userservice_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserByIdRequest.ProtoReflect.Descriptor instead.
func (*UserByIdRequest) Descriptor() ([]byte, []int) {
	return file_userservice_proto_rawDescGZIP(), []int{1}
}

func (x *UserByIdRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type UserByIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId       string       `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	PrivacyLevel PrivacyLevel `protobuf:"varint,2,opt,name=privacyLevel,proto3,enum=usergrpc.PrivacyLevel" json:"privacyLevel,omitempty"`
	Active       bool         `protobuf:"varint,3,opt,name=active,proto3" json:"active,omitempty"`
}

func (x *UserByIdResponse) Reset() {
	*x = UserByIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userservice_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserByIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserByIdResponse) ProtoMessage() {}

func (x *UserByIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_userservice_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserByIdResponse.ProtoReflect.Descriptor instead.
func (*UserByIdResponse) Descriptor() ([]byte, []int) {
	return file_userservice_proto_rawDescGZIP(), []int{2}
}

func (x *UserByIdResponse) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UserByIdResponse) GetPrivacyLevel() PrivacyLevel {
	if x != nil {
		return x.PrivacyLevel
	}
	return PrivacyLevel_PUBLIC
}

func (x *UserByIdResponse) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

type ActivateUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status bool   `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error  string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *ActivateUserResponse) Reset() {
	*x = ActivateUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userservice_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActivateUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivateUserResponse) ProtoMessage() {}

func (x *ActivateUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_userservice_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivateUserResponse.ProtoReflect.Descriptor instead.
func (*ActivateUserResponse) Descriptor() ([]byte, []int) {
	return file_userservice_proto_rawDescGZIP(), []int{3}
}

func (x *ActivateUserResponse) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *ActivateUserResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type LoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Password string `protobuf:"bytes,1,opt,name=password,proto3" json:"password,omitempty"`
	// Types that are assignable to UserCredential:
	//	*LoginRequest_Username
	//	*LoginRequest_Email
	UserCredential isLoginRequest_UserCredential `protobuf_oneof:"userCredential"`
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userservice_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_userservice_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_userservice_proto_rawDescGZIP(), []int{4}
}

func (x *LoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (m *LoginRequest) GetUserCredential() isLoginRequest_UserCredential {
	if m != nil {
		return m.UserCredential
	}
	return nil
}

func (x *LoginRequest) GetUsername() string {
	if x, ok := x.GetUserCredential().(*LoginRequest_Username); ok {
		return x.Username
	}
	return ""
}

func (x *LoginRequest) GetEmail() string {
	if x, ok := x.GetUserCredential().(*LoginRequest_Email); ok {
		return x.Email
	}
	return ""
}

type isLoginRequest_UserCredential interface {
	isLoginRequest_UserCredential()
}

type LoginRequest_Username struct {
	Username string `protobuf:"bytes,2,opt,name=username,proto3,oneof"`
}

type LoginRequest_Email struct {
	Email string `protobuf:"bytes,3,opt,name=email,proto3,oneof"`
}

func (*LoginRequest_Username) isLoginRequest_UserCredential() {}

func (*LoginRequest_Email) isLoginRequest_UserCredential() {}

type LoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status LoginResponse_LoginStatus `protobuf:"varint,1,opt,name=status,proto3,enum=usergrpc.LoginResponse_LoginStatus" json:"status,omitempty"`
	UserId string                    `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
	Role   string                    `protobuf:"bytes,3,opt,name=role,proto3" json:"role,omitempty"`
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userservice_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_userservice_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse.ProtoReflect.Descriptor instead.
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return file_userservice_proto_rawDescGZIP(), []int{5}
}

func (x *LoginResponse) GetStatus() LoginResponse_LoginStatus {
	if x != nil {
		return x.Status
	}
	return LoginResponse_SUCCESS
}

func (x *LoginResponse) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *LoginResponse) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

type CreateUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string                         `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Username  string                         `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Email     string                         `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Firstname string                         `protobuf:"bytes,4,opt,name=firstname,proto3" json:"firstname,omitempty"`
	Lastname  string                         `protobuf:"bytes,5,opt,name=lastname,proto3" json:"lastname,omitempty"`
	Level     CreateUserRequest_PrivacyLevel `protobuf:"varint,6,opt,name=level,proto3,enum=usergrpc.CreateUserRequest_PrivacyLevel" json:"level,omitempty"`
}

func (x *CreateUserRequest) Reset() {
	*x = CreateUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userservice_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserRequest) ProtoMessage() {}

func (x *CreateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_userservice_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserRequest.ProtoReflect.Descriptor instead.
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return file_userservice_proto_rawDescGZIP(), []int{6}
}

func (x *CreateUserRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateUserRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *CreateUserRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreateUserRequest) GetFirstname() string {
	if x != nil {
		return x.Firstname
	}
	return ""
}

func (x *CreateUserRequest) GetLastname() string {
	if x != nil {
		return x.Lastname
	}
	return ""
}

func (x *CreateUserRequest) GetLevel() CreateUserRequest_PrivacyLevel {
	if x != nil {
		return x.Level
	}
	return CreateUserRequest_PUBLIC
}

type CreateUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status       CreateUserResponse_CreationStatus `protobuf:"varint,1,opt,name=status,proto3,enum=usergrpc.CreateUserResponse_CreationStatus" json:"status,omitempty"`
	ErrorMessage string                            `protobuf:"bytes,2,opt,name=errorMessage,proto3" json:"errorMessage,omitempty"`
}

func (x *CreateUserResponse) Reset() {
	*x = CreateUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userservice_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserResponse) ProtoMessage() {}

func (x *CreateUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_userservice_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserResponse.ProtoReflect.Descriptor instead.
func (*CreateUserResponse) Descriptor() ([]byte, []int) {
	return file_userservice_proto_rawDescGZIP(), []int{7}
}

func (x *CreateUserResponse) GetStatus() CreateUserResponse_CreationStatus {
	if x != nil {
		return x.Status
	}
	return CreateUserResponse_SUCCESS
}

func (x *CreateUserResponse) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

var File_userservice_proto protoreflect.FileDescriptor

var file_userservice_proto_rawDesc = []byte{
	0x0a, 0x11, 0x75, 0x73, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x08, 0x75, 0x73, 0x65, 0x72, 0x67, 0x72, 0x70, 0x63, 0x22, 0x2d, 0x0a,
	0x13, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x29, 0x0a, 0x0f,
	0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x7e, 0x0a, 0x10, 0x55, 0x73, 0x65, 0x72, 0x42,
	0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x3a, 0x0a, 0x0c, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x4c, 0x65,
	0x76, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x4c, 0x65, 0x76, 0x65,
	0x6c, 0x52, 0x0c, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x22, 0x44, 0x0a, 0x14, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x72, 0x0a,
	0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1c, 0x0a, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x42,
	0x10, 0x0a, 0x0e, 0x75, 0x73, 0x65, 0x72, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61,
	0x6c, 0x22, 0xc0, 0x01, 0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x22, 0x46, 0x0a, 0x0b,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x53,
	0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x49, 0x4e, 0x56, 0x41,
	0x4c, 0x49, 0x44, 0x5f, 0x43, 0x52, 0x45, 0x44, 0x45, 0x4e, 0x54, 0x49, 0x41, 0x4c, 0x53, 0x10,
	0x01, 0x12, 0x11, 0x0a, 0x0d, 0x4e, 0x4f, 0x54, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x41, 0x54,
	0x45, 0x44, 0x10, 0x02, 0x22, 0x87, 0x02, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1c, 0x0a, 0x09,
	0x66, 0x69, 0x72, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61,
	0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61,
	0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3e, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x28, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x2e, 0x50, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52,
	0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x22, 0x36, 0x0a, 0x0c, 0x50, 0x72, 0x69, 0x76, 0x61, 0x63,
	0x79, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x0a, 0x0a, 0x06, 0x50, 0x55, 0x42, 0x4c, 0x49, 0x43,
	0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x50, 0x52, 0x4f, 0x54, 0x45, 0x43, 0x54, 0x45, 0x44, 0x10,
	0x01, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x52, 0x49, 0x56, 0x41, 0x54, 0x45, 0x10, 0x02, 0x22, 0xd5,
	0x01, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x56,
	0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x00, 0x12, 0x10, 0x0a,
	0x0c, 0x45, 0x4d, 0x41, 0x49, 0x4c, 0x5f, 0x45, 0x58, 0x49, 0x53, 0x54, 0x53, 0x10, 0x01, 0x12,
	0x13, 0x0a, 0x0f, 0x55, 0x53, 0x45, 0x52, 0x4e, 0x41, 0x4d, 0x45, 0x5f, 0x45, 0x58, 0x49, 0x53,
	0x54, 0x53, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f,
	0x44, 0x41, 0x54, 0x41, 0x10, 0x03, 0x2a, 0x36, 0x0a, 0x0c, 0x50, 0x72, 0x69, 0x76, 0x61, 0x63,
	0x79, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x0a, 0x0a, 0x06, 0x50, 0x55, 0x42, 0x4c, 0x49, 0x43,
	0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x50, 0x52, 0x4f, 0x54, 0x45, 0x43, 0x54, 0x45, 0x44, 0x10,
	0x01, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x52, 0x49, 0x56, 0x41, 0x54, 0x45, 0x10, 0x02, 0x32, 0xb4,
	0x02, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4f,
	0x0a, 0x0c, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1d,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x3e, 0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x4b, 0x0a, 0x0c, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x1b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x0c,
	0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x12, 0x5a, 0x10, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_userservice_proto_rawDescOnce sync.Once
	file_userservice_proto_rawDescData = file_userservice_proto_rawDesc
)

func file_userservice_proto_rawDescGZIP() []byte {
	file_userservice_proto_rawDescOnce.Do(func() {
		file_userservice_proto_rawDescData = protoimpl.X.CompressGZIP(file_userservice_proto_rawDescData)
	})
	return file_userservice_proto_rawDescData
}

var file_userservice_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_userservice_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_userservice_proto_goTypes = []interface{}{
	(PrivacyLevel)(0),                      // 0: usergrpc.PrivacyLevel
	(LoginResponse_LoginStatus)(0),         // 1: usergrpc.LoginResponse.LoginStatus
	(CreateUserRequest_PrivacyLevel)(0),    // 2: usergrpc.CreateUserRequest.PrivacyLevel
	(CreateUserResponse_CreationStatus)(0), // 3: usergrpc.CreateUserResponse.CreationStatus
	(*ActivateUserRequest)(nil),            // 4: usergrpc.ActivateUserRequest
	(*UserByIdRequest)(nil),                // 5: usergrpc.UserByIdRequest
	(*UserByIdResponse)(nil),               // 6: usergrpc.UserByIdResponse
	(*ActivateUserResponse)(nil),           // 7: usergrpc.ActivateUserResponse
	(*LoginRequest)(nil),                   // 8: usergrpc.LoginRequest
	(*LoginResponse)(nil),                  // 9: usergrpc.LoginResponse
	(*CreateUserRequest)(nil),              // 10: usergrpc.CreateUserRequest
	(*CreateUserResponse)(nil),             // 11: usergrpc.CreateUserResponse
}
var file_userservice_proto_depIdxs = []int32{
	0,  // 0: usergrpc.UserByIdResponse.privacyLevel:type_name -> usergrpc.PrivacyLevel
	1,  // 1: usergrpc.LoginResponse.status:type_name -> usergrpc.LoginResponse.LoginStatus
	2,  // 2: usergrpc.CreateUserRequest.level:type_name -> usergrpc.CreateUserRequest.PrivacyLevel
	3,  // 3: usergrpc.CreateUserResponse.status:type_name -> usergrpc.CreateUserResponse.CreationStatus
	4,  // 4: usergrpc.UserService.ActivateUser:input_type -> usergrpc.ActivateUserRequest
	8,  // 5: usergrpc.UserService.LoginUser:input_type -> usergrpc.LoginRequest
	10, // 6: usergrpc.UserService.RegisterUser:input_type -> usergrpc.CreateUserRequest
	5,  // 7: usergrpc.UserService.UserByUserId:input_type -> usergrpc.UserByIdRequest
	7,  // 8: usergrpc.UserService.ActivateUser:output_type -> usergrpc.ActivateUserResponse
	9,  // 9: usergrpc.UserService.LoginUser:output_type -> usergrpc.LoginResponse
	11, // 10: usergrpc.UserService.RegisterUser:output_type -> usergrpc.CreateUserResponse
	6,  // 11: usergrpc.UserService.UserByUserId:output_type -> usergrpc.UserByIdResponse
	8,  // [8:12] is the sub-list for method output_type
	4,  // [4:8] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_userservice_proto_init() }
func file_userservice_proto_init() {
	if File_userservice_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_userservice_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActivateUserRequest); i {
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
		file_userservice_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserByIdRequest); i {
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
		file_userservice_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserByIdResponse); i {
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
		file_userservice_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActivateUserResponse); i {
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
		file_userservice_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRequest); i {
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
		file_userservice_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginResponse); i {
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
		file_userservice_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserRequest); i {
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
		file_userservice_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserResponse); i {
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
	file_userservice_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*LoginRequest_Username)(nil),
		(*LoginRequest_Email)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_userservice_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_userservice_proto_goTypes,
		DependencyIndexes: file_userservice_proto_depIdxs,
		EnumInfos:         file_userservice_proto_enumTypes,
		MessageInfos:      file_userservice_proto_msgTypes,
	}.Build()
	File_userservice_proto = out.File
	file_userservice_proto_rawDesc = nil
	file_userservice_proto_goTypes = nil
	file_userservice_proto_depIdxs = nil
}
