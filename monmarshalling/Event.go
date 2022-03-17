package monmarshalling

import (
	"os"
	"time"
)

const HeartBeatMessageTypeName = "heartbeat"
const StatsBeatMessageTypeName = "stats"

type MetaData struct {
	MessageType          string
	StatType             string
	HostName             string
	InstanceName         string
	AgentTimestampUnixMs int64
	PollRateMs           int
}

var hostName = func() string {
	h, _ := os.Hostname()
	return h
}()

func NewMetaDataWithTs() *MetaData {
	return &MetaData{AgentTimestampUnixMs: time.Now().UnixMilli(), InstanceName: hostName}
}

func NewHeartBeatMetaDataWithTs() *MetaData {
	d := NewMetaDataWithTs()
	d.MessageType = HeartBeatMessageTypeName
	return d
}

func NewStatsMetaDataWithTs() *MetaData {
	d := NewMetaDataWithTs()
	d.MessageType = StatsBeatMessageTypeName
	return d
}

type Stat struct {
	MetaData MetaData
	Payload  string
}
