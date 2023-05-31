# dbconfig
### dbconfig is a project that provides structures and interfaces to store SQL and NoSQL database access configurations. With this project, you can easily manage your database connection configurations in a single place and access them in your projects.

## Installation
You can install dbconfig via go get:

```bash
go get github.com/fkmatsuda-dev/dbconfig
```
## Usage
To use dbconfig, you can load the configuration from a JSON file or directly from environment variables. Here is an example JSON configuration file:

```json
{
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
}
```
You can also set environment variables for each database configuration parmeter, following the naming convention DB_<PARAMETER>. For example, to set the host and port, you would set DB_HOST and DB_PORT, to set SSL parameters you can use DB_SSL_<PARAMETER>.

After loading the configuration, you can use it in your project to connect to the database. Here is an example of how to use dbconfig to get the MySQL database connection configurations:

```go
import (
    "github.com/fkmatsuda-dev/dbconfig"
)

func main() {
    // Load configuration from a dbconfig.json file if exists or from environment variables
    config, err := dbconfig.LoadConfig("path/to/config/")
    if err != nil {
    // handle error
    }
    
    // Load configuration from environment variables
    config, err := dbconfig.LoadFromEnv()
    if err != nil {
    // handle error
    }

}
```
The mysqlConfig object contains the MySQL database connection information. You can use it to connect to the database in your project.

## License
This project is licensed under the MIT License. See the LICENSE file for more details.

## Contributing
If you want to contribute to dbconfig library, feel free to send a pull request or create an issue on GitHub. Your feedback is very important for us to improve the project together.

