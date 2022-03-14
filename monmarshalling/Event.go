package monmarshalling

import "time"

const HeartBeatMessageTypeName = "heartbeat"
const StatsBeatMessageTypeName = "stats"

type MetaData struct {
	MessageType          string
	StatType             string
	HostName             string
	AgentTimestampUnixMs int64
	PollRateMs           int
}

func NewMetaDataWithTs() *MetaData {
	return &MetaData{AgentTimestampUnixMs: time.Now().UnixMilli()}
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
