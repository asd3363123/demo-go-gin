package database

import (
	"database/sql"
	"log"
)

//因为我们需要在别的地方使用SqlDB这个变量，因此依照golang的习惯，变量名必须大写开头
var SqlDB *sql.DB

func init() {
	var err error
	SqlDB, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test?parseTime=true")
	if err != nil {
		log.Fatal(err.Error())
	}
	err = SqlDB.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
}
