package configs

import "fmt"

type DatabaseCredential struct {
	Username string
	Password string
	Hostname string
	Port     int
	Database string
	Options  string
}

func SetupDatabase() (connectionString string) {
	var connectionData = DatabaseCredential{
		Username: "root",
		Password: "",
		Hostname: "localhost",
		Port:     3306,
		Database: "praktikum",
		Options:  "",
	}

	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", connectionData.Username, connectionData.Password, connectionData.Hostname, connectionData.Port, connectionData.Database, connectionData.Options)
}
