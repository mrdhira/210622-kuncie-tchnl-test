package configs

type Configs struct {
	HttpConfigs     HttpConfigs     `mapstructure:"http_configs"`
	PostgresConfigs PostgresConfigs `mapstructure:"postgres_configs"`
}

type HttpConfigs struct {
}

type PostgresConfigs struct {
	DSN string `mapstructure:"dsn"`
}
