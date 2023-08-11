package config

type AppConfig struct {
	Addr string
}

func New(addr string) *AppConfig {
	return &AppConfig{
		Addr: "8080",
	}
}
