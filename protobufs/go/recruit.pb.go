// Code generated by protoc-gen-go. DO NOT EDIT.
// source: recruit.proto

/*
Package protob is a generated protocol buffer package.

It is generated from these files:
	recruit.proto
	test.proto

It has these top-level messages:
	Tag
	Recruit
	Person
	Tag2
	Recruit2
	Person2
*/
package protob

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Tag struct {
	TagId int32  `protobuf:"varint,1,opt,name=tag_id,json=tagId" json:"tag_id,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *Tag) Reset()                    { *m = Tag{} }
func (m *Tag) String() string            { return proto.CompactTextString(m) }
func (*Tag) ProtoMessage()               {}
func (*Tag) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Tag) GetTagId() int32 {
	if m != nil {
		return m.TagId
	}
	return 0
}

func (m *Tag) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Recruit struct {
	Person *Person `protobuf:"bytes,1,opt,name=person" json:"person,omitempty"`
	Email  string  `protobuf:"bytes,3,opt,name=email" json:"email,omitempty"`
	Phone  string  `protobuf:"bytes,4,opt,name=phone" json:"phone,omitempty"`
	Tags   []*Tag  `protobuf:"bytes,5,rep,name=tags" json:"tags,omitempty"`
}

func (m *Recruit) Reset()                    { *m = Recruit{} }
func (m *Recruit) String() string            { return proto.CompactTextString(m) }
func (*Recruit) ProtoMessage()               {}
func (*Recruit) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Recruit) GetPerson() *Person {
	if m != nil {
		return m.Person
	}
	return nil
}

func (m *Recruit) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Recruit) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *Recruit) GetTags() []*Tag {
	if m != nil {
		return m.Tags
	}
	return nil
}

type Person struct {
	Id             int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	FirstName      string `protobuf:"bytes,2,opt,name=first_name,json=firstName" json:"first_name,omitempty"`
	LastName       string `protobuf:"bytes,3,opt,name=last_name,json=lastName" json:"last_name,omitempty"`
	PrimaryEmail   string `protobuf:"bytes,4,opt,name=primary_email,json=primaryEmail" json:"primary_email,omitempty"`
	PrimaryAddress string `protobuf:"bytes,5,opt,name=primary_address,json=primaryAddress" json:"primary_address,omitempty"`
	PrimaryPhone   string `protobuf:"bytes,6,opt,name=primary_phone,json=primaryPhone" json:"primary_phone,omitempty"`
	ProfilePic     string `protobuf:"bytes,7,opt,name=profile_pic,json=profilePic" json:"profile_pic,omitempty"`
}

func (m *Person) Reset()                    { *m = Person{} }
func (m *Person) String() string            { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()               {}
func (*Person) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Person) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Person) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *Person) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *Person) GetPrimaryEmail() string {
	if m != nil {
		return m.PrimaryEmail
	}
	return ""
}

func (m *Person) GetPrimaryAddress() string {
	if m != nil {
		return m.PrimaryAddress
	}
	return ""
}

func (m *Person) GetPrimaryPhone() string {
	if m != nil {
		return m.PrimaryPhone
	}
	return ""
}

func (m *Person) GetProfilePic() string {
	if m != nil {
		return m.ProfilePic
	}
	return ""
}

func init() {
	proto.RegisterType((*Tag)(nil), "protob.Tag")
	proto.RegisterType((*Recruit)(nil), "protob.Recruit")
	proto.RegisterType((*Person)(nil), "protob.Person")
}

func init() { proto.RegisterFile("recruit.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 288 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0xcb, 0x6a, 0xac, 0x40,
	0x10, 0x86, 0x71, 0xbc, 0xcc, 0xb1, 0x3c, 0x63, 0xa0, 0x49, 0xa0, 0x21, 0x84, 0x88, 0x81, 0xc4,
	0xd5, 0x10, 0x26, 0x4f, 0x90, 0x45, 0x16, 0xd9, 0x04, 0x91, 0xd9, 0x4b, 0xab, 0x3d, 0xa6, 0xc1,
	0x4b, 0xd3, 0x76, 0x16, 0xd9, 0xe4, 0x79, 0xf3, 0x18, 0xc1, 0x2a, 0x0d, 0x93, 0x95, 0xd6, 0xf7,
	0x97, 0x3f, 0x5f, 0x21, 0xec, 0x8c, 0xac, 0xcd, 0x87, 0xb2, 0x7b, 0x6d, 0x46, 0x3b, 0xb2, 0x00,
	0x1f, 0x55, 0xfa, 0x08, 0xee, 0x51, 0xb4, 0xec, 0x0a, 0x02, 0x2b, 0xda, 0x52, 0x35, 0xdc, 0x49,
	0x9c, 0xcc, 0x2f, 0x7c, 0x2b, 0xda, 0xd7, 0x86, 0x31, 0xf0, 0x06, 0xd1, 0x4b, 0xbe, 0x49, 0x9c,
	0x2c, 0x2c, 0xf0, 0x3d, 0xfd, 0x82, 0x6d, 0x41, 0x55, 0xec, 0x1e, 0x02, 0x2d, 0xcd, 0x34, 0x0e,
	0xf8, 0x55, 0x74, 0x88, 0xa9, 0xbc, 0xda, 0xe7, 0x48, 0x8b, 0x25, 0x65, 0x97, 0xe0, 0xcb, 0x5e,
	0xa8, 0x8e, 0xbb, 0xd8, 0x43, 0xc3, 0x4c, 0xf5, 0xfb, 0x38, 0x48, 0xee, 0x11, 0xc5, 0x81, 0xdd,
	0x82, 0x67, 0x45, 0x3b, 0x71, 0x3f, 0x71, 0xb3, 0xe8, 0x10, 0xad, 0x8d, 0x47, 0xd1, 0x16, 0x18,
	0xa4, 0xdf, 0x0e, 0x04, 0xd4, 0xcf, 0x62, 0xd8, 0xfc, 0x1a, 0x6f, 0x54, 0xc3, 0x6e, 0x00, 0x4e,
	0xca, 0x4c, 0xb6, 0x3c, 0x93, 0x0e, 0x91, 0xbc, 0x89, 0x5e, 0xb2, 0x6b, 0x08, 0x3b, 0xb1, 0xa6,
	0xa4, 0xf2, 0x6f, 0x06, 0x18, 0xde, 0xc1, 0x4e, 0x1b, 0xd5, 0x0b, 0xf3, 0x59, 0x92, 0x2b, 0x59,
	0xfd, 0x5f, 0xe0, 0x0b, 0x2a, 0x3f, 0xc0, 0xc5, 0xba, 0x24, 0x9a, 0xc6, 0xc8, 0x69, 0xf6, 0x9c,
	0xd7, 0xe2, 0x05, 0x3f, 0x13, 0x3d, 0x6f, 0xa3, 0x1b, 0x83, 0x3f, 0x6d, 0xf9, 0x72, 0x6a, 0xa4,
	0xcd, 0x78, 0x52, 0x9d, 0x2c, 0xb5, 0xaa, 0xf9, 0x16, 0x57, 0x60, 0x41, 0xb9, 0xaa, 0x2b, 0xfa,
	0x49, 0x4f, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xfc, 0xe9, 0xd8, 0x00, 0xbc, 0x01, 0x00, 0x00,
}
