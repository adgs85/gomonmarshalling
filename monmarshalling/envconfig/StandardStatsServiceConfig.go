package envconfig

type StatsConfig struct {
	ServerUrl string `mapstructure:"stats_server_url"`
	HostName  string `mapstructure:"host_name"`
}
