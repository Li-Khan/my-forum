package config

// Config ...
type Config struct {
	BindAddr string `toml:"bind_addr"`

	HostDB     string `toml:"host_db"`
	PortDB     string `toml:"port_db"`
	NameDB     string `toml:"name_db"`
	UserDB     string `toml:"user_db"`
	PasswordDB string `toml:"password_db"`

	CacheHost     string `toml:"cache_host"`
	CachePassword string `toml:"cache_password"`
	CacheAddr     string `toml:"cache_addr"`
}

// * NewConfig - returns a config with standard values
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",

		HostDB:     "localhost",
		PortDB:     ":5432",
		NameDB:     "postgres",
		UserDB:     "postgres",
		PasswordDB: "postgres",

		CacheHost:     "localhost",
		CachePassword: "redis",
		CacheAddr:     ":6379",
	}
}
