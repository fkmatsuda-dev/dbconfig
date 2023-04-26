package dbconfig

import (
	"github.com/fkmatsuda-dev/commons/errorex"
	"github.com/fkmatsuda-dev/commons/files"
	"os"
	"testing"
)

func (c Config) compare(c2 Config) bool {
	return c.Type == c2.Type &&
		c.Host == c2.Host &&
		c.Port == c2.Port &&
		c.User == c2.User &&
		c.Password == c2.Password &&
		c.Database == c2.Database &&
		compareSSLConfig(c.SSL, c2.SSL)
}

func compareSSLConfig(ssl *SSLConfig, ssl2 *SSLConfig) bool {
	if ssl == nil && ssl2 == nil {
		return true
	}
	if ssl == nil || ssl2 == nil {
		return false
	}
	return ssl.Mode == ssl2.Mode &&
		ssl.Cert == ssl2.Cert &&
		ssl.Key == ssl2.Key &&
		ssl.Ca == ssl2.Ca
}

// TestLoadConfig The load must be tested by configuration files in json, yaml, toml, env, hcl, ini, and properties format in addition to the load by the environment variables
func TestLoadConfig(t *testing.T) {

	// Create a temporary directory inside system temp directory
	dirName, err := files.CreateTempDir()
	if err != nil {
		t.Errorf("Error creating temporary directory: %s", err.Error())
		return
	}
	defer func() {
		_ = files.CleanupTempDirs()
	}()

	// Create a sample Config struct
	config := Config{
		Type:     DbTypePostgres,
		Host:     "localhost",
		Port:     5432,
		User:     "postgres",
		Password: "postgres",
		Database: "postgres",
		SSL: &SSLConfig{
			Mode: SSLModeVerifyFull,
			Ca:   "ca.crt",
			Cert: "client.crt",
			Key:  "client.key",
		},
	}

	// Test missing file and environment variables
	t.Run("Test missing file and environment variables", func(t *testing.T) {
		// load the configuration file
		_, err := LoadConfig(dirName)
		if err == nil {
			t.Errorf("Error expected")
			return
		}
		ex, ok := err.(errorex.EX)
		if !ok {
			t.Errorf("Error type expected")
			return
		}
		if ex.Code() != ErrorCodeEnvConfigNotLoaded {
			t.Errorf("Error code expected")
			return
		}
		if ex.Detail() != "DB_TYPE environment variable not found" {
			t.Errorf("Error detail expected")
			return
		}
	})

	// Test json configuration file
	t.Run("Test json configuration file", func(t *testing.T) {

		dbconfigstr := `{
			"type": "POSTGRESQL",
			"host": "localhost",
			"port": 5432,
			"user": "postgres",
			"password": "postgres",
			"database": "postgres",
			"ssl": {
				"mode": "verify-full",
				"ca": "ca.crt",
				"key": "client.key",
				"cert": "client.crt"
			}
		}`

		// write the json configuration file
		err := files.WriteFile(dirName+"/dbconfig.json", dbconfigstr)

		if err != nil {
			t.Errorf("Error writing json configuration file: %s", err.Error())
			return
		}

		// load the configuration file
		loadedConfig, err := LoadConfig(dirName)
		if err != nil {
			ex, ok := err.(errorex.EX)
			if !ok {
				t.Errorf("Error loading json configuration file: %s", err.Error())
			} else {
				t.Errorf("Error loading json configuration file: %s; Detail: %s", ex.Error(), ex.Detail())
			}
			return
		}
		// compare the loaded configuration with the sample Config struct
		if !config.compare(loadedConfig) {
			t.Errorf("The loaded configuration is different from the sample configuration")
			return
		}

		// delete the json configuration file
		err = os.Remove(dirName + "/dbconfig.json")
		if err != nil {
			t.Errorf("Error deleting json configuration file: %s", err.Error())
			return
		}

	})

	// Test load by environment variables
	t.Run("Test load by environment variables", func(t *testing.T) {
		t.Setenv("DB_TYPE", "POSTGRESQL")
		t.Setenv("DB_HOST", "localhost")
		t.Setenv("DB_PORT", "5432")
		t.Setenv("DB_USER", "postgres")
		t.Setenv("DB_PASSWORD", "postgres")
		t.Setenv("DB_DATABASE", "postgres")
		t.Setenv("DB_SSL_MODE", "disable")

		configWithoutSSL := Config{
			Type:     DbTypePostgres,
			Host:     "localhost",
			Port:     5432,
			Database: "postgres",
			User:     "postgres",
			Password: "postgres",
		}

		// load the configuration file
		loadedConfig, err := LoadConfig(dirName)
		if err != nil {
			ex, ok := err.(errorex.EX)
			if !ok {
				t.Errorf("Error loading json configuration file: %s", err.Error())
			} else {
				t.Errorf("Error loading json configuration file: %s; Detail: %s", ex.Error(), ex.Detail())
			}
			return
		}
		// compare the loaded configuration with the sample Config struct
		if !configWithoutSSL.compare(loadedConfig) {
			t.Errorf("The loaded configuration is different from the sample configuration")
			return
		}
	})

	// Test load by environment variables with SSL verify-ca
	t.Run("Test load by environment variables with SSL verify-ca", func(t *testing.T) {
		t.Setenv("DB_TYPE", "POSTGRESQL")
		t.Setenv("DB_HOST", "localhost")
		t.Setenv("DB_PORT", "5432")
		t.Setenv("DB_USER", "postgres")
		t.Setenv("DB_PASSWORD", "postgres")
		t.Setenv("DB_DATABASE", "postgres")
		t.Setenv("DB_SSL_MODE", "verify-ca")
		t.Setenv("DB_SSL_CA", "ca.crt")

		configWithSSL := Config{
			Type:     DbTypePostgres,
			Host:     "localhost",
			Port:     5432,
			Database: "postgres",
			User:     "postgres",
			Password: "postgres",
			SSL: &SSLConfig{
				Mode: SSLModeVerifyCA,
				Ca:   "ca.crt",
			},
		}

		// load the configuration file
		loadedConfig, err := LoadConfig(dirName)
		if err != nil {
			ex, ok := err.(errorex.EX)
			if !ok {
				t.Errorf("Error loading json configuration file: %s", err.Error())
			} else {
				t.Errorf("Error loading json configuration file: %s; Detail: %s", ex.Error(), ex.Detail())
			}
			return
		}

		// compare the loaded configuration with the sample Config struct
		if !configWithSSL.compare(loadedConfig) {
			t.Errorf("The loaded configuration is different from the sample configuration")
			return
		}

	})

	// Test load by environment variables with SSL verify-full
	t.Run("Test load by environment variables with SSL verify-full", func(t *testing.T) {
		t.Setenv("DB_TYPE", "POSTGRESQL")
		t.Setenv("DB_HOST", "localhost")
		t.Setenv("DB_PORT", "5432")
		t.Setenv("DB_USER", "postgres")
		t.Setenv("DB_PASSWORD", "postgres")
		t.Setenv("DB_DATABASE", "postgres")
		t.Setenv("DB_SSL_MODE", "verify-full")
		t.Setenv("DB_SSL_CA", "ca.crt")
		t.Setenv("DB_SSL_KEY", "client.key")
		t.Setenv("DB_SSL_CERT", "client.crt")

		configWithSSL := Config{
			Type:     DbTypePostgres,
			Host:     "localhost",
			Port:     5432,
			Database: "postgres",
			User:     "postgres",
			Password: "postgres",
			SSL: &SSLConfig{
				Mode: SSLModeVerifyFull,
				Ca:   "ca.crt",
				Cert: "client.crt",
				Key:  "client.key",
			},
		}

		// load the configuration file
		loadedConfig, err := LoadConfig(dirName)
		if err != nil {
			ex, ok := err.(errorex.EX)
			if !ok {
				t.Errorf("Error loading json configuration file: %s", err.Error())
			} else {
				t.Errorf("Error loading json configuration file: %s; Detail: %s", ex.Error(), ex.Detail())
			}
			return
		}

		// compare the loaded configuration with the sample Config struct
		if !configWithSSL.compare(loadedConfig) {
			t.Errorf("The loaded configuration is different from the sample configuration")
			return
		}

	})

}