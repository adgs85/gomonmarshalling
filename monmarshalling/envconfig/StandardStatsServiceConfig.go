package envconfig

type StatsConfig struct {
	ServerUrl         string `mapstructure:"stats_server_url"`
	RequestTimeoutSec string `mapstructure:"request_timeout_sec"`
	HostName          string `mapstructure:"host_name"`
}
