package envconfig

type StatsConfig struct {
	ServerUrl         string `mapstructure:"stats_server_url"`
	RequestTimeoutSec int    `mapstructure:"request_timeout_sec"`
	HostName          string `mapstructure:"host_name"`
	HeartBeatSec      int    `mapstructure:"heart_beat_sec"`
}
