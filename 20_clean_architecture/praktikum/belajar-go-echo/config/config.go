package config

type Config struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_PORT     string
	DB_HOST     string
	DB_NAME     string
	JWT_KEY     string
}

func InitConfiguration() Config {

	return Config{
		DB_USERNAME: "root",
		DB_PASSWORD: "",
		DB_NAME:     "praktikum",
		DB_PORT:     "",
		DB_HOST:     "",
		JWT_KEY:     "b80e3e726a689887f2bcd92b4b2b1182aac0ed2ceba1c456e5f7a9e07abc",
	}
}
