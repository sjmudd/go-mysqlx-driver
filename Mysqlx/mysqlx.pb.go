// Code generated by protoc-gen-go.
// source: mysqlx.proto
// DO NOT EDIT!

/*
Package Mysqlx is a generated protocol buffer package.

It is generated from these files:
	mysqlx.proto

It has these top-level messages:
	ClientMessages
	ServerMessages
	Ok
	Error
*/
package Mysqlx

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type ClientMessages_Type int32

const (
	ClientMessages_CON_CAPABILITIES_GET       ClientMessages_Type = 1
	ClientMessages_CON_CAPABILITIES_SET       ClientMessages_Type = 2
	ClientMessages_CON_CLOSE                  ClientMessages_Type = 3
	ClientMessages_SESS_AUTHENTICATE_START    ClientMessages_Type = 4
	ClientMessages_SESS_AUTHENTICATE_CONTINUE ClientMessages_Type = 5
	ClientMessages_SESS_RESET                 ClientMessages_Type = 6
	ClientMessages_SESS_CLOSE                 ClientMessages_Type = 7
	ClientMessages_SQL_STMT_EXECUTE           ClientMessages_Type = 12
	ClientMessages_CRUD_FIND                  ClientMessages_Type = 17
	ClientMessages_CRUD_INSERT                ClientMessages_Type = 18
	ClientMessages_CRUD_UPDATE                ClientMessages_Type = 19
	ClientMessages_CRUD_DELETE                ClientMessages_Type = 20
	ClientMessages_EXPECT_OPEN                ClientMessages_Type = 24
	ClientMessages_EXPECT_CLOSE               ClientMessages_Type = 25
)

var ClientMessages_Type_name = map[int32]string{
	1:  "CON_CAPABILITIES_GET",
	2:  "CON_CAPABILITIES_SET",
	3:  "CON_CLOSE",
	4:  "SESS_AUTHENTICATE_START",
	5:  "SESS_AUTHENTICATE_CONTINUE",
	6:  "SESS_RESET",
	7:  "SESS_CLOSE",
	12: "SQL_STMT_EXECUTE",
	17: "CRUD_FIND",
	18: "CRUD_INSERT",
	19: "CRUD_UPDATE",
	20: "CRUD_DELETE",
	24: "EXPECT_OPEN",
	25: "EXPECT_CLOSE",
}
var ClientMessages_Type_value = map[string]int32{
	"CON_CAPABILITIES_GET":       1,
	"CON_CAPABILITIES_SET":       2,
	"CON_CLOSE":                  3,
	"SESS_AUTHENTICATE_START":    4,
	"SESS_AUTHENTICATE_CONTINUE": 5,
	"SESS_RESET":                 6,
	"SESS_CLOSE":                 7,
	"SQL_STMT_EXECUTE":           12,
	"CRUD_FIND":                  17,
	"CRUD_INSERT":                18,
	"CRUD_UPDATE":                19,
	"CRUD_DELETE":                20,
	"EXPECT_OPEN":                24,
	"EXPECT_CLOSE":               25,
}

func (x ClientMessages_Type) Enum() *ClientMessages_Type {
	p := new(ClientMessages_Type)
	*p = x
	return p
}
func (x ClientMessages_Type) String() string {
	return proto.EnumName(ClientMessages_Type_name, int32(x))
}
func (x *ClientMessages_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ClientMessages_Type_value, data, "ClientMessages_Type")
	if err != nil {
		return err
	}
	*x = ClientMessages_Type(value)
	return nil
}
func (ClientMessages_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type ServerMessages_Type int32

const (
	ServerMessages_OK                         ServerMessages_Type = 0
	ServerMessages_ERROR                      ServerMessages_Type = 1
	ServerMessages_CONN_CAPABILITIES          ServerMessages_Type = 2
	ServerMessages_SESS_AUTHENTICATE_CONTINUE ServerMessages_Type = 3
	ServerMessages_SESS_AUTHENTICATE_OK       ServerMessages_Type = 4
	// NOTICE has to stay at 11 forever
	ServerMessages_NOTICE                               ServerMessages_Type = 11
	ServerMessages_RESULTSET_COLUMN_META_DATA           ServerMessages_Type = 12
	ServerMessages_RESULTSET_ROW                        ServerMessages_Type = 13
	ServerMessages_RESULTSET_FETCH_DONE                 ServerMessages_Type = 14
	ServerMessages_RESULTSET_FETCH_SUSPENDED            ServerMessages_Type = 15
	ServerMessages_RESULTSET_FETCH_DONE_MORE_RESULTSETS ServerMessages_Type = 16
	ServerMessages_SQL_STMT_EXECUTE_OK                  ServerMessages_Type = 17
	ServerMessages_RESULTSET_FETCH_DONE_MORE_OUT_PARAMS ServerMessages_Type = 18
)

var ServerMessages_Type_name = map[int32]string{
	0:  "OK",
	1:  "ERROR",
	2:  "CONN_CAPABILITIES",
	3:  "SESS_AUTHENTICATE_CONTINUE",
	4:  "SESS_AUTHENTICATE_OK",
	11: "NOTICE",
	12: "RESULTSET_COLUMN_META_DATA",
	13: "RESULTSET_ROW",
	14: "RESULTSET_FETCH_DONE",
	15: "RESULTSET_FETCH_SUSPENDED",
	16: "RESULTSET_FETCH_DONE_MORE_RESULTSETS",
	17: "SQL_STMT_EXECUTE_OK",
	18: "RESULTSET_FETCH_DONE_MORE_OUT_PARAMS",
}
var ServerMessages_Type_value = map[string]int32{
	"OK":                                   0,
	"ERROR":                                1,
	"CONN_CAPABILITIES":                    2,
	"SESS_AUTHENTICATE_CONTINUE":           3,
	"SESS_AUTHENTICATE_OK":                 4,
	"NOTICE":                               11,
	"RESULTSET_COLUMN_META_DATA":           12,
	"RESULTSET_ROW":                        13,
	"RESULTSET_FETCH_DONE":                 14,
	"RESULTSET_FETCH_SUSPENDED":            15,
	"RESULTSET_FETCH_DONE_MORE_RESULTSETS": 16,
	"SQL_STMT_EXECUTE_OK":                  17,
	"RESULTSET_FETCH_DONE_MORE_OUT_PARAMS": 18,
}

func (x ServerMessages_Type) Enum() *ServerMessages_Type {
	p := new(ServerMessages_Type)
	*p = x
	return p
}
func (x ServerMessages_Type) String() string {
	return proto.EnumName(ServerMessages_Type_name, int32(x))
}
func (x *ServerMessages_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ServerMessages_Type_value, data, "ServerMessages_Type")
	if err != nil {
		return err
	}
	*x = ServerMessages_Type(value)
	return nil
}
func (ServerMessages_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

type Error_Severity int32

const (
	Error_ERROR Error_Severity = 0
	Error_FATAL Error_Severity = 1
)

var Error_Severity_name = map[int32]string{
	0: "ERROR",
	1: "FATAL",
}
var Error_Severity_value = map[string]int32{
	"ERROR": 0,
	"FATAL": 1,
}

func (x Error_Severity) Enum() *Error_Severity {
	p := new(Error_Severity)
	*p = x
	return p
}
func (x Error_Severity) String() string {
	return proto.EnumName(Error_Severity_name, int32(x))
}
func (x *Error_Severity) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Error_Severity_value, data, "Error_Severity")
	if err != nil {
		return err
	}
	*x = Error_Severity(value)
	return nil
}
func (Error_Severity) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{3, 0} }

// IDs of messages that can be sent from client to the server
//
// .. note::
//   this message is never sent on the wire. It is only used to let ``protoc``
//
//   * generate constants
//   * check for uniqueness
type ClientMessages struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *ClientMessages) Reset()                    { *m = ClientMessages{} }
func (m *ClientMessages) String() string            { return proto.CompactTextString(m) }
func (*ClientMessages) ProtoMessage()               {}
func (*ClientMessages) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// IDs of messages that can be sent from server to client
//
// .. note::
//   this message is never sent on the wire. It is only used to let ``protoc``
//
//   * generate constants
//   * check for uniqueness
type ServerMessages struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *ServerMessages) Reset()                    { *m = ServerMessages{} }
func (m *ServerMessages) String() string            { return proto.CompactTextString(m) }
func (*ServerMessages) ProtoMessage()               {}
func (*ServerMessages) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// generic Ok message
type Ok struct {
	Msg              *string `protobuf:"bytes,1,opt,name=msg" json:"msg,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Ok) Reset()                    { *m = Ok{} }
func (m *Ok) String() string            { return proto.CompactTextString(m) }
func (*Ok) ProtoMessage()               {}
func (*Ok) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Ok) GetMsg() string {
	if m != nil && m.Msg != nil {
		return *m.Msg
	}
	return ""
}

// generic Error message
//
// A ``severity`` of ``ERROR`` indicates the current message sequence is
// aborted for the given error and the session is ready for more.
//
// In case of a ``FATAL`` error message the client should not expect
// the server to continue handling any further messages and should
// close the connection.
//
// :param severity: severity of the error message
// :param code: error-code
// :param sql_state: SQL state
// :param msg: human readable error message
type Error struct {
	Severity         *Error_Severity `protobuf:"varint,1,opt,name=severity,enum=Mysqlx.Error_Severity,def=0" json:"severity,omitempty"`
	Code             *uint32         `protobuf:"varint,2,req,name=code" json:"code,omitempty"`
	SqlState         *string         `protobuf:"bytes,4,req,name=sql_state" json:"sql_state,omitempty"`
	Msg              *string         `protobuf:"bytes,3,req,name=msg" json:"msg,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *Error) Reset()                    { *m = Error{} }
func (m *Error) String() string            { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()               {}
func (*Error) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

const Default_Error_Severity Error_Severity = Error_ERROR

func (m *Error) GetSeverity() Error_Severity {
	if m != nil && m.Severity != nil {
		return *m.Severity
	}
	return Default_Error_Severity
}

func (m *Error) GetCode() uint32 {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return 0
}

func (m *Error) GetSqlState() string {
	if m != nil && m.SqlState != nil {
		return *m.SqlState
	}
	return ""
}

func (m *Error) GetMsg() string {
	if m != nil && m.Msg != nil {
		return *m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*ClientMessages)(nil), "Mysqlx.ClientMessages")
	proto.RegisterType((*ServerMessages)(nil), "Mysqlx.ServerMessages")
	proto.RegisterType((*Ok)(nil), "Mysqlx.Ok")
	proto.RegisterType((*Error)(nil), "Mysqlx.Error")
	proto.RegisterEnum("Mysqlx.ClientMessages_Type", ClientMessages_Type_name, ClientMessages_Type_value)
	proto.RegisterEnum("Mysqlx.ServerMessages_Type", ServerMessages_Type_name, ServerMessages_Type_value)
	proto.RegisterEnum("Mysqlx.Error_Severity", Error_Severity_name, Error_Severity_value)
}

var fileDescriptor0 = []byte{
	// 581 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x93, 0xdd, 0x52, 0xd3, 0x40,
	0x18, 0x86, 0xed, 0x0f, 0x95, 0x7e, 0xb4, 0x65, 0xbb, 0x54, 0x28, 0xa8, 0x0c, 0xd3, 0xf1, 0x80,
	0xa3, 0x38, 0xe3, 0xa1, 0x67, 0x21, 0xf9, 0x90, 0x8c, 0xf9, 0xa9, 0xd9, 0xcd, 0xc8, 0xd9, 0x0e,
	0x96, 0x95, 0xa9, 0x42, 0x83, 0x49, 0x70, 0xe0, 0x0e, 0xbc, 0x0d, 0xef, 0xc1, 0xfb, 0xf0, 0x1e,
	0xbc, 0x12, 0xbf, 0xfc, 0xb4, 0xd5, 0xc2, 0x70, 0x92, 0x99, 0x7d, 0xde, 0xdd, 0x77, 0xb3, 0xcf,
	0xce, 0x42, 0xe7, 0xea, 0x2e, 0xfd, 0x76, 0x79, 0x6b, 0x5c, 0x27, 0x71, 0x16, 0xf3, 0x96, 0x57,
	0x8c, 0xf6, 0x58, 0x49, 0x15, 0x7d, 0xca, 0x64, 0x6f, 0xbb, 0x22, 0x89, 0x4e, 0x6f, 0x2e, 0xb3,
	0x54, 0x67, 0x15, 0xef, 0x57, 0x7c, 0x92, 0xdc, 0x9c, 0x57, 0x68, 0x30, 0x5f, 0xac, 0xd3, 0x74,
	0x1a, 0xcf, 0x2a, 0xba, 0x33, 0x9f, 0x18, 0xcf, 0x66, 0x7a, 0x92, 0x2d, 0x83, 0xad, 0x2a, 0xd0,
	0xb7, 0xd7, 0xc4, 0x57, 0xe0, 0x2c, 0xce, 0xa6, 0x13, 0x5d, 0xc2, 0xd1, 0xaf, 0x3a, 0xf4, 0xac,
	0xcb, 0xa9, 0x9e, 0x65, 0x1e, 0x55, 0x9f, 0x5d, 0xe8, 0x74, 0xf4, 0xb3, 0x0e, 0x4d, 0x79, 0x77,
	0xad, 0xf9, 0x10, 0x06, 0x56, 0xe0, 0x2b, 0xcb, 0x1c, 0x9b, 0x47, 0x8e, 0xeb, 0x48, 0x07, 0x85,
	0x7a, 0x87, 0x92, 0xd5, 0x1e, 0x4c, 0x04, 0x25, 0x75, 0xde, 0x85, 0x76, 0x91, 0xb8, 0x81, 0x40,
	0xd6, 0xe0, 0xcf, 0x61, 0x47, 0xa0, 0x10, 0xca, 0x8c, 0xe4, 0x09, 0xfa, 0xd2, 0xb1, 0x4c, 0x89,
	0x4a, 0x48, 0x33, 0x94, 0xac, 0xc9, 0xf7, 0x61, 0xef, 0x7e, 0x48, 0xab, 0xa5, 0xe3, 0x47, 0xc8,
	0xd6, 0x78, 0x0f, 0xa0, 0xc8, 0x43, 0xcc, 0xbb, 0x5b, 0x8b, 0x71, 0x59, 0xfe, 0x94, 0x0f, 0x80,
	0x89, 0x0f, 0x2e, 0xd5, 0x79, 0x52, 0xe1, 0x29, 0x5a, 0x91, 0x44, 0xd6, 0x29, 0xfe, 0x20, 0x8c,
	0x6c, 0x75, 0xec, 0xf8, 0x36, 0xeb, 0xf3, 0x4d, 0xd8, 0x28, 0x86, 0x8e, 0x2f, 0x90, 0x76, 0xe5,
	0x0b, 0x10, 0x8d, 0x6d, 0xda, 0x8f, 0x6d, 0x2d, 0x80, 0x8d, 0x2e, 0x12, 0x18, 0xe4, 0x00, 0x4f,
	0xc7, 0x68, 0x49, 0x15, 0x8c, 0xd1, 0x67, 0x43, 0xce, 0xa0, 0x53, 0x81, 0x72, 0xeb, 0xdd, 0xd1,
	0x1f, 0xd2, 0x26, 0x74, 0xf2, 0x5d, 0x27, 0x0b, 0x6d, 0xbf, 0xe7, 0xda, 0x5a, 0x50, 0x0f, 0xde,
	0xb3, 0x27, 0xbc, 0x0d, 0x6b, 0x18, 0x86, 0x41, 0x48, 0xbe, 0x9e, 0x41, 0x9f, 0xce, 0xf5, 0xbf,
	0x30, 0x92, 0xf5, 0xb8, 0x80, 0x46, 0xae, 0xf9, 0x7e, 0x4e, 0xdd, 0x4d, 0x0e, 0xd0, 0xf2, 0x03,
	0x02, 0xc8, 0x36, 0xf2, 0x16, 0x32, 0x14, 0xb9, 0x92, 0x2c, 0xd1, 0x6a, 0x37, 0xf2, 0x7c, 0xe5,
	0xa1, 0x34, 0x15, 0x9d, 0xcf, 0x24, 0x21, 0x7d, 0xe8, 0x2e, 0xf3, 0x30, 0xf8, 0xc8, 0xba, 0x79,
	0xf1, 0x12, 0x1d, 0xa3, 0xb4, 0x4e, 0x94, 0x1d, 0xf8, 0xc8, 0x7a, 0xfc, 0x25, 0xec, 0xae, 0x26,
	0x22, 0x12, 0xa4, 0xc1, 0x46, 0x9b, 0x6d, 0xf2, 0x43, 0x78, 0xf5, 0xd0, 0x42, 0xe5, 0x05, 0x21,
	0xaa, 0x45, 0x22, 0x18, 0xe3, 0x3b, 0xb0, 0xb5, 0x7a, 0x39, 0xf9, 0xaf, 0xf7, 0x1f, 0xaf, 0x08,
	0x22, 0xa9, 0xc6, 0x66, 0x68, 0x7a, 0x82, 0xf1, 0x51, 0x9f, 0x44, 0x7e, 0xe5, 0x1b, 0xd0, 0xb8,
	0x4a, 0x2f, 0x86, 0xb5, 0x83, 0xda, 0x61, 0x7b, 0xf4, 0xa3, 0x46, 0x52, 0x93, 0x24, 0x4e, 0xf8,
	0x6b, 0x58, 0x4f, 0x35, 0xf9, 0x9f, 0x66, 0x77, 0x45, 0xd6, 0x7b, 0xb3, 0x6d, 0x94, 0x2f, 0xcd,
	0x28, 0x26, 0x18, 0xa2, 0x4a, 0xdf, 0x96, 0xb7, 0xc0, 0x3b, 0xd0, 0x9c, 0xc4, 0xe7, 0x7a, 0x58,
	0x3f, 0xa8, 0x1f, 0x76, 0x49, 0x4a, 0x9b, 0xe6, 0xaa, 0x34, 0x3b, 0xcb, 0xf4, 0xb0, 0x49, 0xa8,
	0x3d, 0xdf, 0xa8, 0x91, 0x0f, 0x46, 0x07, 0xb0, 0x3e, 0x2f, 0x58, 0x5e, 0x64, 0x71, 0xa7, 0xc7,
	0x64, 0xd5, 0x65, 0xb5, 0xa3, 0x7d, 0x78, 0x31, 0x89, 0xaf, 0x8c, 0xe2, 0x51, 0x19, 0x93, 0x2f,
	0xc6, 0xbf, 0x8f, 0xfe, 0xd3, 0xcd, 0xe7, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xa4, 0x1a, 0x89,
	0x97, 0x06, 0x04, 0x00, 0x00,
}
