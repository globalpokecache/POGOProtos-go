// Code generated by protoc-gen-go.
// source: data.proto
// DO NOT EDIT!

package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

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

// Ignoring public import of ItemAward from inventory_item.proto

// Ignoring public import of ItemData from inventory_item.proto

// Ignoring public import of ItemType from inventory_item.proto

// Ignoring public import of ItemId from inventory_item.proto

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

type PokedexEntry struct {
	PokemonId            PokemonId `protobuf:"varint,1,opt,name=pokemon_id,enum=POGOProtos.Enums.PokemonId" json:"pokemon_id,omitempty"`
	TimesEncountered     int32     `protobuf:"varint,2,opt,name=times_encountered" json:"times_encountered,omitempty"`
	TimesCaptured        int32     `protobuf:"varint,3,opt,name=times_captured" json:"times_captured,omitempty"`
	EvolutionStonePieces int32     `protobuf:"varint,4,opt,name=evolution_stone_pieces" json:"evolution_stone_pieces,omitempty"`
	EvolutionStones      int32     `protobuf:"varint,5,opt,name=evolution_stones" json:"evolution_stones,omitempty"`
	CapturedCostumes     []Costume `protobuf:"varint,6,rep,packed,name=captured_costumes,enum=POGOProtos.Enums.Costume" json:"captured_costumes,omitempty"`
	CapturedForms        []Form    `protobuf:"varint,7,rep,packed,name=captured_forms,enum=POGOProtos.Enums.Form" json:"captured_forms,omitempty"`
	CapturedGenders      []Gender  `protobuf:"varint,8,rep,packed,name=captured_genders,enum=POGOProtos.Enums.Gender" json:"captured_genders,omitempty"`
	CapturedShiny        bool      `protobuf:"varint,9,opt,name=captured_shiny" json:"captured_shiny,omitempty"`
	EncounteredCostumes  []Costume `protobuf:"varint,10,rep,packed,name=encountered_costumes,enum=POGOProtos.Enums.Costume" json:"encountered_costumes,omitempty"`
	EncounteredForms     []Form    `protobuf:"varint,11,rep,packed,name=encountered_forms,enum=POGOProtos.Enums.Form" json:"encountered_forms,omitempty"`
	EncounteredGenders   []Gender  `protobuf:"varint,12,rep,packed,name=encountered_genders,enum=POGOProtos.Enums.Gender" json:"encountered_genders,omitempty"`
	EncounteredShiny     bool      `protobuf:"varint,13,opt,name=encountered_shiny" json:"encountered_shiny,omitempty"`
}

func (m *PokedexEntry) Reset()                    { *m = PokedexEntry{} }
func (m *PokedexEntry) String() string            { return proto.CompactTextString(m) }
func (*PokedexEntry) ProtoMessage()               {}
func (*PokedexEntry) Descriptor() ([]byte, []int) { return fileDescriptor19, []int{0} }

type DownloadUrlEntry struct {
	AssetId  string `protobuf:"bytes,1,opt,name=asset_id" json:"asset_id,omitempty"`
	Url      string `protobuf:"bytes,2,opt,name=url" json:"url,omitempty"`
	Size     int32  `protobuf:"varint,3,opt,name=size" json:"size,omitempty"`
	Checksum uint32 `protobuf:"fixed32,4,opt,name=checksum" json:"checksum,omitempty"`
}

func (m *DownloadUrlEntry) Reset()                    { *m = DownloadUrlEntry{} }
func (m *DownloadUrlEntry) String() string            { return proto.CompactTextString(m) }
func (*DownloadUrlEntry) ProtoMessage()               {}
func (*DownloadUrlEntry) Descriptor() ([]byte, []int) { return fileDescriptor19, []int{1} }

type PokemonDisplay struct {
	Costume Costume `protobuf:"varint,1,opt,name=costume,enum=POGOProtos.Enums.Costume" json:"costume,omitempty"`
	Gender  Gender  `protobuf:"varint,2,opt,name=gender,enum=POGOProtos.Enums.Gender" json:"gender,omitempty"`
	Shiny   bool    `protobuf:"varint,3,opt,name=shiny" json:"shiny,omitempty"`
	Form    Form    `protobuf:"varint,4,opt,name=form,enum=POGOProtos.Enums.Form" json:"form,omitempty"`
}

func (m *PokemonDisplay) Reset()                    { *m = PokemonDisplay{} }
func (m *PokemonDisplay) String() string            { return proto.CompactTextString(m) }
func (*PokemonDisplay) ProtoMessage()               {}
func (*PokemonDisplay) Descriptor() ([]byte, []int) { return fileDescriptor19, []int{2} }

type PlayerBadge struct {
	BadgeType    BadgeType `protobuf:"varint,1,opt,name=badge_type,enum=POGOProtos.Enums.BadgeType" json:"badge_type,omitempty"`
	Rank         int32     `protobuf:"varint,2,opt,name=rank" json:"rank,omitempty"`
	StartValue   int32     `protobuf:"varint,3,opt,name=start_value" json:"start_value,omitempty"`
	EndValue     int32     `protobuf:"varint,4,opt,name=end_value" json:"end_value,omitempty"`
	CurrentValue float64   `protobuf:"fixed64,5,opt,name=current_value" json:"current_value,omitempty"`
}

func (m *PlayerBadge) Reset()                    { *m = PlayerBadge{} }
func (m *PlayerBadge) String() string            { return proto.CompactTextString(m) }
func (*PlayerBadge) ProtoMessage()               {}
func (*PlayerBadge) Descriptor() ([]byte, []int) { return fileDescriptor19, []int{3} }

type ClientVersion struct {
	MinVersion string `protobuf:"bytes,1,opt,name=min_version" json:"min_version,omitempty"`
}

func (m *ClientVersion) Reset()                    { *m = ClientVersion{} }
func (m *ClientVersion) String() string            { return proto.CompactTextString(m) }
func (*ClientVersion) ProtoMessage()               {}
func (*ClientVersion) Descriptor() ([]byte, []int) { return fileDescriptor19, []int{4} }

type PokemonData struct {
	Id                     uint64          `protobuf:"fixed64,1,opt,name=id" json:"id,omitempty"`
	PokemonId              PokemonId       `protobuf:"varint,2,opt,name=pokemon_id,enum=POGOProtos.Enums.PokemonId" json:"pokemon_id,omitempty"`
	Cp                     int32           `protobuf:"varint,3,opt,name=cp" json:"cp,omitempty"`
	Stamina                int32           `protobuf:"varint,4,opt,name=stamina" json:"stamina,omitempty"`
	StaminaMax             int32           `protobuf:"varint,5,opt,name=stamina_max" json:"stamina_max,omitempty"`
	Move_1                 PokemonMove     `protobuf:"varint,6,opt,name=move_1,enum=POGOProtos.Enums.PokemonMove" json:"move_1,omitempty"`
	Move_2                 PokemonMove     `protobuf:"varint,7,opt,name=move_2,enum=POGOProtos.Enums.PokemonMove" json:"move_2,omitempty"`
	DeployedFortId         string          `protobuf:"bytes,8,opt,name=deployed_fort_id" json:"deployed_fort_id,omitempty"`
	OwnerName              string          `protobuf:"bytes,9,opt,name=owner_name" json:"owner_name,omitempty"`
	IsEgg                  bool            `protobuf:"varint,10,opt,name=is_egg" json:"is_egg,omitempty"`
	EggKmWalkedTarget      float64         `protobuf:"fixed64,11,opt,name=egg_km_walked_target" json:"egg_km_walked_target,omitempty"`
	EggKmWalkedStart       float64         `protobuf:"fixed64,12,opt,name=egg_km_walked_start" json:"egg_km_walked_start,omitempty"`
	Origin                 int32           `protobuf:"varint,14,opt,name=origin" json:"origin,omitempty"`
	HeightM                float32         `protobuf:"fixed32,15,opt,name=height_m" json:"height_m,omitempty"`
	WeightKg               float32         `protobuf:"fixed32,16,opt,name=weight_kg" json:"weight_kg,omitempty"`
	IndividualAttack       int32           `protobuf:"varint,17,opt,name=individual_attack" json:"individual_attack,omitempty"`
	IndividualDefense      int32           `protobuf:"varint,18,opt,name=individual_defense" json:"individual_defense,omitempty"`
	IndividualStamina      int32           `protobuf:"varint,19,opt,name=individual_stamina" json:"individual_stamina,omitempty"`
	CpMultiplier           float32         `protobuf:"fixed32,20,opt,name=cp_multiplier" json:"cp_multiplier,omitempty"`
	Pokeball               ItemId          `protobuf:"varint,21,opt,name=pokeball,enum=POGOProtos.Inventory.Item.ItemId" json:"pokeball,omitempty"`
	CapturedCellId         uint64          `protobuf:"varint,22,opt,name=captured_cell_id" json:"captured_cell_id,omitempty"`
	BattlesAttacked        int32           `protobuf:"varint,23,opt,name=battles_attacked" json:"battles_attacked,omitempty"`
	BattlesDefended        int32           `protobuf:"varint,24,opt,name=battles_defended" json:"battles_defended,omitempty"`
	EggIncubatorId         string          `protobuf:"bytes,25,opt,name=egg_incubator_id" json:"egg_incubator_id,omitempty"`
	CreationTimeMs         uint64          `protobuf:"varint,26,opt,name=creation_time_ms" json:"creation_time_ms,omitempty"`
	NumUpgrades            int32           `protobuf:"varint,27,opt,name=num_upgrades" json:"num_upgrades,omitempty"`
	AdditionalCpMultiplier float32         `protobuf:"fixed32,28,opt,name=additional_cp_multiplier" json:"additional_cp_multiplier,omitempty"`
	Favorite               int32           `protobuf:"varint,29,opt,name=favorite" json:"favorite,omitempty"`
	Nickname               string          `protobuf:"bytes,30,opt,name=nickname" json:"nickname,omitempty"`
	FromFort               int32           `protobuf:"varint,31,opt,name=from_fort" json:"from_fort,omitempty"`
	BuddyCandyAwarded      int32           `protobuf:"varint,32,opt,name=buddy_candy_awarded" json:"buddy_candy_awarded,omitempty"`
	BuddyTotalKmWalked     float32         `protobuf:"fixed32,33,opt,name=buddy_total_km_walked" json:"buddy_total_km_walked,omitempty"`
	DisplayPokemonId       int32           `protobuf:"varint,34,opt,name=display_pokemon_id" json:"display_pokemon_id,omitempty"`
	DisplayCp              int32           `protobuf:"varint,35,opt,name=display_cp" json:"display_cp,omitempty"`
	PokemonDisplay         *PokemonDisplay `protobuf:"bytes,36,opt,name=pokemon_display" json:"pokemon_display,omitempty"`
}

func (m *PokemonData) Reset()                    { *m = PokemonData{} }
func (m *PokemonData) String() string            { return proto.CompactTextString(m) }
func (*PokemonData) ProtoMessage()               {}
func (*PokemonData) Descriptor() ([]byte, []int) { return fileDescriptor19, []int{5} }

func (m *PokemonData) GetPokemonDisplay() *PokemonDisplay {
	if m != nil {
		return m.PokemonDisplay
	}
	return nil
}

type PlayerData struct {
	CreationTimestampMs     int64            `protobuf:"varint,1,opt,name=creation_timestamp_ms" json:"creation_timestamp_ms,omitempty"`
	Username                string           `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
	Team                    TeamColor        `protobuf:"varint,5,opt,name=team,enum=POGOProtos.Enums.TeamColor" json:"team,omitempty"`
	TutorialState           []TutorialState  `protobuf:"varint,7,rep,packed,name=tutorial_state,enum=POGOProtos.Enums.TutorialState" json:"tutorial_state,omitempty"`
	Avatar                  *PlayerAvatar    `protobuf:"bytes,8,opt,name=avatar" json:"avatar,omitempty"`
	MaxPokemonStorage       int32            `protobuf:"varint,9,opt,name=max_pokemon_storage" json:"max_pokemon_storage,omitempty"`
	MaxItemStorage          int32            `protobuf:"varint,10,opt,name=max_item_storage" json:"max_item_storage,omitempty"`
	DailyBonus              *DailyBonus      `protobuf:"bytes,11,opt,name=daily_bonus" json:"daily_bonus,omitempty"`
	EquippedBadge           *EquippedBadge   `protobuf:"bytes,12,opt,name=equipped_badge" json:"equipped_badge,omitempty"`
	ContactSettings         *ContactSettings `protobuf:"bytes,13,opt,name=contact_settings" json:"contact_settings,omitempty"`
	Currencies              []*Currency      `protobuf:"bytes,14,rep,name=currencies" json:"currencies,omitempty"`
	RemainingCodenameClaims int32            `protobuf:"varint,15,opt,name=remaining_codename_claims" json:"remaining_codename_claims,omitempty"`
	BuddyPokemon            *BuddyPokemon    `protobuf:"bytes,16,opt,name=buddy_pokemon" json:"buddy_pokemon,omitempty"`
	BattleLockoutEndMs      int64            `protobuf:"varint,17,opt,name=battle_lockout_end_ms" json:"battle_lockout_end_ms,omitempty"`
	SecondaryPlayerAvatar   *PlayerAvatar    `protobuf:"bytes,18,opt,name=secondary_player_avatar" json:"secondary_player_avatar,omitempty"`
}

func (m *PlayerData) Reset()                    { *m = PlayerData{} }
func (m *PlayerData) String() string            { return proto.CompactTextString(m) }
func (*PlayerData) ProtoMessage()               {}
func (*PlayerData) Descriptor() ([]byte, []int) { return fileDescriptor19, []int{6} }

func (m *PlayerData) GetAvatar() *PlayerAvatar {
	if m != nil {
		return m.Avatar
	}
	return nil
}

func (m *PlayerData) GetDailyBonus() *DailyBonus {
	if m != nil {
		return m.DailyBonus
	}
	return nil
}

func (m *PlayerData) GetEquippedBadge() *EquippedBadge {
	if m != nil {
		return m.EquippedBadge
	}
	return nil
}

func (m *PlayerData) GetContactSettings() *ContactSettings {
	if m != nil {
		return m.ContactSettings
	}
	return nil
}

func (m *PlayerData) GetCurrencies() []*Currency {
	if m != nil {
		return m.Currencies
	}
	return nil
}

func (m *PlayerData) GetBuddyPokemon() *BuddyPokemon {
	if m != nil {
		return m.BuddyPokemon
	}
	return nil
}

func (m *PlayerData) GetSecondaryPlayerAvatar() *PlayerAvatar {
	if m != nil {
		return m.SecondaryPlayerAvatar
	}
	return nil
}

type BackgroundToken struct {
	Token          []byte `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	ExpirationTime int64  `protobuf:"varint,2,opt,name=expiration_time" json:"expiration_time,omitempty"`
	Iv             []byte `protobuf:"bytes,3,opt,name=iv,proto3" json:"iv,omitempty"`
}

func (m *BackgroundToken) Reset()                    { *m = BackgroundToken{} }
func (m *BackgroundToken) String() string            { return proto.CompactTextString(m) }
func (*BackgroundToken) ProtoMessage()               {}
func (*BackgroundToken) Descriptor() ([]byte, []int) { return fileDescriptor19, []int{7} }

type AssetDigestEntry struct {
	AssetId    string `protobuf:"bytes,1,opt,name=asset_id" json:"asset_id,omitempty"`
	BundleName string `protobuf:"bytes,2,opt,name=bundle_name" json:"bundle_name,omitempty"`
	Version    int64  `protobuf:"varint,3,opt,name=version" json:"version,omitempty"`
	Checksum   uint32 `protobuf:"fixed32,4,opt,name=checksum" json:"checksum,omitempty"`
	Size       int32  `protobuf:"varint,5,opt,name=size" json:"size,omitempty"`
	Key        []byte `protobuf:"bytes,6,opt,name=key,proto3" json:"key,omitempty"`
}

func (m *AssetDigestEntry) Reset()                    { *m = AssetDigestEntry{} }
func (m *AssetDigestEntry) String() string            { return proto.CompactTextString(m) }
func (*AssetDigestEntry) ProtoMessage()               {}
func (*AssetDigestEntry) Descriptor() ([]byte, []int) { return fileDescriptor19, []int{8} }

type BuddyPokemon struct {
	Id            uint64  `protobuf:"fixed64,1,opt,name=id" json:"id,omitempty"`
	StartKmWalked float64 `protobuf:"fixed64,2,opt,name=start_km_walked" json:"start_km_walked,omitempty"`
	LastKmAwarded float64 `protobuf:"fixed64,3,opt,name=last_km_awarded" json:"last_km_awarded,omitempty"`
}

func (m *BuddyPokemon) Reset()                    { *m = BuddyPokemon{} }
func (m *BuddyPokemon) String() string            { return proto.CompactTextString(m) }
func (*BuddyPokemon) ProtoMessage()               {}
func (*BuddyPokemon) Descriptor() ([]byte, []int) { return fileDescriptor19, []int{9} }

func init() {
	proto.RegisterType((*PokedexEntry)(nil), "POGOProtos.Data.PokedexEntry")
	proto.RegisterType((*DownloadUrlEntry)(nil), "POGOProtos.Data.DownloadUrlEntry")
	proto.RegisterType((*PokemonDisplay)(nil), "POGOProtos.Data.PokemonDisplay")
	proto.RegisterType((*PlayerBadge)(nil), "POGOProtos.Data.PlayerBadge")
	proto.RegisterType((*ClientVersion)(nil), "POGOProtos.Data.ClientVersion")
	proto.RegisterType((*PokemonData)(nil), "POGOProtos.Data.PokemonData")
	proto.RegisterType((*PlayerData)(nil), "POGOProtos.Data.PlayerData")
	proto.RegisterType((*BackgroundToken)(nil), "POGOProtos.Data.BackgroundToken")
	proto.RegisterType((*AssetDigestEntry)(nil), "POGOProtos.Data.AssetDigestEntry")
	proto.RegisterType((*BuddyPokemon)(nil), "POGOProtos.Data.BuddyPokemon")
}

func init() { proto.RegisterFile("data.proto", fileDescriptor19) }

var fileDescriptor19 = []byte{
	// 1387 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x56, 0x5f, 0x6f, 0x1b, 0xb9,
	0x11, 0xaf, 0x2c, 0x47, 0xb6, 0x67, 0x65, 0xfd, 0x59, 0x3b, 0x31, 0xed, 0x5c, 0x2e, 0x8a, 0x9a,
	0xa2, 0x6a, 0x81, 0xba, 0x38, 0x5d, 0x8a, 0x2b, 0x50, 0xf4, 0x21, 0x8e, 0xdd, 0x43, 0x80, 0x16,
	0xa7, 0xe2, 0xd2, 0x3e, 0xf4, 0x85, 0xa0, 0x76, 0x27, 0x1b, 0x42, 0xbb, 0xe4, 0x96, 0xe4, 0xca,
	0x51, 0xbf, 0x42, 0x3f, 0x42, 0xfb, 0x59, 0xfa, 0xd0, 0x4f, 0x76, 0xe0, 0x70, 0x57, 0x96, 0xe2,
	0x28, 0xc9, 0x8b, 0xbd, 0xe2, 0xcc, 0x70, 0xfe, 0xfd, 0x7e, 0x33, 0x04, 0x48, 0x85, 0x13, 0x97,
	0xa5, 0xd1, 0x4e, 0xc7, 0xfd, 0xd9, 0x0f, 0xdf, 0xff, 0x30, 0xf3, 0x9f, 0xf6, 0xf2, 0x5a, 0x38,
	0x71, 0x11, 0xa1, 0xaa, 0x0a, 0x1b, 0xa4, 0x17, 0xa7, 0x52, 0x2d, 0x51, 0x39, 0x6d, 0x56, 0x5c,
	0x3a, 0x2c, 0xea, 0xd3, 0xa1, 0xb7, 0xe7, 0x65, 0x2e, 0x56, 0x68, 0xc2, 0xd1, 0xf8, 0xbf, 0xfb,
	0xd0, 0x9d, 0xe9, 0x05, 0xa6, 0xf8, 0xfe, 0x46, 0x39, 0xb3, 0x8a, 0x7f, 0x0b, 0x50, 0xea, 0x05,
	0x16, 0x5a, 0x71, 0x99, 0xb2, 0xd6, 0xa8, 0x35, 0xe9, 0x4d, 0x1f, 0x5f, 0x6e, 0x38, 0xbb, 0x21,
	0x37, 0xb3, 0xa0, 0xf3, 0x3a, 0x8d, 0xcf, 0x61, 0xe8, 0x64, 0x81, 0x96, 0xa3, 0x4a, 0x74, 0xa5,
	0x1c, 0x1a, 0x4c, 0xd9, 0xde, 0xa8, 0x35, 0x79, 0x10, 0x3f, 0x82, 0x5e, 0x10, 0x25, 0xa2, 0x74,
	0x95, 0x3f, 0x6f, 0xd3, 0xf9, 0xd7, 0xf0, 0x08, 0x97, 0x3a, 0xaf, 0x9c, 0xd4, 0x8a, 0x5b, 0xa7,
	0x15, 0xf2, 0x52, 0x62, 0x82, 0x96, 0xed, 0x93, 0x9c, 0xc1, 0xe0, 0x03, 0xb9, 0x65, 0x0f, 0x48,
	0xf2, 0x02, 0x86, 0xcd, 0x5d, 0x3c, 0xd1, 0xd6, 0x55, 0x05, 0x5a, 0xd6, 0x19, 0xb5, 0x27, 0xbd,
	0xe9, 0xf9, 0xfd, 0x20, 0x5f, 0x05, 0x8d, 0xf8, 0x12, 0x7a, 0x6b, 0xab, 0xb7, 0xda, 0x14, 0x96,
	0x1d, 0x90, 0xc9, 0xa3, 0xfb, 0x26, 0x7f, 0xd2, 0xa6, 0x88, 0xa7, 0x30, 0x58, 0xeb, 0x67, 0xa8,
	0x52, 0x34, 0x96, 0x1d, 0x92, 0x05, 0xbb, 0x6f, 0xf1, 0x3d, 0x29, 0xf8, 0x5c, 0xd7, 0x36, 0xf6,
	0x9d, 0x54, 0x2b, 0x76, 0x34, 0x6a, 0x4d, 0x0e, 0xe3, 0xef, 0xe0, 0x74, 0xa3, 0x30, 0x77, 0x41,
	0xc3, 0xe7, 0x82, 0xfe, 0x06, 0x86, 0x9b, 0x86, 0x21, 0xee, 0xe8, 0x93, 0x71, 0xff, 0x0e, 0x4e,
	0x36, 0x4d, 0x9a, 0xd0, 0xbb, 0x9f, 0x09, 0xfd, 0x7c, 0xdb, 0x53, 0x88, 0xfe, 0xd8, 0x47, 0x3f,
	0xfe, 0x2b, 0x0c, 0xae, 0xf5, 0xad, 0xca, 0xb5, 0x48, 0xff, 0x66, 0xf2, 0x80, 0x90, 0x01, 0x1c,
	0x0a, 0x6b, 0xd1, 0x35, 0xf8, 0x38, 0x8a, 0x23, 0x68, 0x57, 0x26, 0xa7, 0xa6, 0x1f, 0xc5, 0x5d,
	0xd8, 0xb7, 0xf2, 0x5f, 0x58, 0xb7, 0x7a, 0x00, 0x87, 0xc9, 0x3b, 0x4c, 0x16, 0xb6, 0x2a, 0xa8,
	0xb9, 0x07, 0xe3, 0xff, 0xb4, 0xa0, 0x57, 0xa3, 0xe7, 0x5a, 0x5a, 0x0f, 0xc6, 0xf8, 0xd7, 0x70,
	0x50, 0xd7, 0xa5, 0x06, 0xdc, 0x27, 0xca, 0x32, 0x81, 0x4e, 0xc8, 0x8b, 0xdc, 0x7d, 0x2a, 0xad,
	0x63, 0x78, 0x10, 0x52, 0x69, 0x53, 0x23, 0x9e, 0xc3, 0xbe, 0xaf, 0x21, 0x45, 0xb1, 0xb3, 0x84,
	0xe3, 0x7f, 0xb7, 0x20, 0x9a, 0x11, 0x41, 0xae, 0x44, 0x9a, 0xa1, 0xa7, 0xc3, 0xdc, 0x7f, 0x70,
	0xb7, 0x2a, 0x71, 0x37, 0x1d, 0x48, 0xf9, 0xcd, 0xaa, 0x44, 0x9f, 0xbe, 0x11, 0x6a, 0x51, 0x33,
	0xe0, 0x04, 0x22, 0xeb, 0x84, 0x71, 0x7c, 0x29, 0xf2, 0xaa, 0xa9, 0xc9, 0x10, 0x8e, 0x50, 0xa5,
	0xf5, 0x51, 0x40, 0xfc, 0x43, 0x38, 0x4e, 0x2a, 0x63, 0x50, 0x35, 0x9a, 0x1e, 0xee, 0xad, 0xf1,
	0x73, 0x38, 0x7e, 0x95, 0x4b, 0x54, 0xee, 0xef, 0x68, 0xac, 0xd4, 0xca, 0xdf, 0x57, 0x48, 0xc5,
	0x97, 0xe1, 0x67, 0x28, 0xff, 0xf8, 0x7f, 0x07, 0x10, 0x35, 0x15, 0x15, 0x4e, 0xc4, 0x00, 0x7b,
	0x75, 0x6b, 0x3a, 0x1f, 0xd0, 0x79, 0xef, 0xf3, 0x74, 0x06, 0xd8, 0x4b, 0xca, 0x3a, 0xd0, 0x3e,
	0x1c, 0x58, 0x27, 0x0a, 0xa9, 0x44, 0x1d, 0x66, 0x48, 0xc7, 0x1f, 0xf0, 0x42, 0xbc, 0xaf, 0x39,
	0xf9, 0x1b, 0xe8, 0x14, 0x7a, 0x89, 0xfc, 0x1b, 0xd6, 0xa1, 0xeb, 0x9f, 0xec, 0xbc, 0xfe, 0x2f,
	0x7a, 0x89, 0x6b, 0xf5, 0x29, 0x3b, 0xf8, 0x12, 0x75, 0x06, 0x83, 0x14, 0xcb, 0x5c, 0xaf, 0x02,
	0x07, 0x08, 0x75, 0x87, 0x04, 0xb4, 0x18, 0x40, 0xdf, 0x2a, 0x34, 0x5c, 0x89, 0x02, 0x89, 0x6d,
	0x47, 0x71, 0x0f, 0x3a, 0xd2, 0x72, 0xcc, 0x32, 0x06, 0xd4, 0xf4, 0xaf, 0xe0, 0x14, 0xb3, 0x8c,
	0x2f, 0x0a, 0x7e, 0x2b, 0xf2, 0x05, 0xa6, 0xdc, 0x09, 0x93, 0xa1, 0x63, 0x91, 0x2f, 0x6f, 0xfc,
	0x18, 0x4e, 0xb6, 0xa5, 0xd4, 0x2b, 0xd6, 0x25, 0x61, 0x0f, 0x3a, 0xda, 0xc8, 0x4c, 0x2a, 0xd6,
	0x6b, 0x90, 0xfc, 0x0e, 0x65, 0xf6, 0xce, 0xf1, 0x82, 0xf5, 0x47, 0xad, 0xc9, 0x9e, 0xef, 0xe3,
	0x6d, 0x38, 0x59, 0x64, 0x6c, 0x40, 0x47, 0xe7, 0x30, 0x94, 0x2a, 0x95, 0x4b, 0x99, 0x56, 0x22,
	0xe7, 0xc2, 0x39, 0x91, 0x2c, 0xd8, 0x90, 0xec, 0x2f, 0x20, 0xde, 0x10, 0xa5, 0xf8, 0x16, 0x95,
	0x45, 0x16, 0x7f, 0x44, 0xd6, 0xd4, 0xfc, 0x64, 0x0d, 0x8d, 0x92, 0x17, 0x55, 0xee, 0x64, 0x99,
	0x4b, 0x34, 0xec, 0x94, 0x3c, 0x7d, 0x0b, 0x87, 0xbe, 0xb1, 0x73, 0x91, 0xe7, 0xec, 0x21, 0x15,
	0xf2, 0xd9, 0x66, 0x21, 0x5f, 0x37, 0xf3, 0xff, 0xf2, 0xb5, 0x9f, 0xff, 0xfe, 0xcf, 0xeb, 0xd4,
	0x17, 0xf3, 0x6e, 0x7c, 0x62, 0x9e, 0xfb, 0x62, 0x3e, 0x1a, 0xb5, 0x26, 0xfb, 0x5e, 0x32, 0x17,
	0xce, 0xe5, 0x68, 0xeb, 0xa8, 0x31, 0x65, 0x67, 0xcd, 0x30, 0x6e, 0x24, 0x14, 0x74, 0x8a, 0x29,
	0x63, 0xeb, 0x31, 0x9d, 0x65, 0x5c, 0xaa, 0xa4, 0x9a, 0x0b, 0xa7, 0x8d, 0xbf, 0xed, 0x9c, 0xda,
	0xe0, 0xfd, 0x18, 0x14, 0x34, 0xbf, 0xfd, 0x06, 0xe0, 0x85, 0x65, 0x17, 0xe4, 0xe7, 0x14, 0xba,
	0xaa, 0x2a, 0x78, 0x55, 0x66, 0x46, 0xa4, 0x68, 0xd9, 0x63, 0xba, 0x69, 0x04, 0x4c, 0xa4, 0xa9,
	0xf4, 0xfa, 0x22, 0xe7, 0xdb, 0xe9, 0x7e, 0x45, 0xe9, 0x0e, 0xe0, 0xf0, 0xad, 0x58, 0x6a, 0x23,
	0x1d, 0xb2, 0x27, 0x4d, 0x3f, 0x94, 0x4c, 0x16, 0xd4, 0xfc, 0xaf, 0xc9, 0xeb, 0x10, 0x8e, 0xde,
	0x1a, 0x5d, 0x10, 0x4c, 0xd8, 0x53, 0x52, 0x7a, 0x0c, 0x27, 0xf3, 0x2a, 0x4d, 0x57, 0x3c, 0x11,
	0x2a, 0x5d, 0x71, 0x71, 0x2b, 0x8c, 0x8f, 0x7f, 0x44, 0xc2, 0x27, 0xf0, 0x30, 0x08, 0x9d, 0x76,
	0x22, 0xbf, 0x83, 0x01, 0x7b, 0x46, 0x2e, 0x2f, 0x20, 0x4e, 0xc3, 0x80, 0xe2, 0x1b, 0x14, 0x1a,
	0x93, 0x69, 0x0c, 0xd0, 0xc8, 0x92, 0x92, 0xfd, 0x9c, 0xce, 0x7e, 0x0f, 0xfd, 0x46, 0xaf, 0x96,
	0xb1, 0xe7, 0xa3, 0xd6, 0x24, 0x9a, 0x3e, 0xbd, 0xfc, 0x60, 0x57, 0x5f, 0x6e, 0xcf, 0xbf, 0xf1,
	0xff, 0x1f, 0x00, 0x84, 0xa1, 0x43, 0xfc, 0x7d, 0x02, 0x0f, 0xb7, 0xaa, 0xe7, 0xf1, 0x50, 0xfa,
	0x12, 0x7a, 0x4a, 0xb7, 0x7d, 0xe2, 0x95, 0x45, 0x43, 0x89, 0x87, 0x91, 0xfb, 0x2b, 0xd8, 0x77,
	0x28, 0x0a, 0xe2, 0xe3, 0x47, 0xe9, 0xfd, 0x06, 0x45, 0xf1, 0x4a, 0xe7, 0xda, 0xc4, 0x7f, 0x80,
	0x9e, 0xab, 0x9c, 0x36, 0x32, 0xe0, 0xcc, 0x61, 0xbd, 0x0a, 0x9f, 0x7e, 0xc4, 0xa8, 0xd6, 0xfb,
	0xd1, 0xab, 0x5d, 0xed, 0x0d, 0x5a, 0xf1, 0x0b, 0xe8, 0x88, 0xa5, 0x70, 0xc2, 0x10, 0x03, 0xa3,
	0xe9, 0xf3, 0xfb, 0x89, 0x85, 0xb7, 0x45, 0xf8, 0xf7, 0x92, 0x74, 0x7d, 0x0f, 0x0a, 0xf1, 0x7e,
	0x5d, 0x43, 0xeb, 0xb4, 0x11, 0x59, 0x20, 0x2c, 0x61, 0xc8, 0x0b, 0xfd, 0x23, 0x65, 0x2d, 0x01,
	0x92, 0x7c, 0x07, 0x51, 0x2a, 0x64, 0xbe, 0xe2, 0x73, 0xad, 0x2a, 0x4b, 0x8c, 0x8d, 0xa6, 0xe3,
	0x5d, 0x1e, 0xaf, 0xbd, 0xea, 0x95, 0xd7, 0x8c, 0xff, 0x08, 0x3d, 0xfc, 0x67, 0x25, 0xcb, 0x12,
	0x53, 0x4e, 0xb3, 0x9b, 0x08, 0x1d, 0x4d, 0x7f, 0xb1, 0xcb, 0xf6, 0xa6, 0xd6, 0x0e, 0x13, 0xff,
	0x25, 0x0c, 0x12, 0xad, 0x9c, 0x48, 0x1c, 0xb7, 0xe8, 0x9c, 0x54, 0x99, 0xa5, 0x65, 0x18, 0x4d,
	0x7f, 0xb9, 0xeb, 0x82, 0x57, 0x41, 0xff, 0xc7, 0x5a, 0x3d, 0x7e, 0x01, 0x10, 0xa6, 0x79, 0x22,
	0xd1, 0xb2, 0xde, 0xa8, 0x3d, 0x89, 0xa6, 0xa3, 0x9d, 0xc6, 0x41, 0x73, 0x15, 0x3f, 0x83, 0x73,
	0x83, 0x85, 0x90, 0x4a, 0xaa, 0x8c, 0x27, 0x3a, 0x45, 0xdf, 0x61, 0x9e, 0xe4, 0x42, 0x16, 0x96,
	0x26, 0x8e, 0x7f, 0xfe, 0x1c, 0x07, 0xc4, 0xd6, 0xc5, 0xa4, 0xa9, 0x13, 0x6d, 0x8f, 0x50, 0xba,
	0xfb, 0xca, 0x6b, 0xd5, 0x28, 0x23, 0x9c, 0x13, 0x83, 0x79, 0xae, 0x93, 0x85, 0xae, 0x1c, 0xf7,
	0xeb, 0xa7, 0xb0, 0x34, 0x98, 0xda, 0xf1, 0x0d, 0x9c, 0x59, 0x4c, 0xb4, 0x4a, 0x85, 0x59, 0xd5,
	0x8f, 0x43, 0x5e, 0xb7, 0x39, 0xfe, 0xf2, 0x36, 0x8f, 0x6f, 0xa0, 0x7f, 0x25, 0x92, 0x45, 0x66,
	0x74, 0xa5, 0xd2, 0x37, 0x7a, 0x81, 0xca, 0x6f, 0x60, 0xe7, 0x3f, 0x08, 0xb8, 0xdd, 0xf8, 0x0c,
	0xfa, 0xf8, 0xbe, 0x94, 0xe6, 0x0e, 0xd9, 0x84, 0xdf, 0x36, 0x2d, 0xac, 0x25, 0xed, 0x9c, 0xee,
	0xf8, 0x16, 0x06, 0x2f, 0xfd, 0xeb, 0xe2, 0x5a, 0x66, 0x68, 0xdd, 0xae, 0x17, 0xc7, 0x09, 0x44,
	0xf3, 0x4a, 0xa5, 0x39, 0xf2, 0x0d, 0x1a, 0xf4, 0xe1, 0xa0, 0x59, 0x8c, 0xed, 0x86, 0x29, 0xdb,
	0x8f, 0x8f, 0xf5, 0xe3, 0x24, 0x6c, 0xae, 0x08, 0xda, 0x0b, 0x5c, 0xd1, 0xda, 0xea, 0x8e, 0xff,
	0x0c, 0xdd, 0xad, 0xaa, 0x6d, 0x6e, 0xd1, 0x33, 0xe8, 0x87, 0x35, 0x7e, 0x37, 0x23, 0xf6, 0x68,
	0x49, 0x9c, 0x41, 0x3f, 0x17, 0x96, 0xce, 0x9b, 0xd9, 0xe2, 0x5d, 0xb7, 0xae, 0x0e, 0xff, 0xd1,
	0xa1, 0x07, 0xb6, 0x9d, 0xfd, 0x6c, 0xd6, 0x9a, 0xed, 0xcd, 0xc3, 0xaf, 0x6f, 0x7f, 0x0a, 0x00,
	0x00, 0xff, 0xff, 0x8e, 0xa3, 0x14, 0x44, 0xc3, 0x0b, 0x00, 0x00,
}
