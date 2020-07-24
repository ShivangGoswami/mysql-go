package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// RUN these queries in sql first
// -> create database stackoverflow;
// -> use stackoverflow;
// -> create table example(
//  -> Parm_name varchar(50) NOT NULL,
//  -> Param_val varchar(50) NOT NULL,
//  -> Param_id integer primary key);
// ->insert into example values('IP','127.0.0.1',205);
// ->insert into example values('log level','2',206);

func dbConnect() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "docker"
	dbName := "stackoverflow"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp(localhost:6767)/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	if err := db.Ping(); err != nil {
		panic(err.Error())
	}
	return db
}

type rowData struct {
	ParmName string
	ParamVal string
	ParamId  int64
}

func main() {
	db := dbConnect()
	rows, err := db.Query("select * from example")
	if err != nil {
		fmt.Println(err)
		return
	}
	var values []rowData
	for rows.Next() {
		var row rowData
		err := rows.Scan(&row.ParmName, &row.ParamVal, &row.ParamId)
		if err != nil {
			fmt.Println(err)
			return
		}
		values = append(values, row)
	}
	fmt.Println("Data present in MYSQL", values)
}
