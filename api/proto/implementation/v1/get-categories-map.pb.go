// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: v1/get-categories-map.proto

package v1

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

type GetCategoriesMapRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId      string   `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	CategoryIds []string `protobuf:"bytes,2,rep,name=categoryIds,proto3" json:"categoryIds,omitempty"`
}

func (x *GetCategoriesMapRequest) Reset() {
	*x = GetCategoriesMapRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_get_categories_map_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCategoriesMapRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCategoriesMapRequest) ProtoMessage() {}

func (x *GetCategoriesMapRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_get_categories_map_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCategoriesMapRequest.ProtoReflect.Descriptor instead.
func (*GetCategoriesMapRequest) Descriptor() ([]byte, []int) {
	return file_v1_get_categories_map_proto_rawDescGZIP(), []int{0}
}

func (x *GetCategoriesMapRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *GetCategoriesMapRequest) GetCategoryIds() []string {
	if x != nil {
		return x.CategoryIds
	}
	return nil
}

type GetCategoriesMapResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Categories map[string]*Category `protobuf:"bytes,1,rep,name=categories,proto3" json:"categories,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *GetCategoriesMapResponse) Reset() {
	*x = GetCategoriesMapResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_get_categories_map_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCategoriesMapResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCategoriesMapResponse) ProtoMessage() {}

func (x *GetCategoriesMapResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_get_categories_map_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCategoriesMapResponse.ProtoReflect.Descriptor instead.
func (*GetCategoriesMapResponse) Descriptor() ([]byte, []int) {
	return file_v1_get_categories_map_proto_rawDescGZIP(), []int{1}
}

func (x *GetCategoriesMapResponse) GetCategories() map[string]*Category {
	if x != nil {
		return x.Categories
	}
	return nil
}

var File_v1_get_categories_map_proto protoreflect.FileDescriptor

var file_v1_get_categories_map_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x2d, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x69, 0x65, 0x73, 0x2d, 0x6d, 0x61, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76,
	0x31, 0x1a, 0x1c, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x2d, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x53, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73,
	0x4d, 0x61, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x49, 0x64, 0x73, 0x22, 0xb5, 0x01, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x4d, 0x61, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x4c, 0x0a, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x4d, 0x61, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x1a,
	0x4b, 0x0a, 0x0f, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x22, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x3f, 0x5a, 0x3d,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x65, 0x70, 0x68, 0x69,
	0x73, 0x74, 0x6f, 0x6c, 0x69, 0x65, 0x2f, 0x63, 0x68, 0x65, 0x66, 0x62, 0x6f, 0x6f, 0x6b, 0x2d,
	0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_get_categories_map_proto_rawDescOnce sync.Once
	file_v1_get_categories_map_proto_rawDescData = file_v1_get_categories_map_proto_rawDesc
)

func file_v1_get_categories_map_proto_rawDescGZIP() []byte {
	file_v1_get_categories_map_proto_rawDescOnce.Do(func() {
		file_v1_get_categories_map_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_get_categories_map_proto_rawDescData)
	})
	return file_v1_get_categories_map_proto_rawDescData
}

var file_v1_get_categories_map_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_v1_get_categories_map_proto_goTypes = []interface{}{
	(*GetCategoriesMapRequest)(nil),  // 0: v1.GetCategoriesMapRequest
	(*GetCategoriesMapResponse)(nil), // 1: v1.GetCategoriesMapResponse
	nil,                              // 2: v1.GetCategoriesMapResponse.CategoriesEntry
	(*Category)(nil),                 // 3: v1.Category
}
var file_v1_get_categories_map_proto_depIdxs = []int32{
	2, // 0: v1.GetCategoriesMapResponse.categories:type_name -> v1.GetCategoriesMapResponse.CategoriesEntry
	3, // 1: v1.GetCategoriesMapResponse.CategoriesEntry.value:type_name -> v1.Category
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_v1_get_categories_map_proto_init() }
func file_v1_get_categories_map_proto_init() {
	if File_v1_get_categories_map_proto != nil {
		return
	}
	file_v1_get_user_categories_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_v1_get_categories_map_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCategoriesMapRequest); i {
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
		file_v1_get_categories_map_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCategoriesMapResponse); i {
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
			RawDescriptor: file_v1_get_categories_map_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1_get_categories_map_proto_goTypes,
		DependencyIndexes: file_v1_get_categories_map_proto_depIdxs,
		MessageInfos:      file_v1_get_categories_map_proto_msgTypes,
	}.Build()
	File_v1_get_categories_map_proto = out.File
	file_v1_get_categories_map_proto_rawDesc = nil
	file_v1_get_categories_map_proto_goTypes = nil
	file_v1_get_categories_map_proto_depIdxs = nil
}