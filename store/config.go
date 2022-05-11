package store

// PostgresDB todo support another db
type PostgresDBConfig struct {
	Port     string `yaml:"port"`
	Host     string `yaml:"host"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Dbname   string `yaml:"dbname"`
}