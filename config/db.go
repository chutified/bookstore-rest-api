package config

// DBConfig is a database config struct.
type DBConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DBName   string `yaml:"db_name"`
	User     string `yaml:"username"`
	Password string `yaml:"password"`
}
