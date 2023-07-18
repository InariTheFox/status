package models

type Config struct {
	ServerIP      string `yaml:"server_ip"`
	ServerPort    string `yaml:"server_port"`
	AdminUser     string `yaml:"admin_user"`
	AdminEmail    string `yaml:"admin_email"`
	AdminPassword string `yaml:"admin_password"`
	Language      string `yaml:"language"`
	DbDriver      string `yaml:"db_driver"`
	DbHost        string `yaml:"db_host"`
	DbPort        int    `yaml:"db_port"`
	DbUser        string `yaml:"db_user"`
	DbPassword    string `yaml:"db_password"`
	DbName        string `yaml:"db_name"`
}
