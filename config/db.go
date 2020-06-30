package config

// DBConfig holds database settings.
type DBConfig struct {
	Host     string
	Port     int
	DBName   string
	User     string
	Password string
}
