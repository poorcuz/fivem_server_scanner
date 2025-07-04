// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.31.1
// source: server.proto

package pb

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

type Player struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Identifiers   []string               `protobuf:"bytes,2,rep,name=identifiers,proto3" json:"identifiers,omitempty"`
	Endpoint      string                 `protobuf:"bytes,3,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	Ping          int32                  `protobuf:"varint,4,opt,name=ping,proto3" json:"ping,omitempty"`
	Id            int32                  `protobuf:"varint,5,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Player) Reset() {
	*x = Player{}
	mi := &file_server_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Player) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Player) ProtoMessage() {}

func (x *Player) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Player.ProtoReflect.Descriptor instead.
func (*Player) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{0}
}

func (x *Player) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Player) GetIdentifiers() []string {
	if x != nil {
		return x.Identifiers
	}
	return nil
}

func (x *Player) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

func (x *Player) GetPing() int32 {
	if x != nil {
		return x.Ping
	}
	return 0
}

func (x *Player) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ServerData struct {
	state        protoimpl.MessageState `protogen:"open.v1"`
	SvMaxclients int32                  `protobuf:"varint,1,opt,name=svMaxclients,proto3" json:"svMaxclients,omitempty"`
	Clients      int32                  `protobuf:"varint,2,opt,name=clients,proto3" json:"clients,omitempty"`
	Protocol     int32                  `protobuf:"varint,3,opt,name=protocol,proto3" json:"protocol,omitempty"`
	Hostname     string                 `protobuf:"bytes,4,opt,name=hostname,proto3" json:"hostname,omitempty"`
	Gametype     string                 `protobuf:"bytes,5,opt,name=gametype,proto3" json:"gametype,omitempty"`
	Mapname      string                 `protobuf:"bytes,6,opt,name=mapname,proto3" json:"mapname,omitempty"`
	// 7 is unused
	Resources           []string          `protobuf:"bytes,8,rep,name=resources,proto3" json:"resources,omitempty"`
	Server              string            `protobuf:"bytes,9,opt,name=server,proto3" json:"server,omitempty"`
	Players             []*Player         `protobuf:"bytes,10,rep,name=players,proto3" json:"players,omitempty"`
	IconVersion         int32             `protobuf:"varint,11,opt,name=iconVersion,proto3" json:"iconVersion,omitempty"`
	Vars                map[string]string `protobuf:"bytes,12,rep,name=vars,proto3" json:"vars,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	EnhancedHostSupport bool              `protobuf:"varint,16,opt,name=enhancedHostSupport,proto3" json:"enhancedHostSupport,omitempty"`
	UpvotePower         int32             `protobuf:"varint,17,opt,name=upvotePower,proto3" json:"upvotePower,omitempty"`
	ConnectEndPoints    []string          `protobuf:"bytes,18,rep,name=connectEndPoints,proto3" json:"connectEndPoints,omitempty"`
	BurstPower          int32             `protobuf:"varint,19,opt,name=burstPower,proto3" json:"burstPower,omitempty"`
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *ServerData) Reset() {
	*x = ServerData{}
	mi := &file_server_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ServerData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerData) ProtoMessage() {}

func (x *ServerData) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerData.ProtoReflect.Descriptor instead.
func (*ServerData) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{1}
}

func (x *ServerData) GetSvMaxclients() int32 {
	if x != nil {
		return x.SvMaxclients
	}
	return 0
}

func (x *ServerData) GetClients() int32 {
	if x != nil {
		return x.Clients
	}
	return 0
}

func (x *ServerData) GetProtocol() int32 {
	if x != nil {
		return x.Protocol
	}
	return 0
}

func (x *ServerData) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *ServerData) GetGametype() string {
	if x != nil {
		return x.Gametype
	}
	return ""
}

func (x *ServerData) GetMapname() string {
	if x != nil {
		return x.Mapname
	}
	return ""
}

func (x *ServerData) GetResources() []string {
	if x != nil {
		return x.Resources
	}
	return nil
}

func (x *ServerData) GetServer() string {
	if x != nil {
		return x.Server
	}
	return ""
}

func (x *ServerData) GetPlayers() []*Player {
	if x != nil {
		return x.Players
	}
	return nil
}

func (x *ServerData) GetIconVersion() int32 {
	if x != nil {
		return x.IconVersion
	}
	return 0
}

func (x *ServerData) GetVars() map[string]string {
	if x != nil {
		return x.Vars
	}
	return nil
}

func (x *ServerData) GetEnhancedHostSupport() bool {
	if x != nil {
		return x.EnhancedHostSupport
	}
	return false
}

func (x *ServerData) GetUpvotePower() int32 {
	if x != nil {
		return x.UpvotePower
	}
	return 0
}

func (x *ServerData) GetConnectEndPoints() []string {
	if x != nil {
		return x.ConnectEndPoints
	}
	return nil
}

func (x *ServerData) GetBurstPower() int32 {
	if x != nil {
		return x.BurstPower
	}
	return 0
}

type Server struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	EndPoint      string                 `protobuf:"bytes,1,opt,name=EndPoint,proto3" json:"EndPoint,omitempty"`
	Data          *ServerData            `protobuf:"bytes,2,opt,name=Data,proto3" json:"Data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Server) Reset() {
	*x = Server{}
	mi := &file_server_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Server) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server) ProtoMessage() {}

func (x *Server) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server.ProtoReflect.Descriptor instead.
func (*Server) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{2}
}

func (x *Server) GetEndPoint() string {
	if x != nil {
		return x.EndPoint
	}
	return ""
}

func (x *Server) GetData() *ServerData {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_server_proto protoreflect.FileDescriptor

const file_server_proto_rawDesc = "" +
	"\n" +
	"\fserver.proto\x12\x06master\"~\n" +
	"\x06Player\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12 \n" +
	"\videntifiers\x18\x02 \x03(\tR\videntifiers\x12\x1a\n" +
	"\bendpoint\x18\x03 \x01(\tR\bendpoint\x12\x12\n" +
	"\x04ping\x18\x04 \x01(\x05R\x04ping\x12\x0e\n" +
	"\x02id\x18\x05 \x01(\x05R\x02id\"\xc5\x04\n" +
	"\n" +
	"ServerData\x12\"\n" +
	"\fsvMaxclients\x18\x01 \x01(\x05R\fsvMaxclients\x12\x18\n" +
	"\aclients\x18\x02 \x01(\x05R\aclients\x12\x1a\n" +
	"\bprotocol\x18\x03 \x01(\x05R\bprotocol\x12\x1a\n" +
	"\bhostname\x18\x04 \x01(\tR\bhostname\x12\x1a\n" +
	"\bgametype\x18\x05 \x01(\tR\bgametype\x12\x18\n" +
	"\amapname\x18\x06 \x01(\tR\amapname\x12\x1c\n" +
	"\tresources\x18\b \x03(\tR\tresources\x12\x16\n" +
	"\x06server\x18\t \x01(\tR\x06server\x12(\n" +
	"\aplayers\x18\n" +
	" \x03(\v2\x0e.master.PlayerR\aplayers\x12 \n" +
	"\viconVersion\x18\v \x01(\x05R\viconVersion\x120\n" +
	"\x04vars\x18\f \x03(\v2\x1c.master.ServerData.VarsEntryR\x04vars\x120\n" +
	"\x13enhancedHostSupport\x18\x10 \x01(\bR\x13enhancedHostSupport\x12 \n" +
	"\vupvotePower\x18\x11 \x01(\x05R\vupvotePower\x12*\n" +
	"\x10connectEndPoints\x18\x12 \x03(\tR\x10connectEndPoints\x12\x1e\n" +
	"\n" +
	"burstPower\x18\x13 \x01(\x05R\n" +
	"burstPower\x1a7\n" +
	"\tVarsEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01\"L\n" +
	"\x06Server\x12\x1a\n" +
	"\bEndPoint\x18\x01 \x01(\tR\bEndPoint\x12&\n" +
	"\x04Data\x18\x02 \x01(\v2\x12.master.ServerDataR\x04DataB\tZ\a./pb;pbb\x06proto3"

var (
	file_server_proto_rawDescOnce sync.Once
	file_server_proto_rawDescData []byte
)

func file_server_proto_rawDescGZIP() []byte {
	file_server_proto_rawDescOnce.Do(func() {
		file_server_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_server_proto_rawDesc), len(file_server_proto_rawDesc)))
	})
	return file_server_proto_rawDescData
}

var file_server_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_server_proto_goTypes = []any{
	(*Player)(nil),     // 0: master.Player
	(*ServerData)(nil), // 1: master.ServerData
	(*Server)(nil),     // 2: master.Server
	nil,                // 3: master.ServerData.VarsEntry
}
var file_server_proto_depIdxs = []int32{
	0, // 0: master.ServerData.players:type_name -> master.Player
	3, // 1: master.ServerData.vars:type_name -> master.ServerData.VarsEntry
	1, // 2: master.Server.Data:type_name -> master.ServerData
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_server_proto_init() }
func file_server_proto_init() {
	if File_server_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_server_proto_rawDesc), len(file_server_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_server_proto_goTypes,
		DependencyIndexes: file_server_proto_depIdxs,
		MessageInfos:      file_server_proto_msgTypes,
	}.Build()
	File_server_proto = out.File
	file_server_proto_goTypes = nil
	file_server_proto_depIdxs = nil
}
