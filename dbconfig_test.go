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
	"testing"
)

func TestDbConfig(t *testing.T) {

	// Test marshal and unmarshal
	t.Run("Test marshal and unmarshal", func(t *testing.T) {
		config := Config{
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

		// encode config to json
		configBytes, err := json.Marshal(config)
		if err != nil {
			t.Errorf("Error marshaling config: %s", err.Error())
			return
		}

		// decode config from json
		var config2 Config
		if err := json.Unmarshal(configBytes, &config2); err != nil {
			t.Errorf("Error unmarshaling config: %s", err.Error())
			return
		}

		// compare config and config2
		if !config.compare(config2) {
			t.Errorf("The unmarshaled config is different from the original config")
			return
		}
	})

}
