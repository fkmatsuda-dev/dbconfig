# dbconfig
### dbconfig is a project that provides structures and interfaces to store SQL and NoSQL database access configurations. With this project, you can easily manage your database connection configurations in a single place and access them in your projects.

## Installation
You can install dbconfig via go get:

```bash
go get github.com/fkmatsuda-dev/dbconfig
```
## Usage
To use dbconfig, you can load the configuration from a JSON or YAML file or directly from environment variables. Here is an example YAML configuration file:

```yaml
mysql:
  host: localhost
  port: 3306
  user: root
  password: 12345
  database: mydb

mongodb:
  host: localhost
  port: 27017
  user: root
  password: 12345
  database: mydb
```
You can also set environment variables for each database configuration, following the naming convention DB_<DATABASE_NAME>_<CONFIG_KEY>. For example, to set the MySQL host and port, you would set DB_MYSQL_HOST and DB_MYSQL_PORT.

After loading the configuration, you can use it in your project to connect to the database. Here is an example of how to use dbconfig to get the MySQL database connection configurations:

```go
import (
    "github.com/fkmatsuda-dev/dbconfig"
)

func main() {
// Load configuration from a YAML file
config, err := dbconfig.LoadFromYAMLFile("path/to/config.yml")
if err != nil {
// handle error
}

// Load configuration from environment variables
config, err := dbconfig.LoadFromEnv()
if err != nil {
// handle error
}

mysqlConfig := config.Get("mysql")
// use mysqlConfig to connect to the database
}
```
The mysqlConfig object contains the MySQL database connection information. You can use it to connect to the database in your project.

## License
This project is licensed under the MIT License. See the LICENSE file for more details.

## Contributing
If you want to contribute to dbconfig library, feel free to send a pull request or create an issue on GitHub. Your feedback is very important for us to improve the project together.

