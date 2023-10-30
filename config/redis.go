package config

type Redis struct {
	Host    string `toml:"host"`
	Port    string `toml:"port"`
	User    string `toml:"user"`
	Pwd     string `toml:"pwd"`
	CrtPath string `yaml:"crt_path"`
}

func GetRedisConfig() *Redis {
	return config.Redis
}
