package dbconfig

import (
	"database/sql"
	"fmt"

	"backend/logging"
	env "backend/utils"

	_ "github.com/go-sql-driver/mysql"
)

var User string = env.GetEnvironmentVars("DBUSER")
var Pass string = env.GetEnvironmentVars("DBPASS")
var Host string = env.GetEnvironmentVars("DBHOST")
var DbPort string = env.GetEnvironmentVars("DBPORT")
var Database string = env.GetEnvironmentVars("DBNAME")
var Port string = env.GetEnvironmentVars("PORT")

// DB
type DB struct {
	SQL *sql.DB
}

var dbConn = &DB{}

func Connect() (*DB, error) {

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", User, Pass, Host, DbPort, Database))
	if err != nil {
		logging.LogError(err.Error())
	}
	dbConn.SQL = db
	return dbConn, err
}
