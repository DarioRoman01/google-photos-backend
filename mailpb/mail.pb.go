// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.6.1
// source: mailpb/mail.proto

package mailpb

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

// SendMailRequest is the request message containing the mail to send.
type SendMailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type     string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Receiver string `protobuf:"bytes,2,opt,name=receiver,proto3" json:"receiver,omitempty"`
	User     string `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	Token    string `protobuf:"bytes,4,opt,name=token,proto3" json:"token,omitempty"`
	Subject  string `protobuf:"bytes,5,opt,name=subject,proto3" json:"subject,omitempty"`
	Body     string `protobuf:"bytes,6,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *SendMailRequest) Reset() {
	*x = SendMailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mailpb_mail_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMailRequest) ProtoMessage() {}

func (x *SendMailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mailpb_mail_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMailRequest.ProtoReflect.Descriptor instead.
func (*SendMailRequest) Descriptor() ([]byte, []int) {
	return file_mailpb_mail_proto_rawDescGZIP(), []int{0}
}

func (x *SendMailRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *SendMailRequest) GetReceiver() string {
	if x != nil {
		return x.Receiver
	}
	return ""
}

func (x *SendMailRequest) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *SendMailRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *SendMailRequest) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *SendMailRequest) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

// SendMailResponse is the response message containing the id of the mail.
type SendMailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *SendMailResponse) Reset() {
	*x = SendMailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mailpb_mail_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMailResponse) ProtoMessage() {}

func (x *SendMailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mailpb_mail_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMailResponse.ProtoReflect.Descriptor instead.
func (*SendMailResponse) Descriptor() ([]byte, []int) {
	return file_mailpb_mail_proto_rawDescGZIP(), []int{1}
}

func (x *SendMailResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_mailpb_mail_proto protoreflect.FileDescriptor

var file_mailpb_mail_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6d, 0x61, 0x69, 0x6c, 0x70, 0x62, 0x2f, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6d, 0x61, 0x69, 0x6c, 0x70, 0x62, 0x22, 0x99, 0x01, 0x0a, 0x0f,
	0x53, 0x65, 0x6e, 0x64, 0x4d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x12,
	0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x22, 0x0a, 0x10, 0x53, 0x65, 0x6e, 0x64, 0x4d,
	0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x32, 0x4c, 0x0a, 0x0b, 0x4d,
	0x61, 0x69, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3d, 0x0a, 0x08, 0x53, 0x65,
	0x6e, 0x64, 0x4d, 0x61, 0x69, 0x6c, 0x12, 0x17, 0x2e, 0x6d, 0x61, 0x69, 0x6c, 0x70, 0x62, 0x2e,
	0x53, 0x65, 0x6e, 0x64, 0x4d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x18, 0x2e, 0x6d, 0x61, 0x69, 0x6c, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x61, 0x69,
	0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x25, 0x5a, 0x23, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x61, 0x72, 0x69, 0x6f, 0x52, 0x6f, 0x6d,
	0x61, 0x6e, 0x30, 0x31, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x6d, 0x61, 0x69, 0x6c, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mailpb_mail_proto_rawDescOnce sync.Once
	file_mailpb_mail_proto_rawDescData = file_mailpb_mail_proto_rawDesc
)

func file_mailpb_mail_proto_rawDescGZIP() []byte {
	file_mailpb_mail_proto_rawDescOnce.Do(func() {
		file_mailpb_mail_proto_rawDescData = protoimpl.X.CompressGZIP(file_mailpb_mail_proto_rawDescData)
	})
	return file_mailpb_mail_proto_rawDescData
}

var file_mailpb_mail_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_mailpb_mail_proto_goTypes = []interface{}{
	(*SendMailRequest)(nil),  // 0: mailpb.SendMailRequest
	(*SendMailResponse)(nil), // 1: mailpb.SendMailResponse
}
var file_mailpb_mail_proto_depIdxs = []int32{
	0, // 0: mailpb.MailService.SendMail:input_type -> mailpb.SendMailRequest
	1, // 1: mailpb.MailService.SendMail:output_type -> mailpb.SendMailResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_mailpb_mail_proto_init() }
func file_mailpb_mail_proto_init() {
	if File_mailpb_mail_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mailpb_mail_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMailRequest); i {
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
		file_mailpb_mail_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMailResponse); i {
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
			RawDescriptor: file_mailpb_mail_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mailpb_mail_proto_goTypes,
		DependencyIndexes: file_mailpb_mail_proto_depIdxs,
		MessageInfos:      file_mailpb_mail_proto_msgTypes,
	}.Build()
	File_mailpb_mail_proto = out.File
	file_mailpb_mail_proto_rawDesc = nil
	file_mailpb_mail_proto_goTypes = nil
	file_mailpb_mail_proto_depIdxs = nil
}
