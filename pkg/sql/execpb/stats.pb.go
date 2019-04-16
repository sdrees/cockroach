// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sql/execpb/stats.proto

package execpb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import time "time"

import github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// VectorizedStats represents the stats collected from an operator.
type VectorizedStats struct {
	ID         int32         `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	NumBatches int64         `protobuf:"varint,2,opt,name=num_batches,json=numBatches,proto3" json:"num_batches,omitempty"`
	NumTuples  int64         `protobuf:"varint,3,opt,name=num_tuples,json=numTuples,proto3" json:"num_tuples,omitempty"`
	Time       time.Duration `protobuf:"bytes,4,opt,name=time,proto3,stdduration" json:"time"`
	// stall indicates whether stall time or execution time is being tracked.
	Stall                bool     `protobuf:"varint,5,opt,name=stall,proto3" json:"stall,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VectorizedStats) Reset()         { *m = VectorizedStats{} }
func (m *VectorizedStats) String() string { return proto.CompactTextString(m) }
func (*VectorizedStats) ProtoMessage()    {}
func (*VectorizedStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_stats_3477029c41c5fd5d, []int{0}
}
func (m *VectorizedStats) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VectorizedStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalTo(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (dst *VectorizedStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VectorizedStats.Merge(dst, src)
}
func (m *VectorizedStats) XXX_Size() int {
	return m.Size()
}
func (m *VectorizedStats) XXX_DiscardUnknown() {
	xxx_messageInfo_VectorizedStats.DiscardUnknown(m)
}

var xxx_messageInfo_VectorizedStats proto.InternalMessageInfo

func init() {
	proto.RegisterType((*VectorizedStats)(nil), "cockroach.sql.execpb.VectorizedStats")
}
func (m *VectorizedStats) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VectorizedStats) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.ID != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintStats(dAtA, i, uint64(m.ID))
	}
	if m.NumBatches != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintStats(dAtA, i, uint64(m.NumBatches))
	}
	if m.NumTuples != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintStats(dAtA, i, uint64(m.NumTuples))
	}
	dAtA[i] = 0x22
	i++
	i = encodeVarintStats(dAtA, i, uint64(github_com_gogo_protobuf_types.SizeOfStdDuration(m.Time)))
	n1, err := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.Time, dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	if m.Stall {
		dAtA[i] = 0x28
		i++
		if m.Stall {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	return i, nil
}

func encodeVarintStats(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *VectorizedStats) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovStats(uint64(m.ID))
	}
	if m.NumBatches != 0 {
		n += 1 + sovStats(uint64(m.NumBatches))
	}
	if m.NumTuples != 0 {
		n += 1 + sovStats(uint64(m.NumTuples))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.Time)
	n += 1 + l + sovStats(uint64(l))
	if m.Stall {
		n += 2
	}
	return n
}

func sovStats(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozStats(x uint64) (n int) {
	return sovStats(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *VectorizedStats) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStats
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: VectorizedStats: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VectorizedStats: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ID |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumBatches", wireType)
			}
			m.NumBatches = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NumBatches |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumTuples", wireType)
			}
			m.NumTuples = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NumTuples |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Time", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStats
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.Time, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stall", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Stall = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipStats(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStats
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipStats(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStats
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
					return 0, ErrIntOverflowStats
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowStats
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthStats
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowStats
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipStats(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthStats = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStats   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("sql/execpb/stats.proto", fileDescriptor_stats_3477029c41c5fd5d) }

var fileDescriptor_stats_3477029c41c5fd5d = []byte{
	// 287 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0x90, 0xbd, 0x4e, 0xeb, 0x30,
	0x14, 0xc7, 0xeb, 0xb4, 0x89, 0x7a, 0xdd, 0xe1, 0x4a, 0x56, 0x54, 0x85, 0x4a, 0x38, 0x11, 0x53,
	0x26, 0x47, 0x82, 0x81, 0x3d, 0xea, 0xc2, 0x1a, 0x10, 0x03, 0x4b, 0x95, 0x38, 0x26, 0x8d, 0x70,
	0xe2, 0x34, 0xb6, 0x25, 0xc4, 0x53, 0x30, 0xf2, 0x28, 0x3c, 0x42, 0x46, 0x46, 0xa6, 0x02, 0xe1,
	0x45, 0x50, 0x6c, 0xd8, 0xfc, 0xff, 0xf0, 0x39, 0x3f, 0x1d, 0xb8, 0x96, 0x07, 0x9e, 0xb0, 0x47,
	0x46, 0xbb, 0x22, 0x91, 0x2a, 0x57, 0x92, 0x74, 0xbd, 0x50, 0x02, 0xf9, 0x54, 0xd0, 0x87, 0x5e,
	0xe4, 0x74, 0x4f, 0xe4, 0x81, 0x13, 0xdb, 0xd8, 0xf8, 0x95, 0xa8, 0x84, 0x29, 0x24, 0xd3, 0xcb,
	0x76, 0x37, 0xb8, 0x12, 0xa2, 0xe2, 0x2c, 0x31, 0xaa, 0xd0, 0xf7, 0x49, 0xa9, 0xfb, 0x5c, 0xd5,
	0xa2, 0xb5, 0xf9, 0xd9, 0x2b, 0x80, 0xff, 0x6f, 0x19, 0x55, 0xa2, 0xaf, 0x9f, 0x58, 0x79, 0x3d,
	0x6d, 0x41, 0x6b, 0xe8, 0xd4, 0x65, 0x00, 0x22, 0x10, 0xbb, 0xa9, 0x37, 0x1e, 0x43, 0xe7, 0x6a,
	0x9b, 0x39, 0x75, 0x89, 0x42, 0xb8, 0x6a, 0x75, 0xb3, 0x2b, 0x72, 0x45, 0xf7, 0x4c, 0x06, 0x4e,
	0x04, 0xe2, 0x79, 0x06, 0x5b, 0xdd, 0xa4, 0xd6, 0x41, 0xa7, 0x70, 0x52, 0x3b, 0xa5, 0x3b, 0xce,
	0x64, 0x30, 0x37, 0xf9, 0xbf, 0x56, 0x37, 0x37, 0xc6, 0x40, 0x97, 0x70, 0xa1, 0xea, 0x86, 0x05,
	0x8b, 0x08, 0xc4, 0xab, 0xf3, 0x13, 0x62, 0xd1, 0xc8, 0x1f, 0x1a, 0xd9, 0xfe, 0xa2, 0xa5, 0xcb,
	0xe1, 0x18, 0xce, 0x5e, 0x3e, 0x42, 0x90, 0x99, 0x0f, 0xc8, 0x87, 0xae, 0x54, 0x39, 0xe7, 0x81,
	0x1b, 0x81, 0x78, 0x99, 0x59, 0x91, 0x46, 0xc3, 0x17, 0x9e, 0x0d, 0x23, 0x06, 0x6f, 0x23, 0x06,
	0xef, 0x23, 0x06, 0x9f, 0x23, 0x06, 0xcf, 0xdf, 0x78, 0x76, 0xe7, 0xd9, 0x93, 0x14, 0x9e, 0x19,
	0x7d, 0xf1, 0x13, 0x00, 0x00, 0xff, 0xff, 0x90, 0x30, 0xf4, 0xdf, 0x49, 0x01, 0x00, 0x00,
}
