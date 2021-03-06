// Code generated by protoc-gen-go.
// source: networking_requests.proto
// DO NOT EDIT!

package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type RequestType int32

const (
	RequestType_METHOD_UNSET                   RequestType = 0
	RequestType_GET_PLAYER                     RequestType = 2
	RequestType_GET_INVENTORY                  RequestType = 4
	RequestType_DOWNLOAD_SETTINGS              RequestType = 5
	RequestType_DOWNLOAD_ITEM_TEMPLATES        RequestType = 6
	RequestType_DOWNLOAD_REMOTE_CONFIG_VERSION RequestType = 7
	RequestType_REGISTER_BACKGROUND_DEVICE     RequestType = 8
	RequestType_FORT_SEARCH                    RequestType = 101
	RequestType_ENCOUNTER                      RequestType = 102
	RequestType_CATCH_POKEMON                  RequestType = 103
	RequestType_FORT_DETAILS                   RequestType = 104
	RequestType_GET_MAP_OBJECTS                RequestType = 106
	RequestType_FORT_DEPLOY_POKEMON            RequestType = 110
	RequestType_FORT_RECALL_POKEMON            RequestType = 111
	RequestType_RELEASE_POKEMON                RequestType = 112
	RequestType_USE_ITEM_POTION                RequestType = 113
	RequestType_USE_ITEM_CAPTURE               RequestType = 114
	RequestType_USE_ITEM_FLEE                  RequestType = 115
	RequestType_USE_ITEM_REVIVE                RequestType = 116
	RequestType_GET_PLAYER_PROFILE             RequestType = 121
	RequestType_EVOLVE_POKEMON                 RequestType = 125
	RequestType_GET_HATCHED_EGGS               RequestType = 126
	RequestType_ENCOUNTER_TUTORIAL_COMPLETE    RequestType = 127
	RequestType_LEVEL_UP_REWARDS               RequestType = 128
	RequestType_CHECK_AWARDED_BADGES           RequestType = 129
	RequestType_USE_ITEM_GYM                   RequestType = 133
	RequestType_GET_GYM_DETAILS                RequestType = 134
	RequestType_START_GYM_BATTLE               RequestType = 135
	RequestType_ATTACK_GYM                     RequestType = 136
	RequestType_RECYCLE_INVENTORY_ITEM         RequestType = 137
	RequestType_COLLECT_DAILY_BONUS            RequestType = 138
	RequestType_USE_ITEM_XP_BOOST              RequestType = 139
	RequestType_USE_ITEM_EGG_INCUBATOR         RequestType = 140
	RequestType_USE_INCENSE                    RequestType = 141
	RequestType_GET_INCENSE_POKEMON            RequestType = 142
	RequestType_INCENSE_ENCOUNTER              RequestType = 143
	RequestType_ADD_FORT_MODIFIER              RequestType = 144
	RequestType_DISK_ENCOUNTER                 RequestType = 145
	RequestType_COLLECT_DAILY_DEFENDER_BONUS   RequestType = 146
	RequestType_UPGRADE_POKEMON                RequestType = 147
	RequestType_SET_FAVORITE_POKEMON           RequestType = 148
	RequestType_NICKNAME_POKEMON               RequestType = 149
	RequestType_EQUIP_BADGE                    RequestType = 150
	RequestType_SET_CONTACT_SETTINGS           RequestType = 151
	RequestType_SET_BUDDY_POKEMON              RequestType = 152
	RequestType_GET_BUDDY_WALKED               RequestType = 153
	RequestType_USE_ITEM_ENCOUNTER             RequestType = 154
	RequestType_GET_ASSET_DIGEST               RequestType = 300
	RequestType_GET_DOWNLOAD_URLS              RequestType = 301
	RequestType_CLAIM_CODENAME                 RequestType = 403
	RequestType_SET_AVATAR                     RequestType = 404
	RequestType_SET_PLAYER_TEAM                RequestType = 405
	RequestType_MARK_TUTORIAL_COMPLETE         RequestType = 406
	RequestType_CHECK_CHALLENGE                RequestType = 600
	RequestType_VERIFY_CHALLENGE               RequestType = 601
	RequestType_ECHO                           RequestType = 666
	RequestType_SFIDA_REGISTRATION             RequestType = 800
	RequestType_SFIDA_ACTION_LOG               RequestType = 801
	RequestType_SFIDA_CERTIFICATION            RequestType = 802
	RequestType_SFIDA_UPDATE                   RequestType = 803
	RequestType_SFIDA_ACTION                   RequestType = 804
	RequestType_SFIDA_DOWSER                   RequestType = 805
	RequestType_SFIDA_CAPTURE                  RequestType = 806
	RequestType_LIST_AVATAR_CUSTOMIZATIONS     RequestType = 807
	RequestType_SET_AVATAR_ITEM_AS_VIEWED      RequestType = 808
	RequestType_GET_INBOX                      RequestType = 809
	RequestType_UPDATE_NOTIFICATION_STATUS     RequestType = 810
)

var RequestType_name = map[int32]string{
	0:   "METHOD_UNSET",
	2:   "GET_PLAYER",
	4:   "GET_INVENTORY",
	5:   "DOWNLOAD_SETTINGS",
	6:   "DOWNLOAD_ITEM_TEMPLATES",
	7:   "DOWNLOAD_REMOTE_CONFIG_VERSION",
	8:   "REGISTER_BACKGROUND_DEVICE",
	101: "FORT_SEARCH",
	102: "ENCOUNTER",
	103: "CATCH_POKEMON",
	104: "FORT_DETAILS",
	106: "GET_MAP_OBJECTS",
	110: "FORT_DEPLOY_POKEMON",
	111: "FORT_RECALL_POKEMON",
	112: "RELEASE_POKEMON",
	113: "USE_ITEM_POTION",
	114: "USE_ITEM_CAPTURE",
	115: "USE_ITEM_FLEE",
	116: "USE_ITEM_REVIVE",
	121: "GET_PLAYER_PROFILE",
	125: "EVOLVE_POKEMON",
	126: "GET_HATCHED_EGGS",
	127: "ENCOUNTER_TUTORIAL_COMPLETE",
	128: "LEVEL_UP_REWARDS",
	129: "CHECK_AWARDED_BADGES",
	133: "USE_ITEM_GYM",
	134: "GET_GYM_DETAILS",
	135: "START_GYM_BATTLE",
	136: "ATTACK_GYM",
	137: "RECYCLE_INVENTORY_ITEM",
	138: "COLLECT_DAILY_BONUS",
	139: "USE_ITEM_XP_BOOST",
	140: "USE_ITEM_EGG_INCUBATOR",
	141: "USE_INCENSE",
	142: "GET_INCENSE_POKEMON",
	143: "INCENSE_ENCOUNTER",
	144: "ADD_FORT_MODIFIER",
	145: "DISK_ENCOUNTER",
	146: "COLLECT_DAILY_DEFENDER_BONUS",
	147: "UPGRADE_POKEMON",
	148: "SET_FAVORITE_POKEMON",
	149: "NICKNAME_POKEMON",
	150: "EQUIP_BADGE",
	151: "SET_CONTACT_SETTINGS",
	152: "SET_BUDDY_POKEMON",
	153: "GET_BUDDY_WALKED",
	154: "USE_ITEM_ENCOUNTER",
	300: "GET_ASSET_DIGEST",
	301: "GET_DOWNLOAD_URLS",
	403: "CLAIM_CODENAME",
	404: "SET_AVATAR",
	405: "SET_PLAYER_TEAM",
	406: "MARK_TUTORIAL_COMPLETE",
	600: "CHECK_CHALLENGE",
	601: "VERIFY_CHALLENGE",
	666: "ECHO",
	800: "SFIDA_REGISTRATION",
	801: "SFIDA_ACTION_LOG",
	802: "SFIDA_CERTIFICATION",
	803: "SFIDA_UPDATE",
	804: "SFIDA_ACTION",
	805: "SFIDA_DOWSER",
	806: "SFIDA_CAPTURE",
	807: "LIST_AVATAR_CUSTOMIZATIONS",
	808: "SET_AVATAR_ITEM_AS_VIEWED",
	809: "GET_INBOX",
	810: "UPDATE_NOTIFICATION_STATUS",
}
var RequestType_value = map[string]int32{
	"METHOD_UNSET":                   0,
	"GET_PLAYER":                     2,
	"GET_INVENTORY":                  4,
	"DOWNLOAD_SETTINGS":              5,
	"DOWNLOAD_ITEM_TEMPLATES":        6,
	"DOWNLOAD_REMOTE_CONFIG_VERSION": 7,
	"REGISTER_BACKGROUND_DEVICE":     8,
	"FORT_SEARCH":                    101,
	"ENCOUNTER":                      102,
	"CATCH_POKEMON":                  103,
	"FORT_DETAILS":                   104,
	"GET_MAP_OBJECTS":                106,
	"FORT_DEPLOY_POKEMON":            110,
	"FORT_RECALL_POKEMON":            111,
	"RELEASE_POKEMON":                112,
	"USE_ITEM_POTION":                113,
	"USE_ITEM_CAPTURE":               114,
	"USE_ITEM_FLEE":                  115,
	"USE_ITEM_REVIVE":                116,
	"GET_PLAYER_PROFILE":             121,
	"EVOLVE_POKEMON":                 125,
	"GET_HATCHED_EGGS":               126,
	"ENCOUNTER_TUTORIAL_COMPLETE":    127,
	"LEVEL_UP_REWARDS":               128,
	"CHECK_AWARDED_BADGES":           129,
	"USE_ITEM_GYM":                   133,
	"GET_GYM_DETAILS":                134,
	"START_GYM_BATTLE":               135,
	"ATTACK_GYM":                     136,
	"RECYCLE_INVENTORY_ITEM":         137,
	"COLLECT_DAILY_BONUS":            138,
	"USE_ITEM_XP_BOOST":              139,
	"USE_ITEM_EGG_INCUBATOR":         140,
	"USE_INCENSE":                    141,
	"GET_INCENSE_POKEMON":            142,
	"INCENSE_ENCOUNTER":              143,
	"ADD_FORT_MODIFIER":              144,
	"DISK_ENCOUNTER":                 145,
	"COLLECT_DAILY_DEFENDER_BONUS":   146,
	"UPGRADE_POKEMON":                147,
	"SET_FAVORITE_POKEMON":           148,
	"NICKNAME_POKEMON":               149,
	"EQUIP_BADGE":                    150,
	"SET_CONTACT_SETTINGS":           151,
	"SET_BUDDY_POKEMON":              152,
	"GET_BUDDY_WALKED":               153,
	"USE_ITEM_ENCOUNTER":             154,
	"GET_ASSET_DIGEST":               300,
	"GET_DOWNLOAD_URLS":              301,
	"CLAIM_CODENAME":                 403,
	"SET_AVATAR":                     404,
	"SET_PLAYER_TEAM":                405,
	"MARK_TUTORIAL_COMPLETE":         406,
	"CHECK_CHALLENGE":                600,
	"VERIFY_CHALLENGE":               601,
	"ECHO":                           666,
	"SFIDA_REGISTRATION":             800,
	"SFIDA_ACTION_LOG":               801,
	"SFIDA_CERTIFICATION":            802,
	"SFIDA_UPDATE":                   803,
	"SFIDA_ACTION":                   804,
	"SFIDA_DOWSER":                   805,
	"SFIDA_CAPTURE":                  806,
	"LIST_AVATAR_CUSTOMIZATIONS":     807,
	"SET_AVATAR_ITEM_AS_VIEWED":      808,
	"GET_INBOX":                      809,
	"UPDATE_NOTIFICATION_STATUS":     810,
}

func (x RequestType) String() string {
	return proto.EnumName(RequestType_name, int32(x))
}
func (RequestType) EnumDescriptor() ([]byte, []int) { return fileDescriptor18, []int{0} }

type Request struct {
	RequestType    RequestType `protobuf:"varint,1,opt,name=request_type,enum=POGOProtos.Networking.Requests.RequestType" json:"request_type,omitempty"`
	RequestMessage []byte      `protobuf:"bytes,2,opt,name=request_message,proto3" json:"request_message,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor18, []int{0} }

func init() {
	proto.RegisterType((*Request)(nil), "POGOProtos.Networking.Requests.Request")
	proto.RegisterEnum("POGOProtos.Networking.Requests.RequestType", RequestType_name, RequestType_value)
}

func init() { proto.RegisterFile("networking_requests.proto", fileDescriptor18) }

var fileDescriptor18 = []byte{
	// 999 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x55, 0x49, 0x6f, 0x1c, 0x37,
	0x13, 0xfd, 0x5a, 0x92, 0xc7, 0x16, 0xb5, 0x95, 0x28, 0x59, 0x8b, 0xf5, 0x41, 0x76, 0x7c, 0x32,
	0x12, 0x40, 0x87, 0xe4, 0x17, 0x70, 0xc8, 0x9a, 0x1e, 0x66, 0xd8, 0xcd, 0x0e, 0xc9, 0x6e, 0x79,
	0x7c, 0x21, 0x12, 0x60, 0xa2, 0x2c, 0x88, 0x24, 0x6b, 0x14, 0x04, 0x3a, 0x64, 0x3b, 0x64, 0xdf,
	0xe3, 0x25, 0x89, 0x4e, 0xd9, 0x37, 0x24, 0xff, 0x27, 0xf9, 0x37, 0x01, 0xd9, 0xa3, 0xee, 0x31,
	0x10, 0xe4, 0xa4, 0xd6, 0xab, 0xaa, 0x57, 0x55, 0x8f, 0x8f, 0x1c, 0xb2, 0x7d, 0x38, 0x3a, 0x7d,
	0xed, 0xe8, 0xe4, 0xe5, 0x17, 0x0f, 0x0f, 0xfc, 0xc9, 0xe8, 0xee, 0xab, 0xa3, 0xf1, 0xe9, 0x78,
	0xef, 0xf8, 0xe4, 0xe8, 0xf4, 0x88, 0xee, 0x16, 0x3a, 0xd5, 0x45, 0xf8, 0x1c, 0xef, 0xe5, 0x4d,
	0xd6, 0x9e, 0x99, 0x64, 0xdd, 0x1c, 0x91, 0xcb, 0x93, 0x6f, 0xca, 0xc8, 0xe2, 0xa4, 0xd8, 0x9f,
	0x9e, 0x1d, 0x8f, 0xb6, 0x92, 0x1b, 0xc9, 0xad, 0xe5, 0x27, 0x9f, 0xd8, 0xfb, 0x6f, 0x86, 0x8b,
	0x0f, 0x77, 0x76, 0x3c, 0xa2, 0x9b, 0x64, 0xe5, 0x82, 0xe2, 0x95, 0xd1, 0x78, 0xfc, 0xec, 0xc1,
	0x68, 0x6b, 0xe6, 0x46, 0x72, 0x6b, 0xf1, 0xf1, 0xf3, 0x45, 0xb2, 0x30, 0x9d, 0x08, 0x64, 0x31,
	0x43, 0xd7, 0xd7, 0xc2, 0x97, 0xb9, 0x45, 0x07, 0xff, 0xa3, 0xcb, 0x84, 0xa4, 0xe8, 0x7c, 0xa1,
	0xd8, 0x10, 0x0d, 0xcc, 0xd0, 0x55, 0xb2, 0x14, 0xfe, 0x97, 0x79, 0x85, 0xb9, 0xd3, 0x66, 0x08,
	0x73, 0xf4, 0x2a, 0x59, 0x15, 0x7a, 0x3f, 0x57, 0x9a, 0x09, 0x6f, 0xd1, 0x39, 0x99, 0xa7, 0x16,
	0x2e, 0xd1, 0x1d, 0xb2, 0xd9, 0xc0, 0xd2, 0x61, 0xe6, 0x1d, 0x66, 0x85, 0x62, 0x0e, 0x2d, 0x74,
	0xe8, 0x4d, 0xb2, 0xdb, 0x04, 0x0d, 0x66, 0xda, 0xa1, 0xe7, 0x3a, 0xef, 0xc9, 0xd4, 0x57, 0x68,
	0xac, 0xd4, 0x39, 0x5c, 0xa6, 0xbb, 0xe4, 0x9a, 0xc1, 0x54, 0x5a, 0x87, 0xc6, 0x77, 0x19, 0x1f,
	0xa4, 0x46, 0x97, 0xb9, 0xf0, 0x02, 0x2b, 0xc9, 0x11, 0xae, 0xd0, 0x15, 0xb2, 0xd0, 0xd3, 0xc6,
	0x79, 0x8b, 0xcc, 0xf0, 0x3e, 0x8c, 0xe8, 0x12, 0x99, 0xc7, 0x9c, 0xeb, 0x32, 0x77, 0x68, 0xe0,
	0xf9, 0x30, 0x2a, 0x67, 0x8e, 0xf7, 0x7d, 0xa1, 0x07, 0x98, 0xe9, 0x1c, 0x0e, 0xc2, 0x7e, 0xb1,
	0x44, 0xa0, 0x63, 0x52, 0x59, 0x78, 0x81, 0xae, 0x91, 0x95, 0xb0, 0x4f, 0xc6, 0x0a, 0xaf, 0xbb,
	0x4f, 0x23, 0x77, 0x16, 0x5e, 0xa2, 0x9b, 0x64, 0x6d, 0x92, 0x56, 0x28, 0x3d, 0x6c, 0xea, 0x0f,
	0x9b, 0x80, 0x41, 0xce, 0x94, 0x6a, 0x02, 0x47, 0x81, 0xc6, 0xa0, 0x42, 0x66, 0xb1, 0x01, 0x8f,
	0x03, 0x58, 0x5a, 0xac, 0x97, 0x2f, 0xb4, 0x0b, 0x5b, 0xdd, 0xa5, 0xeb, 0x04, 0x1a, 0x90, 0xb3,
	0xc2, 0x95, 0x06, 0xe1, 0x24, 0xcc, 0xda, 0xa0, 0x3d, 0x85, 0x08, 0xe3, 0x47, 0xaa, 0x0d, 0x56,
	0xb2, 0x42, 0x38, 0xa5, 0x1b, 0x84, 0xb6, 0xc7, 0xe1, 0x0b, 0xa3, 0x7b, 0x52, 0x21, 0x9c, 0x51,
	0x4a, 0x96, 0xb1, 0xd2, 0xaa, 0x6a, 0xdb, 0xbf, 0x1e, 0x3a, 0x85, 0xdc, 0x7e, 0xd0, 0x00, 0x85,
	0xc7, 0x34, 0xb5, 0xf0, 0x06, 0xbd, 0x4e, 0x76, 0x1a, 0x91, 0xbc, 0x2b, 0x9d, 0x36, 0x92, 0x29,
	0xcf, 0x75, 0x56, 0x28, 0x74, 0x08, 0x6f, 0xd2, 0xab, 0x04, 0x14, 0x56, 0xa8, 0x7c, 0x59, 0x78,
	0x83, 0xfb, 0xcc, 0x08, 0x0b, 0x6f, 0x25, 0x74, 0x9b, 0xac, 0xf3, 0x3e, 0xf2, 0x81, 0x67, 0x01,
	0x42, 0xe1, 0xbb, 0x4c, 0xa4, 0x68, 0xe1, 0xed, 0x84, 0xae, 0x92, 0xc5, 0x66, 0xd2, 0x74, 0x98,
	0xc1, 0x3b, 0x09, 0x5d, 0xaf, 0x65, 0x4d, 0x87, 0x59, 0xa3, 0xf5, 0xbb, 0x49, 0xa0, 0xb6, 0x8e,
	0x99, 0x1a, 0xef, 0x32, 0xe7, 0x14, 0xc2, 0x7b, 0x09, 0x5d, 0x21, 0x84, 0x39, 0xc7, 0xf8, 0x20,
	0x56, 0xbf, 0x9f, 0xd0, 0x1d, 0xb2, 0x61, 0x90, 0x0f, 0xb9, 0xc2, 0xd6, 0x68, 0x91, 0x1e, 0x3e,
	0x48, 0xe8, 0x16, 0x59, 0xe3, 0x5a, 0x29, 0xe4, 0xce, 0x0b, 0x26, 0xd5, 0xd0, 0x77, 0x75, 0x5e,
	0x5a, 0xf8, 0x30, 0xa1, 0x1b, 0x64, 0xb5, 0x99, 0xe3, 0x76, 0xe1, 0xbb, 0x5a, 0x5b, 0x07, 0x1f,
	0x45, 0xba, 0x06, 0xc7, 0x34, 0xf5, 0x32, 0xe7, 0x65, 0x97, 0x39, 0x6d, 0xe0, 0xe3, 0x84, 0x02,
	0x59, 0x88, 0xc1, 0x9c, 0x63, 0x6e, 0x11, 0x3e, 0x89, 0x0d, 0x6a, 0x8b, 0x47, 0xa4, 0x11, 0xf4,
	0xd3, 0xd8, 0xe0, 0x02, 0x6d, 0x8d, 0xf6, 0x59, 0xc4, 0x99, 0x10, 0x3e, 0x5a, 0x23, 0xd3, 0x42,
	0xf6, 0x24, 0x1a, 0xf8, 0x3c, 0xa1, 0x6b, 0x64, 0x59, 0x48, 0x3b, 0x98, 0x4a, 0xfe, 0x22, 0xa1,
	0x8f, 0x91, 0xff, 0x3f, 0x3a, 0xbf, 0xc0, 0x1e, 0xe6, 0x22, 0x98, 0x3c, 0x2e, 0xf2, 0x65, 0x54,
	0xaf, 0x2c, 0x52, 0xc3, 0x44, 0xdb, 0xfd, 0x5e, 0x3c, 0x01, 0x8b, 0xce, 0xf7, 0x58, 0xa5, 0x8d,
	0x74, 0x6d, 0xe8, 0x7e, 0x14, 0x36, 0x97, 0x7c, 0x90, 0xb3, 0xac, 0x85, 0x1f, 0xc4, 0xdd, 0xf0,
	0x99, 0x52, 0x16, 0xf5, 0x59, 0xc1, 0xc3, 0x86, 0x83, 0xeb, 0xdc, 0x31, 0xee, 0xda, 0xeb, 0xfa,
	0x55, 0x5c, 0x22, 0x84, 0xba, 0xa5, 0x10, 0xad, 0xe5, 0xbf, 0x8e, 0xdc, 0x69, 0x83, 0xef, 0x33,
	0x35, 0x40, 0x01, 0xdf, 0x24, 0x74, 0x93, 0xd0, 0x56, 0xd4, 0x66, 0xbf, 0xf3, 0x26, 0x9f, 0xd9,
	0xc0, 0x26, 0x64, 0x8a, 0xd6, 0xc1, 0x1f, 0x33, 0x81, 0x3e, 0xc0, 0xcd, 0xad, 0x2f, 0x8d, 0xb2,
	0xf0, 0xe7, 0x4c, 0xd0, 0x88, 0x2b, 0x26, 0x33, 0xcf, 0xb5, 0xc0, 0xb0, 0x00, 0xdc, 0x9b, 0x0d,
	0x8e, 0x08, 0xd5, 0xac, 0x62, 0x8e, 0x19, 0xb8, 0x3f, 0x1b, 0x14, 0xb1, 0xad, 0xef, 0x1d, 0xb2,
	0x0c, 0x1e, 0xcc, 0x86, 0x83, 0xcd, 0x98, 0x19, 0xfc, 0x8b, 0x8d, 0x1f, 0xc6, 0x92, 0xda, 0xb0,
	0xbc, 0xcf, 0x94, 0xc2, 0x3c, 0x45, 0xf8, 0x2b, 0x3c, 0x56, 0x50, 0xa1, 0x91, 0xbd, 0xe1, 0x14,
	0xfc, 0xf7, 0x1c, 0x9d, 0x27, 0x73, 0xc8, 0xfb, 0x1a, 0xce, 0x2f, 0x85, 0xc5, 0x6c, 0x4f, 0x0a,
	0xe6, 0xeb, 0xc7, 0xc7, 0xb0, 0x78, 0x71, 0xbf, 0xed, 0x44, 0xf7, 0xc6, 0x00, 0xe3, 0x01, 0xf2,
	0x4a, 0xa7, 0xf0, 0x5d, 0x27, 0xd8, 0xa5, 0x86, 0x39, 0x1a, 0x27, 0x7b, 0x92, 0xd7, 0x05, 0xdf,
	0x77, 0xc2, 0xbd, 0xa8, 0x23, 0x65, 0x21, 0x98, 0x43, 0xf8, 0x61, 0x0a, 0xaa, 0x39, 0xe0, 0xc7,
	0x29, 0x48, 0xe8, 0x7d, 0x8b, 0x06, 0x7e, 0xea, 0x50, 0x4a, 0x96, 0x26, 0x94, 0x93, 0x07, 0xe2,
	0xe7, 0x0e, 0xbd, 0x4e, 0xae, 0x29, 0x69, 0x2f, 0x34, 0xf1, 0xbc, 0xb4, 0x4e, 0x67, 0xf2, 0x4e,
	0x6c, 0x66, 0xe1, 0x97, 0x0e, 0xdd, 0x25, 0xdb, 0xad, 0x66, 0xf5, 0xb9, 0x30, 0xeb, 0x2b, 0x89,
	0xfb, 0x28, 0xe0, 0xd7, 0x0e, 0x5d, 0x26, 0xf3, 0xb5, 0xad, 0xbb, 0xfa, 0x36, 0xfc, 0x16, 0x09,
	0xeb, 0xb9, 0x7c, 0xae, 0xdb, 0xb9, 0xbd, 0x75, 0xcc, 0x95, 0x16, 0x7e, 0xef, 0x74, 0xaf, 0xdc,
	0xe9, 0xc4, 0x1f, 0xab, 0xf1, 0x73, 0xf5, 0xdf, 0xa7, 0xfe, 0x09, 0x00, 0x00, 0xff, 0xff, 0x03,
	0xa0, 0xea, 0x03, 0xd1, 0x06, 0x00, 0x00,
}
