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
