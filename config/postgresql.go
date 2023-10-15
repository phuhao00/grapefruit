package config

type Postgresql struct {
	Host     string `toml:"host"`
	Port     string `toml:"port"`
	User     string `toml:"user"`
	DataBase string `toml:"dataBase"`
	Pwd      string `toml:"pwd"`
}

func GetPSQLConfig() *Postgresql {
	return config.Postgresql
}
