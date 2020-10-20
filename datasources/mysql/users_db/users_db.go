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

	_ "github.com/go-sql-driver/mysql"
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

var (
	hosName        string
	dataSourceName string
)

func init() {
	fmt.Println("host name is")
	hostName, _ := os.Hostname()
	fmt.Println(hostName)

	if hostName == "Apple-Macintosh-2.local" {
		// for dev only
		dataSourceName = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", "root", "1980#$Life", "127.0.0.1", "gogolang")

	} else {
		// for production
		dataSourceName = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", username, password, mysql_host, mysql_schema)

	}

	var err error
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Printf("mySQL connect successfully on %s", hostName)
}
