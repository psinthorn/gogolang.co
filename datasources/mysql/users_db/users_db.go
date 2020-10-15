//TODO:
//- package name
//- Import mysql drivers
//- Create const and var for DB environment connection username, password, host, database schema
//- Create func init() for connection
//- Handling connection result error or success

package mysql_db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "gihub.com/go-sql-driver/mysql"
)

const (
	mysql_users_username = "mysql_users_username"
	mysql_users_password = "mysql_users_password"
	mysql_users_host     = "mysql_users_host"
	mysql_users_schema   = "mysql_users_schema"
)

var (
	Client *sql.DB

	username     = os.Getenv(mysql_users_username)
	password     = os.Getenv(mysql_users_password)
	mysql_host   = os.Getenv(mysql_users_host)
	mysql_schema = os.Getenv(mysql_users_schema)
)

func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", username, password, mysql_host, mysql_schema)

	var err error
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("mySQL connect successfully ")
}
