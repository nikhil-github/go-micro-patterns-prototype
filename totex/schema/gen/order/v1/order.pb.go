// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: order/v1/order.proto

package orderv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// CreateOrderRequest contains the data needed to create an order
type CreateOrderRequest struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	UserId          string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Items           []*OrderItem           `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
	ShippingAddress string                 `protobuf:"bytes,3,opt,name=shipping_address,json=shippingAddress,proto3" json:"shipping_address,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *CreateOrderRequest) Reset() {
	*x = CreateOrderRequest{}
	mi := &file_order_v1_order_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderRequest) ProtoMessage() {}

func (x *CreateOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_order_v1_order_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrderRequest.ProtoReflect.Descriptor instead.
func (*CreateOrderRequest) Descriptor() ([]byte, []int) {
	return file_order_v1_order_proto_rawDescGZIP(), []int{0}
}

func (x *CreateOrderRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CreateOrderRequest) GetItems() []*OrderItem {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *CreateOrderRequest) GetShippingAddress() string {
	if x != nil {
		return x.ShippingAddress
	}
	return ""
}

// CreateOrderResponse contains the result of order creation
type CreateOrderResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OrderId       string                 `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	UserId        string                 `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Items         []*OrderItem           `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Status        string                 `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	CreatedAt     string                 `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateOrderResponse) Reset() {
	*x = CreateOrderResponse{}
	mi := &file_order_v1_order_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderResponse) ProtoMessage() {}

func (x *CreateOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_order_v1_order_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrderResponse.ProtoReflect.Descriptor instead.
func (*CreateOrderResponse) Descriptor() ([]byte, []int) {
	return file_order_v1_order_proto_rawDescGZIP(), []int{1}
}

func (x *CreateOrderResponse) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *CreateOrderResponse) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CreateOrderResponse) GetItems() []*OrderItem {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *CreateOrderResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *CreateOrderResponse) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

// OrderItem represents a single item in an order
type OrderItem struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ProductId     string                 `protobuf:"bytes,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Quantity      int32                  `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Price         float64                `protobuf:"fixed64,3,opt,name=price,proto3" json:"price,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OrderItem) Reset() {
	*x = OrderItem{}
	mi := &file_order_v1_order_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OrderItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderItem) ProtoMessage() {}

func (x *OrderItem) ProtoReflect() protoreflect.Message {
	mi := &file_order_v1_order_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderItem.ProtoReflect.Descriptor instead.
func (*OrderItem) Descriptor() ([]byte, []int) {
	return file_order_v1_order_proto_rawDescGZIP(), []int{2}
}

func (x *OrderItem) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *OrderItem) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *OrderItem) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

var File_order_v1_order_proto protoreflect.FileDescriptor

const file_order_v1_order_proto_rawDesc = "" +
	"\n" +
	"\x14order/v1/order.proto\x12\border.v1\"\x83\x01\n" +
	"\x12CreateOrderRequest\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\tR\x06userId\x12)\n" +
	"\x05items\x18\x02 \x03(\v2\x13.order.v1.OrderItemR\x05items\x12)\n" +
	"\x10shipping_address\x18\x03 \x01(\tR\x0fshippingAddress\"\xab\x01\n" +
	"\x13CreateOrderResponse\x12\x19\n" +
	"\border_id\x18\x01 \x01(\tR\aorderId\x12\x17\n" +
	"\auser_id\x18\x02 \x01(\tR\x06userId\x12)\n" +
	"\x05items\x18\x03 \x03(\v2\x13.order.v1.OrderItemR\x05items\x12\x16\n" +
	"\x06status\x18\x04 \x01(\tR\x06status\x12\x1d\n" +
	"\n" +
	"created_at\x18\x05 \x01(\tR\tcreatedAt\"\\\n" +
	"\tOrderItem\x12\x1d\n" +
	"\n" +
	"product_id\x18\x01 \x01(\tR\tproductId\x12\x1a\n" +
	"\bquantity\x18\x02 \x01(\x05R\bquantity\x12\x14\n" +
	"\x05price\x18\x03 \x01(\x01R\x05price2Z\n" +
	"\fOrderService\x12J\n" +
	"\vCreateOrder\x12\x1c.order.v1.CreateOrderRequest\x1a\x1d.order.v1.CreateOrderResponseB\x90\x01\n" +
	"\fcom.order.v1B\n" +
	"OrderProtoP\x01Z3github.com/yourusername/schema/gen/order/v1;orderv1\xa2\x02\x03OXX\xaa\x02\bOrder.V1\xca\x02\bOrder\\V1\xe2\x02\x14Order\\V1\\GPBMetadata\xea\x02\tOrder::V1b\x06proto3"

var (
	file_order_v1_order_proto_rawDescOnce sync.Once
	file_order_v1_order_proto_rawDescData []byte
)

func file_order_v1_order_proto_rawDescGZIP() []byte {
	file_order_v1_order_proto_rawDescOnce.Do(func() {
		file_order_v1_order_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_order_v1_order_proto_rawDesc), len(file_order_v1_order_proto_rawDesc)))
	})
	return file_order_v1_order_proto_rawDescData
}

var file_order_v1_order_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_order_v1_order_proto_goTypes = []any{
	(*CreateOrderRequest)(nil),  // 0: order.v1.CreateOrderRequest
	(*CreateOrderResponse)(nil), // 1: order.v1.CreateOrderResponse
	(*OrderItem)(nil),           // 2: order.v1.OrderItem
}
var file_order_v1_order_proto_depIdxs = []int32{
	2, // 0: order.v1.CreateOrderRequest.items:type_name -> order.v1.OrderItem
	2, // 1: order.v1.CreateOrderResponse.items:type_name -> order.v1.OrderItem
	0, // 2: order.v1.OrderService.CreateOrder:input_type -> order.v1.CreateOrderRequest
	1, // 3: order.v1.OrderService.CreateOrder:output_type -> order.v1.CreateOrderResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_order_v1_order_proto_init() }
func file_order_v1_order_proto_init() {
	if File_order_v1_order_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_order_v1_order_proto_rawDesc), len(file_order_v1_order_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_order_v1_order_proto_goTypes,
		DependencyIndexes: file_order_v1_order_proto_depIdxs,
		MessageInfos:      file_order_v1_order_proto_msgTypes,
	}.Build()
	File_order_v1_order_proto = out.File
	file_order_v1_order_proto_goTypes = nil
	file_order_v1_order_proto_depIdxs = nil
}
