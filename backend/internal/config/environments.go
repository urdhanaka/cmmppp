package config

import (
	"errors"
	"strings"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

// Deps holds every external connection the server needs. Add a field here
// (Redis, Kafka, ...) and wire it in cmd/server/server.go — no router
// rewiring needed elsewhere.
type Deps struct {
	DB           *gorm.DB
	TokenManager *TokenManager
}

// EnvConfig struct holds the application configuration mappings
type EnvConfig struct {
	// Database configuration
	Database struct {
		Host string `mapstructure:"host"`
		Port int    `mapstructure:"port"`
		User string `mapstructure:"user"`
		Pass string `mapstructure:"pass"`
		Name string `mapstructure:"name"`
	} `mapstructure:"database"`

	// JWT configuration
	JWT struct {
		Secret string `mapstructure:"secret"`
	} `mapstructure:"jwt"`

	// Port that will be used by the backend
	Port int `mapstructure:"port"`
}

// BindEnv wires each Config field to its environment variable name.
// On Azure App Service, App Settings are exposed as env vars — including
// ones backed by Key Vault references, which App Service resolves to the
// secret value before the process ever sees them. No Key Vault SDK needed.
func BindEnv(v *viper.Viper) {
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	v.BindEnv("port", "PORT")
	v.BindEnv("env", "ENV")
	v.BindEnv("database.host", "DATABASE_HOST")
	v.BindEnv("database.port", "DATABASE_PORT")
	v.BindEnv("database.user", "DATABASE_USER")
	v.BindEnv("database.pass", "DATABASE_PASS")
	v.BindEnv("database.name", "DATABASE_NAME")
	v.BindEnv("jwt.secret", "JWT_SECRET")
}

func Load(v *viper.Viper) (*EnvConfig, error) {
	cfg := &EnvConfig{}
	if err := v.Unmarshal(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

// NewConfig builds the application config from config.yaml (local-dev
// convenience) and environment variables (App Settings on Azure App
// Service), and returns it. Call it once at each composition root
// (cmd/app, cmd/database, ...) and pass the result down via constructors —
// don't fetch it through a package-level getter from arbitrary call sites.
func NewConfig() (*EnvConfig, error) {
	v := viper.New()
	v.SetConfigType("yaml")
	v.SetConfigName("config")
	v.AddConfigPath(".")

	BindEnv(v)

	// config.yaml is a local-dev convenience only. On Azure App Service
	// there is no file — config comes entirely from App Settings env vars.
	if err := v.ReadInConfig(); err != nil {
		if _, ok := errors.AsType[viper.ConfigFileNotFoundError](err); !ok {
			return nil, err
		}
	}

	return Load(v)
}
