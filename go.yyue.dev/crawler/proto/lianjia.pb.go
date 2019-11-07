// Code generated by protoc-gen-go. DO NOT EDIT.
// source: crawler/lianjia.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type House struct {
	Id                   string                `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	HouseNo              string                `protobuf:"bytes,2,opt,name=house_no,json=houseNo,proto3" json:"house_no,omitempty"`
	HouseInfo            *HouseInfo            `protobuf:"bytes,3,opt,name=house_info,json=houseInfo,proto3" json:"house_info,omitempty"`
	HouseBaseInfo        *HouseBaseInfo        `protobuf:"bytes,4,opt,name=house_base_info,json=houseBaseInfo,proto3" json:"house_base_info,omitempty"`
	HouseTransactionInfo *HouseTransactionInfo `protobuf:"bytes,5,opt,name=house_transaction_info,json=houseTransactionInfo,proto3" json:"house_transaction_info,omitempty"`
	HousePics            []*HousePic           `protobuf:"bytes,6,rep,name=house_pics,json=housePics,proto3" json:"house_pics,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *House) Reset()         { *m = House{} }
func (m *House) String() string { return proto.CompactTextString(m) }
func (*House) ProtoMessage()    {}
func (*House) Descriptor() ([]byte, []int) {
	return fileDescriptor_af7531cd28e23262, []int{0}
}

func (m *House) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_House.Unmarshal(m, b)
}
func (m *House) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_House.Marshal(b, m, deterministic)
}
func (m *House) XXX_Merge(src proto.Message) {
	xxx_messageInfo_House.Merge(m, src)
}
func (m *House) XXX_Size() int {
	return xxx_messageInfo_House.Size(m)
}
func (m *House) XXX_DiscardUnknown() {
	xxx_messageInfo_House.DiscardUnknown(m)
}

var xxx_messageInfo_House proto.InternalMessageInfo

func (m *House) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *House) GetHouseNo() string {
	if m != nil {
		return m.HouseNo
	}
	return ""
}

func (m *House) GetHouseInfo() *HouseInfo {
	if m != nil {
		return m.HouseInfo
	}
	return nil
}

func (m *House) GetHouseBaseInfo() *HouseBaseInfo {
	if m != nil {
		return m.HouseBaseInfo
	}
	return nil
}

func (m *House) GetHouseTransactionInfo() *HouseTransactionInfo {
	if m != nil {
		return m.HouseTransactionInfo
	}
	return nil
}

func (m *House) GetHousePics() []*HousePic {
	if m != nil {
		return m.HousePics
	}
	return nil
}

// 房子信息
type HouseInfo struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	HouseNo              string   `protobuf:"bytes,2,opt,name=house_no,json=houseNo,proto3" json:"house_no,omitempty"`
	Url                  string   `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	Title                string   `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	SubTitle             string   `protobuf:"bytes,5,opt,name=sub_title,json=subTitle,proto3" json:"sub_title,omitempty"`
	Region               string   `protobuf:"bytes,6,opt,name=region,proto3" json:"region,omitempty"`
	TotalPrice           string   `protobuf:"bytes,7,opt,name=total_price,json=totalPrice,proto3" json:"total_price,omitempty"`
	UnitPrice            string   `protobuf:"bytes,8,opt,name=unit_price,json=unitPrice,proto3" json:"unit_price,omitempty"`
	RoomInfo             string   `protobuf:"bytes,9,opt,name=room_info,json=roomInfo,proto3" json:"room_info,omitempty"`
	RoomSub              string   `protobuf:"bytes,10,opt,name=room_sub,json=roomSub,proto3" json:"room_sub,omitempty"`
	TypeInfo             string   `protobuf:"bytes,11,opt,name=type_info,json=typeInfo,proto3" json:"type_info,omitempty"`
	TypeSub              string   `protobuf:"bytes,12,opt,name=type_sub,json=typeSub,proto3" json:"type_sub,omitempty"`
	AreaInfo             string   `protobuf:"bytes,13,opt,name=area_info,json=areaInfo,proto3" json:"area_info,omitempty"`
	AreaSub              string   `protobuf:"bytes,14,opt,name=area_sub,json=areaSub,proto3" json:"area_sub,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HouseInfo) Reset()         { *m = HouseInfo{} }
func (m *HouseInfo) String() string { return proto.CompactTextString(m) }
func (*HouseInfo) ProtoMessage()    {}
func (*HouseInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_af7531cd28e23262, []int{1}
}

func (m *HouseInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HouseInfo.Unmarshal(m, b)
}
func (m *HouseInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HouseInfo.Marshal(b, m, deterministic)
}
func (m *HouseInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HouseInfo.Merge(m, src)
}
func (m *HouseInfo) XXX_Size() int {
	return xxx_messageInfo_HouseInfo.Size(m)
}
func (m *HouseInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_HouseInfo.DiscardUnknown(m)
}

var xxx_messageInfo_HouseInfo proto.InternalMessageInfo

func (m *HouseInfo) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *HouseInfo) GetHouseNo() string {
	if m != nil {
		return m.HouseNo
	}
	return ""
}

func (m *HouseInfo) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *HouseInfo) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *HouseInfo) GetSubTitle() string {
	if m != nil {
		return m.SubTitle
	}
	return ""
}

func (m *HouseInfo) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

func (m *HouseInfo) GetTotalPrice() string {
	if m != nil {
		return m.TotalPrice
	}
	return ""
}

func (m *HouseInfo) GetUnitPrice() string {
	if m != nil {
		return m.UnitPrice
	}
	return ""
}

func (m *HouseInfo) GetRoomInfo() string {
	if m != nil {
		return m.RoomInfo
	}
	return ""
}

func (m *HouseInfo) GetRoomSub() string {
	if m != nil {
		return m.RoomSub
	}
	return ""
}

func (m *HouseInfo) GetTypeInfo() string {
	if m != nil {
		return m.TypeInfo
	}
	return ""
}

func (m *HouseInfo) GetTypeSub() string {
	if m != nil {
		return m.TypeSub
	}
	return ""
}

func (m *HouseInfo) GetAreaInfo() string {
	if m != nil {
		return m.AreaInfo
	}
	return ""
}

func (m *HouseInfo) GetAreaSub() string {
	if m != nil {
		return m.AreaSub
	}
	return ""
}

// 基本信息
type HouseBaseInfo struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	HouseNo              string   `protobuf:"bytes,2,opt,name=house_no,json=houseNo,proto3" json:"house_no,omitempty"`
	Layout               string   `protobuf:"bytes,3,opt,name=layout,proto3" json:"layout,omitempty"`
	Floor                string   `protobuf:"bytes,4,opt,name=floor,proto3" json:"floor,omitempty"`
	AreaBuild            string   `protobuf:"bytes,5,opt,name=area_build,json=areaBuild,proto3" json:"area_build,omitempty"`
	StructHouse          string   `protobuf:"bytes,6,opt,name=struct_house,json=structHouse,proto3" json:"struct_house,omitempty"`
	AreaInner            string   `protobuf:"bytes,7,opt,name=area_inner,json=areaInner,proto3" json:"area_inner,omitempty"`
	BuildType            string   `protobuf:"bytes,8,opt,name=build_type,json=buildType,proto3" json:"build_type,omitempty"`
	Face                 string   `protobuf:"bytes,9,opt,name=face,proto3" json:"face,omitempty"`
	StructBuild          string   `protobuf:"bytes,10,opt,name=struct_build,json=structBuild,proto3" json:"struct_build,omitempty"`
	Decoration           string   `protobuf:"bytes,11,opt,name=decoration,proto3" json:"decoration,omitempty"`
	ElevatorRatio        string   `protobuf:"bytes,12,opt,name=elevator_ratio,json=elevatorRatio,proto3" json:"elevator_ratio,omitempty"`
	Elevator             string   `protobuf:"bytes,13,opt,name=elevator,proto3" json:"elevator,omitempty"`
	Property             string   `protobuf:"bytes,14,opt,name=property,proto3" json:"property,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HouseBaseInfo) Reset()         { *m = HouseBaseInfo{} }
func (m *HouseBaseInfo) String() string { return proto.CompactTextString(m) }
func (*HouseBaseInfo) ProtoMessage()    {}
func (*HouseBaseInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_af7531cd28e23262, []int{2}
}

func (m *HouseBaseInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HouseBaseInfo.Unmarshal(m, b)
}
func (m *HouseBaseInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HouseBaseInfo.Marshal(b, m, deterministic)
}
func (m *HouseBaseInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HouseBaseInfo.Merge(m, src)
}
func (m *HouseBaseInfo) XXX_Size() int {
	return xxx_messageInfo_HouseBaseInfo.Size(m)
}
func (m *HouseBaseInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_HouseBaseInfo.DiscardUnknown(m)
}

var xxx_messageInfo_HouseBaseInfo proto.InternalMessageInfo

func (m *HouseBaseInfo) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *HouseBaseInfo) GetHouseNo() string {
	if m != nil {
		return m.HouseNo
	}
	return ""
}

func (m *HouseBaseInfo) GetLayout() string {
	if m != nil {
		return m.Layout
	}
	return ""
}

func (m *HouseBaseInfo) GetFloor() string {
	if m != nil {
		return m.Floor
	}
	return ""
}

func (m *HouseBaseInfo) GetAreaBuild() string {
	if m != nil {
		return m.AreaBuild
	}
	return ""
}

func (m *HouseBaseInfo) GetStructHouse() string {
	if m != nil {
		return m.StructHouse
	}
	return ""
}

func (m *HouseBaseInfo) GetAreaInner() string {
	if m != nil {
		return m.AreaInner
	}
	return ""
}

func (m *HouseBaseInfo) GetBuildType() string {
	if m != nil {
		return m.BuildType
	}
	return ""
}

func (m *HouseBaseInfo) GetFace() string {
	if m != nil {
		return m.Face
	}
	return ""
}

func (m *HouseBaseInfo) GetStructBuild() string {
	if m != nil {
		return m.StructBuild
	}
	return ""
}

func (m *HouseBaseInfo) GetDecoration() string {
	if m != nil {
		return m.Decoration
	}
	return ""
}

func (m *HouseBaseInfo) GetElevatorRatio() string {
	if m != nil {
		return m.ElevatorRatio
	}
	return ""
}

func (m *HouseBaseInfo) GetElevator() string {
	if m != nil {
		return m.Elevator
	}
	return ""
}

func (m *HouseBaseInfo) GetProperty() string {
	if m != nil {
		return m.Property
	}
	return ""
}

// 交易信息
type HouseTransactionInfo struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	HouseNo              string   `protobuf:"bytes,2,opt,name=house_no,json=houseNo,proto3" json:"house_no,omitempty"`
	ListingTime          string   `protobuf:"bytes,3,opt,name=listing_time,json=listingTime,proto3" json:"listing_time,omitempty"`
	TradingAuthority     string   `protobuf:"bytes,4,opt,name=trading_authority,json=tradingAuthority,proto3" json:"trading_authority,omitempty"`
	LastTransaction      string   `protobuf:"bytes,5,opt,name=last_transaction,json=lastTransaction,proto3" json:"last_transaction,omitempty"`
	HousingPurposes      string   `protobuf:"bytes,6,opt,name=housing_purposes,json=housingPurposes,proto3" json:"housing_purposes,omitempty"`
	HouseYear            string   `protobuf:"bytes,7,opt,name=house_year,json=houseYear,proto3" json:"house_year,omitempty"`
	PropertyRights       string   `protobuf:"bytes,8,opt,name=property_rights,json=propertyRights,proto3" json:"property_rights,omitempty"`
	MortgageInfo         string   `protobuf:"bytes,9,opt,name=mortgage_info,json=mortgageInfo,proto3" json:"mortgage_info,omitempty"`
	DocumentPhoto        string   `protobuf:"bytes,10,opt,name=document_photo,json=documentPhoto,proto3" json:"document_photo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HouseTransactionInfo) Reset()         { *m = HouseTransactionInfo{} }
func (m *HouseTransactionInfo) String() string { return proto.CompactTextString(m) }
func (*HouseTransactionInfo) ProtoMessage()    {}
func (*HouseTransactionInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_af7531cd28e23262, []int{3}
}

func (m *HouseTransactionInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HouseTransactionInfo.Unmarshal(m, b)
}
func (m *HouseTransactionInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HouseTransactionInfo.Marshal(b, m, deterministic)
}
func (m *HouseTransactionInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HouseTransactionInfo.Merge(m, src)
}
func (m *HouseTransactionInfo) XXX_Size() int {
	return xxx_messageInfo_HouseTransactionInfo.Size(m)
}
func (m *HouseTransactionInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_HouseTransactionInfo.DiscardUnknown(m)
}

var xxx_messageInfo_HouseTransactionInfo proto.InternalMessageInfo

func (m *HouseTransactionInfo) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *HouseTransactionInfo) GetHouseNo() string {
	if m != nil {
		return m.HouseNo
	}
	return ""
}

func (m *HouseTransactionInfo) GetListingTime() string {
	if m != nil {
		return m.ListingTime
	}
	return ""
}

func (m *HouseTransactionInfo) GetTradingAuthority() string {
	if m != nil {
		return m.TradingAuthority
	}
	return ""
}

func (m *HouseTransactionInfo) GetLastTransaction() string {
	if m != nil {
		return m.LastTransaction
	}
	return ""
}

func (m *HouseTransactionInfo) GetHousingPurposes() string {
	if m != nil {
		return m.HousingPurposes
	}
	return ""
}

func (m *HouseTransactionInfo) GetHouseYear() string {
	if m != nil {
		return m.HouseYear
	}
	return ""
}

func (m *HouseTransactionInfo) GetPropertyRights() string {
	if m != nil {
		return m.PropertyRights
	}
	return ""
}

func (m *HouseTransactionInfo) GetMortgageInfo() string {
	if m != nil {
		return m.MortgageInfo
	}
	return ""
}

func (m *HouseTransactionInfo) GetDocumentPhoto() string {
	if m != nil {
		return m.DocumentPhoto
	}
	return ""
}

// 照片
type HousePic struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	HouseNo              string   `protobuf:"bytes,2,opt,name=house_no,json=houseNo,proto3" json:"house_no,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	PicSmallUrl          string   `protobuf:"bytes,4,opt,name=pic_small_url,json=picSmallUrl,proto3" json:"pic_small_url,omitempty"`
	PicNormalUrl         string   `protobuf:"bytes,5,opt,name=pic_normal_url,json=picNormalUrl,proto3" json:"pic_normal_url,omitempty"`
	PicLargeUrl          string   `protobuf:"bytes,6,opt,name=pic_large_url,json=picLargeUrl,proto3" json:"pic_large_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HousePic) Reset()         { *m = HousePic{} }
func (m *HousePic) String() string { return proto.CompactTextString(m) }
func (*HousePic) ProtoMessage()    {}
func (*HousePic) Descriptor() ([]byte, []int) {
	return fileDescriptor_af7531cd28e23262, []int{4}
}

func (m *HousePic) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HousePic.Unmarshal(m, b)
}
func (m *HousePic) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HousePic.Marshal(b, m, deterministic)
}
func (m *HousePic) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HousePic.Merge(m, src)
}
func (m *HousePic) XXX_Size() int {
	return xxx_messageInfo_HousePic.Size(m)
}
func (m *HousePic) XXX_DiscardUnknown() {
	xxx_messageInfo_HousePic.DiscardUnknown(m)
}

var xxx_messageInfo_HousePic proto.InternalMessageInfo

func (m *HousePic) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *HousePic) GetHouseNo() string {
	if m != nil {
		return m.HouseNo
	}
	return ""
}

func (m *HousePic) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *HousePic) GetPicSmallUrl() string {
	if m != nil {
		return m.PicSmallUrl
	}
	return ""
}

func (m *HousePic) GetPicNormalUrl() string {
	if m != nil {
		return m.PicNormalUrl
	}
	return ""
}

func (m *HousePic) GetPicLargeUrl() string {
	if m != nil {
		return m.PicLargeUrl
	}
	return ""
}

func init() {
	proto.RegisterType((*House)(nil), "lianjia.House")
	proto.RegisterType((*HouseInfo)(nil), "lianjia.HouseInfo")
	proto.RegisterType((*HouseBaseInfo)(nil), "lianjia.HouseBaseInfo")
	proto.RegisterType((*HouseTransactionInfo)(nil), "lianjia.HouseTransactionInfo")
	proto.RegisterType((*HousePic)(nil), "lianjia.HousePic")
}

func init() { proto.RegisterFile("crawler/lianjia.proto", fileDescriptor_af7531cd28e23262) }

var fileDescriptor_af7531cd28e23262 = []byte{
	// 782 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x95, 0xdf, 0x8e, 0xe3, 0x34,
	0x14, 0xc6, 0xd5, 0x76, 0xdb, 0x69, 0x4e, 0xfa, 0x67, 0xd6, 0x1a, 0xaa, 0x0e, 0xa3, 0x85, 0x6e,
	0x61, 0xc5, 0x20, 0xa4, 0x2e, 0x2c, 0xf7, 0x48, 0xcc, 0x15, 0x48, 0x68, 0x55, 0xb5, 0xdd, 0x0b,
	0xb8, 0x89, 0x9c, 0xc4, 0x6d, 0x8d, 0xdc, 0x38, 0x72, 0x9c, 0x45, 0x79, 0x19, 0x9e, 0x80, 0xd7,
	0xe0, 0xa9, 0xb8, 0x41, 0xe7, 0xd8, 0xee, 0xa6, 0xd2, 0xdc, 0xf4, 0xaa, 0x39, 0xbf, 0xe3, 0xef,
	0xd8, 0xe7, 0x7c, 0x75, 0x02, 0x9f, 0x65, 0x86, 0xff, 0xa5, 0x84, 0x79, 0xab, 0x24, 0x2f, 0xfe,
	0x94, 0x7c, 0x55, 0x1a, 0x6d, 0x35, 0xbb, 0xf1, 0xe1, 0xf2, 0x9f, 0x2e, 0xf4, 0x7f, 0xd1, 0x75,
	0x25, 0xd8, 0x04, 0xba, 0x32, 0x9f, 0x77, 0x16, 0x9d, 0xc7, 0x68, 0xd3, 0x95, 0x39, 0xbb, 0x87,
	0xe1, 0x11, 0x13, 0x49, 0xa1, 0xe7, 0x5d, 0xa2, 0x37, 0x14, 0xbf, 0xd7, 0xec, 0x07, 0x00, 0x97,
	0x92, 0xc5, 0x5e, 0xcf, 0x7b, 0x8b, 0xce, 0x63, 0xfc, 0x8e, 0xad, 0xc2, 0x0e, 0x54, 0xee, 0xd7,
	0x62, 0xaf, 0x37, 0xd1, 0x31, 0x3c, 0xb2, 0x9f, 0x60, 0xea, 0x24, 0x29, 0x0f, 0xba, 0x17, 0xa4,
	0x9b, 0x5d, 0xea, 0x9e, 0xb8, 0xd7, 0x8e, 0x8f, 0xed, 0x90, 0x6d, 0x61, 0xe6, 0xf4, 0xd6, 0xf0,
	0xa2, 0xe2, 0x99, 0x95, 0xba, 0x70, 0x65, 0xfa, 0x54, 0xe6, 0xd5, 0x65, 0x99, 0xdd, 0xa7, 0x55,
	0x54, 0xed, 0xee, 0xf8, 0x0c, 0x65, 0xdf, 0x87, 0x3e, 0x4a, 0x99, 0x55, 0xf3, 0xc1, 0xa2, 0xf7,
	0x18, 0xbf, 0x7b, 0x79, 0x59, 0x68, 0x2d, 0x33, 0xdf, 0xc6, 0x5a, 0x66, 0xd5, 0xf2, 0xbf, 0x2e,
	0x44, 0xe7, 0xfe, 0xae, 0x19, 0xd9, 0x2d, 0xf4, 0x6a, 0xa3, 0x68, 0x56, 0xd1, 0x06, 0x1f, 0xd9,
	0x1d, 0xf4, 0xad, 0xb4, 0x4a, 0xd0, 0x1c, 0xa2, 0x8d, 0x0b, 0xd8, 0x03, 0x44, 0x55, 0x9d, 0x26,
	0x2e, 0xd3, 0xa7, 0xcc, 0xb0, 0xaa, 0xd3, 0x1d, 0x25, 0x67, 0x30, 0x30, 0xe2, 0x20, 0x75, 0x31,
	0x1f, 0x50, 0xc6, 0x47, 0xec, 0x4b, 0x88, 0xad, 0xb6, 0x5c, 0x25, 0xa5, 0x91, 0x99, 0x98, 0xdf,
	0x50, 0x12, 0x08, 0xad, 0x91, 0xb0, 0x57, 0x00, 0x75, 0x21, 0xad, 0xcf, 0x0f, 0x29, 0x1f, 0x21,
	0x71, 0xe9, 0x07, 0x88, 0x8c, 0xd6, 0x27, 0x37, 0xcf, 0xc8, 0x6d, 0x8a, 0x80, 0x9a, 0xbc, 0x07,
	0x7a, 0x4e, 0xaa, 0x3a, 0x9d, 0x83, 0x6b, 0x0a, 0xe3, 0x6d, 0x9d, 0xa2, 0xce, 0x36, 0xa5, 0xb7,
	0x33, 0x76, 0x3a, 0x04, 0x41, 0x47, 0x49, 0xd4, 0x8d, 0x9c, 0x0e, 0x63, 0xaf, 0xe3, 0x46, 0x70,
	0xa7, 0x1b, 0x3b, 0x1d, 0x82, 0xa0, 0xa3, 0x24, 0xea, 0x26, 0x4e, 0x87, 0xf1, 0xb6, 0x4e, 0x97,
	0x7f, 0xf7, 0x60, 0x7c, 0xf1, 0x2f, 0xb9, 0xc6, 0x81, 0x19, 0x0c, 0x14, 0x6f, 0x74, 0x6d, 0xbd,
	0x09, 0x3e, 0x42, 0x1f, 0xf6, 0x4a, 0x6b, 0x13, 0x7c, 0xa0, 0x00, 0x27, 0x46, 0xa7, 0x48, 0x6b,
	0xa9, 0x72, 0x6f, 0x04, 0x1d, 0xfa, 0x09, 0x01, 0x7b, 0x0d, 0xa3, 0xca, 0x9a, 0x3a, 0xb3, 0x09,
	0x95, 0xf7, 0x7e, 0xc4, 0x8e, 0xb9, 0xfb, 0x14, 0x2a, 0xc8, 0xa2, 0x10, 0xc6, 0x7b, 0x12, 0xb9,
	0x2e, 0x0b, 0x41, 0x1b, 0x50, 0xed, 0x04, 0x87, 0x12, 0x2c, 0x21, 0xb2, 0x6b, 0x4a, 0xc1, 0x18,
	0xbc, 0xd8, 0xf3, 0x4c, 0x78, 0x37, 0xe8, 0xb9, 0xb5, 0xa9, 0x3b, 0x15, 0xb4, 0x37, 0x75, 0xe7,
	0xfa, 0x02, 0x20, 0x17, 0x99, 0x36, 0x1c, 0xff, 0xe3, 0xde, 0x92, 0x16, 0x61, 0x6f, 0x60, 0x22,
	0x94, 0xf8, 0xc8, 0xad, 0x36, 0x09, 0x21, 0x6f, 0xcd, 0x38, 0xd0, 0x0d, 0x42, 0xf6, 0x39, 0x0c,
	0x03, 0x08, 0xfe, 0x84, 0x18, 0x73, 0xa5, 0xd1, 0xa5, 0x30, 0xb6, 0xf1, 0xfe, 0x9c, 0x63, 0xbc,
	0x1e, 0x77, 0xcf, 0xdd, 0xbf, 0x6b, 0x7c, 0x7a, 0x0d, 0x23, 0x25, 0x2b, 0x2b, 0x8b, 0x43, 0x62,
	0xe5, 0x49, 0x78, 0xb7, 0x62, 0xcf, 0x76, 0xf2, 0x24, 0xd8, 0x77, 0xf0, 0xd2, 0x1a, 0x9e, 0xe3,
	0x12, 0x5e, 0xdb, 0xa3, 0x36, 0xd2, 0x36, 0xde, 0xbe, 0x5b, 0x9f, 0xf8, 0x39, 0x70, 0xf6, 0x2d,
	0xdc, 0x2a, 0x5e, 0xd9, 0xf6, 0x8b, 0xc3, 0xfb, 0x39, 0x45, 0xde, 0x3a, 0x29, 0x2e, 0xc5, 0x53,
	0x60, 0xdd, 0xb2, 0x36, 0xa5, 0xae, 0x44, 0xe5, 0x9d, 0x9d, 0x7a, 0xbe, 0xf6, 0x18, 0xed, 0x73,
	0x0d, 0x34, 0x82, 0x9f, 0xdd, 0x25, 0xf2, 0xbb, 0xe0, 0x86, 0x7d, 0x03, 0xd3, 0x30, 0x94, 0xc4,
	0xc8, 0xc3, 0xd1, 0x56, 0xde, 0xe2, 0x49, 0xc0, 0x1b, 0xa2, 0xec, 0x2b, 0x18, 0x9f, 0xb4, 0xb1,
	0x07, 0x7e, 0x10, 0xed, 0xeb, 0x37, 0x0a, 0x90, 0xa6, 0xf7, 0x06, 0x26, 0xb9, 0xce, 0xea, 0x93,
	0x28, 0x6c, 0x52, 0x1e, 0xb5, 0xd5, 0xde, 0xfa, 0x71, 0xa0, 0x6b, 0x84, 0xcb, 0x7f, 0x3b, 0x30,
	0x0c, 0x2f, 0xad, 0x6b, 0x26, 0xbe, 0x80, 0x38, 0x17, 0x55, 0x66, 0x64, 0x49, 0xc3, 0xf1, 0x03,
	0x6f, 0x21, 0xb6, 0x84, 0x71, 0x29, 0xb3, 0xa4, 0x3a, 0x71, 0xa5, 0x12, 0x7c, 0x8f, 0xb9, 0x61,
	0xc7, 0xa5, 0xcc, 0xb6, 0xc8, 0x3e, 0x18, 0xc5, 0xbe, 0x86, 0x09, 0xae, 0x29, 0xb4, 0x39, 0x71,
	0xb7, 0xc8, 0x4d, 0x79, 0x54, 0xca, 0xec, 0x3d, 0x41, 0x5c, 0xe5, 0x2b, 0x29, 0x6e, 0x0e, 0x82,
	0x16, 0x0d, 0xce, 0x95, 0x7e, 0x43, 0xf6, 0xc1, 0xa8, 0xa7, 0x87, 0x3f, 0xee, 0x0f, 0x7a, 0xd5,
	0x34, 0xb5, 0x58, 0xe5, 0xe2, 0xe3, 0xdb, 0xf0, 0x05, 0xa3, 0x2f, 0x57, 0x3a, 0xa0, 0x9f, 0x1f,
	0xff, 0x0f, 0x00, 0x00, 0xff, 0xff, 0xac, 0x70, 0xb4, 0x19, 0xd9, 0x06, 0x00, 0x00,
}
