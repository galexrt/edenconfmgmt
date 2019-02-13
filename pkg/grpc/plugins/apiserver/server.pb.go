// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/galexrt/edenrun/pkg/grpc/plugins/apiserver/server.proto

package apiserver

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	descriptor "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	io "io"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type MessageOptions struct {
	Generate             *bool    `protobuf:"varint,1,opt,name=generate" json:"generate,omitempty"`
	Namespaced           *bool    `protobuf:"varint,2,opt,name=namespaced" json:"namespaced,omitempty"`
	CoreAPIPkg           *string  `protobuf:"bytes,3,opt,name=coreAPIPkg" json:"coreAPIPkg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MessageOptions) Reset()         { *m = MessageOptions{} }
func (m *MessageOptions) String() string { return proto.CompactTextString(m) }
func (*MessageOptions) ProtoMessage()    {}
func (*MessageOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_25862b7b5f04ef0b, []int{0}
}
func (m *MessageOptions) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MessageOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MessageOptions.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MessageOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageOptions.Merge(m, src)
}
func (m *MessageOptions) XXX_Size() int {
	return m.Size()
}
func (m *MessageOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageOptions.DiscardUnknown(m)
}

var xxx_messageInfo_MessageOptions proto.InternalMessageInfo

func (m *MessageOptions) GetGenerate() bool {
	if m != nil && m.Generate != nil {
		return *m.Generate
	}
	return false
}

func (m *MessageOptions) GetNamespaced() bool {
	if m != nil && m.Namespaced != nil {
		return *m.Namespaced
	}
	return false
}

func (m *MessageOptions) GetCoreAPIPkg() string {
	if m != nil && m.CoreAPIPkg != nil {
		return *m.CoreAPIPkg
	}
	return ""
}

var E_Options = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*MessageOptions)(nil),
	Field:         1337,
	Name:          "apiserver.options",
	Tag:           "bytes,1337,opt,name=options",
	Filename:      "github.com/galexrt/edenrun/pkg/grpc/plugins/apiserver/server.proto",
}

func init() {
	proto.RegisterType((*MessageOptions)(nil), "apiserver.MessageOptions")
	proto.RegisterExtension(E_Options)
}

func init() {
	proto.RegisterFile("github.com/galexrt/edenrun/pkg/grpc/plugins/apiserver/server.proto", fileDescriptor_25862b7b5f04ef0b)
}

var fileDescriptor_25862b7b5f04ef0b = []byte{
	// 270 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x50, 0x3d, 0x4e, 0xc3, 0x30,
	0x14, 0x96, 0x61, 0xa0, 0x35, 0x12, 0x43, 0xa6, 0x90, 0x21, 0x44, 0x4c, 0x59, 0xb0, 0x25, 0xc6,
	0x6e, 0xb0, 0x20, 0x24, 0x10, 0x55, 0x24, 0x16, 0x36, 0xd7, 0x79, 0x7d, 0x44, 0x4d, 0xfc, 0x2c,
	0xdb, 0x41, 0x1c, 0x85, 0xab, 0x70, 0x03, 0x46, 0x8e, 0x80, 0xc2, 0x45, 0x10, 0x49, 0x1b, 0x02,
	0x23, 0x93, 0xfd, 0xfd, 0xf8, 0xf3, 0xf7, 0x1e, 0xbf, 0xc2, 0x2a, 0x3c, 0xb6, 0x2b, 0xa1, 0xa9,
	0x91, 0xa8, 0x6a, 0x78, 0x76, 0x41, 0x42, 0x09, 0x46, 0x93, 0x59, 0x37, 0xd8, 0x04, 0x69, 0x37,
	0x28, 0xd1, 0x59, 0x2d, 0x6d, 0xdd, 0x62, 0x65, 0xbc, 0x54, 0xb6, 0xf2, 0xe0, 0x9e, 0xc0, 0xc9,
	0xe1, 0x10, 0xd6, 0x51, 0xa0, 0x68, 0x3e, 0xf2, 0xc9, 0xd9, 0x34, 0x93, 0x90, 0x64, 0xef, 0x58,
	0xb5, 0xeb, 0x1e, 0xf5, 0xa0, 0xbf, 0x0d, 0x2f, 0x93, 0x0c, 0x89, 0xb0, 0x86, 0x1f, 0x57, 0x09,
	0x5e, 0xbb, 0xca, 0x06, 0xda, 0x66, 0x9f, 0xd6, 0xfc, 0xe8, 0x16, 0xbc, 0x57, 0x08, 0x77, 0x36,
	0x54, 0x64, 0x7c, 0x94, 0xf0, 0x19, 0x82, 0x01, 0xa7, 0x02, 0xc4, 0x2c, 0x63, 0xf9, 0xac, 0x18,
	0x71, 0x94, 0x72, 0x6e, 0x54, 0x03, 0xde, 0x2a, 0x0d, 0x65, 0xbc, 0xd7, 0xab, 0x13, 0xe6, 0x5b,
	0xd7, 0xe4, 0xe0, 0x62, 0x79, 0xbd, 0xdc, 0x60, 0xbc, 0x9f, 0xb1, 0x7c, 0x5e, 0x4c, 0x98, 0xc5,
	0x3d, 0x3f, 0xa0, 0xed, 0x37, 0x27, 0x62, 0xe8, 0x26, 0x76, 0xdd, 0xc4, 0xef, 0x1e, 0xf1, 0x2b,
	0xcf, 0x58, 0x7e, 0x78, 0x7e, 0x2c, 0xc6, 0xe9, 0xff, 0x38, 0x8a, 0x5d, 0xd6, 0xe5, 0xcd, 0x5b,
	0x97, 0xb2, 0xf7, 0x2e, 0x65, 0x1f, 0x5d, 0xca, 0x5e, 0x3e, 0x53, 0xf6, 0xb0, 0xf8, 0xff, 0xee,
	0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0xbd, 0xa7, 0xdd, 0x0c, 0xb8, 0x01, 0x00, 0x00,
}

func (m *MessageOptions) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MessageOptions) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Generate != nil {
		dAtA[i] = 0x8
		i++
		if *m.Generate {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.Namespaced != nil {
		dAtA[i] = 0x10
		i++
		if *m.Namespaced {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.CoreAPIPkg != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintServer(dAtA, i, uint64(len(*m.CoreAPIPkg)))
		i += copy(dAtA[i:], *m.CoreAPIPkg)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintServer(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *MessageOptions) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Generate != nil {
		n += 2
	}
	if m.Namespaced != nil {
		n += 2
	}
	if m.CoreAPIPkg != nil {
		l = len(*m.CoreAPIPkg)
		n += 1 + l + sovServer(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovServer(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozServer(x uint64) (n int) {
	return sovServer(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MessageOptions) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowServer
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
			return fmt.Errorf("proto: MessageOptions: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MessageOptions: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Generate", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServer
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
			b := bool(v != 0)
			m.Generate = &b
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Namespaced", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServer
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
			b := bool(v != 0)
			m.Namespaced = &b
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CoreAPIPkg", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthServer
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.CoreAPIPkg = &s
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipServer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthServer
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
func skipServer(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowServer
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
					return 0, ErrIntOverflowServer
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
					return 0, ErrIntOverflowServer
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
				return 0, ErrInvalidLengthServer
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowServer
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
				next, err := skipServer(dAtA[start:])
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
	ErrInvalidLengthServer = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowServer   = fmt.Errorf("proto: integer overflow")
)
