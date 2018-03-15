// Code generated by protoc-gen-go. DO NOT EDIT.
// source: model.proto

/*
Package snomed is a generated protocol buffer package.

It is generated from these files:
	model.proto

It has these top-level messages:
	Concept
	Description
	Relationship
	ReferenceSetHeader
	RefSetDescriptorReferenceSet
	SimpleReferenceSet
	LanguageReferenceSet
*/
package snomed

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// A Concept represents a SNOMED-CT concept.
// The RF2 release allows multiple duplicate entries per concept identifier to permit versioning.
// As such, we have a compound primary key made up of the concept identifier and the effective time.
// Only one concept with a specified identifier will be active at any time point.
// See https://confluence.ihtsdotools.org/display/DOCRELFMT/3.2.1.+Concept+File+Specification
type Concept struct {
	Id                 int64                      `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	EffectiveTime      *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=effective_time,json=effectiveTime" json:"effective_time,omitempty"`
	Active             bool                       `protobuf:"varint,3,opt,name=active" json:"active,omitempty"`
	ModuleId           int64                      `protobuf:"varint,4,opt,name=module_id,json=moduleId" json:"module_id,omitempty"`
	DefinitionStatusId int64                      `protobuf:"varint,5,opt,name=definition_status_id,json=definitionStatusId" json:"definition_status_id,omitempty"`
}

func (m *Concept) Reset()                    { *m = Concept{} }
func (m *Concept) String() string            { return proto.CompactTextString(m) }
func (*Concept) ProtoMessage()               {}
func (*Concept) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Concept) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Concept) GetEffectiveTime() *google_protobuf.Timestamp {
	if m != nil {
		return m.EffectiveTime
	}
	return nil
}

func (m *Concept) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

func (m *Concept) GetModuleId() int64 {
	if m != nil {
		return m.ModuleId
	}
	return 0
}

func (m *Concept) GetDefinitionStatusId() int64 {
	if m != nil {
		return m.DefinitionStatusId
	}
	return 0
}

// A Description holds descriptions that describe SNOMED CT concepts.
// A description is used to give meaning to a concept and provide well-understood and standard ways of referring to a concept.
// See https://confluence.ihtsdotools.org/display/DOCRELFMT/3.2.2.+Description+File+Specification
type Description struct {
	Id               int64                      `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	EffectiveTime    *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=effective_time,json=effectiveTime" json:"effective_time,omitempty"`
	Active           bool                       `protobuf:"varint,3,opt,name=active" json:"active,omitempty"`
	ModuleId         int64                      `protobuf:"varint,4,opt,name=module_id,json=moduleId" json:"module_id,omitempty"`
	ConceptId        int64                      `protobuf:"varint,5,opt,name=concept_id,json=conceptId" json:"concept_id,omitempty"`
	LanguageCode     string                     `protobuf:"bytes,6,opt,name=LanguageCode" json:"LanguageCode,omitempty"`
	TypeId           int64                      `protobuf:"varint,7,opt,name=type_id,json=typeId" json:"type_id,omitempty"`
	Term             string                     `protobuf:"bytes,8,opt,name=term" json:"term,omitempty"`
	CaseSignificance int64                      `protobuf:"varint,9,opt,name=case_significance,json=caseSignificance" json:"case_significance,omitempty"`
}

func (m *Description) Reset()                    { *m = Description{} }
func (m *Description) String() string            { return proto.CompactTextString(m) }
func (*Description) ProtoMessage()               {}
func (*Description) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Description) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Description) GetEffectiveTime() *google_protobuf.Timestamp {
	if m != nil {
		return m.EffectiveTime
	}
	return nil
}

func (m *Description) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

func (m *Description) GetModuleId() int64 {
	if m != nil {
		return m.ModuleId
	}
	return 0
}

func (m *Description) GetConceptId() int64 {
	if m != nil {
		return m.ConceptId
	}
	return 0
}

func (m *Description) GetLanguageCode() string {
	if m != nil {
		return m.LanguageCode
	}
	return ""
}

func (m *Description) GetTypeId() int64 {
	if m != nil {
		return m.TypeId
	}
	return 0
}

func (m *Description) GetTerm() string {
	if m != nil {
		return m.Term
	}
	return ""
}

func (m *Description) GetCaseSignificance() int64 {
	if m != nil {
		return m.CaseSignificance
	}
	return 0
}

// Relationship defines a relationship between two concepts as a type itself defined as a concept
type Relationship struct {
	Id                   int64                      `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	EffectiveTime        *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=effective_time,json=effectiveTime" json:"effective_time,omitempty"`
	Active               bool                       `protobuf:"varint,3,opt,name=active" json:"active,omitempty"`
	ModuleId             int64                      `protobuf:"varint,4,opt,name=module_id,json=moduleId" json:"module_id,omitempty"`
	SourceId             int64                      `protobuf:"varint,5,opt,name=source_id,json=sourceId" json:"source_id,omitempty"`
	DestinationId        int64                      `protobuf:"varint,6,opt,name=destination_id,json=destinationId" json:"destination_id,omitempty"`
	RelationshipGroup    int64                      `protobuf:"varint,7,opt,name=relationship_group,json=relationshipGroup" json:"relationship_group,omitempty"`
	TypeId               int64                      `protobuf:"varint,8,opt,name=type_id,json=typeId" json:"type_id,omitempty"`
	CharacteristicTypeId int64                      `protobuf:"varint,9,opt,name=characteristic_type_id,json=characteristicTypeId" json:"characteristic_type_id,omitempty"`
	ModifierId           int64                      `protobuf:"varint,10,opt,name=modifier_id,json=modifierId" json:"modifier_id,omitempty"`
}

func (m *Relationship) Reset()                    { *m = Relationship{} }
func (m *Relationship) String() string            { return proto.CompactTextString(m) }
func (*Relationship) ProtoMessage()               {}
func (*Relationship) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Relationship) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Relationship) GetEffectiveTime() *google_protobuf.Timestamp {
	if m != nil {
		return m.EffectiveTime
	}
	return nil
}

func (m *Relationship) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

func (m *Relationship) GetModuleId() int64 {
	if m != nil {
		return m.ModuleId
	}
	return 0
}

func (m *Relationship) GetSourceId() int64 {
	if m != nil {
		return m.SourceId
	}
	return 0
}

func (m *Relationship) GetDestinationId() int64 {
	if m != nil {
		return m.DestinationId
	}
	return 0
}

func (m *Relationship) GetRelationshipGroup() int64 {
	if m != nil {
		return m.RelationshipGroup
	}
	return 0
}

func (m *Relationship) GetTypeId() int64 {
	if m != nil {
		return m.TypeId
	}
	return 0
}

func (m *Relationship) GetCharacteristicTypeId() int64 {
	if m != nil {
		return m.CharacteristicTypeId
	}
	return 0
}

func (m *Relationship) GetModifierId() int64 {
	if m != nil {
		return m.ModifierId
	}
	return 0
}

// ReferenceSet support customization and enhancement of SNOMED CT content. These include representation of subsets,
// language preferences maps for or from other code systems.
// There are multiple reference set types which extend this structure
// In the specification, the referenced component ID can be a SCT identifier or a UUID which is... problematic.
// In this structure, the referenced component ID is a SCT identifier... only. For now.
// Fortunately, in concrete types of reference set ("patterns"), it is made explicit.
type ReferenceSetHeader struct {
	Id                    string                     `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	EffectiveTime         *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=effective_time,json=effectiveTime" json:"effective_time,omitempty"`
	Active                bool                       `protobuf:"varint,3,opt,name=active" json:"active,omitempty"`
	ModuleId              int64                      `protobuf:"varint,4,opt,name=module_id,json=moduleId" json:"module_id,omitempty"`
	RefsetId              int64                      `protobuf:"varint,5,opt,name=refset_id,json=refsetId" json:"refset_id,omitempty"`
	ReferencedComponentId int64                      `protobuf:"varint,6,opt,name=referenced_component_id,json=referencedComponentId" json:"referenced_component_id,omitempty"`
}

func (m *ReferenceSetHeader) Reset()                    { *m = ReferenceSetHeader{} }
func (m *ReferenceSetHeader) String() string            { return proto.CompactTextString(m) }
func (*ReferenceSetHeader) ProtoMessage()               {}
func (*ReferenceSetHeader) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ReferenceSetHeader) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ReferenceSetHeader) GetEffectiveTime() *google_protobuf.Timestamp {
	if m != nil {
		return m.EffectiveTime
	}
	return nil
}

func (m *ReferenceSetHeader) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

func (m *ReferenceSetHeader) GetModuleId() int64 {
	if m != nil {
		return m.ModuleId
	}
	return 0
}

func (m *ReferenceSetHeader) GetRefsetId() int64 {
	if m != nil {
		return m.RefsetId
	}
	return 0
}

func (m *ReferenceSetHeader) GetReferencedComponentId() int64 {
	if m != nil {
		return m.ReferencedComponentId
	}
	return 0
}

// RefSetDescriptorReferenceSet is a type of reference set that provides information about a different reference set
// See https://confluence.ihtsdotools.org/display/DOCRELFMT/4.2.11.+Reference+Set+Descriptor
// It provides the additional structure for a given reference set.
type RefSetDescriptorReferenceSet struct {
	Header                 *ReferenceSetHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	AttributeDescriptionId int64               `protobuf:"varint,2,opt,name=attribute_description_id,json=attributeDescriptionId" json:"attribute_description_id,omitempty"`
	AttributeTypeId        int64               `protobuf:"varint,3,opt,name=attribute_type_id,json=attributeTypeId" json:"attribute_type_id,omitempty"`
	AttributeOrder         uint32              `protobuf:"varint,4,opt,name=attribute_order,json=attributeOrder" json:"attribute_order,omitempty"`
}

func (m *RefSetDescriptorReferenceSet) Reset()                    { *m = RefSetDescriptorReferenceSet{} }
func (m *RefSetDescriptorReferenceSet) String() string            { return proto.CompactTextString(m) }
func (*RefSetDescriptorReferenceSet) ProtoMessage()               {}
func (*RefSetDescriptorReferenceSet) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *RefSetDescriptorReferenceSet) GetHeader() *ReferenceSetHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *RefSetDescriptorReferenceSet) GetAttributeDescriptionId() int64 {
	if m != nil {
		return m.AttributeDescriptionId
	}
	return 0
}

func (m *RefSetDescriptorReferenceSet) GetAttributeTypeId() int64 {
	if m != nil {
		return m.AttributeTypeId
	}
	return 0
}

func (m *RefSetDescriptorReferenceSet) GetAttributeOrder() uint32 {
	if m != nil {
		return m.AttributeOrder
	}
	return 0
}

// SimpleReferenceSet is a simple reference set usable for defining subsets
// See https://confluence.ihtsdotools.org/display/DOCRELFMT/4.2.1.+Simple+Reference+Set
type SimpleReferenceSet struct {
	Header *ReferenceSetHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
}

func (m *SimpleReferenceSet) Reset()                    { *m = SimpleReferenceSet{} }
func (m *SimpleReferenceSet) String() string            { return proto.CompactTextString(m) }
func (*SimpleReferenceSet) ProtoMessage()               {}
func (*SimpleReferenceSet) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *SimpleReferenceSet) GetHeader() *ReferenceSetHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

// LanguageReferenceSet is a A 900000000000506000 |Language type reference set| supporting the representation of
// language and dialects preferences for the use of particular descriptions.
// "The most common use case for this type of reference set is to specify the acceptable and preferred terms
// for use within a particular country or region. However, the same type of reference set can also be used to
// represent preferences for use of descriptions in a more specific context such as a clinical specialty,
// organization or department.
//
// No more than one description of a specific description type associated with a single concept may have the acceptabilityId value 900000000000548007 |Preferred|.
// Every active concept should have one preferred synonym in each language.
// This means that a language reference set should assign the acceptabilityId  900000000000548007 |Preferred|  to one  synonym (a  description with  typeId value 900000000000013009 |synonym|) associated with each concept .
// This description is the preferred term for that concept in the specified language or dialect.
// Any  description which is not referenced by an active row in the   reference set is regarded as unacceptable (i.e. not a valid  synonym in the language or  dialect ).
// If a description becomes unacceptable, the relevant language reference set member is inactivated by adding a new row with the same id, the effectiveTime of the the change and the value active=0.
// For this reason there is no requirement for an "unacceptable" value."
// See https://confluence.ihtsdotools.org/display/DOCRELFMT/4.2.4.+Language+Reference+Set
//
type LanguageReferenceSet struct {
	Header          *ReferenceSetHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	AcceptabilityId int64               `protobuf:"varint,2,opt,name=acceptability_id,json=acceptabilityId" json:"acceptability_id,omitempty"`
}

func (m *LanguageReferenceSet) Reset()                    { *m = LanguageReferenceSet{} }
func (m *LanguageReferenceSet) String() string            { return proto.CompactTextString(m) }
func (*LanguageReferenceSet) ProtoMessage()               {}
func (*LanguageReferenceSet) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *LanguageReferenceSet) GetHeader() *ReferenceSetHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *LanguageReferenceSet) GetAcceptabilityId() int64 {
	if m != nil {
		return m.AcceptabilityId
	}
	return 0
}

func init() {
	proto.RegisterType((*Concept)(nil), "snomed.Concept")
	proto.RegisterType((*Description)(nil), "snomed.Description")
	proto.RegisterType((*Relationship)(nil), "snomed.Relationship")
	proto.RegisterType((*ReferenceSetHeader)(nil), "snomed.ReferenceSetHeader")
	proto.RegisterType((*RefSetDescriptorReferenceSet)(nil), "snomed.RefSetDescriptorReferenceSet")
	proto.RegisterType((*SimpleReferenceSet)(nil), "snomed.SimpleReferenceSet")
	proto.RegisterType((*LanguageReferenceSet)(nil), "snomed.LanguageReferenceSet")
}

func init() { proto.RegisterFile("model.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 604 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x54, 0xdf, 0x6a, 0xd4, 0x4c,
	0x14, 0x27, 0xdb, 0x7e, 0xe9, 0xe6, 0xb4, 0xdd, 0xb6, 0x43, 0xbf, 0x36, 0xb4, 0x4a, 0x97, 0x80,
	0xb8, 0x2a, 0x6e, 0xa5, 0x8a, 0x78, 0x2b, 0x15, 0x6c, 0x40, 0x10, 0xb2, 0xbd, 0x0f, 0xb3, 0x33,
	0x27, 0xe9, 0x40, 0x92, 0x09, 0x33, 0x13, 0xa1, 0xaf, 0xe5, 0x2b, 0xf8, 0x24, 0x3e, 0x82, 0x97,
	0xde, 0xc9, 0x4c, 0x36, 0x9b, 0x2c, 0x5e, 0xf6, 0xc2, 0xde, 0xed, 0xfc, 0xfe, 0xcc, 0x9e, 0xf3,
	0x3b, 0x99, 0x03, 0xbb, 0xa5, 0xe4, 0x58, 0xcc, 0x6b, 0x25, 0x8d, 0x24, 0xbe, 0xae, 0x64, 0x89,
	0xfc, 0xec, 0x22, 0x97, 0x32, 0x2f, 0xf0, 0xd2, 0xa1, 0xcb, 0x26, 0xbb, 0x34, 0xa2, 0x44, 0x6d,
	0x68, 0x59, 0xb7, 0xc2, 0xe8, 0x87, 0x07, 0x3b, 0xd7, 0xb2, 0x62, 0x58, 0x1b, 0x32, 0x81, 0x91,
	0xe0, 0xa1, 0x37, 0xf5, 0x66, 0x5b, 0xc9, 0x48, 0x70, 0xf2, 0x11, 0x26, 0x98, 0x65, 0xc8, 0x8c,
	0xf8, 0x86, 0xa9, 0x35, 0x86, 0xa3, 0xa9, 0x37, 0xdb, 0xbd, 0x3a, 0x9b, 0xb7, 0xb7, 0xce, 0xbb,
	0x5b, 0xe7, 0xb7, 0xdd, 0xad, 0xc9, 0xfe, 0xda, 0x61, 0x31, 0x72, 0x02, 0x3e, 0x75, 0xa7, 0x70,
	0x6b, 0xea, 0xcd, 0xc6, 0xc9, 0xea, 0x44, 0xce, 0x21, 0x28, 0x25, 0x6f, 0x0a, 0x4c, 0x05, 0x0f,
	0xb7, 0xdd, 0x3f, 0x8e, 0x5b, 0x20, 0xe6, 0xe4, 0x0d, 0x1c, 0x73, 0xcc, 0x44, 0x25, 0x8c, 0x90,
	0x55, 0xaa, 0x0d, 0x35, 0x8d, 0xb6, 0xba, 0xff, 0x9c, 0x8e, 0xf4, 0xdc, 0xc2, 0x51, 0x31, 0x8f,
	0xbe, 0x8f, 0x60, 0xf7, 0x13, 0x6a, 0xa6, 0x44, 0x6d, 0xf1, 0x47, 0xd3, 0xc9, 0x53, 0x00, 0xd6,
	0x86, 0xdb, 0xd7, 0x1f, 0xac, 0x90, 0x98, 0x93, 0x08, 0xf6, 0xbe, 0xd0, 0x2a, 0x6f, 0x68, 0x8e,
	0xd7, 0x92, 0x63, 0xe8, 0x4f, 0xbd, 0x59, 0x90, 0x6c, 0x60, 0xe4, 0x14, 0x76, 0xcc, 0x7d, 0xed,
	0x6e, 0xdf, 0x71, 0x7e, 0xdf, 0x1e, 0x63, 0x4e, 0x08, 0x6c, 0x1b, 0x54, 0x65, 0x38, 0x76, 0x26,
	0xf7, 0x9b, 0xbc, 0x82, 0x23, 0x46, 0x35, 0xa6, 0x5a, 0xe4, 0x95, 0xc8, 0x04, 0xa3, 0x15, 0xc3,
	0x30, 0x70, 0xb6, 0x43, 0x4b, 0x2c, 0x06, 0x78, 0xf4, 0x7b, 0x04, 0x7b, 0x09, 0x16, 0xd4, 0x26,
	0xa6, 0xef, 0x44, 0xfd, 0x68, 0x52, 0x3b, 0x87, 0x40, 0xcb, 0x46, 0x31, 0xec, 0x43, 0x1b, 0xb7,
	0x40, 0xcc, 0xc9, 0x33, 0x98, 0x70, 0xd4, 0x46, 0x54, 0xae, 0x6e, 0xab, 0xf0, 0x9d, 0x62, 0x7f,
	0x80, 0xc6, 0x9c, 0xbc, 0x06, 0xa2, 0x06, 0xbd, 0xa5, 0xb9, 0x92, 0x4d, 0xbd, 0x4a, 0xf0, 0x68,
	0xc8, 0x7c, 0xb6, 0xc4, 0x30, 0xe5, 0xf1, 0x46, 0xca, 0xef, 0xe0, 0x84, 0xdd, 0x51, 0x45, 0x99,
	0x41, 0x25, 0xb4, 0x11, 0x2c, 0xed, 0x74, 0x6d, 0xac, 0xc7, 0x9b, 0xec, 0x6d, 0xeb, 0xba, 0x70,
	0xaf, 0x51, 0x64, 0x02, 0x95, 0x95, 0x82, 0x93, 0x42, 0x07, 0xc5, 0x3c, 0xfa, 0xe5, 0x01, 0x49,
	0x30, 0x43, 0x85, 0x15, 0xc3, 0x05, 0x9a, 0x1b, 0xa4, 0x1c, 0xd5, 0x60, 0x02, 0xc1, 0xbf, 0x9e,
	0x80, 0xc2, 0x4c, 0xe3, 0xe0, 0xb3, 0x1d, 0xb7, 0x40, 0xcc, 0xc9, 0x7b, 0x38, 0x55, 0x5d, 0xe9,
	0x3c, 0x65, 0xb2, 0xac, 0x65, 0x85, 0x95, 0xe9, 0x47, 0xf1, 0x7f, 0x4f, 0x5f, 0x77, 0x6c, 0xcc,
	0xa3, 0x9f, 0x1e, 0x3c, 0x49, 0x30, 0x5b, 0xa0, 0xe9, 0x9e, 0xaa, 0x54, 0xc3, 0x0c, 0xc8, 0x15,
	0xf8, 0x77, 0x2e, 0x07, 0x97, 0x80, 0xed, 0xb2, 0xdd, 0x62, 0xf3, 0xbf, 0x93, 0x4a, 0x56, 0x4a,
	0xf2, 0x01, 0x42, 0x6a, 0x8c, 0x12, 0xcb, 0xc6, 0x60, 0xca, 0xfb, 0x15, 0x60, 0xab, 0x19, 0xb9,
	0x6a, 0x4e, 0xd6, 0xfc, 0x60, 0x43, 0xc4, 0x9c, 0xbc, 0x84, 0xa3, 0xde, 0xd9, 0x0d, 0x75, 0xcb,
	0x59, 0x0e, 0xd6, 0xc4, 0x6a, 0x9e, 0xcf, 0xa1, 0x87, 0x52, 0xa9, 0x6c, 0x89, 0x36, 0xb2, 0xfd,
	0x64, 0xb2, 0x86, 0xbf, 0x5a, 0x34, 0xba, 0x01, 0xb2, 0x10, 0x65, 0x5d, 0xe0, 0x43, 0x1b, 0x8b,
	0x1a, 0x38, 0xee, 0xf6, 0xc0, 0x83, 0x43, 0x7a, 0x01, 0x87, 0x94, 0xd9, 0x9d, 0x43, 0x97, 0xa2,
	0x10, 0xe6, 0xbe, 0x0f, 0xe7, 0x60, 0x03, 0x8f, 0xf9, 0xd2, 0x77, 0x5f, 0xd4, 0xdb, 0x3f, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x87, 0x7f, 0x4f, 0xfe, 0x4e, 0x06, 0x00, 0x00,
}
