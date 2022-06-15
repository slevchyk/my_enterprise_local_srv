package models

type Config struct {
	ServerConfig   ServerConfig   `json:"server"`
	ServiceConfig  ServiceConfig  `json:"service_cfg"`
	DatabaseConfig DatabaseConfig `json:"database_cfg"`
	MainAuth       AuthConfig     `json:"main_auth"`
	MobileAuth     AuthConfig     `json:"mobile_auth"`
}

type ServerConfig struct {
	Port int `json:"port"`
}

type ServiceConfig struct {
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
	Description string `json:"description"`
}

type DatabaseConfig struct {
	Path string `json:"path"`
}

type AuthConfig struct {
	User     string `json:"user"`
	Password string `json:"password"`
}
