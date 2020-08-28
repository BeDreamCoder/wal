// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: wal.proto

package wal

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
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

type EntryType int32

const (
	EntryType_EntryNormal     EntryType = 0
	EntryType_EntryConfChange EntryType = 1
)

var EntryType_name = map[int32]string{
	0: "EntryNormal",
	1: "EntryConfChange",
}

var EntryType_value = map[string]int32{
	"EntryNormal":     0,
	"EntryConfChange": 1,
}

func (x EntryType) String() string {
	return proto.EnumName(EntryType_name, int32(x))
}

func (EntryType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ae6364fc8077884f, []int{0}
}

type Record struct {
	Type                 int64    `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	Crc                  uint32   `protobuf:"varint,2,opt,name=crc,proto3" json:"crc,omitempty"`
	Data                 []byte   `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Record) Reset()         { *m = Record{} }
func (m *Record) String() string { return proto.CompactTextString(m) }
func (*Record) ProtoMessage()    {}
func (*Record) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae6364fc8077884f, []int{0}
}
func (m *Record) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Record) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Record.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Record) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Record.Merge(m, src)
}
func (m *Record) XXX_Size() int {
	return m.Size()
}
func (m *Record) XXX_DiscardUnknown() {
	xxx_messageInfo_Record.DiscardUnknown(m)
}

var xxx_messageInfo_Record proto.InternalMessageInfo

func (m *Record) GetType() int64 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *Record) GetCrc() uint32 {
	if m != nil {
		return m.Crc
	}
	return 0
}

func (m *Record) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type Entry struct {
	Index                uint64    `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	Type                 EntryType `protobuf:"varint,2,opt,name=type,proto3,enum=wal.EntryType" json:"type,omitempty"`
	Data                 []byte    `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Entry) Reset()         { *m = Entry{} }
func (m *Entry) String() string { return proto.CompactTextString(m) }
func (*Entry) ProtoMessage()    {}
func (*Entry) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae6364fc8077884f, []int{1}
}
func (m *Entry) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Entry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Entry.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Entry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Entry.Merge(m, src)
}
func (m *Entry) XXX_Size() int {
	return m.Size()
}
func (m *Entry) XXX_DiscardUnknown() {
	xxx_messageInfo_Entry.DiscardUnknown(m)
}

var xxx_messageInfo_Entry proto.InternalMessageInfo

func (m *Entry) GetIndex() uint64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *Entry) GetType() EntryType {
	if m != nil {
		return m.Type
	}
	return EntryType_EntryNormal
}

func (m *Entry) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type Snapshot struct {
	Index                uint64   `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	Term                 uint64   `protobuf:"varint,2,opt,name=term,proto3" json:"term,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Snapshot) Reset()         { *m = Snapshot{} }
func (m *Snapshot) String() string { return proto.CompactTextString(m) }
func (*Snapshot) ProtoMessage()    {}
func (*Snapshot) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae6364fc8077884f, []int{2}
}
func (m *Snapshot) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Snapshot) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Snapshot.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Snapshot) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Snapshot.Merge(m, src)
}
func (m *Snapshot) XXX_Size() int {
	return m.Size()
}
func (m *Snapshot) XXX_DiscardUnknown() {
	xxx_messageInfo_Snapshot.DiscardUnknown(m)
}

var xxx_messageInfo_Snapshot proto.InternalMessageInfo

func (m *Snapshot) GetIndex() uint64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *Snapshot) GetTerm() uint64 {
	if m != nil {
		return m.Term
	}
	return 0
}

type HardState struct {
	Term                 uint64   `protobuf:"varint,1,opt,name=term,proto3" json:"term,omitempty"`
	Vote                 uint64   `protobuf:"varint,2,opt,name=vote,proto3" json:"vote,omitempty"`
	Commit               uint64   `protobuf:"varint,3,opt,name=commit,proto3" json:"commit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HardState) Reset()         { *m = HardState{} }
func (m *HardState) String() string { return proto.CompactTextString(m) }
func (*HardState) ProtoMessage()    {}
func (*HardState) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae6364fc8077884f, []int{3}
}
func (m *HardState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HardState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HardState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HardState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HardState.Merge(m, src)
}
func (m *HardState) XXX_Size() int {
	return m.Size()
}
func (m *HardState) XXX_DiscardUnknown() {
	xxx_messageInfo_HardState.DiscardUnknown(m)
}

var xxx_messageInfo_HardState proto.InternalMessageInfo

func (m *HardState) GetTerm() uint64 {
	if m != nil {
		return m.Term
	}
	return 0
}

func (m *HardState) GetVote() uint64 {
	if m != nil {
		return m.Vote
	}
	return 0
}

func (m *HardState) GetCommit() uint64 {
	if m != nil {
		return m.Commit
	}
	return 0
}

func init() {
	proto.RegisterEnum("wal.EntryType", EntryType_name, EntryType_value)
	proto.RegisterType((*Record)(nil), "wal.Record")
	proto.RegisterType((*Entry)(nil), "wal.Entry")
	proto.RegisterType((*Snapshot)(nil), "wal.Snapshot")
	proto.RegisterType((*HardState)(nil), "wal.HardState")
}

func init() { proto.RegisterFile("wal.proto", fileDescriptor_ae6364fc8077884f) }

var fileDescriptor_ae6364fc8077884f = []byte{
	// 296 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xdf, 0x4a, 0xb4, 0x40,
	0x18, 0xc6, 0x77, 0x3e, 0x5d, 0xf9, 0x7c, 0xab, 0x5d, 0x99, 0x22, 0x84, 0x40, 0xc4, 0x23, 0x09,
	0x72, 0xe9, 0xcf, 0x15, 0x68, 0x41, 0x10, 0x74, 0x30, 0x5b, 0x27, 0x9d, 0xcd, 0xea, 0xb4, 0x0a,
	0x8e, 0x23, 0xd3, 0x94, 0x79, 0x27, 0x5d, 0x52, 0x87, 0x5d, 0x42, 0xd8, 0x8d, 0xc4, 0x4c, 0xb6,
	0x9d, 0xd4, 0xd9, 0xef, 0x79, 0x79, 0xfc, 0xf9, 0xc0, 0x80, 0xdb, 0xd1, 0x3a, 0x69, 0xa5, 0x50,
	0x02, 0x5b, 0x1d, 0xad, 0xa3, 0x14, 0x1c, 0xc2, 0x72, 0x21, 0x0b, 0x8c, 0xc1, 0x56, 0x7d, 0xcb,
	0x7c, 0x14, 0xa2, 0xd8, 0x22, 0x86, 0xb1, 0x07, 0x56, 0x2e, 0x73, 0xff, 0x5f, 0x88, 0xe2, 0x1d,
	0xa2, 0x51, 0xb7, 0x0a, 0xaa, 0xa8, 0x6f, 0x85, 0x28, 0xde, 0x26, 0x86, 0xa3, 0x5b, 0x98, 0x5e,
	0x34, 0x4a, 0xf6, 0x78, 0x0f, 0xa6, 0x55, 0x53, 0xb0, 0x67, 0xe3, 0xb0, 0xc9, 0x57, 0xc0, 0xd1,
	0x28, 0xd6, 0x96, 0xd9, 0xc9, 0x2c, 0xd1, 0x0b, 0x4c, 0xff, 0xa6, 0x6f, 0xd9, 0xf8, 0xa3, 0xdf,
	0xb4, 0x67, 0xf0, 0x7f, 0xd9, 0xd0, 0xf6, 0xa1, 0x14, 0xea, 0x0f, 0xb3, 0x9e, 0xcc, 0x24, 0x37,
	0x66, 0x9b, 0x18, 0x8e, 0xae, 0xc0, 0xbd, 0xa4, 0xb2, 0x58, 0x2a, 0xaa, 0xd8, 0xa6, 0x80, 0x7e,
	0x0a, 0xfa, 0xf6, 0x24, 0x14, 0xfb, 0xfe, 0x48, 0x33, 0xde, 0x07, 0x27, 0x17, 0x9c, 0x57, 0xca,
	0x0c, 0xb0, 0xc9, 0x98, 0x0e, 0x8f, 0xc1, 0xdd, 0x2c, 0xc5, 0x73, 0xd8, 0x32, 0xe1, 0x5a, 0x48,
	0x4e, 0x6b, 0x6f, 0x82, 0x77, 0x61, 0x6e, 0x0e, 0x99, 0x68, 0xee, 0xb3, 0x92, 0x36, 0x6b, 0xe6,
	0xa1, 0xf4, 0xe8, 0x75, 0x08, 0xd0, 0xdb, 0x10, 0xa0, 0xf7, 0x21, 0x40, 0x2f, 0x1f, 0xc1, 0xe4,
	0xee, 0x60, 0x5d, 0xa9, 0xf2, 0x71, 0x95, 0xe4, 0x82, 0x2f, 0x52, 0x76, 0x2e, 0x19, 0xe5, 0x99,
	0x28, 0x98, 0x5c, 0x74, 0xb4, 0x5e, 0x39, 0xe6, 0x2d, 0x4e, 0x3f, 0x03, 0x00, 0x00, 0xff, 0xff,
	0x13, 0xe4, 0x0e, 0x30, 0x98, 0x01, 0x00, 0x00,
}

func (m *Record) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Record) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Record) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintWal(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Crc != 0 {
		i = encodeVarintWal(dAtA, i, uint64(m.Crc))
		i--
		dAtA[i] = 0x10
	}
	if m.Type != 0 {
		i = encodeVarintWal(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Entry) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Entry) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Entry) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintWal(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Type != 0 {
		i = encodeVarintWal(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x10
	}
	if m.Index != 0 {
		i = encodeVarintWal(dAtA, i, uint64(m.Index))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Snapshot) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Snapshot) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Snapshot) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Term != 0 {
		i = encodeVarintWal(dAtA, i, uint64(m.Term))
		i--
		dAtA[i] = 0x10
	}
	if m.Index != 0 {
		i = encodeVarintWal(dAtA, i, uint64(m.Index))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *HardState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HardState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HardState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Commit != 0 {
		i = encodeVarintWal(dAtA, i, uint64(m.Commit))
		i--
		dAtA[i] = 0x18
	}
	if m.Vote != 0 {
		i = encodeVarintWal(dAtA, i, uint64(m.Vote))
		i--
		dAtA[i] = 0x10
	}
	if m.Term != 0 {
		i = encodeVarintWal(dAtA, i, uint64(m.Term))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintWal(dAtA []byte, offset int, v uint64) int {
	offset -= sovWal(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Record) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovWal(uint64(m.Type))
	}
	if m.Crc != 0 {
		n += 1 + sovWal(uint64(m.Crc))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovWal(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Entry) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Index != 0 {
		n += 1 + sovWal(uint64(m.Index))
	}
	if m.Type != 0 {
		n += 1 + sovWal(uint64(m.Type))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovWal(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Snapshot) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Index != 0 {
		n += 1 + sovWal(uint64(m.Index))
	}
	if m.Term != 0 {
		n += 1 + sovWal(uint64(m.Term))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *HardState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Term != 0 {
		n += 1 + sovWal(uint64(m.Term))
	}
	if m.Vote != 0 {
		n += 1 + sovWal(uint64(m.Vote))
	}
	if m.Commit != 0 {
		n += 1 + sovWal(uint64(m.Commit))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovWal(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozWal(x uint64) (n int) {
	return sovWal(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Record) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWal
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Record: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Record: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Crc", wireType)
			}
			m.Crc = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Crc |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthWal
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthWal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthWal
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthWal
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Entry) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWal
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Entry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Entry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			m.Index = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Index |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= EntryType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthWal
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthWal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthWal
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthWal
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Snapshot) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWal
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Snapshot: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Snapshot: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			m.Index = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Index |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Term", wireType)
			}
			m.Term = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Term |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipWal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthWal
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthWal
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *HardState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWal
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: HardState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HardState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Term", wireType)
			}
			m.Term = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Term |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Vote", wireType)
			}
			m.Vote = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Vote |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Commit", wireType)
			}
			m.Commit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Commit |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipWal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthWal
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthWal
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipWal(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowWal
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowWal
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowWal
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthWal
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupWal
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthWal
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthWal        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowWal          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupWal = fmt.Errorf("proto: unexpected end of group")
)