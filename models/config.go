package models

type Config struct {
	Server        ServerConfig  `json:"server"`
	MainAuth      AuthConfig    `json:"main_auth"`
	MobileAuth    AuthConfig    `json:"mobile_auth"`
	ServiceConfig ServiceConfig `json:"service_config"`
}

type ServerConfig struct {
	Port int `json:"port"`
}

type ServiceConfig struct {
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
	Description string `json:"description"`
}

type AuthConfig struct {
	User     string `json:"user"`
	Password string `json:"password"`
}
