package apiserver

type Config struct {
	BindAddr string `toml:"bind-addr"`
	LogLevel string `toml:"log_level"`
}

func NewConfir() *Config {
	return &Config{
		BindAddr: ":8090",
		LogLevel: "debug",
	}
}
