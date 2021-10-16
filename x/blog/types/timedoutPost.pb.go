// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: blog/timedoutPost.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type TimedoutPost struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Id      uint64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Title   string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Chain   string `protobuf:"bytes,4,opt,name=chain,proto3" json:"chain,omitempty"`
}

func (m *TimedoutPost) Reset()         { *m = TimedoutPost{} }
func (m *TimedoutPost) String() string { return proto.CompactTextString(m) }
func (*TimedoutPost) ProtoMessage()    {}
func (*TimedoutPost) Descriptor() ([]byte, []int) {
	return fileDescriptor_77030da5d8316270, []int{0}
}
func (m *TimedoutPost) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TimedoutPost) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TimedoutPost.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TimedoutPost) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimedoutPost.Merge(m, src)
}
func (m *TimedoutPost) XXX_Size() int {
	return m.Size()
}
func (m *TimedoutPost) XXX_DiscardUnknown() {
	xxx_messageInfo_TimedoutPost.DiscardUnknown(m)
}

var xxx_messageInfo_TimedoutPost proto.InternalMessageInfo

func (m *TimedoutPost) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *TimedoutPost) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *TimedoutPost) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *TimedoutPost) GetChain() string {
	if m != nil {
		return m.Chain
	}
	return ""
}

func init() {
	proto.RegisterType((*TimedoutPost)(nil), "schnetzlerjoe.planet.blog.TimedoutPost")
}

func init() { proto.RegisterFile("blog/timedoutPost.proto", fileDescriptor_77030da5d8316270) }

var fileDescriptor_77030da5d8316270 = []byte{
	// 218 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4f, 0xca, 0xc9, 0x4f,
	0xd7, 0x2f, 0xc9, 0xcc, 0x4d, 0x4d, 0xc9, 0x2f, 0x2d, 0x09, 0xc8, 0x2f, 0x2e, 0xd1, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0x92, 0x2c, 0x4e, 0xce, 0xc8, 0x4b, 0x2d, 0xa9, 0xca, 0x49, 0x2d, 0xca,
	0xca, 0x4f, 0xd5, 0x2b, 0xc8, 0x49, 0xcc, 0x4b, 0x2d, 0xd1, 0x03, 0xa9, 0x96, 0x12, 0x49, 0xcf,
	0x4f, 0xcf, 0x07, 0xab, 0xd2, 0x07, 0xb1, 0x20, 0x1a, 0x94, 0x52, 0xb8, 0x78, 0x42, 0x90, 0x8c,
	0x11, 0x92, 0xe0, 0x62, 0x4f, 0x2e, 0x4a, 0x4d, 0x2c, 0xc9, 0x2f, 0x92, 0x60, 0x54, 0x60, 0xd4,
	0xe0, 0x0c, 0x82, 0x71, 0x85, 0xf8, 0xb8, 0x98, 0x32, 0x53, 0x24, 0x98, 0x14, 0x18, 0x35, 0x58,
	0x82, 0x98, 0x32, 0x53, 0x84, 0x44, 0xb8, 0x58, 0x4b, 0x32, 0x4b, 0x72, 0x52, 0x25, 0x98, 0xc1,
	0xea, 0x20, 0x1c, 0x90, 0x68, 0x72, 0x46, 0x62, 0x66, 0x9e, 0x04, 0x0b, 0x44, 0x14, 0xcc, 0x71,
	0x72, 0x3b, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x27, 0x3c,
	0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86, 0x28, 0x9d, 0xf4, 0xcc, 0x92,
	0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0x7d, 0x14, 0xb7, 0xeb, 0x43, 0xdc, 0xae, 0x5f, 0xa1,
	0x0f, 0xf1, 0x6b, 0x65, 0x41, 0x6a, 0x71, 0x12, 0x1b, 0xd8, 0xd1, 0xc6, 0x80, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x1c, 0x6e, 0x89, 0x88, 0x00, 0x01, 0x00, 0x00,
}

func (m *TimedoutPost) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TimedoutPost) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TimedoutPost) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Chain) > 0 {
		i -= len(m.Chain)
		copy(dAtA[i:], m.Chain)
		i = encodeVarintTimedoutPost(dAtA, i, uint64(len(m.Chain)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintTimedoutPost(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Id != 0 {
		i = encodeVarintTimedoutPost(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTimedoutPost(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTimedoutPost(dAtA []byte, offset int, v uint64) int {
	offset -= sovTimedoutPost(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TimedoutPost) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTimedoutPost(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovTimedoutPost(uint64(m.Id))
	}
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovTimedoutPost(uint64(l))
	}
	l = len(m.Chain)
	if l > 0 {
		n += 1 + l + sovTimedoutPost(uint64(l))
	}
	return n
}

func sovTimedoutPost(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTimedoutPost(x uint64) (n int) {
	return sovTimedoutPost(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TimedoutPost) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTimedoutPost
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
			return fmt.Errorf("proto: TimedoutPost: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TimedoutPost: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTimedoutPost
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTimedoutPost
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTimedoutPost
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTimedoutPost
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTimedoutPost
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTimedoutPost
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTimedoutPost
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTimedoutPost
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTimedoutPost
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTimedoutPost
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Chain = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTimedoutPost(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTimedoutPost
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
func skipTimedoutPost(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTimedoutPost
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
					return 0, ErrIntOverflowTimedoutPost
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
					return 0, ErrIntOverflowTimedoutPost
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
				return 0, ErrInvalidLengthTimedoutPost
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTimedoutPost
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTimedoutPost
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTimedoutPost        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTimedoutPost          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTimedoutPost = fmt.Errorf("proto: unexpected end of group")
)
