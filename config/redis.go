package config

type Redis struct {
	Host string `toml:"host"`
	Port string `toml:"port"`
	User string `toml:"user"`
	Pwd  string `toml:"pwd"`
}

func (c *Config) GetRedisConfig() *Redis {
	return c.Redis
}
