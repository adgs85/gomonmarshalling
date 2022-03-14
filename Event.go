package monmarshalling

type MetaData struct {
	Type          string
	StatType      string
	CollectUnixMs int64
	PollRateMs    int64
}

type Stat struct {
	MetaData MetaData
	Payload  string
}
