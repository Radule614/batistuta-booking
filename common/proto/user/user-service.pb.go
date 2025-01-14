// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.3
// source: user/user-service.proto

package user

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type UserRole int32

const (
	UserRole_Host  UserRole = 0
	UserRole_Guest UserRole = 1
)

// Enum value maps for UserRole.
var (
	UserRole_name = map[int32]string{
		0: "Host",
		1: "Guest",
	}
	UserRole_value = map[string]int32{
		"Host":  0,
		"Guest": 1,
	}
)

func (x UserRole) Enum() *UserRole {
	p := new(UserRole)
	*p = x
	return p
}

func (x UserRole) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserRole) Descriptor() protoreflect.EnumDescriptor {
	return file_user_user_service_proto_enumTypes[0].Descriptor()
}

func (UserRole) Type() protoreflect.EnumType {
	return &file_user_user_service_proto_enumTypes[0]
}

func (x UserRole) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserRole.Descriptor instead.
func (UserRole) EnumDescriptor() ([]byte, []int) {
	return file_user_user_service_proto_rawDescGZIP(), []int{0}
}

type Empty_Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty_Message) Reset() {
	*x = Empty_Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_user_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty_Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty_Message) ProtoMessage() {}

func (x *Empty_Message) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty_Message.ProtoReflect.Descriptor instead.
func (*Empty_Message) Descriptor() ([]byte, []int) {
	return file_user_user_service_proto_rawDescGZIP(), []int{0}
}

type UserLessInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Role        UserRole `protobuf:"varint,1,opt,name=Role,proto3,enum=user.UserRole" json:"Role,omitempty"`
	Username    string   `protobuf:"bytes,2,opt,name=Username,proto3" json:"Username,omitempty"`
	FirstName   string   `protobuf:"bytes,3,opt,name=FirstName,proto3" json:"FirstName,omitempty"`
	LastName    string   `protobuf:"bytes,4,opt,name=LastName,proto3" json:"LastName,omitempty"`
	Email       string   `protobuf:"bytes,5,opt,name=Email,proto3" json:"Email,omitempty"`
	LivingPlace string   `protobuf:"bytes,6,opt,name=LivingPlace,proto3" json:"LivingPlace,omitempty"`
}

func (x *UserLessInfo) Reset() {
	*x = UserLessInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_user_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserLessInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserLessInfo) ProtoMessage() {}

func (x *UserLessInfo) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserLessInfo.ProtoReflect.Descriptor instead.
func (*UserLessInfo) Descriptor() ([]byte, []int) {
	return file_user_user_service_proto_rawDescGZIP(), []int{1}
}

func (x *UserLessInfo) GetRole() UserRole {
	if x != nil {
		return x.Role
	}
	return UserRole_Host
}

func (x *UserLessInfo) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserLessInfo) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *UserLessInfo) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *UserLessInfo) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserLessInfo) GetLivingPlace() string {
	if x != nil {
		return x.LivingPlace
	}
	return ""
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string   `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Role        UserRole `protobuf:"varint,2,opt,name=Role,proto3,enum=user.UserRole" json:"Role,omitempty"`
	Username    string   `protobuf:"bytes,3,opt,name=Username,proto3" json:"Username,omitempty"`
	FirstName   string   `protobuf:"bytes,4,opt,name=FirstName,proto3" json:"FirstName,omitempty"`
	LastName    string   `protobuf:"bytes,5,opt,name=LastName,proto3" json:"LastName,omitempty"`
	Email       string   `protobuf:"bytes,6,opt,name=Email,proto3" json:"Email,omitempty"`
	LivingPlace string   `protobuf:"bytes,7,opt,name=LivingPlace,proto3" json:"LivingPlace,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_user_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_user_user_service_proto_rawDescGZIP(), []int{2}
}

func (x *User) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *User) GetRole() UserRole {
	if x != nil {
		return x.Role
	}
	return UserRole_Host
}

func (x *User) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *User) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *User) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetLivingPlace() string {
	if x != nil {
		return x.LivingPlace
	}
	return ""
}

type GetAllUsers_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []*User `protobuf:"bytes,1,rep,name=Users,proto3" json:"Users,omitempty"`
}

func (x *GetAllUsers_Response) Reset() {
	*x = GetAllUsers_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_user_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllUsers_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllUsers_Response) ProtoMessage() {}

func (x *GetAllUsers_Response) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllUsers_Response.ProtoReflect.Descriptor instead.
func (*GetAllUsers_Response) Descriptor() ([]byte, []int) {
	return file_user_user_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetAllUsers_Response) GetUsers() []*User {
	if x != nil {
		return x.Users
	}
	return nil
}

type ChangeUserInfo_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName   *string `protobuf:"bytes,1,opt,name=FirstName,proto3,oneof" json:"FirstName,omitempty"`
	LastName    *string `protobuf:"bytes,2,opt,name=LastName,proto3,oneof" json:"LastName,omitempty"`
	LivingPlace *string `protobuf:"bytes,3,opt,name=LivingPlace,proto3,oneof" json:"LivingPlace,omitempty"`
	Username    string  `protobuf:"bytes,4,opt,name=Username,proto3" json:"Username,omitempty"`
}

func (x *ChangeUserInfo_Request) Reset() {
	*x = ChangeUserInfo_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_user_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeUserInfo_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeUserInfo_Request) ProtoMessage() {}

func (x *ChangeUserInfo_Request) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeUserInfo_Request.ProtoReflect.Descriptor instead.
func (*ChangeUserInfo_Request) Descriptor() ([]byte, []int) {
	return file_user_user_service_proto_rawDescGZIP(), []int{4}
}

func (x *ChangeUserInfo_Request) GetFirstName() string {
	if x != nil && x.FirstName != nil {
		return *x.FirstName
	}
	return ""
}

func (x *ChangeUserInfo_Request) GetLastName() string {
	if x != nil && x.LastName != nil {
		return *x.LastName
	}
	return ""
}

func (x *ChangeUserInfo_Request) GetLivingPlace() string {
	if x != nil && x.LivingPlace != nil {
		return *x.LivingPlace
	}
	return ""
}

func (x *ChangeUserInfo_Request) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type ChangeUserInfo_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success      bool          `protobuf:"varint,1,opt,name=Success,proto3" json:"Success,omitempty"`
	ErrorMessage string        `protobuf:"bytes,2,opt,name=ErrorMessage,proto3" json:"ErrorMessage,omitempty"`
	User         *UserLessInfo `protobuf:"bytes,3,opt,name=User,proto3,oneof" json:"User,omitempty"`
}

func (x *ChangeUserInfo_Response) Reset() {
	*x = ChangeUserInfo_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_user_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeUserInfo_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeUserInfo_Response) ProtoMessage() {}

func (x *ChangeUserInfo_Response) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeUserInfo_Response.ProtoReflect.Descriptor instead.
func (*ChangeUserInfo_Response) Descriptor() ([]byte, []int) {
	return file_user_user_service_proto_rawDescGZIP(), []int{5}
}

func (x *ChangeUserInfo_Response) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *ChangeUserInfo_Response) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

func (x *ChangeUserInfo_Response) GetUser() *UserLessInfo {
	if x != nil {
		return x.User
	}
	return nil
}

type ChangePassword_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CurrentPassword string `protobuf:"bytes,1,opt,name=CurrentPassword,proto3" json:"CurrentPassword,omitempty"`
	NewPassword     string `protobuf:"bytes,2,opt,name=NewPassword,proto3" json:"NewPassword,omitempty"`
}

func (x *ChangePassword_Request) Reset() {
	*x = ChangePassword_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_user_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangePassword_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangePassword_Request) ProtoMessage() {}

func (x *ChangePassword_Request) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangePassword_Request.ProtoReflect.Descriptor instead.
func (*ChangePassword_Request) Descriptor() ([]byte, []int) {
	return file_user_user_service_proto_rawDescGZIP(), []int{6}
}

func (x *ChangePassword_Request) GetCurrentPassword() string {
	if x != nil {
		return x.CurrentPassword
	}
	return ""
}

func (x *ChangePassword_Request) GetNewPassword() string {
	if x != nil {
		return x.NewPassword
	}
	return ""
}

type RatingDTO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	UserId     string `protobuf:"bytes,2,opt,name=UserId,proto3" json:"UserId,omitempty"`
	TargetId   string `protobuf:"bytes,3,opt,name=TargetId,proto3" json:"TargetId,omitempty"`
	TargetType uint32 `protobuf:"varint,4,opt,name=TargetType,proto3" json:"TargetType,omitempty"`
	Value      uint32 `protobuf:"varint,5,opt,name=Value,proto3" json:"Value,omitempty"`
}

func (x *RatingDTO) Reset() {
	*x = RatingDTO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_user_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RatingDTO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RatingDTO) ProtoMessage() {}

func (x *RatingDTO) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RatingDTO.ProtoReflect.Descriptor instead.
func (*RatingDTO) Descriptor() ([]byte, []int) {
	return file_user_user_service_proto_rawDescGZIP(), []int{7}
}

func (x *RatingDTO) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RatingDTO) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *RatingDTO) GetTargetId() string {
	if x != nil {
		return x.TargetId
	}
	return ""
}

func (x *RatingDTO) GetTargetType() uint32 {
	if x != nil {
		return x.TargetType
	}
	return 0
}

func (x *RatingDTO) GetValue() uint32 {
	if x != nil {
		return x.Value
	}
	return 0
}

type HostWithRating struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string       `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Role        UserRole     `protobuf:"varint,2,opt,name=Role,proto3,enum=user.UserRole" json:"Role,omitempty"`
	Username    string       `protobuf:"bytes,3,opt,name=Username,proto3" json:"Username,omitempty"`
	FirstName   string       `protobuf:"bytes,4,opt,name=FirstName,proto3" json:"FirstName,omitempty"`
	LastName    string       `protobuf:"bytes,5,opt,name=LastName,proto3" json:"LastName,omitempty"`
	Email       string       `protobuf:"bytes,6,opt,name=Email,proto3" json:"Email,omitempty"`
	LivingPlace string       `protobuf:"bytes,7,opt,name=LivingPlace,proto3" json:"LivingPlace,omitempty"`
	Ratings     []*RatingDTO `protobuf:"bytes,8,rep,name=Ratings,proto3" json:"Ratings,omitempty"`
}

func (x *HostWithRating) Reset() {
	*x = HostWithRating{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_user_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HostWithRating) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HostWithRating) ProtoMessage() {}

func (x *HostWithRating) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HostWithRating.ProtoReflect.Descriptor instead.
func (*HostWithRating) Descriptor() ([]byte, []int) {
	return file_user_user_service_proto_rawDescGZIP(), []int{8}
}

func (x *HostWithRating) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *HostWithRating) GetRole() UserRole {
	if x != nil {
		return x.Role
	}
	return UserRole_Host
}

func (x *HostWithRating) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *HostWithRating) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *HostWithRating) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *HostWithRating) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *HostWithRating) GetLivingPlace() string {
	if x != nil {
		return x.LivingPlace
	}
	return ""
}

func (x *HostWithRating) GetRatings() []*RatingDTO {
	if x != nil {
		return x.Ratings
	}
	return nil
}

type GetAllHostsWithRatings_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hosts []*HostWithRating `protobuf:"bytes,1,rep,name=Hosts,proto3" json:"Hosts,omitempty"`
}

func (x *GetAllHostsWithRatings_Response) Reset() {
	*x = GetAllHostsWithRatings_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_user_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllHostsWithRatings_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllHostsWithRatings_Response) ProtoMessage() {}

func (x *GetAllHostsWithRatings_Response) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllHostsWithRatings_Response.ProtoReflect.Descriptor instead.
func (*GetAllHostsWithRatings_Response) Descriptor() ([]byte, []int) {
	return file_user_user_service_proto_rawDescGZIP(), []int{9}
}

func (x *GetAllHostsWithRatings_Response) GetHosts() []*HostWithRating {
	if x != nil {
		return x.Hosts
	}
	return nil
}

var File_user_user_service_proto protoreflect.FileDescriptor

var file_user_user_service_proto_rawDesc = []byte{
	0x0a, 0x17, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x75, 0x73, 0x65, 0x72, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x0f, 0x0a,
	0x0d, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x5f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xc0,
	0x01, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x22, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04, 0x52,
	0x6f, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x46, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x46, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x4c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x4c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x20, 0x0a, 0x0b, 0x4c, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x4c, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x50, 0x6c, 0x61, 0x63,
	0x65, 0x22, 0xc8, 0x01, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x04, 0x52, 0x6f,
	0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x46, 0x69,
	0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x46,
	0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x4c, 0x61, 0x73, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4c, 0x61, 0x73, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x4c, 0x69,
	0x76, 0x69, 0x6e, 0x67, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x4c, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x22, 0x38, 0x0a, 0x14,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x73, 0x5f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x05, 0x55, 0x73, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x05, 0x55, 0x73, 0x65, 0x72, 0x73, 0x22, 0xca, 0x01, 0x0a, 0x16, 0x43, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x21, 0x0a, 0x09, 0x46, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x09, 0x46, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x4c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x08, 0x4c, 0x61, 0x73, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x25, 0x0a, 0x0b, 0x4c, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x50,
	0x6c, 0x61, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x0b, 0x4c, 0x69,
	0x76, 0x69, 0x6e, 0x67, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1a, 0x0a, 0x08,
	0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x46, 0x69, 0x72,
	0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x4c, 0x61, 0x73, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x4c, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x50, 0x6c,
	0x61, 0x63, 0x65, 0x22, 0x8d, 0x01, 0x0a, 0x17, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2b, 0x0a,
	0x04, 0x55, 0x73, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x48,
	0x00, 0x52, 0x04, 0x55, 0x73, 0x65, 0x72, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x55,
	0x73, 0x65, 0x72, 0x22, 0x64, 0x0a, 0x16, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a,
	0x0f, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x50,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x4e, 0x65, 0x77, 0x50, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x4e, 0x65,
	0x77, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x85, 0x01, 0x0a, 0x09, 0x52, 0x61,
	0x74, 0x69, 0x6e, 0x67, 0x44, 0x54, 0x4f, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x54,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0a, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0xfd, 0x01, 0x0a, 0x0e, 0x48, 0x6f, 0x73, 0x74, 0x57, 0x69, 0x74, 0x68, 0x52, 0x61,
	0x74, 0x69, 0x6e, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f,
	0x6c, 0x65, 0x52, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x46, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x46, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x4c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x4c, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x50, 0x6c,
	0x61, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x4c, 0x69, 0x76, 0x69, 0x6e,
	0x67, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x12, 0x29, 0x0a, 0x07, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67,
	0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x52,
	0x61, 0x74, 0x69, 0x6e, 0x67, 0x44, 0x54, 0x4f, 0x52, 0x07, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67,
	0x73, 0x22, 0x4d, 0x0a, 0x1f, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x48, 0x6f, 0x73, 0x74, 0x73,
	0x57, 0x69, 0x74, 0x68, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x5f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x05, 0x48, 0x6f, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x57,
	0x69, 0x74, 0x68, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x05, 0x48, 0x6f, 0x73, 0x74, 0x73,
	0x2a, 0x1f, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x08, 0x0a, 0x04,
	0x48, 0x6f, 0x73, 0x74, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x47, 0x75, 0x65, 0x73, 0x74, 0x10,
	0x01, 0x32, 0xa1, 0x03, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x52, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x73,
	0x12, 0x13, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x5f, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x1a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x73, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x12, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x12, 0x0a, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x12, 0x6f, 0x0a, 0x0e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1c, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x43,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x5f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x5f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x3a, 0x01, 0x2a, 0x32,
	0x15, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x55, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x12, 0x63, 0x0a, 0x0e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1c, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x5f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x5f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x1e, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x18, 0x3a, 0x01, 0x2a, 0x32, 0x13, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x2f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x68, 0x0a, 0x16, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x48, 0x6f, 0x73, 0x74, 0x73, 0x57, 0x69, 0x74, 0x68, 0x52, 0x61,
	0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x13, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x5f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x25, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x48, 0x6f, 0x73, 0x74, 0x73, 0x57, 0x69, 0x74,
	0x68, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x12, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x12, 0x0a, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x68, 0x6f, 0x73, 0x74, 0x73, 0x42, 0x0c, 0x5a, 0x0a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_user_service_proto_rawDescOnce sync.Once
	file_user_user_service_proto_rawDescData = file_user_user_service_proto_rawDesc
)

func file_user_user_service_proto_rawDescGZIP() []byte {
	file_user_user_service_proto_rawDescOnce.Do(func() {
		file_user_user_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_user_service_proto_rawDescData)
	})
	return file_user_user_service_proto_rawDescData
}

var file_user_user_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_user_user_service_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_user_user_service_proto_goTypes = []interface{}{
	(UserRole)(0),                           // 0: user.UserRole
	(*Empty_Message)(nil),                   // 1: user.Empty_Message
	(*UserLessInfo)(nil),                    // 2: user.UserLessInfo
	(*User)(nil),                            // 3: user.User
	(*GetAllUsers_Response)(nil),            // 4: user.GetAllUsers_Response
	(*ChangeUserInfo_Request)(nil),          // 5: user.ChangeUserInfo_Request
	(*ChangeUserInfo_Response)(nil),         // 6: user.ChangeUserInfo_Response
	(*ChangePassword_Request)(nil),          // 7: user.ChangePassword_Request
	(*RatingDTO)(nil),                       // 8: user.RatingDTO
	(*HostWithRating)(nil),                  // 9: user.HostWithRating
	(*GetAllHostsWithRatings_Response)(nil), // 10: user.GetAllHostsWithRatings_Response
}
var file_user_user_service_proto_depIdxs = []int32{
	0,  // 0: user.UserLessInfo.Role:type_name -> user.UserRole
	0,  // 1: user.User.Role:type_name -> user.UserRole
	3,  // 2: user.GetAllUsers_Response.Users:type_name -> user.User
	2,  // 3: user.ChangeUserInfo_Response.User:type_name -> user.UserLessInfo
	0,  // 4: user.HostWithRating.Role:type_name -> user.UserRole
	8,  // 5: user.HostWithRating.Ratings:type_name -> user.RatingDTO
	9,  // 6: user.GetAllHostsWithRatings_Response.Hosts:type_name -> user.HostWithRating
	1,  // 7: user.UserService.GetAllUsers:input_type -> user.Empty_Message
	5,  // 8: user.UserService.ChangeUserInfo:input_type -> user.ChangeUserInfo_Request
	7,  // 9: user.UserService.ChangePassword:input_type -> user.ChangePassword_Request
	1,  // 10: user.UserService.GetAllHostsWithRatings:input_type -> user.Empty_Message
	4,  // 11: user.UserService.GetAllUsers:output_type -> user.GetAllUsers_Response
	6,  // 12: user.UserService.ChangeUserInfo:output_type -> user.ChangeUserInfo_Response
	1,  // 13: user.UserService.ChangePassword:output_type -> user.Empty_Message
	10, // 14: user.UserService.GetAllHostsWithRatings:output_type -> user.GetAllHostsWithRatings_Response
	11, // [11:15] is the sub-list for method output_type
	7,  // [7:11] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_user_user_service_proto_init() }
func file_user_user_service_proto_init() {
	if File_user_user_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_user_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty_Message); i {
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
		file_user_user_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserLessInfo); i {
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
		file_user_user_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_user_user_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllUsers_Response); i {
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
		file_user_user_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeUserInfo_Request); i {
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
		file_user_user_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeUserInfo_Response); i {
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
		file_user_user_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangePassword_Request); i {
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
		file_user_user_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RatingDTO); i {
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
		file_user_user_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HostWithRating); i {
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
		file_user_user_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllHostsWithRatings_Response); i {
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
	file_user_user_service_proto_msgTypes[4].OneofWrappers = []interface{}{}
	file_user_user_service_proto_msgTypes[5].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_user_user_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_user_service_proto_goTypes,
		DependencyIndexes: file_user_user_service_proto_depIdxs,
		EnumInfos:         file_user_user_service_proto_enumTypes,
		MessageInfos:      file_user_user_service_proto_msgTypes,
	}.Build()
	File_user_user_service_proto = out.File
	file_user_user_service_proto_rawDesc = nil
	file_user_user_service_proto_goTypes = nil
	file_user_user_service_proto_depIdxs = nil
}
