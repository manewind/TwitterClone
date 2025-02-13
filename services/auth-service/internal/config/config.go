package config

type AppConfig struct {
	Port string
}

func NewAppConfig() AppConfig {
	return AppConfig{
		Port: ":8080",
	}
}
