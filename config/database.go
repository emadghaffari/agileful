package config

// database struct
type Database struct {
	Username    string `yaml:"username"`
	Password    string `yaml:"password"`
	Host        string `yaml:"host"`
	Schema      string `yaml:"schema"`
	Automigrate bool   `yaml:"automigrate"`
	Logger      bool   `yaml:"logger"`
	Namespace   string
}
