// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nodeclient/nodeclientproto/protos/nodeclient.proto

package nodeclientproto

import (
	fmt "fmt"
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

type ErrCodes int32

const (
	ErrCodes_Unexpected    ErrCodes = 0
	ErrCodes_InvalidRecord ErrCodes = 1
	ErrCodes_ErrorOffset   ErrCodes = 1100
)

var ErrCodes_name = map[int32]string{
	0:    "Unexpected",
	1:    "InvalidRecord",
	1100: "ErrorOffset",
}

var ErrCodes_value = map[string]int32{
	"Unexpected":    0,
	"InvalidRecord": 1,
	"ErrorOffset":   1100,
}

func (x ErrCodes) String() string {
	return proto.EnumName(ErrCodes_name, int32(x))
}

func (ErrCodes) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3d5653db8f44420b, []int{0}
}

type AclAddRequest struct {
	ClientRecord []byte `protobuf:"bytes,1,opt,name=clientRecord,proto3" json:"clientRecord,omitempty"`
	SpaceId      []byte `protobuf:"bytes,2,opt,name=spaceId,proto3" json:"spaceId,omitempty"`
}

func (m *AclAddRequest) Reset()         { *m = AclAddRequest{} }
func (m *AclAddRequest) String() string { return proto.CompactTextString(m) }
func (*AclAddRequest) ProtoMessage()    {}
func (*AclAddRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d5653db8f44420b, []int{0}
}
func (m *AclAddRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AclAddRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AclAddRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AclAddRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AclAddRequest.Merge(m, src)
}
func (m *AclAddRequest) XXX_Size() int {
	return m.Size()
}
func (m *AclAddRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AclAddRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AclAddRequest proto.InternalMessageInfo

func (m *AclAddRequest) GetClientRecord() []byte {
	if m != nil {
		return m.ClientRecord
	}
	return nil
}

func (m *AclAddRequest) GetSpaceId() []byte {
	if m != nil {
		return m.SpaceId
	}
	return nil
}

type AclAddResponse struct {
	AddedRecord []byte `protobuf:"bytes,1,opt,name=addedRecord,proto3" json:"addedRecord,omitempty"`
	RecordId    []byte `protobuf:"bytes,2,opt,name=recordId,proto3" json:"recordId,omitempty"`
}

func (m *AclAddResponse) Reset()         { *m = AclAddResponse{} }
func (m *AclAddResponse) String() string { return proto.CompactTextString(m) }
func (*AclAddResponse) ProtoMessage()    {}
func (*AclAddResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d5653db8f44420b, []int{1}
}
func (m *AclAddResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AclAddResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AclAddResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AclAddResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AclAddResponse.Merge(m, src)
}
func (m *AclAddResponse) XXX_Size() int {
	return m.Size()
}
func (m *AclAddResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AclAddResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AclAddResponse proto.InternalMessageInfo

func (m *AclAddResponse) GetAddedRecord() []byte {
	if m != nil {
		return m.AddedRecord
	}
	return nil
}

func (m *AclAddResponse) GetRecordId() []byte {
	if m != nil {
		return m.RecordId
	}
	return nil
}

func init() {
	proto.RegisterEnum("anyNodeClient.ErrCodes", ErrCodes_name, ErrCodes_value)
	proto.RegisterType((*AclAddRequest)(nil), "anyNodeClient.AclAddRequest")
	proto.RegisterType((*AclAddResponse)(nil), "anyNodeClient.AclAddResponse")
}

func init() {
	proto.RegisterFile("nodeclient/nodeclientproto/protos/nodeclient.proto", fileDescriptor_3d5653db8f44420b)
}

var fileDescriptor_3d5653db8f44420b = []byte{
	// 278 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xbd, 0x4e, 0xc3, 0x30,
	0x10, 0x80, 0x63, 0x86, 0x12, 0x5d, 0x9b, 0x2a, 0x78, 0x8a, 0x22, 0xb0, 0xaa, 0x4c, 0x88, 0x21,
	0x95, 0x0a, 0x3b, 0x2a, 0x55, 0x86, 0x0e, 0x14, 0x29, 0x88, 0x85, 0x2d, 0xe4, 0xae, 0x52, 0xa5,
	0xc8, 0x0e, 0x76, 0x40, 0xf0, 0x16, 0x3c, 0x16, 0x03, 0x43, 0x47, 0x46, 0x94, 0xbc, 0x08, 0x92,
	0xd3, 0xbf, 0x0c, 0x5d, 0xec, 0xbb, 0xef, 0xac, 0x4f, 0xbe, 0x3b, 0x98, 0x48, 0x85, 0x94, 0x17,
	0x2b, 0x92, 0xd5, 0x78, 0x1f, 0x96, 0x5a, 0x55, 0x6a, 0x6c, 0x4f, 0x73, 0x80, 0x63, 0x4b, 0xb8,
	0x97, 0xc9, 0xcf, 0x85, 0x42, 0x9a, 0x59, 0x18, 0xdd, 0x83, 0x37, 0xcd, 0x8b, 0x29, 0x62, 0x4a,
	0xaf, 0x6f, 0x64, 0x2a, 0x1e, 0xc1, 0xa0, 0x7d, 0x9f, 0x52, 0xae, 0x34, 0x06, 0x6c, 0xc4, 0x2e,
	0x07, 0x69, 0x87, 0xf1, 0x00, 0x4e, 0x4d, 0x99, 0xe5, 0x34, 0xc7, 0xe0, 0xc4, 0x96, 0xb7, 0x69,
	0xb4, 0x80, 0xe1, 0x56, 0x67, 0x4a, 0x25, 0x0d, 0xf1, 0x11, 0xf4, 0x33, 0x44, 0xc2, 0x8e, 0xee,
	0x10, 0xf1, 0x10, 0x5c, 0x6d, 0xa3, 0x9d, 0x6e, 0x97, 0x5f, 0xdd, 0x82, 0x9b, 0x68, 0x3d, 0x53,
	0x48, 0x86, 0x0f, 0x01, 0x9e, 0x24, 0x7d, 0x94, 0x94, 0x57, 0x84, 0xbe, 0xc3, 0xcf, 0xc0, 0x9b,
	0xcb, 0xf7, 0xac, 0x58, 0x6d, 0x44, 0x3e, 0xe3, 0x3e, 0xf4, 0x13, 0xad, 0x95, 0x7e, 0x58, 0x2e,
	0x0d, 0x55, 0xfe, 0x8f, 0x3b, 0x79, 0x04, 0xd8, 0x77, 0xcb, 0x13, 0xe8, 0xb5, 0xdf, 0xe3, 0xe7,
	0x71, 0x67, 0x0e, 0x71, 0x67, 0x08, 0xe1, 0xc5, 0x91, 0x6a, 0xdb, 0xd3, 0xdd, 0xcd, 0x77, 0x2d,
	0xd8, 0xba, 0x16, 0xec, 0xaf, 0x16, 0xec, 0xab, 0x11, 0xce, 0xba, 0x11, 0xce, 0x6f, 0x23, 0x9c,
	0xe7, 0xf0, 0xf8, 0x46, 0x5e, 0x7a, 0xf6, 0xba, 0xfe, 0x0f, 0x00, 0x00, 0xff, 0xff, 0x00, 0x60,
	0x18, 0x24, 0xb6, 0x01, 0x00, 0x00,
}

func (m *AclAddRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AclAddRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AclAddRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SpaceId) > 0 {
		i -= len(m.SpaceId)
		copy(dAtA[i:], m.SpaceId)
		i = encodeVarintNodeclient(dAtA, i, uint64(len(m.SpaceId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ClientRecord) > 0 {
		i -= len(m.ClientRecord)
		copy(dAtA[i:], m.ClientRecord)
		i = encodeVarintNodeclient(dAtA, i, uint64(len(m.ClientRecord)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *AclAddResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AclAddResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AclAddResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.RecordId) > 0 {
		i -= len(m.RecordId)
		copy(dAtA[i:], m.RecordId)
		i = encodeVarintNodeclient(dAtA, i, uint64(len(m.RecordId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.AddedRecord) > 0 {
		i -= len(m.AddedRecord)
		copy(dAtA[i:], m.AddedRecord)
		i = encodeVarintNodeclient(dAtA, i, uint64(len(m.AddedRecord)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintNodeclient(dAtA []byte, offset int, v uint64) int {
	offset -= sovNodeclient(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AclAddRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ClientRecord)
	if l > 0 {
		n += 1 + l + sovNodeclient(uint64(l))
	}
	l = len(m.SpaceId)
	if l > 0 {
		n += 1 + l + sovNodeclient(uint64(l))
	}
	return n
}

func (m *AclAddResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.AddedRecord)
	if l > 0 {
		n += 1 + l + sovNodeclient(uint64(l))
	}
	l = len(m.RecordId)
	if l > 0 {
		n += 1 + l + sovNodeclient(uint64(l))
	}
	return n
}

func sovNodeclient(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozNodeclient(x uint64) (n int) {
	return sovNodeclient(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AclAddRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNodeclient
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
			return fmt.Errorf("proto: AclAddRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AclAddRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientRecord", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNodeclient
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
				return ErrInvalidLengthNodeclient
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthNodeclient
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClientRecord = append(m.ClientRecord[:0], dAtA[iNdEx:postIndex]...)
			if m.ClientRecord == nil {
				m.ClientRecord = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SpaceId", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNodeclient
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
				return ErrInvalidLengthNodeclient
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthNodeclient
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SpaceId = append(m.SpaceId[:0], dAtA[iNdEx:postIndex]...)
			if m.SpaceId == nil {
				m.SpaceId = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNodeclient(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNodeclient
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
func (m *AclAddResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNodeclient
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
			return fmt.Errorf("proto: AclAddResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AclAddResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AddedRecord", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNodeclient
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
				return ErrInvalidLengthNodeclient
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthNodeclient
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AddedRecord = append(m.AddedRecord[:0], dAtA[iNdEx:postIndex]...)
			if m.AddedRecord == nil {
				m.AddedRecord = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RecordId", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNodeclient
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
				return ErrInvalidLengthNodeclient
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthNodeclient
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RecordId = append(m.RecordId[:0], dAtA[iNdEx:postIndex]...)
			if m.RecordId == nil {
				m.RecordId = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNodeclient(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNodeclient
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
func skipNodeclient(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNodeclient
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
					return 0, ErrIntOverflowNodeclient
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
					return 0, ErrIntOverflowNodeclient
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
				return 0, ErrInvalidLengthNodeclient
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupNodeclient
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthNodeclient
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthNodeclient        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNodeclient          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupNodeclient = fmt.Errorf("proto: unexpected end of group")
)
