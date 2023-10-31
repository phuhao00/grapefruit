package config

//Mongo mongodb://localhost:27017
type Mongo struct {
	Host       string `toml:"host"`
	Port       string `toml:"port"`
	CtxTimeout int64  `toml:"ctx_timeout"`
}

func GetMongoConfig() *Mongo {
	return config.Mongo
}
