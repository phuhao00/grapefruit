package config

//Mongo mongodb://localhost:27017
type Mongo struct {
	Host string `toml:"host"`
	Port string `toml:"port"`
}

func (c *Config) GetMongoConfig() *Mongo {
	return c.Mongo
}
