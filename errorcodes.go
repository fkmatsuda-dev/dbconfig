package dbconfig

import "github.com/fkmatsuda-dev/commons/errorex"

const (
	ErrorCodeDbTypeParseError     = "DBCONFIG-1001"
	ErrorCodeSSLModeParseError    = "DBCONFIG-1005"
	ErrorCodeConfigFileNotFound   = "DBCONFIG-1011"
	ErrorCodeConfigFileNotLoaded  = "DBCONFIG-1012"
	ErrorCodeEnvConfigNotLoaded   = "DBCONFIG-1013"
	ErrorCodeConfigFileParseError = "DBCONFIG-1014"
	ErrorCodeEnvConfigParseError  = "DBCONFIG-1015"
)

func init() {
	// Register error codes
	errorex.RegisterErrorCode(ErrorCodeDbTypeParseError, "DbType parse error")
	errorex.RegisterErrorCode(ErrorCodeSSLModeParseError, "SSLMode parse error")
	errorex.RegisterErrorCode(ErrorCodeConfigFileNotFound, "Configuration file not found")
	errorex.RegisterErrorCode(ErrorCodeConfigFileNotLoaded, "Configuration file not loaded")
	errorex.RegisterErrorCode(ErrorCodeEnvConfigNotLoaded, "Environment configuration not loaded")
	errorex.RegisterErrorCode(ErrorCodeEnvConfigParseError, "Environment configuration cannot be parsed")
	errorex.RegisterErrorCode(ErrorCodeConfigFileParseError, "Configuration file parse error")
}
