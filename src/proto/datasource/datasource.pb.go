// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: datasource.proto

package datasource

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

//Request for Create Datasource
type AddRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DataDomain  string `protobuf:"bytes,1,opt,name=data_domain,json=dataDomain,proto3" json:"data_domain,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Type        string `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	Version     string `protobuf:"bytes,5,opt,name=version,proto3" json:"version,omitempty"`
	Host        string `protobuf:"bytes,7,opt,name=host,proto3" json:"host,omitempty"`
	Port        string `protobuf:"bytes,8,opt,name=port,proto3" json:"port,omitempty"`
	User        string `protobuf:"bytes,9,opt,name=user,proto3" json:"user,omitempty"`
	Password    string `protobuf:"bytes,10,opt,name=password,proto3" json:"password,omitempty"`
	Policy_Id   int64  `protobuf:"varint,11,opt,name=policy_Id,json=policyId,proto3" json:"policy_Id,omitempty"`
}

func (x *AddRequest) Reset() {
	*x = AddRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datasource_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRequest) ProtoMessage() {}

func (x *AddRequest) ProtoReflect() protoreflect.Message {
	mi := &file_datasource_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRequest.ProtoReflect.Descriptor instead.
func (*AddRequest) Descriptor() ([]byte, []int) {
	return file_datasource_proto_rawDescGZIP(), []int{0}
}

func (x *AddRequest) GetDataDomain() string {
	if x != nil {
		return x.DataDomain
	}
	return ""
}

func (x *AddRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AddRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *AddRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *AddRequest) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *AddRequest) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *AddRequest) GetPort() string {
	if x != nil {
		return x.Port
	}
	return ""
}

func (x *AddRequest) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *AddRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *AddRequest) GetPolicy_Id() int64 {
	if x != nil {
		return x.Policy_Id
	}
	return 0
}

//Request for List Datasource
type ListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	DataDomain []string `protobuf:"bytes,2,rep,name=data_domain,json=dataDomain,proto3" json:"data_domain,omitempty"`
	Type       []string `protobuf:"bytes,3,rep,name=type,proto3" json:"type,omitempty"`
	First      int64    `protobuf:"varint,4,opt,name=first,proto3" json:"first,omitempty"`
	Last       int64    `protobuf:"varint,5,opt,name=last,proto3" json:"last,omitempty"`
	Limit      int64    `protobuf:"varint,6,opt,name=limit,proto3" json:"limit,omitempty"`
	Count      bool     `protobuf:"varint,7,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *ListRequest) Reset() {
	*x = ListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datasource_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequest) ProtoMessage() {}

func (x *ListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_datasource_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRequest.ProtoReflect.Descriptor instead.
func (*ListRequest) Descriptor() ([]byte, []int) {
	return file_datasource_proto_rawDescGZIP(), []int{1}
}

func (x *ListRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ListRequest) GetDataDomain() []string {
	if x != nil {
		return x.DataDomain
	}
	return nil
}

func (x *ListRequest) GetType() []string {
	if x != nil {
		return x.Type
	}
	return nil
}

func (x *ListRequest) GetFirst() int64 {
	if x != nil {
		return x.First
	}
	return 0
}

func (x *ListRequest) GetLast() int64 {
	if x != nil {
		return x.Last
	}
	return 0
}

func (x *ListRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListRequest) GetCount() bool {
	if x != nil {
		return x.Count
	}
	return false
}

//Response for List Datasource
type ListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ListAllDatasources []*ListAll `protobuf:"bytes,1,rep,name=list_all_datasources,json=listAllDatasources,proto3" json:"list_all_datasources,omitempty"`
	Count              int64      `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *ListResponse) Reset() {
	*x = ListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datasource_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResponse) ProtoMessage() {}

func (x *ListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_datasource_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResponse.ProtoReflect.Descriptor instead.
func (*ListResponse) Descriptor() ([]byte, []int) {
	return file_datasource_proto_rawDescGZIP(), []int{2}
}

func (x *ListResponse) GetListAllDatasources() []*ListAll {
	if x != nil {
		return x.ListAllDatasources
	}
	return nil
}

func (x *ListResponse) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

//Request for Delete Datasource
type DeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          []string `protobuf:"bytes,1,rep,name=id,proto3" json:"id,omitempty"`
	IsDeleteAll bool     `protobuf:"varint,2,opt,name=isDeleteAll,proto3" json:"isDeleteAll,omitempty"`
}

func (x *DeleteRequest) Reset() {
	*x = DeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datasource_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_datasource_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRequest.ProtoReflect.Descriptor instead.
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return file_datasource_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteRequest) GetId() []string {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *DeleteRequest) GetIsDeleteAll() bool {
	if x != nil {
		return x.IsDeleteAll
	}
	return false
}

//Response for  Datasource
type MessageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *MessageResponse) Reset() {
	*x = MessageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datasource_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageResponse) ProtoMessage() {}

func (x *MessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_datasource_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageResponse.ProtoReflect.Descriptor instead.
func (*MessageResponse) Descriptor() ([]byte, []int) {
	return file_datasource_proto_rawDescGZIP(), []int{4}
}

func (x *MessageResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ListAll struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Datadomain  string `protobuf:"bytes,2,opt,name=datadomain,proto3" json:"datadomain,omitempty"`
	Name        string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Type        string `protobuf:"bytes,5,opt,name=type,proto3" json:"type,omitempty"`
	Version     string `protobuf:"bytes,6,opt,name=version,proto3" json:"version,omitempty"`
	Key         string `protobuf:"bytes,7,opt,name=key,proto3" json:"key,omitempty"`
	CreatedAt   string `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *ListAll) Reset() {
	*x = ListAll{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datasource_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAll) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAll) ProtoMessage() {}

func (x *ListAll) ProtoReflect() protoreflect.Message {
	mi := &file_datasource_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAll.ProtoReflect.Descriptor instead.
func (*ListAll) Descriptor() ([]byte, []int) {
	return file_datasource_proto_rawDescGZIP(), []int{5}
}

func (x *ListAll) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ListAll) GetDatadomain() string {
	if x != nil {
		return x.Datadomain
	}
	return ""
}

func (x *ListAll) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ListAll) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ListAll) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *ListAll) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *ListAll) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *ListAll) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

//Request for datasource log
type DatasourceLogRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Datasource string   `protobuf:"bytes,1,opt,name=datasource,proto3" json:"datasource,omitempty"`
	Database   string   `protobuf:"bytes,2,opt,name=database,proto3" json:"database,omitempty"`
	Table      string   `protobuf:"bytes,3,opt,name=table,proto3" json:"table,omitempty"`
	Columns    []string `protobuf:"bytes,4,rep,name=columns,proto3" json:"columns,omitempty"`
}

func (x *DatasourceLogRequest) Reset() {
	*x = DatasourceLogRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datasource_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DatasourceLogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DatasourceLogRequest) ProtoMessage() {}

func (x *DatasourceLogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_datasource_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DatasourceLogRequest.ProtoReflect.Descriptor instead.
func (*DatasourceLogRequest) Descriptor() ([]byte, []int) {
	return file_datasource_proto_rawDescGZIP(), []int{6}
}

func (x *DatasourceLogRequest) GetDatasource() string {
	if x != nil {
		return x.Datasource
	}
	return ""
}

func (x *DatasourceLogRequest) GetDatabase() string {
	if x != nil {
		return x.Database
	}
	return ""
}

func (x *DatasourceLogRequest) GetTable() string {
	if x != nil {
		return x.Table
	}
	return ""
}

func (x *DatasourceLogRequest) GetColumns() []string {
	if x != nil {
		return x.Columns
	}
	return nil
}

//datasource log Response Structure
type DatasourceLogResponseStruct struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Datasource   string   `protobuf:"bytes,1,opt,name=datasource,proto3" json:"datasource,omitempty"`
	Database     string   `protobuf:"bytes,2,opt,name=database,proto3" json:"database,omitempty"`
	Table        string   `protobuf:"bytes,3,opt,name=table,proto3" json:"table,omitempty"`
	Column       string   `protobuf:"bytes,4,opt,name=column,proto3" json:"column,omitempty"`
	Tags         []string `protobuf:"bytes,5,rep,name=tags,proto3" json:"tags,omitempty"`
	Classes      []string `protobuf:"bytes,6,rep,name=classes,proto3" json:"classes,omitempty"`
	LastScanTime string   `protobuf:"bytes,7,opt,name=lastScanTime,proto3" json:"lastScanTime,omitempty"`
}

func (x *DatasourceLogResponseStruct) Reset() {
	*x = DatasourceLogResponseStruct{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datasource_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DatasourceLogResponseStruct) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DatasourceLogResponseStruct) ProtoMessage() {}

func (x *DatasourceLogResponseStruct) ProtoReflect() protoreflect.Message {
	mi := &file_datasource_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DatasourceLogResponseStruct.ProtoReflect.Descriptor instead.
func (*DatasourceLogResponseStruct) Descriptor() ([]byte, []int) {
	return file_datasource_proto_rawDescGZIP(), []int{7}
}

func (x *DatasourceLogResponseStruct) GetDatasource() string {
	if x != nil {
		return x.Datasource
	}
	return ""
}

func (x *DatasourceLogResponseStruct) GetDatabase() string {
	if x != nil {
		return x.Database
	}
	return ""
}

func (x *DatasourceLogResponseStruct) GetTable() string {
	if x != nil {
		return x.Table
	}
	return ""
}

func (x *DatasourceLogResponseStruct) GetColumn() string {
	if x != nil {
		return x.Column
	}
	return ""
}

func (x *DatasourceLogResponseStruct) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *DatasourceLogResponseStruct) GetClasses() []string {
	if x != nil {
		return x.Classes
	}
	return nil
}

func (x *DatasourceLogResponseStruct) GetLastScanTime() string {
	if x != nil {
		return x.LastScanTime
	}
	return ""
}

//log Response
type DatasourceLogResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DatasourceLog []*DatasourceLogResponseStruct `protobuf:"bytes,1,rep,name=datasourceLog,proto3" json:"datasourceLog,omitempty"`
}

func (x *DatasourceLogResponse) Reset() {
	*x = DatasourceLogResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datasource_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DatasourceLogResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DatasourceLogResponse) ProtoMessage() {}

func (x *DatasourceLogResponse) ProtoReflect() protoreflect.Message {
	mi := &file_datasource_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DatasourceLogResponse.ProtoReflect.Descriptor instead.
func (*DatasourceLogResponse) Descriptor() ([]byte, []int) {
	return file_datasource_proto_rawDescGZIP(), []int{8}
}

func (x *DatasourceLogResponse) GetDatasourceLog() []*DatasourceLogResponseStruct {
	if x != nil {
		return x.DatasourceLog
	}
	return nil
}

//Datasource Name
type DatasourceName struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DatasourceName) Reset() {
	*x = DatasourceName{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datasource_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DatasourceName) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DatasourceName) ProtoMessage() {}

func (x *DatasourceName) ProtoReflect() protoreflect.Message {
	mi := &file_datasource_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DatasourceName.ProtoReflect.Descriptor instead.
func (*DatasourceName) Descriptor() ([]byte, []int) {
	return file_datasource_proto_rawDescGZIP(), []int{9}
}

func (x *DatasourceName) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

//ScanMessage
type ScanResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int64  `protobuf:"varint,1,opt,name=StatusCode,proto3" json:"StatusCode,omitempty"`
	Message    string `protobuf:"bytes,2,opt,name=Message,proto3" json:"Message,omitempty"`
}

func (x *ScanResponse) Reset() {
	*x = ScanResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datasource_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScanResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScanResponse) ProtoMessage() {}

func (x *ScanResponse) ProtoReflect() protoreflect.Message {
	mi := &file_datasource_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScanResponse.ProtoReflect.Descriptor instead.
func (*ScanResponse) Descriptor() ([]byte, []int) {
	return file_datasource_proto_rawDescGZIP(), []int{10}
}

func (x *ScanResponse) GetStatusCode() int64 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *ScanResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_datasource_proto protoreflect.FileDescriptor

var file_datasource_proto_rawDesc = []byte{
	0x0a, 0x10, 0x64, 0x61, 0x74, 0x61, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0x86,
	0x02, 0x0a, 0x0a, 0x61, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a,
	0x0b, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x5f, 0x49, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x49, 0x64, 0x22, 0xac, 0x01, 0x0a, 0x0b, 0x6c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x64,
	0x61, 0x74, 0x61, 0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x72, 0x73, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x66, 0x69, 0x72, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x61, 0x73, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x6c, 0x61, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x6b, 0x0a, 0x0c, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x14, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x61,
	0x6c, 0x6c, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x12, 0x6c, 0x69, 0x73, 0x74, 0x41,
	0x6c, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x22, 0x41, 0x0a, 0x0d, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x69, 0x73, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x41, 0x6c, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x41, 0x6c, 0x6c, 0x22, 0x2b, 0x0a, 0x0f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0xce, 0x01, 0x0a, 0x07, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x1e, 0x0a, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x22, 0x82, 0x01, 0x0a, 0x14, 0x64, 0x61, 0x74, 0x61, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a,
	0x0a, 0x64, 0x61, 0x74, 0x61, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x61, 0x62,
	0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x07, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x22, 0xd9, 0x01, 0x0a, 0x1b, 0x64, 0x61,
	0x74, 0x61, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x61, 0x74,
	0x61, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64,
	0x61, 0x74, 0x61, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x61, 0x74,
	0x61, 0x62, 0x61, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x61, 0x74,
	0x61, 0x62, 0x61, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x63,
	0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6f, 0x6c,
	0x75, 0x6d, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6c, 0x61, 0x73, 0x73,
	0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x65,
	0x73, 0x12, 0x22, 0x0a, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x53, 0x63, 0x61, 0x6e, 0x54, 0x69, 0x6d,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x53, 0x63, 0x61,
	0x6e, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x66, 0x0a, 0x15, 0x64, 0x61, 0x74, 0x61, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d,
	0x0a, 0x0d, 0x64, 0x61, 0x74, 0x61, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c, 0x6f, 0x67, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c, 0x6f, 0x67,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x0d,
	0x64, 0x61, 0x74, 0x61, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c, 0x6f, 0x67, 0x22, 0x24, 0x0a,
	0x0e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x48, 0x0a, 0x0c, 0x73, 0x63, 0x61, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xf7, 0x02,
	0x0a, 0x0a, 0x44, 0x61, 0x74, 0x61, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x44, 0x0a, 0x0d,
	0x41, 0x64, 0x64, 0x44, 0x61, 0x74, 0x61, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x16, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x61, 0x64, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x43, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x12, 0x17, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x44, 0x61, 0x74, 0x61, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x19, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x54, 0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x44, 0x61, 0x74, 0x61, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x12, 0x20, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c, 0x6f, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c, 0x6f,
	0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x04, 0x53, 0x63, 0x61,
	0x6e, 0x12, 0x1a, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x18, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x73, 0x63, 0x61, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x61, 0x67, 0x65, 0x2d, 0x69,
	0x6f, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_datasource_proto_rawDescOnce sync.Once
	file_datasource_proto_rawDescData = file_datasource_proto_rawDesc
)

func file_datasource_proto_rawDescGZIP() []byte {
	file_datasource_proto_rawDescOnce.Do(func() {
		file_datasource_proto_rawDescData = protoimpl.X.CompressGZIP(file_datasource_proto_rawDescData)
	})
	return file_datasource_proto_rawDescData
}

var file_datasource_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_datasource_proto_goTypes = []interface{}{
	(*AddRequest)(nil),                  // 0: datasource.addRequest
	(*ListRequest)(nil),                 // 1: datasource.listRequest
	(*ListResponse)(nil),                // 2: datasource.listResponse
	(*DeleteRequest)(nil),               // 3: datasource.deleteRequest
	(*MessageResponse)(nil),             // 4: datasource.messageResponse
	(*ListAll)(nil),                     // 5: datasource.ListAll
	(*DatasourceLogRequest)(nil),        // 6: datasource.datasourceLogRequest
	(*DatasourceLogResponseStruct)(nil), // 7: datasource.datasourceLogResponseStruct
	(*DatasourceLogResponse)(nil),       // 8: datasource.datasourceLogResponse
	(*DatasourceName)(nil),              // 9: datasource.datasourceName
	(*ScanResponse)(nil),                // 10: datasource.scanResponse
}
var file_datasource_proto_depIdxs = []int32{
	5,  // 0: datasource.listResponse.list_all_datasources:type_name -> datasource.ListAll
	7,  // 1: datasource.datasourceLogResponse.datasourceLog:type_name -> datasource.datasourceLogResponseStruct
	0,  // 2: datasource.Datasource.AddDatasource:input_type -> datasource.addRequest
	1,  // 3: datasource.Datasource.ListDatasource:input_type -> datasource.listRequest
	3,  // 4: datasource.Datasource.DeleteDatasource:input_type -> datasource.deleteRequest
	6,  // 5: datasource.Datasource.LogDatasource:input_type -> datasource.datasourceLogRequest
	9,  // 6: datasource.Datasource.Scan:input_type -> datasource.datasourceName
	4,  // 7: datasource.Datasource.AddDatasource:output_type -> datasource.messageResponse
	2,  // 8: datasource.Datasource.ListDatasource:output_type -> datasource.listResponse
	4,  // 9: datasource.Datasource.DeleteDatasource:output_type -> datasource.messageResponse
	8,  // 10: datasource.Datasource.LogDatasource:output_type -> datasource.datasourceLogResponse
	10, // 11: datasource.Datasource.Scan:output_type -> datasource.scanResponse
	7,  // [7:12] is the sub-list for method output_type
	2,  // [2:7] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_datasource_proto_init() }
func file_datasource_proto_init() {
	if File_datasource_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_datasource_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddRequest); i {
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
		file_datasource_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRequest); i {
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
		file_datasource_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListResponse); i {
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
		file_datasource_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRequest); i {
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
		file_datasource_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageResponse); i {
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
		file_datasource_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAll); i {
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
		file_datasource_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DatasourceLogRequest); i {
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
		file_datasource_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DatasourceLogResponseStruct); i {
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
		file_datasource_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DatasourceLogResponse); i {
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
		file_datasource_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DatasourceName); i {
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
		file_datasource_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScanResponse); i {
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
			RawDescriptor: file_datasource_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_datasource_proto_goTypes,
		DependencyIndexes: file_datasource_proto_depIdxs,
		MessageInfos:      file_datasource_proto_msgTypes,
	}.Build()
	File_datasource_proto = out.File
	file_datasource_proto_rawDesc = nil
	file_datasource_proto_goTypes = nil
	file_datasource_proto_depIdxs = nil
}
