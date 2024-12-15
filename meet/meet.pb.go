// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.3
// source: meet/meet.proto

package meet

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

type CreateMeetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstId  int64  `protobuf:"varint,1,opt,name=first_id,json=firstId,proto3" json:"first_id,omitempty"`
	SecondId int64  `protobuf:"varint,2,opt,name=second_id,json=secondId,proto3" json:"second_id,omitempty"`
	Date     string `protobuf:"bytes,3,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *CreateMeetRequest) Reset() {
	*x = CreateMeetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meet_meet_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMeetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMeetRequest) ProtoMessage() {}

func (x *CreateMeetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_meet_meet_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMeetRequest.ProtoReflect.Descriptor instead.
func (*CreateMeetRequest) Descriptor() ([]byte, []int) {
	return file_meet_meet_proto_rawDescGZIP(), []int{0}
}

func (x *CreateMeetRequest) GetFirstId() int64 {
	if x != nil {
		return x.FirstId
	}
	return 0
}

func (x *CreateMeetRequest) GetSecondId() int64 {
	if x != nil {
		return x.SecondId
	}
	return 0
}

func (x *CreateMeetRequest) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

type CreateMeetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MeetId int64 `protobuf:"varint,1,opt,name=meet_id,json=meetId,proto3" json:"meet_id,omitempty"`
}

func (x *CreateMeetResponse) Reset() {
	*x = CreateMeetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meet_meet_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMeetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMeetResponse) ProtoMessage() {}

func (x *CreateMeetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_meet_meet_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMeetResponse.ProtoReflect.Descriptor instead.
func (*CreateMeetResponse) Descriptor() ([]byte, []int) {
	return file_meet_meet_proto_rawDescGZIP(), []int{1}
}

func (x *CreateMeetResponse) GetMeetId() int64 {
	if x != nil {
		return x.MeetId
	}
	return 0
}

type ApproveMeetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MeetId int64 `protobuf:"varint,1,opt,name=meet_id,json=meetId,proto3" json:"meet_id,omitempty"`
}

func (x *ApproveMeetRequest) Reset() {
	*x = ApproveMeetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meet_meet_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApproveMeetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApproveMeetRequest) ProtoMessage() {}

func (x *ApproveMeetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_meet_meet_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApproveMeetRequest.ProtoReflect.Descriptor instead.
func (*ApproveMeetRequest) Descriptor() ([]byte, []int) {
	return file_meet_meet_proto_rawDescGZIP(), []int{2}
}

func (x *ApproveMeetRequest) GetMeetId() int64 {
	if x != nil {
		return x.MeetId
	}
	return 0
}

type ApproveMeetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ApproveMeetResponse) Reset() {
	*x = ApproveMeetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meet_meet_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApproveMeetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApproveMeetResponse) ProtoMessage() {}

func (x *ApproveMeetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_meet_meet_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApproveMeetResponse.ProtoReflect.Descriptor instead.
func (*ApproveMeetResponse) Descriptor() ([]byte, []int) {
	return file_meet_meet_proto_rawDescGZIP(), []int{3}
}

type GetFutureRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetFutureRequest) Reset() {
	*x = GetFutureRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meet_meet_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFutureRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFutureRequest) ProtoMessage() {}

func (x *GetFutureRequest) ProtoReflect() protoreflect.Message {
	mi := &file_meet_meet_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFutureRequest.ProtoReflect.Descriptor instead.
func (*GetFutureRequest) Descriptor() ([]byte, []int) {
	return file_meet_meet_proto_rawDescGZIP(), []int{4}
}

func (x *GetFutureRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetFutureResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meets []*Meet `protobuf:"bytes,1,rep,name=meets,proto3" json:"meets,omitempty"`
}

func (x *GetFutureResponse) Reset() {
	*x = GetFutureResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meet_meet_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFutureResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFutureResponse) ProtoMessage() {}

func (x *GetFutureResponse) ProtoReflect() protoreflect.Message {
	mi := &file_meet_meet_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFutureResponse.ProtoReflect.Descriptor instead.
func (*GetFutureResponse) Descriptor() ([]byte, []int) {
	return file_meet_meet_proto_rawDescGZIP(), []int{5}
}

func (x *GetFutureResponse) GetMeets() []*Meet {
	if x != nil {
		return x.Meets
	}
	return nil
}

type GetHistoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    int64  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	StartDate string `protobuf:"bytes,2,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	EndDate   string `protobuf:"bytes,3,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
}

func (x *GetHistoryRequest) Reset() {
	*x = GetHistoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meet_meet_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHistoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHistoryRequest) ProtoMessage() {}

func (x *GetHistoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_meet_meet_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHistoryRequest.ProtoReflect.Descriptor instead.
func (*GetHistoryRequest) Descriptor() ([]byte, []int) {
	return file_meet_meet_proto_rawDescGZIP(), []int{6}
}

func (x *GetHistoryRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *GetHistoryRequest) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *GetHistoryRequest) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

type GetHistoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meets []*Meet `protobuf:"bytes,1,rep,name=meets,proto3" json:"meets,omitempty"`
}

func (x *GetHistoryResponse) Reset() {
	*x = GetHistoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meet_meet_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHistoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHistoryResponse) ProtoMessage() {}

func (x *GetHistoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_meet_meet_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHistoryResponse.ProtoReflect.Descriptor instead.
func (*GetHistoryResponse) Descriptor() ([]byte, []int) {
	return file_meet_meet_proto_rawDescGZIP(), []int{7}
}

func (x *GetHistoryResponse) GetMeets() []*Meet {
	if x != nil {
		return x.Meets
	}
	return nil
}

type Meet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstId    int64  `protobuf:"varint,1,opt,name=first_id,json=firstId,proto3" json:"first_id,omitempty"`
	SecondId   int64  `protobuf:"varint,2,opt,name=second_id,json=secondId,proto3" json:"second_id,omitempty"`
	Date       string `protobuf:"bytes,3,opt,name=date,proto3" json:"date,omitempty"`
	IsApproved bool   `protobuf:"varint,4,opt,name=is_approved,json=isApproved,proto3" json:"is_approved,omitempty"`
}

func (x *Meet) Reset() {
	*x = Meet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meet_meet_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Meet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Meet) ProtoMessage() {}

func (x *Meet) ProtoReflect() protoreflect.Message {
	mi := &file_meet_meet_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Meet.ProtoReflect.Descriptor instead.
func (*Meet) Descriptor() ([]byte, []int) {
	return file_meet_meet_proto_rawDescGZIP(), []int{8}
}

func (x *Meet) GetFirstId() int64 {
	if x != nil {
		return x.FirstId
	}
	return 0
}

func (x *Meet) GetSecondId() int64 {
	if x != nil {
		return x.SecondId
	}
	return 0
}

func (x *Meet) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *Meet) GetIsApproved() bool {
	if x != nil {
		return x.IsApproved
	}
	return false
}

var File_meet_meet_proto protoreflect.FileDescriptor

var file_meet_meet_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x6d, 0x65, 0x65, 0x74, 0x2f, 0x6d, 0x65, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x6d, 0x65, 0x65, 0x74, 0x22, 0x5f, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4d, 0x65, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08,
	0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07,
	0x66, 0x69, 0x72, 0x73, 0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x63, 0x6f, 0x6e,
	0x64, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x65, 0x63, 0x6f,
	0x6e, 0x64, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x22, 0x2d, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4d, 0x65, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17,
	0x0a, 0x07, 0x6d, 0x65, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x6d, 0x65, 0x65, 0x74, 0x49, 0x64, 0x22, 0x2d, 0x0a, 0x12, 0x41, 0x70, 0x70, 0x72, 0x6f,
	0x76, 0x65, 0x4d, 0x65, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a,
	0x07, 0x6d, 0x65, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x6d, 0x65, 0x65, 0x74, 0x49, 0x64, 0x22, 0x15, 0x0a, 0x13, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76,
	0x65, 0x4d, 0x65, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2b, 0x0a,
	0x10, 0x47, 0x65, 0x74, 0x46, 0x75, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x35, 0x0a, 0x11, 0x47, 0x65,
	0x74, 0x46, 0x75, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x20, 0x0a, 0x05, 0x6d, 0x65, 0x65, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a,
	0x2e, 0x6d, 0x65, 0x65, 0x74, 0x2e, 0x4d, 0x65, 0x65, 0x74, 0x52, 0x05, 0x6d, 0x65, 0x65, 0x74,
	0x73, 0x22, 0x66, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x19,
	0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x22, 0x36, 0x0a, 0x12, 0x47, 0x65, 0x74,
	0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x20, 0x0a, 0x05, 0x6d, 0x65, 0x65, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a,
	0x2e, 0x6d, 0x65, 0x65, 0x74, 0x2e, 0x4d, 0x65, 0x65, 0x74, 0x52, 0x05, 0x6d, 0x65, 0x65, 0x74,
	0x73, 0x22, 0x73, 0x0a, 0x04, 0x4d, 0x65, 0x65, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x69, 0x72,
	0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x66, 0x69, 0x72,
	0x73, 0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x49,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x61, 0x70, 0x70, 0x72,
	0x6f, 0x76, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x41, 0x70,
	0x70, 0x72, 0x6f, 0x76, 0x65, 0x64, 0x32, 0x91, 0x02, 0x0a, 0x0b, 0x4d, 0x65, 0x65, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3f, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x48, 0x69, 0x73,
	0x74, 0x6f, 0x72, 0x79, 0x12, 0x17, 0x2e, 0x6d, 0x65, 0x65, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x48,
	0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e,
	0x6d, 0x65, 0x65, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x46, 0x75,
	0x74, 0x75, 0x72, 0x65, 0x12, 0x16, 0x2e, 0x6d, 0x65, 0x65, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x46,
	0x75, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x6d,
	0x65, 0x65, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x75, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x0b, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65,
	0x4d, 0x65, 0x65, 0x74, 0x12, 0x18, 0x2e, 0x6d, 0x65, 0x65, 0x74, 0x2e, 0x41, 0x70, 0x70, 0x72,
	0x6f, 0x76, 0x65, 0x4d, 0x65, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19,
	0x2e, 0x6d, 0x65, 0x65, 0x74, 0x2e, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x4d, 0x65, 0x65,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x0a, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4d, 0x65, 0x65, 0x74, 0x12, 0x17, 0x2e, 0x6d, 0x65, 0x65, 0x74, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x18, 0x2e, 0x6d, 0x65, 0x65, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65,
	0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x10, 0x5a, 0x0e, 0x61, 0x70,
	0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x65, 0x65, 0x74, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_meet_meet_proto_rawDescOnce sync.Once
	file_meet_meet_proto_rawDescData = file_meet_meet_proto_rawDesc
)

func file_meet_meet_proto_rawDescGZIP() []byte {
	file_meet_meet_proto_rawDescOnce.Do(func() {
		file_meet_meet_proto_rawDescData = protoimpl.X.CompressGZIP(file_meet_meet_proto_rawDescData)
	})
	return file_meet_meet_proto_rawDescData
}

var file_meet_meet_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_meet_meet_proto_goTypes = []any{
	(*CreateMeetRequest)(nil),   // 0: meet.CreateMeetRequest
	(*CreateMeetResponse)(nil),  // 1: meet.CreateMeetResponse
	(*ApproveMeetRequest)(nil),  // 2: meet.ApproveMeetRequest
	(*ApproveMeetResponse)(nil), // 3: meet.ApproveMeetResponse
	(*GetFutureRequest)(nil),    // 4: meet.GetFutureRequest
	(*GetFutureResponse)(nil),   // 5: meet.GetFutureResponse
	(*GetHistoryRequest)(nil),   // 6: meet.GetHistoryRequest
	(*GetHistoryResponse)(nil),  // 7: meet.GetHistoryResponse
	(*Meet)(nil),                // 8: meet.Meet
}
var file_meet_meet_proto_depIdxs = []int32{
	8, // 0: meet.GetFutureResponse.meets:type_name -> meet.Meet
	8, // 1: meet.GetHistoryResponse.meets:type_name -> meet.Meet
	6, // 2: meet.MeetService.GetHistory:input_type -> meet.GetHistoryRequest
	4, // 3: meet.MeetService.GetFuture:input_type -> meet.GetFutureRequest
	2, // 4: meet.MeetService.ApproveMeet:input_type -> meet.ApproveMeetRequest
	0, // 5: meet.MeetService.CreateMeet:input_type -> meet.CreateMeetRequest
	7, // 6: meet.MeetService.GetHistory:output_type -> meet.GetHistoryResponse
	5, // 7: meet.MeetService.GetFuture:output_type -> meet.GetFutureResponse
	3, // 8: meet.MeetService.ApproveMeet:output_type -> meet.ApproveMeetResponse
	1, // 9: meet.MeetService.CreateMeet:output_type -> meet.CreateMeetResponse
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_meet_meet_proto_init() }
func file_meet_meet_proto_init() {
	if File_meet_meet_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_meet_meet_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CreateMeetRequest); i {
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
		file_meet_meet_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CreateMeetResponse); i {
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
		file_meet_meet_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ApproveMeetRequest); i {
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
		file_meet_meet_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*ApproveMeetResponse); i {
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
		file_meet_meet_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GetFutureRequest); i {
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
		file_meet_meet_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*GetFutureResponse); i {
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
		file_meet_meet_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*GetHistoryRequest); i {
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
		file_meet_meet_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*GetHistoryResponse); i {
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
		file_meet_meet_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*Meet); i {
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
			RawDescriptor: file_meet_meet_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_meet_meet_proto_goTypes,
		DependencyIndexes: file_meet_meet_proto_depIdxs,
		MessageInfos:      file_meet_meet_proto_msgTypes,
	}.Build()
	File_meet_meet_proto = out.File
	file_meet_meet_proto_rawDesc = nil
	file_meet_meet_proto_goTypes = nil
	file_meet_meet_proto_depIdxs = nil
}
