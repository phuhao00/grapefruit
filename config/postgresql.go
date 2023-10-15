package config

type Postgresql struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	DataBase string `json:"dataBase"`
	Pwd      string `json:"pwd"`
}

func GetPSQLConfig() *Postgresql {
	return config.Postgresql
}
