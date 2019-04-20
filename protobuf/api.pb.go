// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protobuf/api.proto

package evileye

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import empty "github.com/golang/protobuf/ptypes/empty"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Tarekomi struct {
	TargetUserId         int64    `protobuf:"varint,1,opt,name=target_user_id,json=targetUserId,proto3" json:"target_user_id,omitempty"`
	Url                  string   `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	Desc                 string   `protobuf:"bytes,4,opt,name=desc,proto3" json:"desc,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Tarekomi) Reset()         { *m = Tarekomi{} }
func (m *Tarekomi) String() string { return proto.CompactTextString(m) }
func (*Tarekomi) ProtoMessage()    {}
func (*Tarekomi) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_c4b1180d2ab7905e, []int{0}
}
func (m *Tarekomi) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tarekomi.Unmarshal(m, b)
}
func (m *Tarekomi) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tarekomi.Marshal(b, m, deterministic)
}
func (dst *Tarekomi) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tarekomi.Merge(dst, src)
}
func (m *Tarekomi) XXX_Size() int {
	return xxx_messageInfo_Tarekomi.Size(m)
}
func (m *Tarekomi) XXX_DiscardUnknown() {
	xxx_messageInfo_Tarekomi.DiscardUnknown(m)
}

var xxx_messageInfo_Tarekomi proto.InternalMessageInfo

func (m *Tarekomi) GetTargetUserId() int64 {
	if m != nil {
		return m.TargetUserId
	}
	return 0
}

func (m *Tarekomi) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *Tarekomi) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

type User struct {
	UserId               int64       `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	UserName             string      `protobuf:"bytes,2,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	Tarekomis            []*Tarekomi `protobuf:"bytes,3,rep,name=tarekomis,proto3" json:"tarekomis,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_c4b1180d2ab7905e, []int{1}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (dst *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(dst, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *User) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *User) GetTarekomis() []*Tarekomi {
	if m != nil {
		return m.Tarekomis
	}
	return nil
}

type HealthCheckRes struct {
	CommitHash           string   `protobuf:"bytes,1,opt,name=commit_hash,json=commitHash,proto3" json:"commit_hash,omitempty"`
	BuildTime            uint64   `protobuf:"varint,2,opt,name=build_time,json=buildTime,proto3" json:"build_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HealthCheckRes) Reset()         { *m = HealthCheckRes{} }
func (m *HealthCheckRes) String() string { return proto.CompactTextString(m) }
func (*HealthCheckRes) ProtoMessage()    {}
func (*HealthCheckRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_c4b1180d2ab7905e, []int{2}
}
func (m *HealthCheckRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheckRes.Unmarshal(m, b)
}
func (m *HealthCheckRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheckRes.Marshal(b, m, deterministic)
}
func (dst *HealthCheckRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheckRes.Merge(dst, src)
}
func (m *HealthCheckRes) XXX_Size() int {
	return xxx_messageInfo_HealthCheckRes.Size(m)
}
func (m *HealthCheckRes) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheckRes.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheckRes proto.InternalMessageInfo

func (m *HealthCheckRes) GetCommitHash() string {
	if m != nil {
		return m.CommitHash
	}
	return ""
}

func (m *HealthCheckRes) GetBuildTime() uint64 {
	if m != nil {
		return m.BuildTime
	}
	return 0
}

type TarekomiReq struct {
	Tarekomi             *Tarekomi `protobuf:"bytes,1,opt,name=tarekomi,proto3" json:"tarekomi,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *TarekomiReq) Reset()         { *m = TarekomiReq{} }
func (m *TarekomiReq) String() string { return proto.CompactTextString(m) }
func (*TarekomiReq) ProtoMessage()    {}
func (*TarekomiReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_c4b1180d2ab7905e, []int{3}
}
func (m *TarekomiReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TarekomiReq.Unmarshal(m, b)
}
func (m *TarekomiReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TarekomiReq.Marshal(b, m, deterministic)
}
func (dst *TarekomiReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TarekomiReq.Merge(dst, src)
}
func (m *TarekomiReq) XXX_Size() int {
	return xxx_messageInfo_TarekomiReq.Size(m)
}
func (m *TarekomiReq) XXX_DiscardUnknown() {
	xxx_messageInfo_TarekomiReq.DiscardUnknown(m)
}

var xxx_messageInfo_TarekomiReq proto.InternalMessageInfo

func (m *TarekomiReq) GetTarekomi() *Tarekomi {
	if m != nil {
		return m.Tarekomi
	}
	return nil
}

type VoteReq struct {
	TarekomiId           int64    `protobuf:"varint,1,opt,name=tarekomi_id,json=tarekomiId,proto3" json:"tarekomi_id,omitempty"`
	Desc                 string   `protobuf:"bytes,2,opt,name=desc,proto3" json:"desc,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VoteReq) Reset()         { *m = VoteReq{} }
func (m *VoteReq) String() string { return proto.CompactTextString(m) }
func (*VoteReq) ProtoMessage()    {}
func (*VoteReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_c4b1180d2ab7905e, []int{4}
}
func (m *VoteReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VoteReq.Unmarshal(m, b)
}
func (m *VoteReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VoteReq.Marshal(b, m, deterministic)
}
func (dst *VoteReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VoteReq.Merge(dst, src)
}
func (m *VoteReq) XXX_Size() int {
	return xxx_messageInfo_VoteReq.Size(m)
}
func (m *VoteReq) XXX_DiscardUnknown() {
	xxx_messageInfo_VoteReq.DiscardUnknown(m)
}

var xxx_messageInfo_VoteReq proto.InternalMessageInfo

func (m *VoteReq) GetTarekomiId() int64 {
	if m != nil {
		return m.TarekomiId
	}
	return 0
}

func (m *VoteReq) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

type TarekomiBoardReq struct {
	Limit                int64    `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset               int64    `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TarekomiBoardReq) Reset()         { *m = TarekomiBoardReq{} }
func (m *TarekomiBoardReq) String() string { return proto.CompactTextString(m) }
func (*TarekomiBoardReq) ProtoMessage()    {}
func (*TarekomiBoardReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_c4b1180d2ab7905e, []int{5}
}
func (m *TarekomiBoardReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TarekomiBoardReq.Unmarshal(m, b)
}
func (m *TarekomiBoardReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TarekomiBoardReq.Marshal(b, m, deterministic)
}
func (dst *TarekomiBoardReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TarekomiBoardReq.Merge(dst, src)
}
func (m *TarekomiBoardReq) XXX_Size() int {
	return xxx_messageInfo_TarekomiBoardReq.Size(m)
}
func (m *TarekomiBoardReq) XXX_DiscardUnknown() {
	xxx_messageInfo_TarekomiBoardReq.DiscardUnknown(m)
}

var xxx_messageInfo_TarekomiBoardReq proto.InternalMessageInfo

func (m *TarekomiBoardReq) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *TarekomiBoardReq) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

type TarekomiSummary struct {
	Tarekomi             *Tarekomi `protobuf:"bytes,1,opt,name=tarekomi,proto3" json:"tarekomi,omitempty"`
	UserName             string    `protobuf:"bytes,2,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *TarekomiSummary) Reset()         { *m = TarekomiSummary{} }
func (m *TarekomiSummary) String() string { return proto.CompactTextString(m) }
func (*TarekomiSummary) ProtoMessage()    {}
func (*TarekomiSummary) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_c4b1180d2ab7905e, []int{6}
}
func (m *TarekomiSummary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TarekomiSummary.Unmarshal(m, b)
}
func (m *TarekomiSummary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TarekomiSummary.Marshal(b, m, deterministic)
}
func (dst *TarekomiSummary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TarekomiSummary.Merge(dst, src)
}
func (m *TarekomiSummary) XXX_Size() int {
	return xxx_messageInfo_TarekomiSummary.Size(m)
}
func (m *TarekomiSummary) XXX_DiscardUnknown() {
	xxx_messageInfo_TarekomiSummary.DiscardUnknown(m)
}

var xxx_messageInfo_TarekomiSummary proto.InternalMessageInfo

func (m *TarekomiSummary) GetTarekomi() *Tarekomi {
	if m != nil {
		return m.Tarekomi
	}
	return nil
}

func (m *TarekomiSummary) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

type TarekomiSummaries struct {
	Tarekomis            []*TarekomiSummary `protobuf:"bytes,1,rep,name=tarekomis,proto3" json:"tarekomis,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *TarekomiSummaries) Reset()         { *m = TarekomiSummaries{} }
func (m *TarekomiSummaries) String() string { return proto.CompactTextString(m) }
func (*TarekomiSummaries) ProtoMessage()    {}
func (*TarekomiSummaries) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_c4b1180d2ab7905e, []int{7}
}
func (m *TarekomiSummaries) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TarekomiSummaries.Unmarshal(m, b)
}
func (m *TarekomiSummaries) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TarekomiSummaries.Marshal(b, m, deterministic)
}
func (dst *TarekomiSummaries) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TarekomiSummaries.Merge(dst, src)
}
func (m *TarekomiSummaries) XXX_Size() int {
	return xxx_messageInfo_TarekomiSummaries.Size(m)
}
func (m *TarekomiSummaries) XXX_DiscardUnknown() {
	xxx_messageInfo_TarekomiSummaries.DiscardUnknown(m)
}

var xxx_messageInfo_TarekomiSummaries proto.InternalMessageInfo

func (m *TarekomiSummaries) GetTarekomis() []*TarekomiSummary {
	if m != nil {
		return m.Tarekomis
	}
	return nil
}

type TarekomiBoardRes struct {
	Tarekomis            []*TarekomiSummary `protobuf:"bytes,1,rep,name=tarekomis,proto3" json:"tarekomis,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *TarekomiBoardRes) Reset()         { *m = TarekomiBoardRes{} }
func (m *TarekomiBoardRes) String() string { return proto.CompactTextString(m) }
func (*TarekomiBoardRes) ProtoMessage()    {}
func (*TarekomiBoardRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_c4b1180d2ab7905e, []int{8}
}
func (m *TarekomiBoardRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TarekomiBoardRes.Unmarshal(m, b)
}
func (m *TarekomiBoardRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TarekomiBoardRes.Marshal(b, m, deterministic)
}
func (dst *TarekomiBoardRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TarekomiBoardRes.Merge(dst, src)
}
func (m *TarekomiBoardRes) XXX_Size() int {
	return xxx_messageInfo_TarekomiBoardRes.Size(m)
}
func (m *TarekomiBoardRes) XXX_DiscardUnknown() {
	xxx_messageInfo_TarekomiBoardRes.DiscardUnknown(m)
}

var xxx_messageInfo_TarekomiBoardRes proto.InternalMessageInfo

func (m *TarekomiBoardRes) GetTarekomis() []*TarekomiSummary {
	if m != nil {
		return m.Tarekomis
	}
	return nil
}

type UserInfoReq struct {
	UserName             int64    `protobuf:"varint,1,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserInfoReq) Reset()         { *m = UserInfoReq{} }
func (m *UserInfoReq) String() string { return proto.CompactTextString(m) }
func (*UserInfoReq) ProtoMessage()    {}
func (*UserInfoReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_c4b1180d2ab7905e, []int{9}
}
func (m *UserInfoReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfoReq.Unmarshal(m, b)
}
func (m *UserInfoReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfoReq.Marshal(b, m, deterministic)
}
func (dst *UserInfoReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfoReq.Merge(dst, src)
}
func (m *UserInfoReq) XXX_Size() int {
	return xxx_messageInfo_UserInfoReq.Size(m)
}
func (m *UserInfoReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfoReq.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfoReq proto.InternalMessageInfo

func (m *UserInfoReq) GetUserName() int64 {
	if m != nil {
		return m.UserName
	}
	return 0
}

type GetUserListReq struct {
	Limit                int64    `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset               int64    `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserListReq) Reset()         { *m = GetUserListReq{} }
func (m *GetUserListReq) String() string { return proto.CompactTextString(m) }
func (*GetUserListReq) ProtoMessage()    {}
func (*GetUserListReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_c4b1180d2ab7905e, []int{10}
}
func (m *GetUserListReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserListReq.Unmarshal(m, b)
}
func (m *GetUserListReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserListReq.Marshal(b, m, deterministic)
}
func (dst *GetUserListReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserListReq.Merge(dst, src)
}
func (m *GetUserListReq) XXX_Size() int {
	return xxx_messageInfo_GetUserListReq.Size(m)
}
func (m *GetUserListReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserListReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserListReq proto.InternalMessageInfo

func (m *GetUserListReq) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *GetUserListReq) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

type Users struct {
	Users                []*User  `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Users) Reset()         { *m = Users{} }
func (m *Users) String() string { return proto.CompactTextString(m) }
func (*Users) ProtoMessage()    {}
func (*Users) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_c4b1180d2ab7905e, []int{11}
}
func (m *Users) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Users.Unmarshal(m, b)
}
func (m *Users) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Users.Marshal(b, m, deterministic)
}
func (dst *Users) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Users.Merge(dst, src)
}
func (m *Users) XXX_Size() int {
	return xxx_messageInfo_Users.Size(m)
}
func (m *Users) XXX_DiscardUnknown() {
	xxx_messageInfo_Users.DiscardUnknown(m)
}

var xxx_messageInfo_Users proto.InternalMessageInfo

func (m *Users) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

type AddStarReq struct {
	TarekomiId           int64    `protobuf:"varint,1,opt,name=tarekomi_id,json=tarekomiId,proto3" json:"tarekomi_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddStarReq) Reset()         { *m = AddStarReq{} }
func (m *AddStarReq) String() string { return proto.CompactTextString(m) }
func (*AddStarReq) ProtoMessage()    {}
func (*AddStarReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_c4b1180d2ab7905e, []int{12}
}
func (m *AddStarReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddStarReq.Unmarshal(m, b)
}
func (m *AddStarReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddStarReq.Marshal(b, m, deterministic)
}
func (dst *AddStarReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddStarReq.Merge(dst, src)
}
func (m *AddStarReq) XXX_Size() int {
	return xxx_messageInfo_AddStarReq.Size(m)
}
func (m *AddStarReq) XXX_DiscardUnknown() {
	xxx_messageInfo_AddStarReq.DiscardUnknown(m)
}

var xxx_messageInfo_AddStarReq proto.InternalMessageInfo

func (m *AddStarReq) GetTarekomiId() int64 {
	if m != nil {
		return m.TarekomiId
	}
	return 0
}

func init() {
	proto.RegisterType((*Tarekomi)(nil), "evileye.Tarekomi")
	proto.RegisterType((*User)(nil), "evileye.User")
	proto.RegisterType((*HealthCheckRes)(nil), "evileye.HealthCheckRes")
	proto.RegisterType((*TarekomiReq)(nil), "evileye.TarekomiReq")
	proto.RegisterType((*VoteReq)(nil), "evileye.VoteReq")
	proto.RegisterType((*TarekomiBoardReq)(nil), "evileye.TarekomiBoardReq")
	proto.RegisterType((*TarekomiSummary)(nil), "evileye.TarekomiSummary")
	proto.RegisterType((*TarekomiSummaries)(nil), "evileye.TarekomiSummaries")
	proto.RegisterType((*TarekomiBoardRes)(nil), "evileye.TarekomiBoardRes")
	proto.RegisterType((*UserInfoReq)(nil), "evileye.UserInfoReq")
	proto.RegisterType((*GetUserListReq)(nil), "evileye.GetUserListReq")
	proto.RegisterType((*Users)(nil), "evileye.Users")
	proto.RegisterType((*AddStarReq)(nil), "evileye.AddStarReq")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PrivateClient is the client API for Private service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PrivateClient interface {
	//
	// HealthCheck
	//
	// ヘルスチェック用
	HealthCheck(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*HealthCheckRes, error)
	//
	// Tarekomi
	//
	// Tarekomi returns some status
	// Invalid      ... invalid parameter
	// Internal     ... an error occured in the server
	// Unauthorized ... JWT is invalid
	// タレコミをする
	Tarekomi(ctx context.Context, in *TarekomiReq, opts ...grpc.CallOption) (*empty.Empty, error)
	//
	// Vote
	//
	// 投票は、匿名で行う
	// descは、投票が完了して、ユーザ画面で表示される際には誰が投票したかは伏せて公開する
	Vote(ctx context.Context, in *VoteReq, opts ...grpc.CallOption) (*empty.Empty, error)
	//
	// TarekomiBoard
	//
	// 投票待ちのタレコミを表示する
	TarekomiBoard(ctx context.Context, in *TarekomiBoardReq, opts ...grpc.CallOption) (*TarekomiBoardRes, error)
	//
	// GetUserInfo
	//
	// ユーザの情報を返す
	GetUserInfo(ctx context.Context, in *UserInfoReq, opts ...grpc.CallOption) (*User, error)
	//
	// GetUserList
	//
	// ユーザ一覧を返す
	GetUserList(ctx context.Context, in *GetUserListReq, opts ...grpc.CallOption) (*Users, error)
	//
	// AddStar
	//
	// タレコミにStarをつける
	// Approve済みのものだけ可能
	AddStar(ctx context.Context, in *AddStarReq, opts ...grpc.CallOption) (*empty.Empty, error)
	//
	// GetStaredTarekomi
	//
	// 自分のスターしたタレコミ一覧
	// JWTからユーザ情報を取得する
	GetStaredTarekomi(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*TarekomiSummaries, error)
}

type privateClient struct {
	cc *grpc.ClientConn
}

func NewPrivateClient(cc *grpc.ClientConn) PrivateClient {
	return &privateClient{cc}
}

func (c *privateClient) HealthCheck(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*HealthCheckRes, error) {
	out := new(HealthCheckRes)
	err := c.cc.Invoke(ctx, "/evileye.Private/HealthCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *privateClient) Tarekomi(ctx context.Context, in *TarekomiReq, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/evileye.Private/Tarekomi", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *privateClient) Vote(ctx context.Context, in *VoteReq, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/evileye.Private/Vote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *privateClient) TarekomiBoard(ctx context.Context, in *TarekomiBoardReq, opts ...grpc.CallOption) (*TarekomiBoardRes, error) {
	out := new(TarekomiBoardRes)
	err := c.cc.Invoke(ctx, "/evileye.Private/TarekomiBoard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *privateClient) GetUserInfo(ctx context.Context, in *UserInfoReq, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/evileye.Private/GetUserInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *privateClient) GetUserList(ctx context.Context, in *GetUserListReq, opts ...grpc.CallOption) (*Users, error) {
	out := new(Users)
	err := c.cc.Invoke(ctx, "/evileye.Private/GetUserList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *privateClient) AddStar(ctx context.Context, in *AddStarReq, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/evileye.Private/AddStar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *privateClient) GetStaredTarekomi(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*TarekomiSummaries, error) {
	out := new(TarekomiSummaries)
	err := c.cc.Invoke(ctx, "/evileye.Private/GetStaredTarekomi", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PrivateServer is the server API for Private service.
type PrivateServer interface {
	//
	// HealthCheck
	//
	// ヘルスチェック用
	HealthCheck(context.Context, *empty.Empty) (*HealthCheckRes, error)
	//
	// Tarekomi
	//
	// Tarekomi returns some status
	// Invalid      ... invalid parameter
	// Internal     ... an error occured in the server
	// Unauthorized ... JWT is invalid
	// タレコミをする
	Tarekomi(context.Context, *TarekomiReq) (*empty.Empty, error)
	//
	// Vote
	//
	// 投票は、匿名で行う
	// descは、投票が完了して、ユーザ画面で表示される際には誰が投票したかは伏せて公開する
	Vote(context.Context, *VoteReq) (*empty.Empty, error)
	//
	// TarekomiBoard
	//
	// 投票待ちのタレコミを表示する
	TarekomiBoard(context.Context, *TarekomiBoardReq) (*TarekomiBoardRes, error)
	//
	// GetUserInfo
	//
	// ユーザの情報を返す
	GetUserInfo(context.Context, *UserInfoReq) (*User, error)
	//
	// GetUserList
	//
	// ユーザ一覧を返す
	GetUserList(context.Context, *GetUserListReq) (*Users, error)
	//
	// AddStar
	//
	// タレコミにStarをつける
	// Approve済みのものだけ可能
	AddStar(context.Context, *AddStarReq) (*empty.Empty, error)
	//
	// GetStaredTarekomi
	//
	// 自分のスターしたタレコミ一覧
	// JWTからユーザ情報を取得する
	GetStaredTarekomi(context.Context, *empty.Empty) (*TarekomiSummaries, error)
}

func RegisterPrivateServer(s *grpc.Server, srv PrivateServer) {
	s.RegisterService(&_Private_serviceDesc, srv)
}

func _Private_HealthCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivateServer).HealthCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/evileye.Private/HealthCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivateServer).HealthCheck(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Private_Tarekomi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TarekomiReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivateServer).Tarekomi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/evileye.Private/Tarekomi",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivateServer).Tarekomi(ctx, req.(*TarekomiReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Private_Vote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VoteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivateServer).Vote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/evileye.Private/Vote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivateServer).Vote(ctx, req.(*VoteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Private_TarekomiBoard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TarekomiBoardReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivateServer).TarekomiBoard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/evileye.Private/TarekomiBoard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivateServer).TarekomiBoard(ctx, req.(*TarekomiBoardReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Private_GetUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivateServer).GetUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/evileye.Private/GetUserInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivateServer).GetUserInfo(ctx, req.(*UserInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Private_GetUserList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivateServer).GetUserList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/evileye.Private/GetUserList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivateServer).GetUserList(ctx, req.(*GetUserListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Private_AddStar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddStarReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivateServer).AddStar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/evileye.Private/AddStar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivateServer).AddStar(ctx, req.(*AddStarReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Private_GetStaredTarekomi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivateServer).GetStaredTarekomi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/evileye.Private/GetStaredTarekomi",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivateServer).GetStaredTarekomi(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Private_serviceDesc = grpc.ServiceDesc{
	ServiceName: "evileye.Private",
	HandlerType: (*PrivateServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HealthCheck",
			Handler:    _Private_HealthCheck_Handler,
		},
		{
			MethodName: "Tarekomi",
			Handler:    _Private_Tarekomi_Handler,
		},
		{
			MethodName: "Vote",
			Handler:    _Private_Vote_Handler,
		},
		{
			MethodName: "TarekomiBoard",
			Handler:    _Private_TarekomiBoard_Handler,
		},
		{
			MethodName: "GetUserInfo",
			Handler:    _Private_GetUserInfo_Handler,
		},
		{
			MethodName: "GetUserList",
			Handler:    _Private_GetUserList_Handler,
		},
		{
			MethodName: "AddStar",
			Handler:    _Private_AddStar_Handler,
		},
		{
			MethodName: "GetStaredTarekomi",
			Handler:    _Private_GetStaredTarekomi_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobuf/api.proto",
}

func init() { proto.RegisterFile("protobuf/api.proto", fileDescriptor_api_c4b1180d2ab7905e) }

var fileDescriptor_api_c4b1180d2ab7905e = []byte{
	// 605 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0xed, 0x4f, 0xd3, 0x40,
	0x18, 0xcf, 0xe8, 0xa0, 0xec, 0xa9, 0x4c, 0x38, 0x09, 0xd4, 0x12, 0x03, 0xa9, 0x7e, 0x58, 0x8c,
	0x14, 0x33, 0x0d, 0xc1, 0xc4, 0x10, 0xc5, 0x10, 0x40, 0x8d, 0x59, 0x0a, 0xf2, 0xcd, 0x2c, 0xb7,
	0xf5, 0xd9, 0x76, 0xa1, 0x47, 0x47, 0xef, 0x4a, 0xb2, 0x3f, 0xd9, 0xff, 0xc2, 0xdc, 0xf5, 0x65,
	0xad, 0xb3, 0x46, 0xfc, 0x76, 0xf7, 0xbc, 0xfc, 0x9e, 0x97, 0xdf, 0xef, 0x01, 0x32, 0x8d, 0x23,
	0x19, 0x0d, 0x92, 0xd1, 0x01, 0x9d, 0x32, 0x4f, 0x7f, 0x88, 0x89, 0xf7, 0x2c, 0xc4, 0x19, 0x3a,
	0x3b, 0xe3, 0x28, 0x1a, 0x87, 0x78, 0x50, 0xc4, 0x20, 0x9f, 0xca, 0x59, 0x1a, 0xe5, 0x5e, 0xc3,
	0xea, 0x15, 0x8d, 0xf1, 0x26, 0xe2, 0x8c, 0xbc, 0x80, 0xb6, 0xa4, 0xf1, 0x18, 0x65, 0x3f, 0x11,
	0x18, 0xf7, 0x59, 0x60, 0x37, 0xf6, 0x1a, 0x1d, 0xc3, 0x7f, 0x94, 0x5a, 0xbf, 0x0b, 0x8c, 0x2f,
	0x02, 0xb2, 0x0e, 0x46, 0x12, 0x87, 0xb6, 0xb1, 0xd7, 0xe8, 0xb4, 0x7c, 0xf5, 0x24, 0x04, 0x9a,
	0x01, 0x8a, 0xa1, 0xdd, 0xd4, 0x26, 0xfd, 0x76, 0x39, 0x34, 0x55, 0x3c, 0xd9, 0x06, 0xb3, 0x0a,
	0xb6, 0x92, 0xa4, 0x30, 0x3b, 0xd0, 0xd2, 0x8e, 0x5b, 0xca, 0xd1, 0x5e, 0xd2, 0x99, 0xab, 0xca,
	0xf0, 0x8d, 0x72, 0x24, 0x07, 0xd0, 0x92, 0x59, 0x57, 0xc2, 0x36, 0xf6, 0x8c, 0x8e, 0xd5, 0xdd,
	0xf0, 0xb2, 0x79, 0xbc, 0xbc, 0x5f, 0x7f, 0x1e, 0xe3, 0xf6, 0xa0, 0x7d, 0x8e, 0x34, 0x94, 0x93,
	0x4f, 0x13, 0x1c, 0xde, 0xf8, 0x28, 0xc8, 0x2e, 0x58, 0xc3, 0x88, 0x73, 0x26, 0xfb, 0x13, 0x2a,
	0x26, 0xba, 0x78, 0xcb, 0x87, 0xd4, 0x74, 0x4e, 0xc5, 0x84, 0x3c, 0x03, 0x18, 0x24, 0x2c, 0x0c,
	0xfa, 0x92, 0x65, 0x1d, 0x34, 0xfd, 0x96, 0xb6, 0x5c, 0x31, 0x8e, 0xee, 0x7b, 0xb0, 0x8a, 0x42,
	0x78, 0x47, 0xf6, 0x61, 0x35, 0xaf, 0xa6, 0xb1, 0xfe, 0xd8, 0x50, 0x11, 0xe2, 0x1e, 0x83, 0x79,
	0x1d, 0x49, 0x54, 0x99, 0xbb, 0x60, 0xe5, 0xe6, 0xf9, 0x16, 0x20, 0x37, 0x5d, 0x04, 0xc5, 0xfa,
	0x96, 0x4a, 0xeb, 0xfb, 0x00, 0xeb, 0x39, 0xea, 0x49, 0x44, 0xe3, 0x40, 0x01, 0x6d, 0xc2, 0x72,
	0xc8, 0x38, 0x93, 0x19, 0x44, 0xfa, 0x21, 0x5b, 0xb0, 0x12, 0x8d, 0x46, 0x02, 0xa5, 0xce, 0x37,
	0xfc, 0xec, 0xe7, 0xfe, 0x80, 0xc7, 0x39, 0xc2, 0x65, 0xc2, 0x39, 0x8d, 0x67, 0x0f, 0x9c, 0xe1,
	0xaf, 0x0c, 0xb9, 0x5f, 0x60, 0xa3, 0x0a, 0xcf, 0x50, 0x90, 0xc3, 0x32, 0x6d, 0x0d, 0x4d, 0x9b,
	0xbd, 0x50, 0x21, 0xeb, 0xa6, 0xcc, 0xde, 0xe7, 0x85, 0x69, 0xff, 0x1f, 0xeb, 0x25, 0x58, 0x5a,
	0xa8, 0xb7, 0xa3, 0x48, 0x2d, 0xad, 0x32, 0x44, 0xba, 0xb8, 0xf9, 0x10, 0xc7, 0xd0, 0x3e, 0x4b,
	0x75, 0xfd, 0x95, 0x09, 0xf9, 0xf0, 0x1d, 0xbf, 0x82, 0x65, 0x95, 0x2c, 0xc8, 0x73, 0x58, 0x56,
	0xa0, 0x79, 0xa3, 0x6b, 0x45, 0xa3, 0xca, 0xed, 0xa7, 0x3e, 0x77, 0x1f, 0xe0, 0x63, 0x10, 0x5c,
	0x4a, 0x1a, 0xff, 0x8b, 0x2c, 0xba, 0x3f, 0x0d, 0x30, 0x7b, 0x31, 0xbb, 0xa7, 0x12, 0xc9, 0x31,
	0x58, 0x25, 0x79, 0x93, 0x2d, 0x2f, 0x3d, 0x69, 0x2f, 0x3f, 0x69, 0xef, 0x54, 0x9d, 0xb4, 0xb3,
	0x5d, 0xd4, 0xfd, 0xed, 0x18, 0x8e, 0x4a, 0x57, 0xbe, 0xb9, 0xc8, 0x39, 0xde, 0x39, 0x35, 0x90,
	0xe4, 0x35, 0x34, 0x95, 0x90, 0xc9, 0x7a, 0x91, 0x95, 0xe9, 0xba, 0x36, 0xe3, 0x14, 0xd6, 0x2a,
	0x64, 0x92, 0xa7, 0x0b, 0x05, 0x73, 0x49, 0x3b, 0xb5, 0x2e, 0x41, 0xba, 0x60, 0x65, 0xdc, 0x28,
	0x2a, 0x4b, 0x5d, 0x97, 0xd8, 0x75, 0xaa, 0x8b, 0x26, 0x87, 0x45, 0x8e, 0xe2, 0x93, 0xcc, 0xd7,
	0x51, 0x65, 0xd9, 0x69, 0x57, 0xd2, 0x94, 0xd6, 0xcc, 0x8c, 0x19, 0xf2, 0xa4, 0x70, 0xcd, 0xb9,
	0xaa, 0x1d, 0xf5, 0x0c, 0x36, 0xce, 0x50, 0xaa, 0x28, 0x0c, 0x8a, 0xfd, 0xd6, 0x91, 0xe3, 0xd4,
	0xa8, 0x97, 0xa1, 0x38, 0xe9, 0xc0, 0xce, 0x30, 0xe2, 0xde, 0x4c, 0x1e, 0xbd, 0x7d, 0xd7, 0x2d,
	0xe2, 0x72, 0x80, 0x13, 0xf3, 0x14, 0x7b, 0xea, 0xdd, 0x6b, 0x0c, 0x56, 0xb4, 0xf1, 0xcd, 0xaf,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x40, 0x19, 0xcb, 0xd7, 0xf2, 0x05, 0x00, 0x00,
}