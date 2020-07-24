package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func dbConnect() (db *sqlx.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "docker"
	dbName := "stackoverflow"
	db, err := sqlx.Open(dbDriver, dbUser+":"+dbPass+"@tcp(localhost:6767)/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	if err := db.Ping(); err != nil {
		panic(err.Error())
	}
	return db
}

type rowData struct {
	ParmName string `db:"Parm_name"`
	ParamVal string `db:"Param_val"`
	ParamID  int64  `db:"Param_id"`
}

func main() {
	db := dbConnect()
	rows, err := db.Queryx("select * from example")
	if err != nil {
		fmt.Println(err)
		return
	}
	var values []rowData
	for rows.Next() {
		var row rowData
		err := rows.StructScan(&row)
		if err != nil {
			fmt.Println(err)
			return
		}
		values = append(values, row)
	}
	fmt.Println("Data present in MYSQL", values)
}
