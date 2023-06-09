// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.3
// source: hub.io/api/plugin/v1alpha1/plugin.proto

package v1alpha1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	types "hub.io/api/types"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Plugin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Domain       string   `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain,omitempty"`                                 // plugin domain
	Name         string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`                                     // plugin name
	Description  string   `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`                       // plugin description
	AuthType     string   `protobuf:"bytes,5,opt,name=auth_type,json=authType,proto3" json:"auth_type,omitempty"`             // plugin auth type
	LogoUrl      string   `protobuf:"bytes,6,opt,name=logo_url,json=logoUrl,proto3" json:"logo_url,omitempty"`                // plugin logo url
	ContactEmail string   `protobuf:"bytes,7,opt,name=contact_email,json=contactEmail,proto3" json:"contact_email,omitempty"` // plugin contact email
	Organization string   `protobuf:"bytes,8,opt,name=organization,proto3" json:"organization,omitempty"`                     // plugin organization
	ApiType      string   `protobuf:"bytes,9,opt,name=api_type,json=apiType,proto3" json:"api_type,omitempty"`                // plugin api type
	ApiUrl       string   `protobuf:"bytes,10,opt,name=api_url,json=apiUrl,proto3" json:"api_url,omitempty"`                  // plugin api url
	Labels       []string `protobuf:"bytes,11,rep,name=labels,proto3" json:"labels,omitempty"`                                // plugin label
	State        string   `protobuf:"bytes,12,opt,name=state,proto3" json:"state,omitempty"`                                  // plugin state,published or unpublished
	InstallNum   int32    `protobuf:"varint,13,opt,name=install_num,json=installNum,proto3" json:"install_num,omitempty"`     // plugin install num
	Score        float32  `protobuf:"fixed32,14,opt,name=score,proto3" json:"score,omitempty"`                                // plugin score
	Heat         int32    `protobuf:"varint,15,opt,name=heat,proto3" json:"heat,omitempty"`                                   // plugin heat
	CreatedAt    string   `protobuf:"bytes,16,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt    string   `protobuf:"bytes,17,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Plugin) Reset() {
	*x = Plugin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hub_io_api_plugin_v1alpha1_plugin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Plugin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Plugin) ProtoMessage() {}

func (x *Plugin) ProtoReflect() protoreflect.Message {
	mi := &file_hub_io_api_plugin_v1alpha1_plugin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Plugin.ProtoReflect.Descriptor instead.
func (*Plugin) Descriptor() ([]byte, []int) {
	return file_hub_io_api_plugin_v1alpha1_plugin_proto_rawDescGZIP(), []int{0}
}

func (x *Plugin) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Plugin) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *Plugin) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Plugin) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Plugin) GetAuthType() string {
	if x != nil {
		return x.AuthType
	}
	return ""
}

func (x *Plugin) GetLogoUrl() string {
	if x != nil {
		return x.LogoUrl
	}
	return ""
}

func (x *Plugin) GetContactEmail() string {
	if x != nil {
		return x.ContactEmail
	}
	return ""
}

func (x *Plugin) GetOrganization() string {
	if x != nil {
		return x.Organization
	}
	return ""
}

func (x *Plugin) GetApiType() string {
	if x != nil {
		return x.ApiType
	}
	return ""
}

func (x *Plugin) GetApiUrl() string {
	if x != nil {
		return x.ApiUrl
	}
	return ""
}

func (x *Plugin) GetLabels() []string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *Plugin) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *Plugin) GetInstallNum() int32 {
	if x != nil {
		return x.InstallNum
	}
	return 0
}

func (x *Plugin) GetScore() float32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *Plugin) GetHeat() int32 {
	if x != nil {
		return x.Heat
	}
	return 0
}

func (x *Plugin) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Plugin) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type ListPluginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page     int32 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// order by
	OrderBy types.OrderBy `protobuf:"varint,3,opt,name=order_by,json=orderBy,proto3,enum=hub.io.api.types.OrderBy" json:"order_by,omitempty"`
	// fuzzy search
	FuzzyName string `protobuf:"bytes,4,opt,name=fuzzy_name,json=fuzzyName,proto3" json:"fuzzy_name,omitempty"`
	// sort by field name
	SortByFieldName string `protobuf:"bytes,5,opt,name=sort_by_field_name,json=sortByFieldName,proto3" json:"sort_by_field_name,omitempty"`
}

func (x *ListPluginRequest) Reset() {
	*x = ListPluginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hub_io_api_plugin_v1alpha1_plugin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPluginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPluginRequest) ProtoMessage() {}

func (x *ListPluginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hub_io_api_plugin_v1alpha1_plugin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPluginRequest.ProtoReflect.Descriptor instead.
func (*ListPluginRequest) Descriptor() ([]byte, []int) {
	return file_hub_io_api_plugin_v1alpha1_plugin_proto_rawDescGZIP(), []int{1}
}

func (x *ListPluginRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListPluginRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListPluginRequest) GetOrderBy() types.OrderBy {
	if x != nil {
		return x.OrderBy
	}
	return types.OrderBy(0)
}

func (x *ListPluginRequest) GetFuzzyName() string {
	if x != nil {
		return x.FuzzyName
	}
	return ""
}

func (x *ListPluginRequest) GetSortByFieldName() string {
	if x != nil {
		return x.SortByFieldName
	}
	return ""
}

type ListPluginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page *types.Page `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"`
	Item []*Plugin   `protobuf:"bytes,2,rep,name=item,proto3" json:"item,omitempty"`
}

func (x *ListPluginResponse) Reset() {
	*x = ListPluginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hub_io_api_plugin_v1alpha1_plugin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPluginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPluginResponse) ProtoMessage() {}

func (x *ListPluginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_hub_io_api_plugin_v1alpha1_plugin_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPluginResponse.ProtoReflect.Descriptor instead.
func (*ListPluginResponse) Descriptor() ([]byte, []int) {
	return file_hub_io_api_plugin_v1alpha1_plugin_proto_rawDescGZIP(), []int{2}
}

func (x *ListPluginResponse) GetPage() *types.Page {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *ListPluginResponse) GetItem() []*Plugin {
	if x != nil {
		return x.Item
	}
	return nil
}

type CreatePluginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Domain string `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`
	Label  string `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
}

func (x *CreatePluginRequest) Reset() {
	*x = CreatePluginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hub_io_api_plugin_v1alpha1_plugin_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePluginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePluginRequest) ProtoMessage() {}

func (x *CreatePluginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hub_io_api_plugin_v1alpha1_plugin_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePluginRequest.ProtoReflect.Descriptor instead.
func (*CreatePluginRequest) Descriptor() ([]byte, []int) {
	return file_hub_io_api_plugin_v1alpha1_plugin_proto_rawDescGZIP(), []int{3}
}

func (x *CreatePluginRequest) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *CreatePluginRequest) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

type CreatePluginScoreRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PluginId uint32  `protobuf:"varint,1,opt,name=plugin_id,json=pluginId,proto3" json:"plugin_id,omitempty"`
	Score    float32 `protobuf:"fixed32,2,opt,name=score,proto3" json:"score,omitempty"`
	Comments string  `protobuf:"bytes,3,opt,name=comments,proto3" json:"comments,omitempty"`
}

func (x *CreatePluginScoreRequest) Reset() {
	*x = CreatePluginScoreRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hub_io_api_plugin_v1alpha1_plugin_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePluginScoreRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePluginScoreRequest) ProtoMessage() {}

func (x *CreatePluginScoreRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hub_io_api_plugin_v1alpha1_plugin_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePluginScoreRequest.ProtoReflect.Descriptor instead.
func (*CreatePluginScoreRequest) Descriptor() ([]byte, []int) {
	return file_hub_io_api_plugin_v1alpha1_plugin_proto_rawDescGZIP(), []int{4}
}

func (x *CreatePluginScoreRequest) GetPluginId() uint32 {
	if x != nil {
		return x.PluginId
	}
	return 0
}

func (x *CreatePluginScoreRequest) GetScore() float32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *CreatePluginScoreRequest) GetComments() string {
	if x != nil {
		return x.Comments
	}
	return ""
}

var File_hub_io_api_plugin_v1alpha1_plugin_proto protoreflect.FileDescriptor

var file_hub_io_api_plugin_v1alpha1_plugin_proto_rawDesc = []byte{
	0x0a, 0x27, 0x68, 0x75, 0x62, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x68, 0x75, 0x62, 0x2e, 0x69,
	0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x2b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65,
	0x6e, 0x2d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1b, 0x68, 0x75, 0x62, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xd2, 0x03, 0x0a, 0x06, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x68,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x75, 0x74,
	0x68, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x6f, 0x67, 0x6f, 0x5f, 0x75, 0x72,
	0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x6f, 0x67, 0x6f, 0x55, 0x72, 0x6c,
	0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x5f, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x22, 0x0a, 0x0c, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6f, 0x72, 0x67,
	0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x70, 0x69,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x70, 0x69,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x61, 0x70, 0x69, 0x5f, 0x75, 0x72, 0x6c, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x70, 0x69, 0x55, 0x72, 0x6c, 0x12, 0x16, 0x0a,
	0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x69,
	0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x4e, 0x75, 0x6d, 0x12, 0x14, 0x0a, 0x05,
	0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x73, 0x63, 0x6f,
	0x72, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x65, 0x61, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x68, 0x65, 0x61, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x22, 0xd8, 0x01, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x34, 0x0a, 0x08, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x5f, 0x62, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e,
	0x68, 0x75, 0x62, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x42,
	0x79, 0x12, 0x26, 0x0a, 0x0a, 0x66, 0x75, 0x7a, 0x7a, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x00, 0x52, 0x09,
	0x66, 0x75, 0x7a, 0x7a, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x34, 0x0a, 0x12, 0x73, 0x6f, 0x72,
	0x74, 0x5f, 0x62, 0x79, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x00, 0x52, 0x0f,
	0x73, 0x6f, 0x72, 0x74, 0x42, 0x79, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x22,
	0x78, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x68, 0x75, 0x62, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x12, 0x36, 0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x22, 0x2e, 0x68, 0x75, 0x62, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x50, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x52, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x22, 0x55, 0x0a, 0x13, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1f, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x12, 0x1d, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x00, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c,
	0x22, 0x8c, 0x01, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a,
	0x09, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x20, 0x00, 0x52, 0x08, 0x70, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x02, 0x42, 0x0f, 0xfa, 0x42, 0x0c, 0x0a, 0x0a, 0x1d, 0x00, 0x00, 0xa0, 0x40, 0x2d, 0x00,
	0x00, 0x00, 0x00, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x23, 0x0a, 0x08, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42,
	0x04, 0x72, 0x02, 0x10, 0x00, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x42,
	0x25, 0x5a, 0x23, 0x68, 0x75, 0x62, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_hub_io_api_plugin_v1alpha1_plugin_proto_rawDescOnce sync.Once
	file_hub_io_api_plugin_v1alpha1_plugin_proto_rawDescData = file_hub_io_api_plugin_v1alpha1_plugin_proto_rawDesc
)

func file_hub_io_api_plugin_v1alpha1_plugin_proto_rawDescGZIP() []byte {
	file_hub_io_api_plugin_v1alpha1_plugin_proto_rawDescOnce.Do(func() {
		file_hub_io_api_plugin_v1alpha1_plugin_proto_rawDescData = protoimpl.X.CompressGZIP(file_hub_io_api_plugin_v1alpha1_plugin_proto_rawDescData)
	})
	return file_hub_io_api_plugin_v1alpha1_plugin_proto_rawDescData
}

var file_hub_io_api_plugin_v1alpha1_plugin_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_hub_io_api_plugin_v1alpha1_plugin_proto_goTypes = []interface{}{
	(*Plugin)(nil),                   // 0: hub.io.api.plugin.v1alpha1.Plugin
	(*ListPluginRequest)(nil),        // 1: hub.io.api.plugin.v1alpha1.ListPluginRequest
	(*ListPluginResponse)(nil),       // 2: hub.io.api.plugin.v1alpha1.ListPluginResponse
	(*CreatePluginRequest)(nil),      // 3: hub.io.api.plugin.v1alpha1.CreatePluginRequest
	(*CreatePluginScoreRequest)(nil), // 4: hub.io.api.plugin.v1alpha1.CreatePluginScoreRequest
	(types.OrderBy)(0),               // 5: hub.io.api.types.OrderBy
	(*types.Page)(nil),               // 6: hub.io.api.types.Page
}
var file_hub_io_api_plugin_v1alpha1_plugin_proto_depIdxs = []int32{
	5, // 0: hub.io.api.plugin.v1alpha1.ListPluginRequest.order_by:type_name -> hub.io.api.types.OrderBy
	6, // 1: hub.io.api.plugin.v1alpha1.ListPluginResponse.page:type_name -> hub.io.api.types.Page
	0, // 2: hub.io.api.plugin.v1alpha1.ListPluginResponse.item:type_name -> hub.io.api.plugin.v1alpha1.Plugin
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_hub_io_api_plugin_v1alpha1_plugin_proto_init() }
func file_hub_io_api_plugin_v1alpha1_plugin_proto_init() {
	if File_hub_io_api_plugin_v1alpha1_plugin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_hub_io_api_plugin_v1alpha1_plugin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Plugin); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_hub_io_api_plugin_v1alpha1_plugin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPluginRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_hub_io_api_plugin_v1alpha1_plugin_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPluginResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_hub_io_api_plugin_v1alpha1_plugin_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePluginRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_hub_io_api_plugin_v1alpha1_plugin_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePluginScoreRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_hub_io_api_plugin_v1alpha1_plugin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_hub_io_api_plugin_v1alpha1_plugin_proto_goTypes,
		DependencyIndexes: file_hub_io_api_plugin_v1alpha1_plugin_proto_depIdxs,
		MessageInfos:      file_hub_io_api_plugin_v1alpha1_plugin_proto_msgTypes,
	}.Build()
	File_hub_io_api_plugin_v1alpha1_plugin_proto = out.File
	file_hub_io_api_plugin_v1alpha1_plugin_proto_rawDesc = nil
	file_hub_io_api_plugin_v1alpha1_plugin_proto_goTypes = nil
	file_hub_io_api_plugin_v1alpha1_plugin_proto_depIdxs = nil
}
