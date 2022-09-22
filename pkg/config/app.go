package config

type AppConfig struct {
	Env                string   `mapstructure:"ENV"`
	HttpPort           string   `mapstructure:"HTTP_PORT"`
	DatabaseDns        string   `mapstructure:"DB_DNS"`
	PersonKeyspace     string   `mapstructure:"PERSON_KEYSPACE"`
	ScyllaHost         string   `mapstructure:"SCYLLA_HOST"`
	KafkaHost          []string `mapstructure:"KAFKA_HOST"`
	DatabaseEnableLogs bool     `mapstructure:"DB_ENABLE_LOGS"`
}
