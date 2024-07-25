package config

type BoilerConfig struct {
	ControllersActived string   `json:"ControllersActived" env:"CONTROLLERS_ACTIVE"`
	Controllers        []string `json:"Controllers" `
	DatabaseURL        string   `json:"DatabaseURL" env:"DATABASE_URL"`
	DatabasePort       string   `json:"DatabasePort" env:"PORT"`
	DatabaseUser       string   `json:"DatabaseUser" env:"DB_User"`
	DatabasePassword   string   `json:"DatabasePassword" env:"DB_PASSWORD"`
	DatabaseName       string   `json:"DatabaseName" env:"DB_NAME"`
	ServiceName        string   `json:"ServiceName" env:"DATABASE_URL"`
	ServicePort        string   `json:"ServicePort" env:"PORT"`
	GRPCPort           string   `json:"GRPCPort" env:"GRPC_PORT"`
}
