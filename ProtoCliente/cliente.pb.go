// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: cliente.proto

package ProtoCliente

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

type Descripcion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdProducto  int32  `protobuf:"varint,1,opt,name=idProducto,proto3" json:"idProducto,omitempty"`
	Descripcion string `protobuf:"bytes,2,opt,name=descripcion,proto3" json:"descripcion,omitempty"`
	Cantidad    int32  `protobuf:"varint,3,opt,name=cantidad,proto3" json:"cantidad,omitempty"`
}

func (x *Descripcion) Reset() {
	*x = Descripcion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cliente_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Descripcion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Descripcion) ProtoMessage() {}

func (x *Descripcion) ProtoReflect() protoreflect.Message {
	mi := &file_cliente_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Descripcion.ProtoReflect.Descriptor instead.
func (*Descripcion) Descriptor() ([]byte, []int) {
	return file_cliente_proto_rawDescGZIP(), []int{0}
}

func (x *Descripcion) GetIdProducto() int32 {
	if x != nil {
		return x.IdProducto
	}
	return 0
}

func (x *Descripcion) GetDescripcion() string {
	if x != nil {
		return x.Descripcion
	}
	return ""
}

func (x *Descripcion) GetCantidad() int32 {
	if x != nil {
		return x.Cantidad
	}
	return 0
}

type PedidoSolicitudRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Descripciones []*Descripcion `protobuf:"bytes,1,rep,name=descripciones,proto3" json:"descripciones,omitempty"`
}

func (x *PedidoSolicitudRequest) Reset() {
	*x = PedidoSolicitudRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cliente_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PedidoSolicitudRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PedidoSolicitudRequest) ProtoMessage() {}

func (x *PedidoSolicitudRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cliente_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PedidoSolicitudRequest.ProtoReflect.Descriptor instead.
func (*PedidoSolicitudRequest) Descriptor() ([]byte, []int) {
	return file_cliente_proto_rawDescGZIP(), []int{1}
}

func (x *PedidoSolicitudRequest) GetDescripciones() []*Descripcion {
	if x != nil {
		return x.Descripciones
	}
	return nil
}

type Pedido struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdPedido          int32                   `protobuf:"varint,1,opt,name=idPedido,proto3" json:"idPedido,omitempty"`
	EstadoRestaurante string                  `protobuf:"bytes,2,opt,name=estadoRestaurante,proto3" json:"estadoRestaurante,omitempty"`
	EstadoRepartidor  string                  `protobuf:"bytes,3,opt,name=estadoRepartidor,proto3" json:"estadoRepartidor,omitempty"`
	PedidoSolicitud   *PedidoSolicitudRequest `protobuf:"bytes,5,opt,name=pedidoSolicitud,proto3" json:"pedidoSolicitud,omitempty"`
}

func (x *Pedido) Reset() {
	*x = Pedido{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cliente_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pedido) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pedido) ProtoMessage() {}

func (x *Pedido) ProtoReflect() protoreflect.Message {
	mi := &file_cliente_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pedido.ProtoReflect.Descriptor instead.
func (*Pedido) Descriptor() ([]byte, []int) {
	return file_cliente_proto_rawDescGZIP(), []int{2}
}

func (x *Pedido) GetIdPedido() int32 {
	if x != nil {
		return x.IdPedido
	}
	return 0
}

func (x *Pedido) GetEstadoRestaurante() string {
	if x != nil {
		return x.EstadoRestaurante
	}
	return ""
}

func (x *Pedido) GetEstadoRepartidor() string {
	if x != nil {
		return x.EstadoRepartidor
	}
	return ""
}

func (x *Pedido) GetPedidoSolicitud() *PedidoSolicitudRequest {
	if x != nil {
		return x.PedidoSolicitud
	}
	return nil
}

type PedidoSolicitudResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mensaje string  `protobuf:"bytes,1,opt,name=mensaje,proto3" json:"mensaje,omitempty"`
	Pedido  *Pedido `protobuf:"bytes,2,opt,name=pedido,proto3" json:"pedido,omitempty"`
}

func (x *PedidoSolicitudResponse) Reset() {
	*x = PedidoSolicitudResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cliente_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PedidoSolicitudResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PedidoSolicitudResponse) ProtoMessage() {}

func (x *PedidoSolicitudResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cliente_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PedidoSolicitudResponse.ProtoReflect.Descriptor instead.
func (*PedidoSolicitudResponse) Descriptor() ([]byte, []int) {
	return file_cliente_proto_rawDescGZIP(), []int{3}
}

func (x *PedidoSolicitudResponse) GetMensaje() string {
	if x != nil {
		return x.Mensaje
	}
	return ""
}

func (x *PedidoSolicitudResponse) GetPedido() *Pedido {
	if x != nil {
		return x.Pedido
	}
	return nil
}

type IdPedido struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *IdPedido) Reset() {
	*x = IdPedido{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cliente_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdPedido) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdPedido) ProtoMessage() {}

func (x *IdPedido) ProtoReflect() protoreflect.Message {
	mi := &file_cliente_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdPedido.ProtoReflect.Descriptor instead.
func (*IdPedido) Descriptor() ([]byte, []int) {
	return file_cliente_proto_rawDescGZIP(), []int{4}
}

func (x *IdPedido) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_cliente_proto protoreflect.FileDescriptor

var file_cliente_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x04, 0x67, 0x72, 0x70, 0x63, 0x22, 0x6b, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x63, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x64, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x69, 0x64, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x6f, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x63,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x63, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x6e, 0x74, 0x69, 0x64,
	0x61, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x63, 0x61, 0x6e, 0x74, 0x69, 0x64,
	0x61, 0x64, 0x22, 0x51, 0x0a, 0x16, 0x50, 0x65, 0x64, 0x69, 0x64, 0x6f, 0x53, 0x6f, 0x6c, 0x69,
	0x63, 0x69, 0x74, 0x75, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x37, 0x0a, 0x0d,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x63, 0x69, 0x6f, 0x6e, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x63, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x63,
	0x69, 0x6f, 0x6e, 0x65, 0x73, 0x22, 0xc6, 0x01, 0x0a, 0x06, 0x50, 0x65, 0x64, 0x69, 0x64, 0x6f,
	0x12, 0x1a, 0x0a, 0x08, 0x69, 0x64, 0x50, 0x65, 0x64, 0x69, 0x64, 0x6f, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x69, 0x64, 0x50, 0x65, 0x64, 0x69, 0x64, 0x6f, 0x12, 0x2c, 0x0a, 0x11,
	0x65, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x65, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x52,
	0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x65, 0x73,
	0x74, 0x61, 0x64, 0x6f, 0x52, 0x65, 0x70, 0x61, 0x72, 0x74, 0x69, 0x64, 0x6f, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x65, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x52, 0x65, 0x70, 0x61,
	0x72, 0x74, 0x69, 0x64, 0x6f, 0x72, 0x12, 0x46, 0x0a, 0x0f, 0x70, 0x65, 0x64, 0x69, 0x64, 0x6f,
	0x53, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x75, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x65, 0x64, 0x69, 0x64, 0x6f, 0x53, 0x6f, 0x6c,
	0x69, 0x63, 0x69, 0x74, 0x75, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0f, 0x70,
	0x65, 0x64, 0x69, 0x64, 0x6f, 0x53, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x75, 0x64, 0x22, 0x59,
	0x0a, 0x17, 0x50, 0x65, 0x64, 0x69, 0x64, 0x6f, 0x53, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x75,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x6e,
	0x73, 0x61, 0x6a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x6e, 0x73,
	0x61, 0x6a, 0x65, 0x12, 0x24, 0x0a, 0x06, 0x70, 0x65, 0x64, 0x69, 0x64, 0x6f, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x65, 0x64, 0x69, 0x64,
	0x6f, 0x52, 0x06, 0x70, 0x65, 0x64, 0x69, 0x64, 0x6f, 0x22, 0x1a, 0x0a, 0x08, 0x69, 0x64, 0x50,
	0x65, 0x64, 0x69, 0x64, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x69, 0x64, 0x32, 0xff, 0x02, 0x0a, 0x0e, 0x50, 0x65, 0x64, 0x69, 0x64, 0x6f,
	0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x0a, 0x50, 0x6f, 0x73, 0x74,
	0x50, 0x65, 0x64, 0x69, 0x64, 0x6f, 0x12, 0x1c, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x65,
	0x64, 0x69, 0x64, 0x6f, 0x53, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x75, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x65, 0x64, 0x69,
	0x64, 0x6f, 0x12, 0x33, 0x0a, 0x15, 0x50, 0x6f, 0x73, 0x74, 0x50, 0x65, 0x64, 0x69, 0x64, 0x6f,
	0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x12, 0x0c, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x50, 0x65, 0x64, 0x69, 0x64, 0x6f, 0x1a, 0x0c, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x50, 0x65, 0x64, 0x69, 0x64, 0x6f, 0x12, 0x36, 0x0a, 0x16, 0x50, 0x6f, 0x73, 0x74, 0x45,
	0x6e, 0x74, 0x72, 0x65, 0x67, 0x61, 0x72, 0x52, 0x65, 0x70, 0x61, 0x72, 0x74, 0x69, 0x64, 0x6f,
	0x72, 0x12, 0x0e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x69, 0x64, 0x50, 0x65, 0x64, 0x69, 0x64,
	0x6f, 0x1a, 0x0c, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x65, 0x64, 0x69, 0x64, 0x6f, 0x12,
	0x2e, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74,
	0x65, 0x12, 0x0e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x69, 0x64, 0x50, 0x65, 0x64, 0x69, 0x64,
	0x6f, 0x1a, 0x0c, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x65, 0x64, 0x69, 0x64, 0x6f, 0x12,
	0x32, 0x0a, 0x14, 0x50, 0x6f, 0x73, 0x74, 0x50, 0x65, 0x64, 0x69, 0x64, 0x6f, 0x52, 0x65, 0x70,
	0x61, 0x72, 0x74, 0x69, 0x64, 0x6f, 0x72, 0x12, 0x0c, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x50,
	0x65, 0x64, 0x69, 0x64, 0x6f, 0x1a, 0x0c, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x65, 0x64,
	0x69, 0x64, 0x6f, 0x12, 0x33, 0x0a, 0x13, 0x50, 0x6f, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x65,
	0x67, 0x61, 0x72, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x65, 0x12, 0x0e, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x69, 0x64, 0x50, 0x65, 0x64, 0x69, 0x64, 0x6f, 0x1a, 0x0c, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x50, 0x65, 0x64, 0x69, 0x64, 0x6f, 0x12, 0x2d, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x70, 0x61, 0x72, 0x74, 0x69, 0x64, 0x6f, 0x72, 0x12, 0x0e, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x69, 0x64, 0x50, 0x65, 0x64, 0x69, 0x64, 0x6f, 0x1a, 0x0c, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x50, 0x65, 0x64, 0x69, 0x64, 0x6f, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4a, 0x6f, 0x72, 0x67, 0x65, 0x41, 0x6d, 0x62, 0x72, 0x6f,
	0x63, 0x69, 0x6f, 0x2f, 0x47, 0x52, 0x50, 0x43, 0x2d, 0x54, 0x45, 0x53, 0x54, 0x2f, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_cliente_proto_rawDescOnce sync.Once
	file_cliente_proto_rawDescData = file_cliente_proto_rawDesc
)

func file_cliente_proto_rawDescGZIP() []byte {
	file_cliente_proto_rawDescOnce.Do(func() {
		file_cliente_proto_rawDescData = protoimpl.X.CompressGZIP(file_cliente_proto_rawDescData)
	})
	return file_cliente_proto_rawDescData
}

var file_cliente_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_cliente_proto_goTypes = []interface{}{
	(*Descripcion)(nil),             // 0: grpc.Descripcion
	(*PedidoSolicitudRequest)(nil),  // 1: grpc.PedidoSolicitudRequest
	(*Pedido)(nil),                  // 2: grpc.Pedido
	(*PedidoSolicitudResponse)(nil), // 3: grpc.PedidoSolicitudResponse
	(*IdPedido)(nil),                // 4: grpc.idPedido
}
var file_cliente_proto_depIdxs = []int32{
	0,  // 0: grpc.PedidoSolicitudRequest.descripciones:type_name -> grpc.Descripcion
	1,  // 1: grpc.Pedido.pedidoSolicitud:type_name -> grpc.PedidoSolicitudRequest
	2,  // 2: grpc.PedidoSolicitudResponse.pedido:type_name -> grpc.Pedido
	1,  // 3: grpc.PedidosService.PostPedido:input_type -> grpc.PedidoSolicitudRequest
	2,  // 4: grpc.PedidosService.PostPedidoRestaurante:input_type -> grpc.Pedido
	4,  // 5: grpc.PedidosService.PostEntregarRepartidor:input_type -> grpc.idPedido
	4,  // 6: grpc.PedidosService.GetRestaurante:input_type -> grpc.idPedido
	2,  // 7: grpc.PedidosService.PostPedidoRepartidor:input_type -> grpc.Pedido
	4,  // 8: grpc.PedidosService.PostEntregarCliente:input_type -> grpc.idPedido
	4,  // 9: grpc.PedidosService.GetRepartidor:input_type -> grpc.idPedido
	2,  // 10: grpc.PedidosService.PostPedido:output_type -> grpc.Pedido
	2,  // 11: grpc.PedidosService.PostPedidoRestaurante:output_type -> grpc.Pedido
	2,  // 12: grpc.PedidosService.PostEntregarRepartidor:output_type -> grpc.Pedido
	2,  // 13: grpc.PedidosService.GetRestaurante:output_type -> grpc.Pedido
	2,  // 14: grpc.PedidosService.PostPedidoRepartidor:output_type -> grpc.Pedido
	2,  // 15: grpc.PedidosService.PostEntregarCliente:output_type -> grpc.Pedido
	2,  // 16: grpc.PedidosService.GetRepartidor:output_type -> grpc.Pedido
	10, // [10:17] is the sub-list for method output_type
	3,  // [3:10] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_cliente_proto_init() }
func file_cliente_proto_init() {
	if File_cliente_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cliente_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Descripcion); i {
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
		file_cliente_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PedidoSolicitudRequest); i {
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
		file_cliente_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pedido); i {
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
		file_cliente_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PedidoSolicitudResponse); i {
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
		file_cliente_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdPedido); i {
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
			RawDescriptor: file_cliente_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cliente_proto_goTypes,
		DependencyIndexes: file_cliente_proto_depIdxs,
		MessageInfos:      file_cliente_proto_msgTypes,
	}.Build()
	File_cliente_proto = out.File
	file_cliente_proto_rawDesc = nil
	file_cliente_proto_goTypes = nil
	file_cliente_proto_depIdxs = nil
}
