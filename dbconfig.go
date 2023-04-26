package dbconfig

import (
	"encoding/json"
	"fmt"
	"github.com/fkmatsuda-dev/commons/errorex"
)

type DbType int8

const (
	DbTypeMysql DbType = iota + 1
	DbTypePostgres
	DbTypeCockroachdb
)

var DbTypeName = map[DbType]string{
	DbTypeMysql:       "MYSQL",
	DbTypePostgres:    "POSTGRESQL",
	DbTypeCockroachdb: "COCKROACHDB",
}

var DbTypeValue = map[string]DbType{
	DbTypeName[DbTypeMysql]:       DbTypeMysql,
	DbTypeName[DbTypePostgres]:    DbTypePostgres,
	DbTypeName[DbTypeCockroachdb]: DbTypeCockroachdb,
}

// String returns the string value of the DbType
func (s DbType) String() string {
	return DbTypeName[s]
}

// ParseDbType parses the string value to a DbType
func ParseDbType(value string) (DbType, error) {
	if s, ok := DbTypeValue[value]; ok {
		return s, nil
	}
	return DbType(0), errorex.New(
		ErrorCodeDbTypeParseError,
		"DbType parse error",
		fmt.Sprintf("\"%s\" value for DbType is invalid",
			value,
		),
	)
}

// MarshalJSON marshals the enum as a quoted json string
func (s DbType) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (s *DbType) UnmarshalJSON(data []byte) error {
	var value string
	if err := json.Unmarshal(data, &value); err != nil {
		return errorex.New(ErrorCodeDbTypeParseError, "DbType parse error", err.Error())
	}
	parsed, err := ParseDbType(value)
	if err != nil {
		return err
	}
	*s = parsed
	return nil
}

type SSLMode int8

const (
	SSLModeDisable SSLMode = iota + 1
	SSLModeAllow
	SSLModePrefer
	SSLModeRequire
	SSLModeVerifyCA
	SSLModeVerifyFull
)

var SSLModeName = map[SSLMode]string{
	SSLModeDisable:    "disable",
	SSLModeAllow:      "allow",
	SSLModePrefer:     "prefer",
	SSLModeRequire:    "require",
	SSLModeVerifyCA:   "verify-ca",
	SSLModeVerifyFull: "verify-full",
}

var SSLModeValue = map[string]SSLMode{
	SSLModeName[SSLModeDisable]:    SSLModeDisable,
	SSLModeName[SSLModeAllow]:      SSLModeAllow,
	SSLModeName[SSLModePrefer]:     SSLModePrefer,
	SSLModeName[SSLModeRequire]:    SSLModeRequire,
	SSLModeName[SSLModeVerifyCA]:   SSLModeVerifyCA,
	SSLModeName[SSLModeVerifyFull]: SSLModeVerifyFull,
}

// String returns the string value of the SSLMode
func (s SSLMode) String() string {
	return SSLModeName[s]
}

// ParseSSLMode parses the string value to a SSLMode
func ParseSSLMode(value string) (SSLMode, error) {
	if s, ok := SSLModeValue[value]; ok {
		return s, nil
	}
	return SSLMode(0), errorex.New(
		ErrorCodeSSLModeParseError,
		"SSLMode parse error",
		fmt.Sprintf("\"%s\" value for SSLMode is invalid",
			value,
		),
	)
}

// MarshalJSON marshals the enum as a quoted json string
func (s SSLMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (s *SSLMode) UnmarshalJSON(data []byte) error {
	var value string
	if err := json.Unmarshal(data, &value); err != nil {
		return errorex.New(ErrorCodeSSLModeParseError, "SSLMode parse error", err.Error())
	}
	parsed, err := ParseSSLMode(value)
	if err != nil {
		return err
	}
	*s = parsed
	return nil
}

type SSLConfig struct {
	Mode SSLMode
	Cert string
	Key  string
	Ca   string
}

type Config struct {
	Type     DbType
	Host     string
	Port     uint16
	User     string
	Password string
	Database string
	SSL      *SSLConfig
}
