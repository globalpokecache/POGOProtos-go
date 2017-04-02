// Code generated by protoc-gen-go.
// source: inventory.proto
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

// Ignoring public import of AvatarCustomization from data_avatar.proto

// Ignoring public import of AvatarItem from data_avatar.proto

// Ignoring public import of Label from data_avatar.proto

// Ignoring public import of PokedexEntry from data.proto

// Ignoring public import of DownloadUrlEntry from data.proto

// Ignoring public import of PokemonDisplay from data.proto

// Ignoring public import of PlayerBadge from data.proto

// Ignoring public import of ClientVersion from data.proto

// Ignoring public import of PokemonData from data.proto

// Ignoring public import of PlayerData from data.proto

// Ignoring public import of BackgroundToken from data.proto

// Ignoring public import of AssetDigestEntry from data.proto

// Ignoring public import of BuddyPokemon from data.proto

// Ignoring public import of RedeemPasscodeReward from data.proto

// Ignoring public import of PlayerPublicProfile from data_player.proto

// Ignoring public import of PlayerCurrency from data_player.proto

// Ignoring public import of Currency from data_player.proto

// Ignoring public import of PlayerAvatar from data_player.proto

// Ignoring public import of ContactSettings from data_player.proto

// Ignoring public import of PlayerCamera from data_player.proto

// Ignoring public import of EquippedBadge from data_player.proto

// Ignoring public import of PlayerStats from data_player.proto

// Ignoring public import of DailyBonus from data_player.proto

// Ignoring public import of PlayerAvatarType from data_player.proto

// Ignoring public import of DailyQuest from data_quests.proto

// Ignoring public import of Quest from data_quests.proto

// Ignoring public import of PokemonType from enums.proto

// Ignoring public import of ItemCategory from enums.proto

// Ignoring public import of PokemonMovementType from enums.proto

// Ignoring public import of QuestType from enums.proto

// Ignoring public import of EncounterType from enums.proto

// Ignoring public import of Filter from enums.proto

// Ignoring public import of Platform from enums.proto

// Ignoring public import of TutorialState from enums.proto

// Ignoring public import of Costume from enums.proto

// Ignoring public import of ItemEffect from enums.proto

// Ignoring public import of BadgeType from enums.proto

// Ignoring public import of CameraTarget from enums.proto

// Ignoring public import of PokemonMove from enums.proto

// Ignoring public import of TeamColor from enums.proto

// Ignoring public import of Form from enums.proto

// Ignoring public import of PokemonId from enums.proto

// Ignoring public import of Gender from enums.proto

// Ignoring public import of ActivityType from enums.proto

// Ignoring public import of Slot from enums.proto

// Ignoring public import of HoloIapItemCategory from enums.proto

// Ignoring public import of CameraInterpolation from enums.proto

// Ignoring public import of PokemonRarity from enums.proto

// Ignoring public import of PokemonFamilyId from enums.proto

type EggIncubatorType int32

const (
	EggIncubatorType_INCUBATOR_UNSET    EggIncubatorType = 0
	EggIncubatorType_INCUBATOR_DISTANCE EggIncubatorType = 1
)

var EggIncubatorType_name = map[int32]string{
	0: "INCUBATOR_UNSET",
	1: "INCUBATOR_DISTANCE",
}
var EggIncubatorType_value = map[string]int32{
	"INCUBATOR_UNSET":    0,
	"INCUBATOR_DISTANCE": 1,
}

func (x EggIncubatorType) String() string {
	return proto.EnumName(EggIncubatorType_name, int32(x))
}
func (EggIncubatorType) EnumDescriptor() ([]byte, []int) { return fileDescriptor7, []int{0} }

type InventoryUpgradeType int32

const (
	InventoryUpgradeType_UPGRADE_UNSET            InventoryUpgradeType = 0
	InventoryUpgradeType_INCREASE_ITEM_STORAGE    InventoryUpgradeType = 1
	InventoryUpgradeType_INCREASE_POKEMON_STORAGE InventoryUpgradeType = 2
)

var InventoryUpgradeType_name = map[int32]string{
	0: "UPGRADE_UNSET",
	1: "INCREASE_ITEM_STORAGE",
	2: "INCREASE_POKEMON_STORAGE",
}
var InventoryUpgradeType_value = map[string]int32{
	"UPGRADE_UNSET":            0,
	"INCREASE_ITEM_STORAGE":    1,
	"INCREASE_POKEMON_STORAGE": 2,
}

func (x InventoryUpgradeType) String() string {
	return proto.EnumName(InventoryUpgradeType_name, int32(x))
}
func (InventoryUpgradeType) EnumDescriptor() ([]byte, []int) { return fileDescriptor7, []int{1} }

type EggIncubator struct {
	Id             string           `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	ItemId         ItemId           `protobuf:"varint,2,opt,name=item_id,enum=POGOProtos.Inventory.Item.ItemId" json:"item_id,omitempty"`
	IncubatorType  EggIncubatorType `protobuf:"varint,3,opt,name=incubator_type,enum=POGOProtos.Inventory.EggIncubatorType" json:"incubator_type,omitempty"`
	UsesRemaining  int32            `protobuf:"varint,4,opt,name=uses_remaining" json:"uses_remaining,omitempty"`
	PokemonId      uint64           `protobuf:"varint,5,opt,name=pokemon_id" json:"pokemon_id,omitempty"`
	StartKmWalked  float64          `protobuf:"fixed64,6,opt,name=start_km_walked" json:"start_km_walked,omitempty"`
	TargetKmWalked float64          `protobuf:"fixed64,7,opt,name=target_km_walked" json:"target_km_walked,omitempty"`
}

func (m *EggIncubator) Reset()                    { *m = EggIncubator{} }
func (m *EggIncubator) String() string            { return proto.CompactTextString(m) }
func (*EggIncubator) ProtoMessage()               {}
func (*EggIncubator) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{0} }

type InventoryUpgrade struct {
	ItemId            ItemId               `protobuf:"varint,1,opt,name=item_id,enum=POGOProtos.Inventory.Item.ItemId" json:"item_id,omitempty"`
	UpgradeType       InventoryUpgradeType `protobuf:"varint,2,opt,name=upgrade_type,enum=POGOProtos.Inventory.InventoryUpgradeType" json:"upgrade_type,omitempty"`
	AdditionalStorage int32                `protobuf:"varint,3,opt,name=additional_storage" json:"additional_storage,omitempty"`
}

func (m *InventoryUpgrade) Reset()                    { *m = InventoryUpgrade{} }
func (m *InventoryUpgrade) String() string            { return proto.CompactTextString(m) }
func (*InventoryUpgrade) ProtoMessage()               {}
func (*InventoryUpgrade) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{1} }

type InventoryUpgrades struct {
	InventoryUpgrades []*InventoryUpgrade `protobuf:"bytes,1,rep,name=inventory_upgrades" json:"inventory_upgrades,omitempty"`
}

func (m *InventoryUpgrades) Reset()                    { *m = InventoryUpgrades{} }
func (m *InventoryUpgrades) String() string            { return proto.CompactTextString(m) }
func (*InventoryUpgrades) ProtoMessage()               {}
func (*InventoryUpgrades) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{2} }

func (m *InventoryUpgrades) GetInventoryUpgrades() []*InventoryUpgrade {
	if m != nil {
		return m.InventoryUpgrades
	}
	return nil
}

type EggIncubators struct {
	EggIncubator []*EggIncubator `protobuf:"bytes,1,rep,name=egg_incubator" json:"egg_incubator,omitempty"`
}

func (m *EggIncubators) Reset()                    { *m = EggIncubators{} }
func (m *EggIncubators) String() string            { return proto.CompactTextString(m) }
func (*EggIncubators) ProtoMessage()               {}
func (*EggIncubators) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{3} }

func (m *EggIncubators) GetEggIncubator() []*EggIncubator {
	if m != nil {
		return m.EggIncubator
	}
	return nil
}

type InventoryDelta struct {
	OriginalTimestampMs int64            `protobuf:"varint,1,opt,name=original_timestamp_ms" json:"original_timestamp_ms,omitempty"`
	NewTimestampMs      int64            `protobuf:"varint,2,opt,name=new_timestamp_ms" json:"new_timestamp_ms,omitempty"`
	InventoryItems      []*InventoryItem `protobuf:"bytes,3,rep,name=inventory_items" json:"inventory_items,omitempty"`
}

func (m *InventoryDelta) Reset()                    { *m = InventoryDelta{} }
func (m *InventoryDelta) String() string            { return proto.CompactTextString(m) }
func (*InventoryDelta) ProtoMessage()               {}
func (*InventoryDelta) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{4} }

func (m *InventoryDelta) GetInventoryItems() []*InventoryItem {
	if m != nil {
		return m.InventoryItems
	}
	return nil
}

type InventoryItemData struct {
	PokemonData       *PokemonData       `protobuf:"bytes,1,opt,name=pokemon_data" json:"pokemon_data,omitempty"`
	Item              *ItemData          `protobuf:"bytes,2,opt,name=item" json:"item,omitempty"`
	PokedexEntry      *PokedexEntry      `protobuf:"bytes,3,opt,name=pokedex_entry" json:"pokedex_entry,omitempty"`
	PlayerStats       *PlayerStats       `protobuf:"bytes,4,opt,name=player_stats" json:"player_stats,omitempty"`
	PlayerCurrency    *PlayerCurrency    `protobuf:"bytes,5,opt,name=player_currency" json:"player_currency,omitempty"`
	PlayerCamera      *PlayerCamera      `protobuf:"bytes,6,opt,name=player_camera" json:"player_camera,omitempty"`
	InventoryUpgrades *InventoryUpgrades `protobuf:"bytes,7,opt,name=inventory_upgrades" json:"inventory_upgrades,omitempty"`
	AppliedItems      *AppliedItems      `protobuf:"bytes,8,opt,name=applied_items" json:"applied_items,omitempty"`
	EggIncubators     *EggIncubators     `protobuf:"bytes,9,opt,name=egg_incubators" json:"egg_incubators,omitempty"`
	Candy             *Candy             `protobuf:"bytes,10,opt,name=candy" json:"candy,omitempty"`
	Quest             *Quest             `protobuf:"bytes,11,opt,name=quest" json:"quest,omitempty"`
	AvatarItem        *AvatarItem        `protobuf:"bytes,12,opt,name=avatar_item" json:"avatar_item,omitempty"`
}

func (m *InventoryItemData) Reset()                    { *m = InventoryItemData{} }
func (m *InventoryItemData) String() string            { return proto.CompactTextString(m) }
func (*InventoryItemData) ProtoMessage()               {}
func (*InventoryItemData) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{5} }

func (m *InventoryItemData) GetPokemonData() *PokemonData {
	if m != nil {
		return m.PokemonData
	}
	return nil
}

func (m *InventoryItemData) GetItem() *ItemData {
	if m != nil {
		return m.Item
	}
	return nil
}

func (m *InventoryItemData) GetPokedexEntry() *PokedexEntry {
	if m != nil {
		return m.PokedexEntry
	}
	return nil
}

func (m *InventoryItemData) GetPlayerStats() *PlayerStats {
	if m != nil {
		return m.PlayerStats
	}
	return nil
}

func (m *InventoryItemData) GetPlayerCurrency() *PlayerCurrency {
	if m != nil {
		return m.PlayerCurrency
	}
	return nil
}

func (m *InventoryItemData) GetPlayerCamera() *PlayerCamera {
	if m != nil {
		return m.PlayerCamera
	}
	return nil
}

func (m *InventoryItemData) GetInventoryUpgrades() *InventoryUpgrades {
	if m != nil {
		return m.InventoryUpgrades
	}
	return nil
}

func (m *InventoryItemData) GetAppliedItems() *AppliedItems {
	if m != nil {
		return m.AppliedItems
	}
	return nil
}

func (m *InventoryItemData) GetEggIncubators() *EggIncubators {
	if m != nil {
		return m.EggIncubators
	}
	return nil
}

func (m *InventoryItemData) GetCandy() *Candy {
	if m != nil {
		return m.Candy
	}
	return nil
}

func (m *InventoryItemData) GetQuest() *Quest {
	if m != nil {
		return m.Quest
	}
	return nil
}

func (m *InventoryItemData) GetAvatarItem() *AvatarItem {
	if m != nil {
		return m.AvatarItem
	}
	return nil
}

type InventoryItem struct {
	ModifiedTimestampMs int64                      `protobuf:"varint,1,opt,name=modified_timestamp_ms" json:"modified_timestamp_ms,omitempty"`
	DeletedItem         *InventoryItem_DeletedItem `protobuf:"bytes,2,opt,name=deleted_item" json:"deleted_item,omitempty"`
	InventoryItemData   *InventoryItemData         `protobuf:"bytes,3,opt,name=inventory_item_data" json:"inventory_item_data,omitempty"`
}

func (m *InventoryItem) Reset()                    { *m = InventoryItem{} }
func (m *InventoryItem) String() string            { return proto.CompactTextString(m) }
func (*InventoryItem) ProtoMessage()               {}
func (*InventoryItem) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{6} }

func (m *InventoryItem) GetDeletedItem() *InventoryItem_DeletedItem {
	if m != nil {
		return m.DeletedItem
	}
	return nil
}

func (m *InventoryItem) GetInventoryItemData() *InventoryItemData {
	if m != nil {
		return m.InventoryItemData
	}
	return nil
}

type InventoryItem_DeletedItem struct {
	PokemonId uint64 `protobuf:"fixed64,1,opt,name=pokemon_id" json:"pokemon_id,omitempty"`
}

func (m *InventoryItem_DeletedItem) Reset()                    { *m = InventoryItem_DeletedItem{} }
func (m *InventoryItem_DeletedItem) String() string            { return proto.CompactTextString(m) }
func (*InventoryItem_DeletedItem) ProtoMessage()               {}
func (*InventoryItem_DeletedItem) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{6, 0} }

type InventoryKey struct {
	PokemonId         uint64          `protobuf:"fixed64,1,opt,name=pokemon_id" json:"pokemon_id,omitempty"`
	Item              ItemId          `protobuf:"varint,2,opt,name=item,enum=POGOProtos.Inventory.Item.ItemId" json:"item,omitempty"`
	PokedexEntryId    int32           `protobuf:"varint,3,opt,name=pokedex_entry_id" json:"pokedex_entry_id,omitempty"`
	PlayerStats       bool            `protobuf:"varint,4,opt,name=player_stats" json:"player_stats,omitempty"`
	PlayerCurrency    bool            `protobuf:"varint,5,opt,name=player_currency" json:"player_currency,omitempty"`
	PlayerCamera      bool            `protobuf:"varint,6,opt,name=player_camera" json:"player_camera,omitempty"`
	InventoryUpgrades bool            `protobuf:"varint,7,opt,name=inventory_upgrades" json:"inventory_upgrades,omitempty"`
	AppliedItems      bool            `protobuf:"varint,8,opt,name=applied_items" json:"applied_items,omitempty"`
	EggIncubators     bool            `protobuf:"varint,9,opt,name=egg_incubators" json:"egg_incubators,omitempty"`
	PokemonFamilyId   PokemonFamilyId `protobuf:"varint,10,opt,name=pokemon_family_id,enum=POGOProtos.Enums.PokemonFamilyId" json:"pokemon_family_id,omitempty"`
	QuestType         QuestType       `protobuf:"varint,11,opt,name=quest_type,enum=POGOProtos.Enums.QuestType" json:"quest_type,omitempty"`
	AvatarTemplateId  string          `protobuf:"bytes,12,opt,name=avatar_template_id" json:"avatar_template_id,omitempty"`
}

func (m *InventoryKey) Reset()                    { *m = InventoryKey{} }
func (m *InventoryKey) String() string            { return proto.CompactTextString(m) }
func (*InventoryKey) ProtoMessage()               {}
func (*InventoryKey) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{7} }

type AppliedItem struct {
	ItemId    ItemId   `protobuf:"varint,1,opt,name=item_id,enum=POGOProtos.Inventory.Item.ItemId" json:"item_id,omitempty"`
	ItemType  ItemType `protobuf:"varint,2,opt,name=item_type,enum=POGOProtos.Inventory.Item.ItemType" json:"item_type,omitempty"`
	ExpireMs  int64    `protobuf:"varint,3,opt,name=expire_ms" json:"expire_ms,omitempty"`
	AppliedMs int64    `protobuf:"varint,4,opt,name=applied_ms" json:"applied_ms,omitempty"`
}

func (m *AppliedItem) Reset()                    { *m = AppliedItem{} }
func (m *AppliedItem) String() string            { return proto.CompactTextString(m) }
func (*AppliedItem) ProtoMessage()               {}
func (*AppliedItem) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{8} }

type AppliedItems struct {
	Item []*AppliedItem `protobuf:"bytes,4,rep,name=item" json:"item,omitempty"`
}

func (m *AppliedItems) Reset()                    { *m = AppliedItems{} }
func (m *AppliedItems) String() string            { return proto.CompactTextString(m) }
func (*AppliedItems) ProtoMessage()               {}
func (*AppliedItems) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{9} }

func (m *AppliedItems) GetItem() []*AppliedItem {
	if m != nil {
		return m.Item
	}
	return nil
}

type Candy struct {
	FamilyId PokemonFamilyId `protobuf:"varint,1,opt,name=family_id,enum=POGOProtos.Enums.PokemonFamilyId" json:"family_id,omitempty"`
	Candy    int32           `protobuf:"varint,2,opt,name=candy" json:"candy,omitempty"`
}

func (m *Candy) Reset()                    { *m = Candy{} }
func (m *Candy) String() string            { return proto.CompactTextString(m) }
func (*Candy) ProtoMessage()               {}
func (*Candy) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{10} }

func init() {
	proto.RegisterType((*EggIncubator)(nil), "POGOProtos.Inventory.EggIncubator")
	proto.RegisterType((*InventoryUpgrade)(nil), "POGOProtos.Inventory.InventoryUpgrade")
	proto.RegisterType((*InventoryUpgrades)(nil), "POGOProtos.Inventory.InventoryUpgrades")
	proto.RegisterType((*EggIncubators)(nil), "POGOProtos.Inventory.EggIncubators")
	proto.RegisterType((*InventoryDelta)(nil), "POGOProtos.Inventory.InventoryDelta")
	proto.RegisterType((*InventoryItemData)(nil), "POGOProtos.Inventory.InventoryItemData")
	proto.RegisterType((*InventoryItem)(nil), "POGOProtos.Inventory.InventoryItem")
	proto.RegisterType((*InventoryItem_DeletedItem)(nil), "POGOProtos.Inventory.InventoryItem.DeletedItem")
	proto.RegisterType((*InventoryKey)(nil), "POGOProtos.Inventory.InventoryKey")
	proto.RegisterType((*AppliedItem)(nil), "POGOProtos.Inventory.AppliedItem")
	proto.RegisterType((*AppliedItems)(nil), "POGOProtos.Inventory.AppliedItems")
	proto.RegisterType((*Candy)(nil), "POGOProtos.Inventory.Candy")
	proto.RegisterEnum("POGOProtos.Inventory.EggIncubatorType", EggIncubatorType_name, EggIncubatorType_value)
	proto.RegisterEnum("POGOProtos.Inventory.InventoryUpgradeType", InventoryUpgradeType_name, InventoryUpgradeType_value)
}

func init() { proto.RegisterFile("inventory.proto", fileDescriptor7) }

var fileDescriptor7 = []byte{
	// 1018 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x56, 0xef, 0x6e, 0xdb, 0x54,
	0x14, 0x9f, 0x9b, 0xa6, 0x4d, 0x8e, 0x93, 0x36, 0xbd, 0x6b, 0x87, 0xe9, 0x36, 0x29, 0x35, 0x08,
	0xa2, 0x0a, 0xa5, 0x22, 0x4c, 0xa0, 0x69, 0x13, 0x25, 0x4d, 0x4c, 0x15, 0xc6, 0x12, 0xd3, 0xa4,
	0x42, 0xe2, 0x8b, 0x75, 0x5b, 0xdf, 0x45, 0x56, 0xe3, 0x3f, 0xd8, 0x37, 0x5b, 0xf3, 0x04, 0x88,
	0xb7, 0xe0, 0x03, 0x2f, 0xc1, 0xcb, 0xf0, 0x06, 0xbc, 0x03, 0xba, 0xe7, 0xda, 0xa9, 0x9d, 0x38,
	0x24, 0xe2, 0xcb, 0xbc, 0xde, 0x7b, 0x7e, 0xf7, 0xfc, 0xfb, 0xfd, 0xce, 0x09, 0xec, 0x3b, 0xde,
	0x7b, 0xe6, 0x71, 0x3f, 0x9c, 0x35, 0x83, 0xd0, 0xe7, 0x3e, 0x39, 0x34, 0x07, 0x97, 0x03, 0x53,
	0xfc, 0x37, 0x6a, 0xf6, 0x92, 0xbb, 0xe3, 0xc3, 0xb9, 0x99, 0xe5, 0x70, 0xe6, 0x4a, 0xdb, 0xe3,
	0x03, 0x9b, 0x72, 0x6a, 0xd1, 0xf7, 0x94, 0xd3, 0x30, 0x3e, 0x02, 0x71, 0x94, 0xb9, 0x0e, 0x26,
	0x74, 0xc6, 0xc2, 0xcc, 0xd1, 0xaf, 0x53, 0x16, 0xf1, 0x28, 0x3e, 0x52, 0x99, 0x37, 0x75, 0xe3,
	0x3f, 0xf4, 0x7f, 0x14, 0xa8, 0x18, 0xe3, 0x71, 0xcf, 0xbb, 0x9d, 0xde, 0x50, 0xee, 0x87, 0x04,
	0x60, 0xcb, 0xb1, 0x35, 0xa5, 0xae, 0x34, 0xca, 0xa4, 0x05, 0xbb, 0xc2, 0xb9, 0xe5, 0xd8, 0xda,
	0x56, 0x5d, 0x69, 0xec, 0xb5, 0x4e, 0x9a, 0x79, 0xc1, 0x36, 0x7b, 0x22, 0x42, 0xf1, 0x4f, 0xcf,
	0x26, 0xdf, 0xc2, 0x9e, 0x93, 0x3c, 0x66, 0xf1, 0x59, 0xc0, 0xb4, 0x02, 0x42, 0x3f, 0xcb, 0x87,
	0xa6, 0x7d, 0x8f, 0x66, 0x01, 0x23, 0x4f, 0x60, 0x6f, 0x1a, 0xb1, 0xc8, 0x0a, 0x99, 0x4b, 0x1d,
	0xcf, 0xf1, 0xc6, 0xda, 0x76, 0x5d, 0x69, 0x14, 0x09, 0x01, 0x08, 0xfc, 0x3b, 0xe6, 0xfa, 0x9e,
	0x08, 0xa7, 0x58, 0x57, 0x1a, 0xdb, 0xe4, 0x23, 0xd8, 0x8f, 0x38, 0x0d, 0xb9, 0x75, 0xe7, 0x5a,
	0x1f, 0xe8, 0xe4, 0x8e, 0xd9, 0xda, 0x4e, 0x5d, 0x69, 0x28, 0x44, 0x83, 0x1a, 0xa7, 0xe1, 0x98,
	0xa5, 0x6f, 0x76, 0xc5, 0x8d, 0xfe, 0xa7, 0x02, 0xb5, 0xb9, 0xf7, 0xeb, 0x60, 0x1c, 0x52, 0x9b,
	0xa5, 0xf3, 0x54, 0x36, 0xcd, 0xf3, 0x3b, 0xa8, 0x4c, 0x25, 0x5c, 0x66, 0x29, 0x0b, 0x74, 0xba,
	0x02, 0xb8, 0xe0, 0x11, 0x33, 0x3d, 0x06, 0x42, 0x6d, 0xdb, 0xe1, 0x8e, 0xef, 0xd1, 0x89, 0x15,
	0x71, 0x3f, 0xa4, 0x63, 0x59, 0xad, 0xa2, 0xfe, 0x33, 0x1c, 0x2c, 0x62, 0x22, 0x72, 0x01, 0xe4,
	0x81, 0x15, 0xb1, 0xf3, 0x48, 0x53, 0xea, 0x85, 0x86, 0xba, 0xaa, 0xbc, 0x8b, 0x8f, 0xe8, 0x3f,
	0x40, 0x35, 0x5d, 0xf2, 0x88, 0xbc, 0x84, 0x2a, 0x1b, 0x8f, 0xad, 0x79, 0xcf, 0xe2, 0xf7, 0xf4,
	0xf5, 0xed, 0xd2, 0x7f, 0x53, 0x60, 0x6f, 0x7e, 0xd5, 0x65, 0x13, 0x4e, 0xc9, 0x73, 0x38, 0xf2,
	0x43, 0x67, 0xec, 0x88, 0x8c, 0xb8, 0xe3, 0xb2, 0x88, 0x53, 0x37, 0xb0, 0xdc, 0x08, 0xeb, 0x5a,
	0x10, 0x7d, 0xf1, 0xd8, 0x87, 0xec, 0xcd, 0x16, 0xde, 0xbc, 0x4e, 0x09, 0x03, 0x19, 0x1f, 0x69,
	0x05, 0x0c, 0xe4, 0x93, 0x35, 0x89, 0x89, 0x76, 0xe8, 0x7f, 0x15, 0x53, 0xf5, 0x12, 0x27, 0x5d,
	0xca, 0x29, 0x69, 0x41, 0x25, 0xa1, 0x8c, 0x50, 0x01, 0xc6, 0xa0, 0xb6, 0x9e, 0xa5, 0x1f, 0x14,
	0x76, 0x4d, 0x53, 0x1a, 0x21, 0xe6, 0x4b, 0xd8, 0x16, 0xde, 0x31, 0xaa, 0xd5, 0xce, 0x13, 0x1e,
	0x20, 0xe4, 0x05, 0x54, 0x85, 0x1b, 0x9b, 0xdd, 0x5b, 0xcc, 0xe3, 0xe1, 0x0c, 0x5b, 0xa8, 0xb6,
	0x9e, 0xe7, 0xfa, 0xb1, 0xd9, 0xbd, 0x21, 0x8c, 0xc8, 0x4b, 0xa8, 0x48, 0xa1, 0x5a, 0x11, 0xa7,
	0x3c, 0x42, 0x96, 0x2f, 0x38, 0x94, 0x20, 0xa9, 0x66, 0xf9, 0x19, 0x0a, 0x53, 0x72, 0x0e, 0xfb,
	0x31, 0xf4, 0x76, 0x1a, 0x86, 0xcc, 0xbb, 0x9d, 0xa1, 0x1e, 0x16, 0x48, 0xb0, 0x8c, 0xee, 0xc4,
	0xd6, 0xe4, 0x15, 0x54, 0x93, 0x07, 0xa8, 0xcb, 0x42, 0x8a, 0xaa, 0x51, 0x5b, 0x9f, 0xae, 0x81,
	0xa3, 0x2d, 0xe9, 0xe4, 0xb2, 0x70, 0x17, 0x5f, 0xf8, 0x7c, 0x33, 0x16, 0x22, 0xeb, 0x68, 0x10,
	0x4c, 0x1c, 0x66, 0xc7, 0xcd, 0x2e, 0x21, 0x7e, 0x05, 0xeb, 0xda, 0xd2, 0x54, 0x54, 0x3c, 0x22,
	0xaf, 0x60, 0x2f, 0x43, 0xd8, 0x48, 0x2b, 0xff, 0x57, 0xaf, 0xb2, 0x6c, 0x3f, 0x85, 0xe2, 0x2d,
	0xf5, 0xec, 0x99, 0x06, 0x88, 0x79, 0x9a, 0x8f, 0xe9, 0x08, 0x13, 0xf2, 0x05, 0x14, 0x71, 0x6e,
	0x6a, 0xea, 0x8a, 0x7e, 0xfe, 0x24, 0xa7, 0x2a, 0x7e, 0xc8, 0x37, 0xa0, 0xca, 0xb9, 0x8c, 0x09,
	0x69, 0x95, 0xe5, 0x7c, 0x10, 0xd3, 0x96, 0xb3, 0x5b, 0x7e, 0x90, 0xbb, 0x7f, 0x2b, 0x50, 0xcd,
	0x70, 0x57, 0x88, 0xc8, 0xf5, 0x6d, 0xe7, 0x9d, 0xa8, 0x4e, 0x8e, 0x88, 0x0c, 0xa8, 0xd8, 0x6c,
	0xc2, 0x78, 0x5c, 0xbb, 0x98, 0xaa, 0x67, 0x1b, 0xe8, 0xa4, 0xd9, 0x95, 0x38, 0xf4, 0xd2, 0x85,
	0xc7, 0x59, 0xc5, 0x49, 0x91, 0x14, 0x36, 0x6a, 0x64, 0x42, 0xfe, 0xe3, 0x13, 0x50, 0xd3, 0x8f,
	0x66, 0xa7, 0xb4, 0x88, 0x77, 0x47, 0xff, 0xbd, 0x00, 0x95, 0x39, 0xf0, 0x0d, 0x9b, 0xe5, 0x19,
	0x91, 0xb3, 0x94, 0xee, 0x36, 0x9a, 0xbf, 0x1a, 0xd4, 0x32, 0xaa, 0x13, 0x4f, 0xe1, 0xec, 0x24,
	0x87, 0x39, 0xca, 0x2a, 0x89, 0x5d, 0x91, 0x27, 0x9a, 0x12, 0x39, 0xca, 0x13, 0x43, 0x49, 0x4c,
	0xe7, 0x15, 0x34, 0x47, 0xc8, 0x32, 0x7b, 0x4b, 0x62, 0x75, 0xe5, 0x30, 0xb3, 0x44, 0x5e, 0xc3,
	0x41, 0x92, 0xef, 0x3b, 0xea, 0x3a, 0x13, 0x8c, 0x15, 0x96, 0x13, 0x35, 0x70, 0x2f, 0xc7, 0xd3,
	0xe8, 0x7b, 0xb4, 0xec, 0xd9, 0xe4, 0x0c, 0x00, 0x69, 0x28, 0xd7, 0x8c, 0x8a, 0xb0, 0xa7, 0xcb,
	0x30, 0x64, 0xe1, 0x7c, 0xaf, 0x48, 0x26, 0x72, 0xe6, 0x06, 0x13, 0xca, 0x99, 0xf0, 0x27, 0x08,
	0x59, 0xd6, 0xff, 0x50, 0x40, 0x4d, 0xa9, 0xe9, 0x7f, 0x6d, 0xbe, 0xaf, 0xa1, 0x8c, 0x98, 0xd4,
	0xda, 0x5b, 0x37, 0x27, 0x31, 0xae, 0x03, 0x28, 0xb3, 0xfb, 0xc0, 0x09, 0x99, 0x85, 0xc3, 0x5d,
	0x50, 0x99, 0x00, 0x24, 0x85, 0x74, 0x65, 0xa3, 0x0a, 0xfa, 0x39, 0x54, 0x32, 0x7a, 0x4f, 0x98,
	0xb1, 0x8d, 0xeb, 0xe0, 0x64, 0xed, 0x84, 0xd0, 0x7f, 0x84, 0xa2, 0x14, 0xf0, 0x0b, 0x28, 0x3f,
	0xd4, 0x5b, 0xd9, 0xb4, 0xde, 0xd5, 0x64, 0x44, 0x88, 0xd4, 0x8a, 0xa7, 0xe7, 0x50, 0x5b, 0xfa,
	0x8d, 0xf2, 0x18, 0xf6, 0x7b, 0xfd, 0xce, 0xf5, 0x45, 0x7b, 0x34, 0xb8, 0xb2, 0xae, 0xfb, 0x43,
	0x63, 0x54, 0x7b, 0x44, 0x9e, 0x00, 0x79, 0x38, 0xec, 0xf6, 0x86, 0xa3, 0x76, 0xbf, 0x63, 0xd4,
	0x94, 0xd3, 0x1b, 0x38, 0xcc, 0x5d, 0xff, 0x07, 0x50, 0xbd, 0x36, 0x2f, 0xaf, 0xda, 0x5d, 0x63,
	0xfe, 0xc4, 0xc7, 0x70, 0xd4, 0xeb, 0x77, 0xae, 0x8c, 0xf6, 0xd0, 0xb0, 0x7a, 0x23, 0xe3, 0xad,
	0x35, 0x1c, 0x0d, 0xae, 0xda, 0x97, 0x46, 0x4d, 0x21, 0xcf, 0x40, 0x9b, 0x5f, 0x99, 0x83, 0x37,
	0xc6, 0xdb, 0x41, 0x7f, 0x7e, 0xbb, 0x75, 0x51, 0xfa, 0x65, 0x07, 0x7f, 0xce, 0x45, 0xe6, 0x23,
	0x53, 0x31, 0xb7, 0xcc, 0x82, 0xb9, 0x6d, 0x16, 0x6f, 0xe4, 0xd9, 0x57, 0xff, 0x06, 0x00, 0x00,
	0xff, 0xff, 0x5a, 0xa3, 0x8c, 0x68, 0x73, 0x0a, 0x00, 0x00,
}
