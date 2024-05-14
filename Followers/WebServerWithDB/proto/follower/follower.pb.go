// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0--rc2
// source: follower/follower.proto

package follower

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UserFollowingDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId            string `protobuf:"bytes,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	Username          string `protobuf:"bytes,2,opt,name=Username,proto3" json:"Username,omitempty"`
	Image             string `protobuf:"bytes,3,opt,name=Image,proto3" json:"Image,omitempty"`
	FollowingUserId   string `protobuf:"bytes,4,opt,name=FollowingUserId,proto3" json:"FollowingUserId,omitempty"`
	FollowingUsername string `protobuf:"bytes,5,opt,name=FollowingUsername,proto3" json:"FollowingUsername,omitempty"`
	FollowingImage    string `protobuf:"bytes,6,opt,name=FollowingImage,proto3" json:"FollowingImage,omitempty"`
}

func (x *UserFollowingDto) Reset() {
	*x = UserFollowingDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_follower_follower_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserFollowingDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserFollowingDto) ProtoMessage() {}

func (x *UserFollowingDto) ProtoReflect() protoreflect.Message {
	mi := &file_follower_follower_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserFollowingDto.ProtoReflect.Descriptor instead.
func (*UserFollowingDto) Descriptor() ([]byte, []int) {
	return file_follower_follower_proto_rawDescGZIP(), []int{0}
}

func (x *UserFollowingDto) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UserFollowingDto) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserFollowingDto) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *UserFollowingDto) GetFollowingUserId() string {
	if x != nil {
		return x.FollowingUserId
	}
	return ""
}

func (x *UserFollowingDto) GetFollowingUsername() string {
	if x != nil {
		return x.FollowingUsername
	}
	return ""
}

func (x *UserFollowingDto) GetFollowingImage() string {
	if x != nil {
		return x.FollowingImage
	}
	return ""
}

type FollowingResponseDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=Username,proto3" json:"Username,omitempty"`
	Image    string `protobuf:"bytes,3,opt,name=Image,proto3" json:"Image,omitempty"`
}

func (x *FollowingResponseDto) Reset() {
	*x = FollowingResponseDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_follower_follower_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FollowingResponseDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FollowingResponseDto) ProtoMessage() {}

func (x *FollowingResponseDto) ProtoReflect() protoreflect.Message {
	mi := &file_follower_follower_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FollowingResponseDto.ProtoReflect.Descriptor instead.
func (*FollowingResponseDto) Descriptor() ([]byte, []int) {
	return file_follower_follower_proto_rawDescGZIP(), []int{1}
}

func (x *FollowingResponseDto) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *FollowingResponseDto) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *FollowingResponseDto) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

type FollowerResponseDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	UserId       int64 `protobuf:"varint,2,opt,name=UserId,proto3" json:"UserId,omitempty"`
	FollowedById int64 `protobuf:"varint,3,opt,name=FollowedById,proto3" json:"FollowedById,omitempty"`
}

func (x *FollowerResponseDto) Reset() {
	*x = FollowerResponseDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_follower_follower_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FollowerResponseDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FollowerResponseDto) ProtoMessage() {}

func (x *FollowerResponseDto) ProtoReflect() protoreflect.Message {
	mi := &file_follower_follower_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FollowerResponseDto.ProtoReflect.Descriptor instead.
func (*FollowerResponseDto) Descriptor() ([]byte, []int) {
	return file_follower_follower_proto_rawDescGZIP(), []int{2}
}

func (x *FollowerResponseDto) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *FollowerResponseDto) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *FollowerResponseDto) GetFollowedById() int64 {
	if x != nil {
		return x.FollowedById
	}
	return 0
}

type Id struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *Id) Reset() {
	*x = Id{}
	if protoimpl.UnsafeEnabled {
		mi := &file_follower_follower_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Id) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Id) ProtoMessage() {}

func (x *Id) ProtoReflect() protoreflect.Message {
	mi := &file_follower_follower_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Id.ProtoReflect.Descriptor instead.
func (*Id) Descriptor() ([]byte, []int) {
	return file_follower_follower_proto_rawDescGZIP(), []int{3}
}

func (x *Id) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ListFollowingResponseDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResponseList []*FollowingResponseDto `protobuf:"bytes,1,rep,name=ResponseList,proto3" json:"ResponseList,omitempty"`
}

func (x *ListFollowingResponseDto) Reset() {
	*x = ListFollowingResponseDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_follower_follower_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFollowingResponseDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFollowingResponseDto) ProtoMessage() {}

func (x *ListFollowingResponseDto) ProtoReflect() protoreflect.Message {
	mi := &file_follower_follower_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFollowingResponseDto.ProtoReflect.Descriptor instead.
func (*ListFollowingResponseDto) Descriptor() ([]byte, []int) {
	return file_follower_follower_proto_rawDescGZIP(), []int{4}
}

func (x *ListFollowingResponseDto) GetResponseList() []*FollowingResponseDto {
	if x != nil {
		return x.ResponseList
	}
	return nil
}

var File_follower_follower_proto protoreflect.FileDescriptor

var file_follower_follower_proto_rawDesc = []byte{
	0x0a, 0x17, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x2f, 0x66, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdc, 0x01, 0x0a, 0x10, 0x55,
	0x73, 0x65, 0x72, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x44, 0x74, 0x6f, 0x12,
	0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x46, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0f, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x11, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67,
	0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11,
	0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x26, 0x0a, 0x0e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x69, 0x6e, 0x67, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x22, 0x58, 0x0a, 0x14, 0x46, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x74,
	0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x22, 0x61, 0x0a, 0x13, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x74, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x42, 0x79,
	0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x65, 0x64, 0x42, 0x79, 0x49, 0x64, 0x22, 0x14, 0x0a, 0x02, 0x69, 0x64, 0x12, 0x0e, 0x0a, 0x02,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x22, 0x55, 0x0a, 0x18,
	0x4c, 0x69, 0x73, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x74, 0x6f, 0x12, 0x39, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x44, 0x74, 0x6f, 0x52, 0x0c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4c,
	0x69, 0x73, 0x74, 0x32, 0xf0, 0x01, 0x0a, 0x08, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72,
	0x12, 0x3f, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x77, 0x46, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x12, 0x11, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x46, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x44, 0x74, 0x6f, 0x1a, 0x14, 0x2e, 0x46, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x74, 0x6f, 0x22,
	0x00, 0x12, 0x3e, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72,
	0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x03, 0x2e, 0x69, 0x64, 0x1a, 0x19, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x74, 0x6f, 0x22,
	0x00, 0x12, 0x31, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e,
	0x67, 0x73, 0x12, 0x03, 0x2e, 0x69, 0x64, 0x1a, 0x19, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44,
	0x74, 0x6f, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x65, 0x72, 0x73, 0x12, 0x03, 0x2e, 0x69, 0x64, 0x1a, 0x19, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x44, 0x74, 0x6f, 0x22, 0x00, 0x42, 0x10, 0x5a, 0x0e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_follower_follower_proto_rawDescOnce sync.Once
	file_follower_follower_proto_rawDescData = file_follower_follower_proto_rawDesc
)

func file_follower_follower_proto_rawDescGZIP() []byte {
	file_follower_follower_proto_rawDescOnce.Do(func() {
		file_follower_follower_proto_rawDescData = protoimpl.X.CompressGZIP(file_follower_follower_proto_rawDescData)
	})
	return file_follower_follower_proto_rawDescData
}

var file_follower_follower_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_follower_follower_proto_goTypes = []interface{}{
	(*UserFollowingDto)(nil),         // 0: UserFollowingDto
	(*FollowingResponseDto)(nil),     // 1: FollowingResponseDto
	(*FollowerResponseDto)(nil),      // 2: FollowerResponseDto
	(*Id)(nil),                       // 3: id
	(*ListFollowingResponseDto)(nil), // 4: ListFollowingResponseDto
}
var file_follower_follower_proto_depIdxs = []int32{
	1, // 0: ListFollowingResponseDto.ResponseList:type_name -> FollowingResponseDto
	0, // 1: Follower.CreateNewFollowing:input_type -> UserFollowingDto
	3, // 2: Follower.GetFollowerRecommendations:input_type -> id
	3, // 3: Follower.GetFollowings:input_type -> id
	3, // 4: Follower.GetFollowers:input_type -> id
	2, // 5: Follower.CreateNewFollowing:output_type -> FollowerResponseDto
	4, // 6: Follower.GetFollowerRecommendations:output_type -> ListFollowingResponseDto
	4, // 7: Follower.GetFollowings:output_type -> ListFollowingResponseDto
	4, // 8: Follower.GetFollowers:output_type -> ListFollowingResponseDto
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_follower_follower_proto_init() }
func file_follower_follower_proto_init() {
	if File_follower_follower_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_follower_follower_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserFollowingDto); i {
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
		file_follower_follower_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FollowingResponseDto); i {
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
		file_follower_follower_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FollowerResponseDto); i {
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
		file_follower_follower_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Id); i {
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
		file_follower_follower_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFollowingResponseDto); i {
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
			RawDescriptor: file_follower_follower_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_follower_follower_proto_goTypes,
		DependencyIndexes: file_follower_follower_proto_depIdxs,
		MessageInfos:      file_follower_follower_proto_msgTypes,
	}.Build()
	File_follower_follower_proto = out.File
	file_follower_follower_proto_rawDesc = nil
	file_follower_follower_proto_goTypes = nil
	file_follower_follower_proto_depIdxs = nil
}
