// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v4.0.0
// source: recipe.proto

package recipe

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Category struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Category) Reset() {
	*x = Category{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recipe_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Category) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Category) ProtoMessage() {}

func (x *Category) ProtoReflect() protoreflect.Message {
	mi := &file_recipe_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Category.ProtoReflect.Descriptor instead.
func (*Category) Descriptor() ([]byte, []int) {
	return file_recipe_proto_rawDescGZIP(), []int{0}
}

func (x *Category) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type RecipeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MainIngredient string `protobuf:"bytes,1,opt,name=mainIngredient,proto3" json:"mainIngredient,omitempty"`
}

func (x *RecipeRequest) Reset() {
	*x = RecipeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recipe_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecipeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecipeRequest) ProtoMessage() {}

func (x *RecipeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_recipe_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecipeRequest.ProtoReflect.Descriptor instead.
func (*RecipeRequest) Descriptor() ([]byte, []int) {
	return file_recipe_proto_rawDescGZIP(), []int{1}
}

func (x *RecipeRequest) GetMainIngredient() string {
	if x != nil {
		return x.MainIngredient
	}
	return ""
}

type RecipeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RecipeName string    `protobuf:"bytes,1,opt,name=recipeName,proto3" json:"recipeName,omitempty"`
	Rating     int32     `protobuf:"varint,2,opt,name=rating,proto3" json:"rating,omitempty"`
	Category   *Category `protobuf:"bytes,3,opt,name=category,proto3" json:"category,omitempty"`
}

func (x *RecipeReply) Reset() {
	*x = RecipeReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recipe_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecipeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecipeReply) ProtoMessage() {}

func (x *RecipeReply) ProtoReflect() protoreflect.Message {
	mi := &file_recipe_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecipeReply.ProtoReflect.Descriptor instead.
func (*RecipeReply) Descriptor() ([]byte, []int) {
	return file_recipe_proto_rawDescGZIP(), []int{2}
}

func (x *RecipeReply) GetRecipeName() string {
	if x != nil {
		return x.RecipeName
	}
	return ""
}

func (x *RecipeReply) GetRating() int32 {
	if x != nil {
		return x.Rating
	}
	return 0
}

func (x *RecipeReply) GetCategory() *Category {
	if x != nil {
		return x.Category
	}
	return nil
}

var File_recipe_proto protoreflect.FileDescriptor

var file_recipe_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x22, 0x1e, 0x0a, 0x08, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x37, 0x0a, 0x0d, 0x52, 0x65, 0x63, 0x69, 0x70, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x6d, 0x61, 0x69, 0x6e, 0x49,
	0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x22,
	0x73, 0x0a, 0x0b, 0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1e,
	0x0a, 0x0a, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x2c, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x72, 0x65, 0x63, 0x69, 0x70,
	0x65, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x32, 0x4a, 0x0a, 0x0d, 0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x69,
	0x70, 0x65, 0x12, 0x15, 0x2e, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x2e, 0x52, 0x65, 0x63, 0x69,
	0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x72, 0x65, 0x63, 0x69,
	0x70, 0x65, 0x2e, 0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_recipe_proto_rawDescOnce sync.Once
	file_recipe_proto_rawDescData = file_recipe_proto_rawDesc
)

func file_recipe_proto_rawDescGZIP() []byte {
	file_recipe_proto_rawDescOnce.Do(func() {
		file_recipe_proto_rawDescData = protoimpl.X.CompressGZIP(file_recipe_proto_rawDescData)
	})
	return file_recipe_proto_rawDescData
}

var file_recipe_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_recipe_proto_goTypes = []interface{}{
	(*Category)(nil),      // 0: recipe.Category
	(*RecipeRequest)(nil), // 1: recipe.RecipeRequest
	(*RecipeReply)(nil),   // 2: recipe.RecipeReply
}
var file_recipe_proto_depIdxs = []int32{
	0, // 0: recipe.RecipeReply.category:type_name -> recipe.Category
	1, // 1: recipe.RecipeService.GetRecipe:input_type -> recipe.RecipeRequest
	2, // 2: recipe.RecipeService.GetRecipe:output_type -> recipe.RecipeReply
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_recipe_proto_init() }
func file_recipe_proto_init() {
	if File_recipe_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_recipe_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Category); i {
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
		file_recipe_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecipeRequest); i {
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
		file_recipe_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecipeReply); i {
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
			RawDescriptor: file_recipe_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_recipe_proto_goTypes,
		DependencyIndexes: file_recipe_proto_depIdxs,
		MessageInfos:      file_recipe_proto_msgTypes,
	}.Build()
	File_recipe_proto = out.File
	file_recipe_proto_rawDesc = nil
	file_recipe_proto_goTypes = nil
	file_recipe_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// RecipeServiceClient is the client API for RecipeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RecipeServiceClient interface {
	GetRecipe(ctx context.Context, in *RecipeRequest, opts ...grpc.CallOption) (*RecipeReply, error)
}

type recipeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRecipeServiceClient(cc grpc.ClientConnInterface) RecipeServiceClient {
	return &recipeServiceClient{cc}
}

func (c *recipeServiceClient) GetRecipe(ctx context.Context, in *RecipeRequest, opts ...grpc.CallOption) (*RecipeReply, error) {
	out := new(RecipeReply)
	err := c.cc.Invoke(ctx, "/recipe.RecipeService/GetRecipe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RecipeServiceServer is the server API for RecipeService service.
type RecipeServiceServer interface {
	GetRecipe(context.Context, *RecipeRequest) (*RecipeReply, error)
}

// UnimplementedRecipeServiceServer can be embedded to have forward compatible implementations.
type UnimplementedRecipeServiceServer struct {
}

func (*UnimplementedRecipeServiceServer) GetRecipe(context.Context, *RecipeRequest) (*RecipeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRecipe not implemented")
}

func RegisterRecipeServiceServer(s *grpc.Server, srv RecipeServiceServer) {
	s.RegisterService(&_RecipeService_serviceDesc, srv)
}

func _RecipeService_GetRecipe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecipeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecipeServiceServer).GetRecipe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/recipe.RecipeService/GetRecipe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecipeServiceServer).GetRecipe(ctx, req.(*RecipeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RecipeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "recipe.RecipeService",
	HandlerType: (*RecipeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRecipe",
			Handler:    _RecipeService_GetRecipe_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "recipe.proto",
}
