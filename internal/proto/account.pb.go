// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.4
// source: internal/proto/account.proto

package proto

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

type GetAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EthereumAddress string `protobuf:"bytes,1,opt,name=ethereum_address,json=ethereumAddress,proto3" json:"ethereum_address,omitempty"`
	CryptoSignature string `protobuf:"bytes,2,opt,name=crypto_signature,json=cryptoSignature,proto3" json:"crypto_signature,omitempty"`
}

func (x *GetAccountRequest) Reset() {
	*x = GetAccountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_account_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccountRequest) ProtoMessage() {}

func (x *GetAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_account_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccountRequest.ProtoReflect.Descriptor instead.
func (*GetAccountRequest) Descriptor() ([]byte, []int) {
	return file_internal_proto_account_proto_rawDescGZIP(), []int{0}
}

func (x *GetAccountRequest) GetEthereumAddress() string {
	if x != nil {
		return x.EthereumAddress
	}
	return ""
}

func (x *GetAccountRequest) GetCryptoSignature() string {
	if x != nil {
		return x.CryptoSignature
	}
	return ""
}

type GetAccountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GasTokenBalance string `protobuf:"bytes,1,opt,name=gas_token_balance,json=gasTokenBalance,proto3" json:"gas_token_balance,omitempty"`
	WalletNonce     uint64 `protobuf:"varint,2,opt,name=wallet_nonce,json=walletNonce,proto3" json:"wallet_nonce,omitempty"`
}

func (x *GetAccountResponse) Reset() {
	*x = GetAccountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_account_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAccountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccountResponse) ProtoMessage() {}

func (x *GetAccountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_account_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccountResponse.ProtoReflect.Descriptor instead.
func (*GetAccountResponse) Descriptor() ([]byte, []int) {
	return file_internal_proto_account_proto_rawDescGZIP(), []int{1}
}

func (x *GetAccountResponse) GetGasTokenBalance() string {
	if x != nil {
		return x.GasTokenBalance
	}
	return ""
}

func (x *GetAccountResponse) GetWalletNonce() uint64 {
	if x != nil {
		return x.WalletNonce
	}
	return 0
}

type GetAccountsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EthereumAddresses []string `protobuf:"bytes,1,rep,name=ethereum_addresses,json=ethereumAddresses,proto3" json:"ethereum_addresses,omitempty"`
	Erc20TokenAddress string   `protobuf:"bytes,2,opt,name=erc20_token_address,json=erc20TokenAddress,proto3" json:"erc20_token_address,omitempty"`
}

func (x *GetAccountsRequest) Reset() {
	*x = GetAccountsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_account_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAccountsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccountsRequest) ProtoMessage() {}

func (x *GetAccountsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_account_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccountsRequest.ProtoReflect.Descriptor instead.
func (*GetAccountsRequest) Descriptor() ([]byte, []int) {
	return file_internal_proto_account_proto_rawDescGZIP(), []int{2}
}

func (x *GetAccountsRequest) GetEthereumAddresses() []string {
	if x != nil {
		return x.EthereumAddresses
	}
	return nil
}

func (x *GetAccountsRequest) GetErc20TokenAddress() string {
	if x != nil {
		return x.Erc20TokenAddress
	}
	return ""
}

type GetAccountsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EthereumAddress string `protobuf:"bytes,1,opt,name=ethereum_address,json=ethereumAddress,proto3" json:"ethereum_address,omitempty"`
	Erc20Balance    string `protobuf:"bytes,2,opt,name=erc20_balance,json=erc20Balance,proto3" json:"erc20_balance,omitempty"`
}

func (x *GetAccountsResponse) Reset() {
	*x = GetAccountsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_account_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAccountsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccountsResponse) ProtoMessage() {}

func (x *GetAccountsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_account_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccountsResponse.ProtoReflect.Descriptor instead.
func (*GetAccountsResponse) Descriptor() ([]byte, []int) {
	return file_internal_proto_account_proto_rawDescGZIP(), []int{3}
}

func (x *GetAccountsResponse) GetEthereumAddress() string {
	if x != nil {
		return x.EthereumAddress
	}
	return ""
}

func (x *GetAccountsResponse) GetErc20Balance() string {
	if x != nil {
		return x.Erc20Balance
	}
	return ""
}

var File_internal_proto_account_proto protoreflect.FileDescriptor

var file_internal_proto_account_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x69, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x65, 0x74,
	0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x29, 0x0a, 0x10, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x5f,
	0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0f, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x22, 0x63, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x67, 0x61, 0x73, 0x5f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0f, 0x67, 0x61, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5f, 0x6e, 0x6f, 0x6e,
	0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74,
	0x4e, 0x6f, 0x6e, 0x63, 0x65, 0x22, 0x73, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x12, 0x65,
	0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x11, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75,
	0x6d, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x12, 0x2e, 0x0a, 0x13, 0x65, 0x72,
	0x63, 0x32, 0x30, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x65, 0x72, 0x63, 0x32, 0x30, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x65, 0x0a, 0x13, 0x47, 0x65,
	0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x29, 0x0a, 0x10, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x5f, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x65, 0x74, 0x68,
	0x65, 0x72, 0x65, 0x75, 0x6d, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x23, 0x0a, 0x0d,
	0x65, 0x72, 0x63, 0x32, 0x30, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x72, 0x63, 0x32, 0x30, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x32, 0x9d, 0x01, 0x0a, 0x0e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x41, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47,
	0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x30,
	0x01, 0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x79, 0x6f, 0x75, 0x72, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x2f, 0x79, 0x6f,
	0x75, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_proto_account_proto_rawDescOnce sync.Once
	file_internal_proto_account_proto_rawDescData = file_internal_proto_account_proto_rawDesc
)

func file_internal_proto_account_proto_rawDescGZIP() []byte {
	file_internal_proto_account_proto_rawDescOnce.Do(func() {
		file_internal_proto_account_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_proto_account_proto_rawDescData)
	})
	return file_internal_proto_account_proto_rawDescData
}

var file_internal_proto_account_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_internal_proto_account_proto_goTypes = []any{
	(*GetAccountRequest)(nil),   // 0: proto.GetAccountRequest
	(*GetAccountResponse)(nil),  // 1: proto.GetAccountResponse
	(*GetAccountsRequest)(nil),  // 2: proto.GetAccountsRequest
	(*GetAccountsResponse)(nil), // 3: proto.GetAccountsResponse
}
var file_internal_proto_account_proto_depIdxs = []int32{
	0, // 0: proto.AccountService.GetAccount:input_type -> proto.GetAccountRequest
	2, // 1: proto.AccountService.GetAccounts:input_type -> proto.GetAccountsRequest
	1, // 2: proto.AccountService.GetAccount:output_type -> proto.GetAccountResponse
	3, // 3: proto.AccountService.GetAccounts:output_type -> proto.GetAccountsResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_internal_proto_account_proto_init() }
func file_internal_proto_account_proto_init() {
	if File_internal_proto_account_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_proto_account_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GetAccountRequest); i {
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
		file_internal_proto_account_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetAccountResponse); i {
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
		file_internal_proto_account_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GetAccountsRequest); i {
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
		file_internal_proto_account_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetAccountsResponse); i {
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
			RawDescriptor: file_internal_proto_account_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_proto_account_proto_goTypes,
		DependencyIndexes: file_internal_proto_account_proto_depIdxs,
		MessageInfos:      file_internal_proto_account_proto_msgTypes,
	}.Build()
	File_internal_proto_account_proto = out.File
	file_internal_proto_account_proto_rawDesc = nil
	file_internal_proto_account_proto_goTypes = nil
	file_internal_proto_account_proto_depIdxs = nil
}
