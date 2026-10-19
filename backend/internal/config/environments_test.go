package config

import (
	"bytes"
	"testing"

	"github.com/spf13/viper"
)

func TestLoad_Valid(t *testing.T) {
	t.Helper()

	v := viper.New()
	v.SetConfigType("yaml")

	if err := v.ReadConfig(bytes.NewBufferString(`
port: 8000
database:
  host: "localhost"
  port: 5432
`)); err != nil {
		t.Fatalf("Failed to prep fixture viper: %v", err)
	}

	cfg, err := Load(v)
	if err != nil {
		t.Fatalf("Load returned error: %v", err)
	}
	if cfg.Port != 8000 {
		t.Errorf("Port = %d, want %d", cfg.Port, 8000)
	}
	if cfg.Database.Host != "localhost" {
		t.Errorf("Database.Host = %q, want %q", cfg.Database.Host, "localhost")
	}
	if cfg.Database.Port != 5432 {
		t.Errorf("Database.Port = %d, want %d", cfg.Database.Port, 5432)
	}
}

func TestLoad_MalformedYAML(t *testing.T) {
	v := viper.New()

	v.SetConfigType("yaml")

	err := v.ReadConfig(bytes.NewBufferString("port: [this is not valid yaml"))
	if err == nil {
		t.Fatalf("expected ReadConfig to fail on malformed yaml")
	}
}

func TestLoad_TypeMismatch(t *testing.T) {
	v := viper.New()
	v.SetConfigType("yaml")

	if err := v.ReadConfig(bytes.NewBufferString(`port: "not-a-number"`)); err != nil {
		t.Fatalf("Failed to prep fixture viper: %v", err)
	}

	if _, err := Load(v); err == nil {
		t.Fatalf("expected Load to return error when port is not an int")
	}
}

// TestBindEnv_NoFile simulates Azure App Service: no config.yaml, values
// come entirely from env vars (App Settings, incl. Key Vault references
// already resolved by the platform before the process sees them).
func TestBindEnv_NoFile(t *testing.T) {
	t.Setenv("PORT", "9090")
	t.Setenv("DATABASE_HOST", "prod-db.example.com")
	t.Setenv("DATABASE_PORT", "5432")

	v := viper.New()
	BindEnv(v)

	cfg, err := Load(v)
	if err != nil {
		t.Fatalf("Load returned error: %v", err)
	}
	if cfg.Port != 9090 {
		t.Errorf("Port = %d, want %d", cfg.Port, 9090)
	}
	if cfg.Database.Host != "prod-db.example.com" {
		t.Errorf("Database.Host = %q, want %q", cfg.Database.Host, "prod-db.example.com")
	}
	if cfg.Database.Port != 5432 {
		t.Errorf("Database.Port = %d, want %d", cfg.Database.Port, 5432)
	}
}

// TestBindEnv_OverridesFile confirms env vars win over config.yaml values —
// matters locally when overriding a secret without editing the file.
func TestBindEnv_OverridesFile(t *testing.T) {
	t.Setenv("DATABASE_HOST", "overridden-host")

	v := viper.New()
	v.SetConfigType("yaml")
	if err := v.ReadConfig(bytes.NewBufferString(`
port: 8000
database:
  host: "from-file"
  port: 5432
`)); err != nil {
		t.Fatalf("Failed to prep fixture viper: %v", err)
	}
	BindEnv(v)

	cfg, err := Load(v)
	if err != nil {
		t.Fatalf("Load returned error: %v", err)
	}
	if cfg.Database.Host != "overridden-host" {
		t.Errorf("Database.Host = %q, want %q", cfg.Database.Host, "overridden-host")
	}
	if cfg.Port != 8000 {
		t.Errorf("Port = %d, want %d (should keep file value when no env override)", cfg.Port, 8000)
	}
}
