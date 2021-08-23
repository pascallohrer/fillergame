// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0-devel
// 	protoc        v3.12.1
// source: FillerGame.proto

package api

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type GameBoard int32

const (
	GameBoard_SQUARE10 GameBoard = 0
)

// Enum value maps for GameBoard.
var (
	GameBoard_name = map[int32]string{
		0: "SQUARE10",
	}
	GameBoard_value = map[string]int32{
		"SQUARE10": 0,
	}
)

func (x GameBoard) Enum() *GameBoard {
	p := new(GameBoard)
	*p = x
	return p
}

func (x GameBoard) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GameBoard) Descriptor() protoreflect.EnumDescriptor {
	return file_FillerGame_proto_enumTypes[0].Descriptor()
}

func (GameBoard) Type() protoreflect.EnumType {
	return &file_FillerGame_proto_enumTypes[0]
}

func (x GameBoard) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GameBoard.Descriptor instead.
func (GameBoard) EnumDescriptor() ([]byte, []int) {
	return file_FillerGame_proto_rawDescGZIP(), []int{0}
}

type Color int32

const (
	Color_RED     Color = 0
	Color_GREEN   Color = 1
	Color_BLUE    Color = 2
	Color_YELLOW  Color = 3
	Color_CYAN    Color = 4
	Color_MAGENTA Color = 5
	Color_ORANGE  Color = 6
	Color_PURPLE  Color = 7
)

// Enum value maps for Color.
var (
	Color_name = map[int32]string{
		0: "RED",
		1: "GREEN",
		2: "BLUE",
		3: "YELLOW",
		4: "CYAN",
		5: "MAGENTA",
		6: "ORANGE",
		7: "PURPLE",
	}
	Color_value = map[string]int32{
		"RED":     0,
		"GREEN":   1,
		"BLUE":    2,
		"YELLOW":  3,
		"CYAN":    4,
		"MAGENTA": 5,
		"ORANGE":  6,
		"PURPLE":  7,
	}
)

func (x Color) Enum() *Color {
	p := new(Color)
	*p = x
	return p
}

func (x Color) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Color) Descriptor() protoreflect.EnumDescriptor {
	return file_FillerGame_proto_enumTypes[1].Descriptor()
}

func (Color) Type() protoreflect.EnumType {
	return &file_FillerGame_proto_enumTypes[1]
}

func (x Color) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Color.Descriptor instead.
func (Color) EnumDescriptor() ([]byte, []int) {
	return file_FillerGame_proto_rawDescGZIP(), []int{1}
}

type StatusCode int32

const (
	StatusCode_DEFAULT             StatusCode = 0
	StatusCode_YOUR_TURN           StatusCode = 1
	StatusCode_OPPONENTS_TURN      StatusCode = 2
	StatusCode_GAME_OVER           StatusCode = 3
	StatusCode_NOT_YOUR_TURN       StatusCode = 4
	StatusCode_INVALID_TOKEN       StatusCode = 5
	StatusCode_ILLEGAL_MOVE        StatusCode = 6
	StatusCode_WAITING_FOR_PLAYERS StatusCode = 7
	StatusCode_UNAVAILABLE_MODE    StatusCode = 8
	StatusCode_AUTH_TOKEN          StatusCode = 9
)

// Enum value maps for StatusCode.
var (
	StatusCode_name = map[int32]string{
		0: "DEFAULT",
		1: "YOUR_TURN",
		2: "OPPONENTS_TURN",
		3: "GAME_OVER",
		4: "NOT_YOUR_TURN",
		5: "INVALID_TOKEN",
		6: "ILLEGAL_MOVE",
		7: "WAITING_FOR_PLAYERS",
		8: "UNAVAILABLE_MODE",
		9: "AUTH_TOKEN",
	}
	StatusCode_value = map[string]int32{
		"DEFAULT":             0,
		"YOUR_TURN":           1,
		"OPPONENTS_TURN":      2,
		"GAME_OVER":           3,
		"NOT_YOUR_TURN":       4,
		"INVALID_TOKEN":       5,
		"ILLEGAL_MOVE":        6,
		"WAITING_FOR_PLAYERS": 7,
		"UNAVAILABLE_MODE":    8,
		"AUTH_TOKEN":          9,
	}
)

func (x StatusCode) Enum() *StatusCode {
	p := new(StatusCode)
	*p = x
	return p
}

func (x StatusCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StatusCode) Descriptor() protoreflect.EnumDescriptor {
	return file_FillerGame_proto_enumTypes[2].Descriptor()
}

func (StatusCode) Type() protoreflect.EnumType {
	return &file_FillerGame_proto_enumTypes[2]
}

func (x StatusCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StatusCode.Descriptor instead.
func (StatusCode) EnumDescriptor() ([]byte, []int) {
	return file_FillerGame_proto_rawDescGZIP(), []int{2}
}

type JoinRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mode         GameBoard `protobuf:"varint,1,opt,name=mode,proto3,enum=fillergame.GameBoard" json:"mode,omitempty"`
	Username     string    `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	PasswordHash string    `protobuf:"bytes,3,opt,name=password_hash,json=passwordHash,proto3" json:"password_hash,omitempty"`
}

func (x *JoinRequest) Reset() {
	*x = JoinRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_FillerGame_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinRequest) ProtoMessage() {}

func (x *JoinRequest) ProtoReflect() protoreflect.Message {
	mi := &file_FillerGame_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinRequest.ProtoReflect.Descriptor instead.
func (*JoinRequest) Descriptor() ([]byte, []int) {
	return file_FillerGame_proto_rawDescGZIP(), []int{0}
}

func (x *JoinRequest) GetMode() GameBoard {
	if x != nil {
		return x.Mode
	}
	return GameBoard_SQUARE10
}

func (x *JoinRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *JoinRequest) GetPasswordHash() string {
	if x != nil {
		return x.PasswordHash
	}
	return ""
}

type MoveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthToken string `protobuf:"bytes,1,opt,name=auth_token,json=authToken,proto3" json:"auth_token,omitempty"`
	NewColor  Color  `protobuf:"varint,2,opt,name=new_color,json=newColor,proto3,enum=fillergame.Color" json:"new_color,omitempty"`
}

func (x *MoveRequest) Reset() {
	*x = MoveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_FillerGame_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MoveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MoveRequest) ProtoMessage() {}

func (x *MoveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_FillerGame_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MoveRequest.ProtoReflect.Descriptor instead.
func (*MoveRequest) Descriptor() ([]byte, []int) {
	return file_FillerGame_proto_rawDescGZIP(), []int{1}
}

func (x *MoveRequest) GetAuthToken() string {
	if x != nil {
		return x.AuthToken
	}
	return ""
}

func (x *MoveRequest) GetNewColor() Color {
	if x != nil {
		return x.NewColor
	}
	return Color_RED
}

type RebuildRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthToken string `protobuf:"bytes,1,opt,name=auth_token,json=authToken,proto3" json:"auth_token,omitempty"`
}

func (x *RebuildRequest) Reset() {
	*x = RebuildRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_FillerGame_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RebuildRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RebuildRequest) ProtoMessage() {}

func (x *RebuildRequest) ProtoReflect() protoreflect.Message {
	mi := &file_FillerGame_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RebuildRequest.ProtoReflect.Descriptor instead.
func (*RebuildRequest) Descriptor() ([]byte, []int) {
	return file_FillerGame_proto_rawDescGZIP(), []int{2}
}

func (x *RebuildRequest) GetAuthToken() string {
	if x != nil {
		return x.AuthToken
	}
	return ""
}

type BoardState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status *Status     `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Tiles  []*Modifier `protobuf:"bytes,2,rep,name=tiles,proto3" json:"tiles,omitempty"`
}

func (x *BoardState) Reset() {
	*x = BoardState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_FillerGame_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BoardState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BoardState) ProtoMessage() {}

func (x *BoardState) ProtoReflect() protoreflect.Message {
	mi := &file_FillerGame_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BoardState.ProtoReflect.Descriptor instead.
func (*BoardState) Descriptor() ([]byte, []int) {
	return file_FillerGame_proto_rawDescGZIP(), []int{3}
}

func (x *BoardState) GetStatus() *Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *BoardState) GetTiles() []*Modifier {
	if x != nil {
		return x.Tiles
	}
	return nil
}

type MoveAcknowledgment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MoveAcknowledgment) Reset() {
	*x = MoveAcknowledgment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_FillerGame_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MoveAcknowledgment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MoveAcknowledgment) ProtoMessage() {}

func (x *MoveAcknowledgment) ProtoReflect() protoreflect.Message {
	mi := &file_FillerGame_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MoveAcknowledgment.ProtoReflect.Descriptor instead.
func (*MoveAcknowledgment) Descriptor() ([]byte, []int) {
	return file_FillerGame_proto_rawDescGZIP(), []int{4}
}

type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode StatusCode `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3,enum=fillergame.StatusCode" json:"status_code,omitempty"`
	Message    string     `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_FillerGame_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_FillerGame_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_FillerGame_proto_rawDescGZIP(), []int{5}
}

func (x *Status) GetStatusCode() StatusCode {
	if x != nil {
		return x.StatusCode
	}
	return StatusCode_DEFAULT
}

func (x *Status) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type Modifier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TileId   int32 `protobuf:"varint,1,opt,name=tile_id,json=tileId,proto3" json:"tile_id,omitempty"`
	NewColor Color `protobuf:"varint,2,opt,name=new_color,json=newColor,proto3,enum=fillergame.Color" json:"new_color,omitempty"`
}

func (x *Modifier) Reset() {
	*x = Modifier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_FillerGame_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Modifier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Modifier) ProtoMessage() {}

func (x *Modifier) ProtoReflect() protoreflect.Message {
	mi := &file_FillerGame_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Modifier.ProtoReflect.Descriptor instead.
func (*Modifier) Descriptor() ([]byte, []int) {
	return file_FillerGame_proto_rawDescGZIP(), []int{6}
}

func (x *Modifier) GetTileId() int32 {
	if x != nil {
		return x.TileId
	}
	return 0
}

func (x *Modifier) GetNewColor() Color {
	if x != nil {
		return x.NewColor
	}
	return Color_RED
}

var File_FillerGame_proto protoreflect.FileDescriptor

var file_FillerGame_proto_rawDesc = []byte{
	0x0a, 0x10, 0x46, 0x69, 0x6c, 0x6c, 0x65, 0x72, 0x47, 0x61, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x66, 0x69, 0x6c, 0x6c, 0x65, 0x72, 0x67, 0x61, 0x6d, 0x65, 0x22, 0x79,
	0x0a, 0x0b, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a,
	0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x66, 0x69,
	0x6c, 0x6c, 0x65, 0x72, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x42, 0x6f, 0x61,
	0x72, 0x64, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x48, 0x61, 0x73, 0x68, 0x22, 0x5c, 0x0a, 0x0b, 0x4d, 0x6f, 0x76,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x75, 0x74, 0x68,
	0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x75,
	0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x2e, 0x0a, 0x09, 0x6e, 0x65, 0x77, 0x5f, 0x63,
	0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x66, 0x69, 0x6c,
	0x6c, 0x65, 0x72, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x52, 0x08, 0x6e,
	0x65, 0x77, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x22, 0x2f, 0x0a, 0x0e, 0x52, 0x65, 0x62, 0x75, 0x69,
	0x6c, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x75, 0x74,
	0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61,
	0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x64, 0x0a, 0x0a, 0x42, 0x6f, 0x61, 0x72,
	0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x2a, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x66, 0x69, 0x6c, 0x6c, 0x65, 0x72, 0x67,
	0x61, 0x6d, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x2a, 0x0a, 0x05, 0x74, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x66, 0x69, 0x6c, 0x6c, 0x65, 0x72, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x4d,
	0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x05, 0x74, 0x69, 0x6c, 0x65, 0x73, 0x22, 0x14,
	0x0a, 0x12, 0x4d, 0x6f, 0x76, 0x65, 0x41, 0x63, 0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67,
	0x6d, 0x65, 0x6e, 0x74, 0x22, 0x5b, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x37,
	0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x66, 0x69, 0x6c, 0x6c, 0x65, 0x72, 0x67, 0x61, 0x6d, 0x65,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x0a, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0x53, 0x0a, 0x08, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x17, 0x0a,
	0x07, 0x74, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x74, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x09, 0x6e, 0x65, 0x77, 0x5f, 0x63, 0x6f,
	0x6c, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x66, 0x69, 0x6c, 0x6c,
	0x65, 0x72, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x52, 0x08, 0x6e, 0x65,
	0x77, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x2a, 0x19, 0x0a, 0x09, 0x47, 0x61, 0x6d, 0x65, 0x42, 0x6f,
	0x61, 0x72, 0x64, 0x12, 0x0c, 0x0a, 0x08, 0x53, 0x51, 0x55, 0x41, 0x52, 0x45, 0x31, 0x30, 0x10,
	0x00, 0x2a, 0x60, 0x0a, 0x05, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x07, 0x0a, 0x03, 0x52, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x47, 0x52, 0x45, 0x45, 0x4e, 0x10, 0x01, 0x12, 0x08,
	0x0a, 0x04, 0x42, 0x4c, 0x55, 0x45, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x59, 0x45, 0x4c, 0x4c,
	0x4f, 0x57, 0x10, 0x03, 0x12, 0x08, 0x0a, 0x04, 0x43, 0x59, 0x41, 0x4e, 0x10, 0x04, 0x12, 0x0b,
	0x0a, 0x07, 0x4d, 0x41, 0x47, 0x45, 0x4e, 0x54, 0x41, 0x10, 0x05, 0x12, 0x0a, 0x0a, 0x06, 0x4f,
	0x52, 0x41, 0x4e, 0x47, 0x45, 0x10, 0x06, 0x12, 0x0a, 0x0a, 0x06, 0x50, 0x55, 0x52, 0x50, 0x4c,
	0x45, 0x10, 0x07, 0x2a, 0xc2, 0x01, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x45, 0x46, 0x41, 0x55, 0x4c, 0x54, 0x10, 0x00, 0x12,
	0x0d, 0x0a, 0x09, 0x59, 0x4f, 0x55, 0x52, 0x5f, 0x54, 0x55, 0x52, 0x4e, 0x10, 0x01, 0x12, 0x12,
	0x0a, 0x0e, 0x4f, 0x50, 0x50, 0x4f, 0x4e, 0x45, 0x4e, 0x54, 0x53, 0x5f, 0x54, 0x55, 0x52, 0x4e,
	0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x47, 0x41, 0x4d, 0x45, 0x5f, 0x4f, 0x56, 0x45, 0x52, 0x10,
	0x03, 0x12, 0x11, 0x0a, 0x0d, 0x4e, 0x4f, 0x54, 0x5f, 0x59, 0x4f, 0x55, 0x52, 0x5f, 0x54, 0x55,
	0x52, 0x4e, 0x10, 0x04, 0x12, 0x11, 0x0a, 0x0d, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f,
	0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x10, 0x05, 0x12, 0x10, 0x0a, 0x0c, 0x49, 0x4c, 0x4c, 0x45, 0x47,
	0x41, 0x4c, 0x5f, 0x4d, 0x4f, 0x56, 0x45, 0x10, 0x06, 0x12, 0x17, 0x0a, 0x13, 0x57, 0x41, 0x49,
	0x54, 0x49, 0x4e, 0x47, 0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x50, 0x4c, 0x41, 0x59, 0x45, 0x52, 0x53,
	0x10, 0x07, 0x12, 0x14, 0x0a, 0x10, 0x55, 0x4e, 0x41, 0x56, 0x41, 0x49, 0x4c, 0x41, 0x42, 0x4c,
	0x45, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x10, 0x08, 0x12, 0x0e, 0x0a, 0x0a, 0x41, 0x55, 0x54, 0x48,
	0x5f, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x10, 0x09, 0x32, 0xd6, 0x01, 0x0a, 0x0a, 0x46, 0x69, 0x6c,
	0x6c, 0x65, 0x72, 0x47, 0x61, 0x6d, 0x65, 0x12, 0x3d, 0x0a, 0x08, 0x4a, 0x6f, 0x69, 0x6e, 0x47,
	0x61, 0x6d, 0x65, 0x12, 0x17, 0x2e, 0x66, 0x69, 0x6c, 0x6c, 0x65, 0x72, 0x67, 0x61, 0x6d, 0x65,
	0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x66,
	0x69, 0x6c, 0x6c, 0x65, 0x72, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x30, 0x01, 0x12, 0x43, 0x0a, 0x08, 0x4d, 0x61, 0x6b, 0x65, 0x4d, 0x6f,
	0x76, 0x65, 0x12, 0x17, 0x2e, 0x66, 0x69, 0x6c, 0x6c, 0x65, 0x72, 0x67, 0x61, 0x6d, 0x65, 0x2e,
	0x4d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x66, 0x69,
	0x6c, 0x6c, 0x65, 0x72, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x4d, 0x6f, 0x76, 0x65, 0x41, 0x63, 0x6b,
	0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x44, 0x0a, 0x0c, 0x52,
	0x65, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x12, 0x1a, 0x2e, 0x66, 0x69,
	0x6c, 0x6c, 0x65, 0x72, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x52, 0x65, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x66, 0x69, 0x6c, 0x6c, 0x65, 0x72,
	0x67, 0x61, 0x6d, 0x65, 0x2e, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x30,
	0x01, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x70, 0x61, 0x73, 0x63, 0x61, 0x6c, 0x6c, 0x6f, 0x68, 0x72, 0x65, 0x72, 0x2f, 0x66, 0x69, 0x6c,
	0x6c, 0x65, 0x72, 0x67, 0x61, 0x6d, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_FillerGame_proto_rawDescOnce sync.Once
	file_FillerGame_proto_rawDescData = file_FillerGame_proto_rawDesc
)

func file_FillerGame_proto_rawDescGZIP() []byte {
	file_FillerGame_proto_rawDescOnce.Do(func() {
		file_FillerGame_proto_rawDescData = protoimpl.X.CompressGZIP(file_FillerGame_proto_rawDescData)
	})
	return file_FillerGame_proto_rawDescData
}

var file_FillerGame_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_FillerGame_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_FillerGame_proto_goTypes = []interface{}{
	(GameBoard)(0),             // 0: fillergame.GameBoard
	(Color)(0),                 // 1: fillergame.Color
	(StatusCode)(0),            // 2: fillergame.StatusCode
	(*JoinRequest)(nil),        // 3: fillergame.JoinRequest
	(*MoveRequest)(nil),        // 4: fillergame.MoveRequest
	(*RebuildRequest)(nil),     // 5: fillergame.RebuildRequest
	(*BoardState)(nil),         // 6: fillergame.BoardState
	(*MoveAcknowledgment)(nil), // 7: fillergame.MoveAcknowledgment
	(*Status)(nil),             // 8: fillergame.Status
	(*Modifier)(nil),           // 9: fillergame.Modifier
}
var file_FillerGame_proto_depIdxs = []int32{
	0, // 0: fillergame.JoinRequest.mode:type_name -> fillergame.GameBoard
	1, // 1: fillergame.MoveRequest.new_color:type_name -> fillergame.Color
	8, // 2: fillergame.BoardState.status:type_name -> fillergame.Status
	9, // 3: fillergame.BoardState.tiles:type_name -> fillergame.Modifier
	2, // 4: fillergame.Status.status_code:type_name -> fillergame.StatusCode
	1, // 5: fillergame.Modifier.new_color:type_name -> fillergame.Color
	3, // 6: fillergame.FillerGame.JoinGame:input_type -> fillergame.JoinRequest
	4, // 7: fillergame.FillerGame.MakeMove:input_type -> fillergame.MoveRequest
	5, // 8: fillergame.FillerGame.RebuildBoard:input_type -> fillergame.RebuildRequest
	6, // 9: fillergame.FillerGame.JoinGame:output_type -> fillergame.BoardState
	7, // 10: fillergame.FillerGame.MakeMove:output_type -> fillergame.MoveAcknowledgment
	6, // 11: fillergame.FillerGame.RebuildBoard:output_type -> fillergame.BoardState
	9, // [9:12] is the sub-list for method output_type
	6, // [6:9] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_FillerGame_proto_init() }
func file_FillerGame_proto_init() {
	if File_FillerGame_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_FillerGame_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinRequest); i {
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
		file_FillerGame_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MoveRequest); i {
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
		file_FillerGame_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RebuildRequest); i {
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
		file_FillerGame_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BoardState); i {
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
		file_FillerGame_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MoveAcknowledgment); i {
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
		file_FillerGame_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
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
		file_FillerGame_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Modifier); i {
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
			RawDescriptor: file_FillerGame_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_FillerGame_proto_goTypes,
		DependencyIndexes: file_FillerGame_proto_depIdxs,
		EnumInfos:         file_FillerGame_proto_enumTypes,
		MessageInfos:      file_FillerGame_proto_msgTypes,
	}.Build()
	File_FillerGame_proto = out.File
	file_FillerGame_proto_rawDesc = nil
	file_FillerGame_proto_goTypes = nil
	file_FillerGame_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// FillerGameClient is the client API for FillerGame service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FillerGameClient interface {
	JoinGame(ctx context.Context, in *JoinRequest, opts ...grpc.CallOption) (FillerGame_JoinGameClient, error)
	MakeMove(ctx context.Context, in *MoveRequest, opts ...grpc.CallOption) (*MoveAcknowledgment, error)
	RebuildBoard(ctx context.Context, in *RebuildRequest, opts ...grpc.CallOption) (FillerGame_RebuildBoardClient, error)
}

type fillerGameClient struct {
	cc grpc.ClientConnInterface
}

func NewFillerGameClient(cc grpc.ClientConnInterface) FillerGameClient {
	return &fillerGameClient{cc}
}

func (c *fillerGameClient) JoinGame(ctx context.Context, in *JoinRequest, opts ...grpc.CallOption) (FillerGame_JoinGameClient, error) {
	stream, err := c.cc.NewStream(ctx, &_FillerGame_serviceDesc.Streams[0], "/fillergame.FillerGame/JoinGame", opts...)
	if err != nil {
		return nil, err
	}
	x := &fillerGameJoinGameClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FillerGame_JoinGameClient interface {
	Recv() (*BoardState, error)
	grpc.ClientStream
}

type fillerGameJoinGameClient struct {
	grpc.ClientStream
}

func (x *fillerGameJoinGameClient) Recv() (*BoardState, error) {
	m := new(BoardState)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *fillerGameClient) MakeMove(ctx context.Context, in *MoveRequest, opts ...grpc.CallOption) (*MoveAcknowledgment, error) {
	out := new(MoveAcknowledgment)
	err := c.cc.Invoke(ctx, "/fillergame.FillerGame/MakeMove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fillerGameClient) RebuildBoard(ctx context.Context, in *RebuildRequest, opts ...grpc.CallOption) (FillerGame_RebuildBoardClient, error) {
	stream, err := c.cc.NewStream(ctx, &_FillerGame_serviceDesc.Streams[1], "/fillergame.FillerGame/RebuildBoard", opts...)
	if err != nil {
		return nil, err
	}
	x := &fillerGameRebuildBoardClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FillerGame_RebuildBoardClient interface {
	Recv() (*BoardState, error)
	grpc.ClientStream
}

type fillerGameRebuildBoardClient struct {
	grpc.ClientStream
}

func (x *fillerGameRebuildBoardClient) Recv() (*BoardState, error) {
	m := new(BoardState)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FillerGameServer is the server API for FillerGame service.
type FillerGameServer interface {
	JoinGame(*JoinRequest, FillerGame_JoinGameServer) error
	MakeMove(context.Context, *MoveRequest) (*MoveAcknowledgment, error)
	RebuildBoard(*RebuildRequest, FillerGame_RebuildBoardServer) error
}

// UnimplementedFillerGameServer can be embedded to have forward compatible implementations.
type UnimplementedFillerGameServer struct {
}

func (*UnimplementedFillerGameServer) JoinGame(*JoinRequest, FillerGame_JoinGameServer) error {
	return status.Errorf(codes.Unimplemented, "method JoinGame not implemented")
}
func (*UnimplementedFillerGameServer) MakeMove(context.Context, *MoveRequest) (*MoveAcknowledgment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MakeMove not implemented")
}
func (*UnimplementedFillerGameServer) RebuildBoard(*RebuildRequest, FillerGame_RebuildBoardServer) error {
	return status.Errorf(codes.Unimplemented, "method RebuildBoard not implemented")
}

func RegisterFillerGameServer(s *grpc.Server, srv FillerGameServer) {
	s.RegisterService(&_FillerGame_serviceDesc, srv)
}

func _FillerGame_JoinGame_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(JoinRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FillerGameServer).JoinGame(m, &fillerGameJoinGameServer{stream})
}

type FillerGame_JoinGameServer interface {
	Send(*BoardState) error
	grpc.ServerStream
}

type fillerGameJoinGameServer struct {
	grpc.ServerStream
}

func (x *fillerGameJoinGameServer) Send(m *BoardState) error {
	return x.ServerStream.SendMsg(m)
}

func _FillerGame_MakeMove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MoveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FillerGameServer).MakeMove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fillergame.FillerGame/MakeMove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FillerGameServer).MakeMove(ctx, req.(*MoveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FillerGame_RebuildBoard_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RebuildRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FillerGameServer).RebuildBoard(m, &fillerGameRebuildBoardServer{stream})
}

type FillerGame_RebuildBoardServer interface {
	Send(*BoardState) error
	grpc.ServerStream
}

type fillerGameRebuildBoardServer struct {
	grpc.ServerStream
}

func (x *fillerGameRebuildBoardServer) Send(m *BoardState) error {
	return x.ServerStream.SendMsg(m)
}

var _FillerGame_serviceDesc = grpc.ServiceDesc{
	ServiceName: "fillergame.FillerGame",
	HandlerType: (*FillerGameServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MakeMove",
			Handler:    _FillerGame_MakeMove_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "JoinGame",
			Handler:       _FillerGame_JoinGame_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "RebuildBoard",
			Handler:       _FillerGame_RebuildBoard_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "FillerGame.proto",
}
