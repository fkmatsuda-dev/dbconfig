/*
 * © 2023 fkmatsuda

 * This file is licensed under the terms of the MIT license. Permission is hereby
 * granted, free of charge, to any person obtaining a copy of this software and
 * associated documentation files (the "Software"), to deal in the Software without
 * restriction, including without limitation the rights to use, copy, modify,
 * merge, publish, and/or distribute copies of the Software, and to permit persons
 * to whom the Software is furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS," WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
 * WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
 * CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 *
 */

package dbconfig

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/fkmatsuda-dev/commons/errorex"
	"github.com/fkmatsuda-dev/commons/files"
	"github.com/fkmatsuda-dev/env"
)

var (
	// configFormarts is the order of the configuration file formats to be loaded
	configFormats = []string{"json"}
)

// LoadConfig loads the database settings and returns a struct Config
// tries to load the configuration from the dbconfig.json file and if the file does not exist it will try to load it from the environment variables
func LoadConfig(path string) (Config, error) {
	// search for the configuration file
	configFile, err := searchConfigFile(path)
	if err != nil {
		if errorex.IS(err, ErrorCodeConfigFileNotFound) {
			// try to load the environment variables
			return LoadFromEnv()
		}
		return Config{}, err
	}
	// load the configuration file
	return loadConfigFile(configFile)
}

func loadConfigFile(file string) (Config, error) {
	// read the configuration file
	fileContent, err := files.ReadFile(file)
	if err != nil {
		return Config{}, errorex.New(
			ErrorCodeConfigFileNotLoaded,
			"Configuration file not loaded",
			err.Error(),
		)
	}
	// Unmarshal the configuration file
	var config Config
	err = json.Unmarshal([]byte(fileContent), &config)
	if err != nil {
		return Config{}, errorex.New(
			ErrorCodeConfigFileParseError,
			"Configuration file parse error",
			err.Error(),
		)
	}
	return config, nil
}

const (
	configurationParseError = "Environment configuration parse error"
	envNotLoaded            = "Environment configuration not loaded"
)

// LoadFromEnv loads the database settings from the environment variables and returns a struct Config
func LoadFromEnv() (Config, error) {
	config := Config{}
	// load the database type
	strType, chk := env.ChkString("DB_TYPE")
	if !chk {
		return Config{}, errorex.New(ErrorCodeEnvConfigNotLoaded, envNotLoaded, "DB_TYPE environment variable not found")
	}
	dbType, err := ParseDbType(strType)
	if err != nil {
		return Config{}, errorex.New(ErrorCodeEnvConfigParseError, configurationParseError, err.Error())
	}

	// load the database host
	dbHost, chk := env.ChkString("DB_HOST")
	if !chk {
		return Config{}, errorex.New(
			ErrorCodeEnvConfigNotLoaded,
			envNotLoaded,
			"DB_HOST environment variable not found",
		)
	}

	// load the database port
	dbPort, err := env.Int("DB_PORT", 5432)
	if err != nil {
		return Config{}, errorex.New(ErrorCodeEnvConfigParseError, configurationParseError, err.Error())
	}

	// load the database name
	dbDatabase, chk := env.ChkString("DB_DATABASE")
	if !chk {
		return Config{}, errorex.New(
			ErrorCodeEnvConfigNotLoaded,
			envNotLoaded,
			"DB_DATABASE environment variable not found",
		)
	}

	// load the database user
	dbUser, chk := env.ChkString("DB_USER")
	if !chk {
		return Config{}, errorex.New(
			ErrorCodeEnvConfigNotLoaded,
			envNotLoaded,
			"DB_USER environment variable not found",
		)
	}

	// load the database password
	dbPassword, chk := env.ChkString("DB_PASSWORD")
	if !chk {
		return Config{}, errorex.New(
			ErrorCodeEnvConfigNotLoaded,
			envNotLoaded,
			"DB_PASSWORD environment variable not found",
		)
	}
	// fill the configuration struct
	config.Type = dbType
	config.Host = dbHost
	config.Port = uint16(dbPort)
	config.Database = dbDatabase
	config.User = dbUser
	config.Password = dbPassword

	return loadSSL(chk, err, config)
}

func loadSSL(chk bool, err error, config Config) (Config, error) {
	// load the database ssl mode
	strSSLMode, chk := env.ChkString("DB_SSL_MODE")
	if !chk {
		strSSLMode = "disable"
	}
	sslMode, err := ParseSSLMode(strSSLMode)
	if err != nil {
		return Config{}, errorex.New(ErrorCodeEnvConfigParseError, configurationParseError, err.Error())
	}

	// if ssl mode is verify-full, load the ssl certificate and key
	var sslCert, sslKey string
	if sslMode == SSLModeVerifyFull {
		sslCert, chk = env.ChkString("DB_SSL_CERT")
		if !chk {
			return Config{}, errorex.New(
				ErrorCodeEnvConfigNotLoaded,
				envNotLoaded,
				"DB_SSL_CERT environment variable not found",
			)
		}
		sslKey, chk = env.ChkString("DB_SSL_KEY")
		if !chk {
			return Config{}, errorex.New(
				ErrorCodeEnvConfigNotLoaded,
				envNotLoaded,
				"DB_SSL_KEY environment variable not found",
			)
		}
	}

	// if ssl mode is verify-ca or verify-full, load the ssl root certificate
	var sslCa string
	if sslMode == SSLModeVerifyCA || sslMode == SSLModeVerifyFull {
		sslCa, chk = env.ChkString("DB_SSL_CA")
		if !chk {
			return Config{}, errorex.New(
				ErrorCodeEnvConfigNotLoaded,
				envNotLoaded,
				"DB_SSL_CA environment variable not found",
			)
		}
	}

	// if ssl mode is verify-ca or verify-full, fill the ssl struct
	if sslMode == SSLModeVerifyCA || sslMode == SSLModeVerifyFull {
		sslConfig := SSLConfig{
			Mode: sslMode,
			Ca:   sslCa,
		}
		if sslMode == SSLModeVerifyFull {
			sslConfig.Cert = sslCert
			sslConfig.Key = sslKey
		}
		config.SSL = &sslConfig
	} else {
		config.SSL = nil
	}

	return config, nil
}

func searchConfigFile(path string) (string, error) {
	// iterate over the configuration file formats
	for _, format := range configFormats {
		// search for the dbconfig file
		configFile := fmt.Sprintf("%s/dbconfig.%s", path, format)
		if _, err := os.Stat(configFile); err == nil {
			return configFile, nil
		}
	}
	return "", errorex.New(ErrorCodeConfigFileNotFound, "Configuration file not found", "")
}
