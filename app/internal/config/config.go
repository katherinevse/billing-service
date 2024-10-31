package config

type Config struct {
	AppCfg       AppConfig      `yaml:"app"`
	PostgresCfg  PostgresConfig `yaml:"db"`
	LoggerConfig LoggerConfig   `yaml:"logger"`
	KafkaConfig  KafkaConfig    `yaml:"kafka"`
}

type PostgresConfig struct {
	Password string `yaml:"password"`
	User     string `yaml:"user"`
	DbName   string `yaml:"name"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
}

type AppConfig struct {
	Host   string `yaml:"host"`
	Port   string `yaml:"port"`
	Scheme string `yaml:"scheme"`
	Domain string `yaml:"domain"`
}

type LoggerConfig struct {
	Level string `yaml:"level"`
}

type KafkaConfig struct {
	Brokers []string `yaml:"brokers"`
}

func New() *Config {
	return &Config{
		PostgresCfg:  PostgresConfig{},
		AppCfg:       AppConfig{},
		LoggerConfig: LoggerConfig{},
		KafkaConfig:  KafkaConfig{},
	}
}
