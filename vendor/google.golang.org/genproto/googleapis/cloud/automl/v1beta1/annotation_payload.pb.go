// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/automl/v1beta1/annotation_payload.proto

package automl // import "google.golang.org/genproto/googleapis/cloud/automl/v1beta1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Contains annotation information that is relevant to AutoML.
type AnnotationPayload struct {
	// Output only . Additional information about the annotation
	// specific to the AutoML solution.
	//
	// Types that are valid to be assigned to Detail:
	//	*AnnotationPayload_Translation
	//	*AnnotationPayload_Classification
	Detail isAnnotationPayload_Detail `protobuf_oneof:"detail"`
	// Output only . The resource ID of the annotation spec that
	// this annotation pertains to. The annotation spec comes from either an
	// ancestor dataset, or the dataset that was used to train the model in use.
	AnnotationSpecId string `protobuf:"bytes,1,opt,name=annotation_spec_id,json=annotationSpecId,proto3" json:"annotation_spec_id,omitempty"`
	// Output only. The value of [AnnotationSpec.display_name][google.cloud.automl.v1beta1.AnnotationSpec.display_name] when the model
	// was trained. Because this field returns a value at model training time,
	// for different models trained using the same dataset, the returned value
	// could be different as model owner could update the display_name between
	// any two model training.
	DisplayName          string   `protobuf:"bytes,5,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AnnotationPayload) Reset()         { *m = AnnotationPayload{} }
func (m *AnnotationPayload) String() string { return proto.CompactTextString(m) }
func (*AnnotationPayload) ProtoMessage()    {}
func (*AnnotationPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_annotation_payload_f7dca833885c28d5, []int{0}
}
func (m *AnnotationPayload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AnnotationPayload.Unmarshal(m, b)
}
func (m *AnnotationPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AnnotationPayload.Marshal(b, m, deterministic)
}
func (dst *AnnotationPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AnnotationPayload.Merge(dst, src)
}
func (m *AnnotationPayload) XXX_Size() int {
	return xxx_messageInfo_AnnotationPayload.Size(m)
}
func (m *AnnotationPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_AnnotationPayload.DiscardUnknown(m)
}

var xxx_messageInfo_AnnotationPayload proto.InternalMessageInfo

type isAnnotationPayload_Detail interface {
	isAnnotationPayload_Detail()
}

type AnnotationPayload_Translation struct {
	Translation *TranslationAnnotation `protobuf:"bytes,2,opt,name=translation,proto3,oneof"`
}

type AnnotationPayload_Classification struct {
	Classification *ClassificationAnnotation `protobuf:"bytes,3,opt,name=classification,proto3,oneof"`
}

func (*AnnotationPayload_Translation) isAnnotationPayload_Detail() {}

func (*AnnotationPayload_Classification) isAnnotationPayload_Detail() {}

func (m *AnnotationPayload) GetDetail() isAnnotationPayload_Detail {
	if m != nil {
		return m.Detail
	}
	return nil
}

func (m *AnnotationPayload) GetTranslation() *TranslationAnnotation {
	if x, ok := m.GetDetail().(*AnnotationPayload_Translation); ok {
		return x.Translation
	}
	return nil
}

func (m *AnnotationPayload) GetClassification() *ClassificationAnnotation {
	if x, ok := m.GetDetail().(*AnnotationPayload_Classification); ok {
		return x.Classification
	}
	return nil
}

func (m *AnnotationPayload) GetAnnotationSpecId() string {
	if m != nil {
		return m.AnnotationSpecId
	}
	return ""
}

func (m *AnnotationPayload) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*AnnotationPayload) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _AnnotationPayload_OneofMarshaler, _AnnotationPayload_OneofUnmarshaler, _AnnotationPayload_OneofSizer, []interface{}{
		(*AnnotationPayload_Translation)(nil),
		(*AnnotationPayload_Classification)(nil),
	}
}

func _AnnotationPayload_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*AnnotationPayload)
	// detail
	switch x := m.Detail.(type) {
	case *AnnotationPayload_Translation:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Translation); err != nil {
			return err
		}
	case *AnnotationPayload_Classification:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Classification); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("AnnotationPayload.Detail has unexpected type %T", x)
	}
	return nil
}

func _AnnotationPayload_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*AnnotationPayload)
	switch tag {
	case 2: // detail.translation
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TranslationAnnotation)
		err := b.DecodeMessage(msg)
		m.Detail = &AnnotationPayload_Translation{msg}
		return true, err
	case 3: // detail.classification
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ClassificationAnnotation)
		err := b.DecodeMessage(msg)
		m.Detail = &AnnotationPayload_Classification{msg}
		return true, err
	default:
		return false, nil
	}
}

func _AnnotationPayload_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*AnnotationPayload)
	// detail
	switch x := m.Detail.(type) {
	case *AnnotationPayload_Translation:
		s := proto.Size(x.Translation)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *AnnotationPayload_Classification:
		s := proto.Size(x.Classification)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*AnnotationPayload)(nil), "google.cloud.automl.v1beta1.AnnotationPayload")
}

func init() {
	proto.RegisterFile("google/cloud/automl/v1beta1/annotation_payload.proto", fileDescriptor_annotation_payload_f7dca833885c28d5)
}

var fileDescriptor_annotation_payload_f7dca833885c28d5 = []byte{
	// 320 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0xed, 0xc4, 0xa1, 0x99, 0x88, 0xf6, 0x54, 0x36, 0xc1, 0xe9, 0x69, 0x07, 0x4d, 0xdd,
	0xd4, 0x93, 0xa7, 0x6d, 0x07, 0xf5, 0xa0, 0x8c, 0x29, 0x3b, 0xc8, 0xa0, 0xbc, 0xb5, 0x31, 0x04,
	0xd2, 0xbc, 0xb0, 0x64, 0xc2, 0xee, 0x7e, 0x17, 0xbf, 0x8b, 0x9f, 0x4a, 0x4c, 0x8a, 0x6b, 0x45,
	0x7a, 0x4c, 0xde, 0xff, 0xf7, 0x7b, 0x2f, 0x79, 0xe4, 0x9a, 0x23, 0x72, 0xc9, 0xe2, 0x54, 0xe2,
	0x2a, 0x8b, 0x61, 0x65, 0x31, 0x97, 0xf1, 0x7b, 0x7f, 0xc1, 0x2c, 0xf4, 0x63, 0x50, 0x0a, 0x2d,
	0x58, 0x81, 0x2a, 0xd1, 0xb0, 0x96, 0x08, 0x19, 0xd5, 0x4b, 0xb4, 0x18, 0x76, 0x3c, 0x45, 0x1d,
	0x45, 0x3d, 0x45, 0x0b, 0xaa, 0x7d, 0x5c, 0x28, 0x41, 0x8b, 0x92, 0xc1, 0x78, 0xb4, 0x7d, 0x59,
	0xd7, 0x30, 0x95, 0x60, 0x8c, 0x78, 0x13, 0xa9, 0x43, 0x0a, 0xe2, 0xa2, 0x8e, 0xb0, 0x4b, 0x50,
	0x46, 0x96, 0xe2, 0x67, 0x9f, 0x0d, 0x72, 0x34, 0xfc, 0x6d, 0x3b, 0xf1, 0x73, 0x87, 0x33, 0xd2,
	0x2a, 0x45, 0xa3, 0x46, 0x37, 0xe8, 0xb5, 0x06, 0x03, 0x5a, 0xf3, 0x0e, 0xfa, 0xb2, 0xc9, 0x6f,
	0x7c, 0xf7, 0x5b, 0xd3, 0xb2, 0x28, 0x4c, 0xc8, 0x41, 0x75, 0xe8, 0x68, 0xdb, 0xa9, 0x6f, 0x6a,
	0xd5, 0xe3, 0x0a, 0x52, 0xb1, 0xff, 0xd1, 0x85, 0xe7, 0x24, 0x2c, 0xad, 0xc1, 0x68, 0x96, 0x26,
	0x22, 0x8b, 0x82, 0x6e, 0xd0, 0xdb, 0x9b, 0x1e, 0x6e, 0x2a, 0xcf, 0x9a, 0xa5, 0x0f, 0x59, 0x78,
	0x4a, 0xf6, 0x33, 0x61, 0xb4, 0x84, 0x75, 0xa2, 0x20, 0x67, 0xd1, 0x8e, 0xcb, 0xb5, 0x8a, 0xbb,
	0x27, 0xc8, 0xd9, 0x68, 0x97, 0x34, 0x33, 0x66, 0x41, 0xc8, 0xd1, 0x47, 0x40, 0x4e, 0x52, 0xcc,
	0xeb, 0x26, 0x9d, 0x04, 0xaf, 0xc3, 0xa2, 0xcc, 0x51, 0x82, 0xe2, 0x14, 0x97, 0x3c, 0xe6, 0x4c,
	0xb9, 0xbf, 0x8e, 0x7d, 0x09, 0xb4, 0x30, 0xff, 0x6e, 0xe7, 0xd6, 0x1f, 0xbf, 0x1a, 0x9d, 0x3b,
	0x17, 0x9c, 0x8f, 0x7f, 0x42, 0xf3, 0xe1, 0xca, 0xe2, 0xa3, 0x9c, 0xcf, 0x7c, 0x68, 0xd1, 0x74,
	0xae, 0xab, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd5, 0x18, 0x90, 0x9a, 0x8b, 0x02, 0x00, 0x00,
}
