package config

type ES struct {
	Endpoint         string `toml:"endpoint"`
	CredentialKey    string `toml:"credential_key"`
	CredentialSecret string `toml:"credential_secret"`
	User             string `toml:"user"`
	Pwd              string `toml:"pwd"`
}

func GetESConfig() *ES {
	return config.Es
}
