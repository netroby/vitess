// Code generated by protoc-gen-go.
// source: replicationdata.proto
// DO NOT EDIT!

/*
Package replicationdata is a generated protocol buffer package.

It is generated from these files:
	replicationdata.proto

It has these top-level messages:
	MariadbGtid
	MysqlGtidSet
	Position
	Status
*/
package replicationdata

import proto "github.com/golang/protobuf/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal

// MariaDB 10.0
type MariadbGtid struct {
	Domain   uint32 `protobuf:"varint,1,opt,name=domain" json:"domain,omitempty"`
	Server   uint32 `protobuf:"varint,2,opt,name=server" json:"server,omitempty"`
	Sequence uint64 `protobuf:"varint,3,opt,name=sequence" json:"sequence,omitempty"`
}

func (m *MariadbGtid) Reset()         { *m = MariadbGtid{} }
func (m *MariadbGtid) String() string { return proto.CompactTextString(m) }
func (*MariadbGtid) ProtoMessage()    {}

// MySQL 5.6
type MysqlGtidSet struct {
	UuidSet []*MysqlGtidSet_MysqlUuidSet `protobuf:"bytes,1,rep,name=uuid_set" json:"uuid_set,omitempty"`
}

func (m *MysqlGtidSet) Reset()         { *m = MysqlGtidSet{} }
func (m *MysqlGtidSet) String() string { return proto.CompactTextString(m) }
func (*MysqlGtidSet) ProtoMessage()    {}

func (m *MysqlGtidSet) GetUuidSet() []*MysqlGtidSet_MysqlUuidSet {
	if m != nil {
		return m.UuidSet
	}
	return nil
}

type MysqlGtidSet_MysqlInterval struct {
	First uint64 `protobuf:"varint,1,opt,name=first" json:"first,omitempty"`
	Last  uint64 `protobuf:"varint,2,opt,name=last" json:"last,omitempty"`
}

func (m *MysqlGtidSet_MysqlInterval) Reset()         { *m = MysqlGtidSet_MysqlInterval{} }
func (m *MysqlGtidSet_MysqlInterval) String() string { return proto.CompactTextString(m) }
func (*MysqlGtidSet_MysqlInterval) ProtoMessage()    {}

type MysqlGtidSet_MysqlUuidSet struct {
	Uuid     []byte                        `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Interval []*MysqlGtidSet_MysqlInterval `protobuf:"bytes,2,rep,name=interval" json:"interval,omitempty"`
}

func (m *MysqlGtidSet_MysqlUuidSet) Reset()         { *m = MysqlGtidSet_MysqlUuidSet{} }
func (m *MysqlGtidSet_MysqlUuidSet) String() string { return proto.CompactTextString(m) }
func (*MysqlGtidSet_MysqlUuidSet) ProtoMessage()    {}

func (m *MysqlGtidSet_MysqlUuidSet) GetInterval() []*MysqlGtidSet_MysqlInterval {
	if m != nil {
		return m.Interval
	}
	return nil
}

// Position represents the information required to specify where to start
// replication. The contents vary depending on the flavor of MySQL in use.
// We define all the fields here and use only the ones we need for each flavor.
type Position struct {
	MariadbGtid  *MariadbGtid  `protobuf:"bytes,1,opt,name=mariadb_gtid" json:"mariadb_gtid,omitempty"`
	MysqlGtidSet *MysqlGtidSet `protobuf:"bytes,2,opt,name=mysql_gtid_set" json:"mysql_gtid_set,omitempty"`
}

func (m *Position) Reset()         { *m = Position{} }
func (m *Position) String() string { return proto.CompactTextString(m) }
func (*Position) ProtoMessage()    {}

func (m *Position) GetMariadbGtid() *MariadbGtid {
	if m != nil {
		return m.MariadbGtid
	}
	return nil
}

func (m *Position) GetMysqlGtidSet() *MysqlGtidSet {
	if m != nil {
		return m.MysqlGtidSet
	}
	return nil
}

// Status is the replication status for MySQL (returned by 'show slave status'
// and parsed into a Position and fields).
type Status struct {
	Position            *Position `protobuf:"bytes,1,opt,name=position" json:"position,omitempty"`
	SlaveIoRunning      bool      `protobuf:"varint,2,opt,name=slave_io_running" json:"slave_io_running,omitempty"`
	SlaveSqlRunning     bool      `protobuf:"varint,3,opt,name=slave_sql_running" json:"slave_sql_running,omitempty"`
	SecondsBehindMaster uint32    `protobuf:"varint,4,opt,name=seconds_behind_master" json:"seconds_behind_master,omitempty"`
	MasterHost          string    `protobuf:"bytes,5,opt,name=master_host" json:"master_host,omitempty"`
	MasterPort          int32     `protobuf:"varint,6,opt,name=master_port" json:"master_port,omitempty"`
	MasterConnectRetry  int32     `protobuf:"varint,7,opt,name=master_connect_retry" json:"master_connect_retry,omitempty"`
}

func (m *Status) Reset()         { *m = Status{} }
func (m *Status) String() string { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()    {}

func (m *Status) GetPosition() *Position {
	if m != nil {
		return m.Position
	}
	return nil
}
