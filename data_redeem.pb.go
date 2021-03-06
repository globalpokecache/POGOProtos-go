// Code generated by protoc-gen-go.
// source: data_redeem.proto
// DO NOT EDIT!

package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Ignoring public import of ItemAward from inventory_item.proto

// Ignoring public import of ItemData from inventory_item.proto

// Ignoring public import of ItemType from inventory_item.proto

// Ignoring public import of ItemId from inventory_item.proto

type PokeCandy struct {
	PokemonId  uint64 `protobuf:"fixed64,1,opt,name=pokemon_id" json:"pokemon_id,omitempty"`
	CandyCount int32  `protobuf:"varint,2,opt,name=candy_count" json:"candy_count,omitempty"`
}

func (m *PokeCandy) Reset()                    { *m = PokeCandy{} }
func (m *PokeCandy) String() string            { return proto.CompactTextString(m) }
func (*PokeCandy) ProtoMessage()               {}
func (*PokeCandy) Descriptor() ([]byte, []int) { return fileDescriptor16, []int{0} }

type RedeemedItem struct {
	Item      ItemId `protobuf:"varint,1,opt,name=item,enum=POGOProtos.Inventory.Item.ItemId" json:"item,omitempty"`
	ItemCount int32  `protobuf:"varint,2,opt,name=item_count" json:"item_count,omitempty"`
}

func (m *RedeemedItem) Reset()                    { *m = RedeemedItem{} }
func (m *RedeemedItem) String() string            { return proto.CompactTextString(m) }
func (*RedeemedItem) ProtoMessage()               {}
func (*RedeemedItem) Descriptor() ([]byte, []int) { return fileDescriptor16, []int{1} }

type RedeemedAvatarItem struct {
	AvatarTemplateId string `protobuf:"bytes,1,opt,name=avatar_template_id" json:"avatar_template_id,omitempty"`
	ItemCount        int32  `protobuf:"varint,2,opt,name=item_count" json:"item_count,omitempty"`
}

func (m *RedeemedAvatarItem) Reset()                    { *m = RedeemedAvatarItem{} }
func (m *RedeemedAvatarItem) String() string            { return proto.CompactTextString(m) }
func (*RedeemedAvatarItem) ProtoMessage()               {}
func (*RedeemedAvatarItem) Descriptor() ([]byte, []int) { return fileDescriptor16, []int{2} }

func init() {
	proto.RegisterType((*PokeCandy)(nil), "POGOProtos.Data.Redeem.PokeCandy")
	proto.RegisterType((*RedeemedItem)(nil), "POGOProtos.Data.Redeem.RedeemedItem")
	proto.RegisterType((*RedeemedAvatarItem)(nil), "POGOProtos.Data.Redeem.RedeemedAvatarItem")
}

func init() { proto.RegisterFile("data_redeem.proto", fileDescriptor16) }

var fileDescriptor16 = []byte{
	// 218 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x8f, 0x41, 0x4b, 0xc4, 0x30,
	0x10, 0x85, 0xad, 0xe8, 0xe2, 0x8e, 0x22, 0x18, 0x45, 0x96, 0x3d, 0xad, 0x7b, 0xea, 0x29, 0x82,
	0xfa, 0x07, 0xd4, 0x82, 0xf4, 0xd4, 0xa0, 0x37, 0x2f, 0x61, 0x6c, 0xe6, 0x50, 0x6a, 0x92, 0x92,
	0x8e, 0x85, 0xfe, 0x7b, 0x49, 0x4a, 0x40, 0x61, 0x2f, 0x61, 0xc2, 0xbc, 0xf7, 0xbd, 0x37, 0x70,
	0x65, 0x90, 0x51, 0x07, 0x32, 0x44, 0x56, 0x0e, 0xc1, 0xb3, 0x17, 0xb7, 0xaa, 0x79, 0x6b, 0x54,
	0x1c, 0x47, 0x59, 0x21, 0xa3, 0x7c, 0x4f, 0xdb, 0xed, 0x4d, 0xe7, 0x26, 0x72, 0xec, 0xc3, 0xac,
	0x3b, 0xce, 0xea, 0xfd, 0x13, 0xac, 0x95, 0xef, 0xe9, 0x15, 0x9d, 0x99, 0x85, 0x00, 0x18, 0x7c,
	0x4f, 0xd6, 0x3b, 0xdd, 0x99, 0x4d, 0xb1, 0x2b, 0xca, 0x95, 0xb8, 0x86, 0xf3, 0x36, 0x2e, 0x75,
	0xeb, 0x7f, 0x1c, 0x6f, 0x8e, 0x77, 0x45, 0x79, 0xba, 0xff, 0x80, 0x8b, 0x85, 0x4a, 0xa6, 0x66,
	0xb2, 0xe2, 0x1e, 0x4e, 0x22, 0x33, 0x59, 0x2e, 0x1f, 0xee, 0xe4, 0x9f, 0x0a, 0x75, 0x4e, 0x95,
	0x51, 0x99, 0x9e, 0xda, 0xc4, 0xa4, 0x68, 0xf8, 0x07, 0xad, 0x40, 0x64, 0xe8, 0xf3, 0x84, 0x8c,
	0x21, 0xa1, 0xb7, 0x20, 0x30, 0xfd, 0x34, 0x93, 0x1d, 0xbe, 0x91, 0x29, 0x77, 0x5b, 0x1f, 0xa2,
	0xbc, 0x9c, 0x7d, 0xae, 0xd2, 0x65, 0xa3, 0x3a, 0xfa, 0x5a, 0xa6, 0xc7, 0xdf, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xc1, 0xa3, 0x67, 0xc3, 0x26, 0x01, 0x00, 0x00,
}
