package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

func main() {
	db, err := sql.Open("mysql", os.Getenv("MYSQL_CONN"))
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		panic(err.Error())
	}

	var cycles = []*Cycle{}

	var sql string
	sql = "SELECT nid, title FROM cycles"
	rows, err := db.Query(sql)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var nid int
		var title string
		if err := rows.Scan(&nid, &title); err != nil {
			panic(err.Error())
		}
		cycle := NewCycle(nid, title)
		if err := cycle.loadMotets(db); err != nil {
			panic(err.Error())
		}
		cycles = append(cycles, cycle)
	}
	if err := rows.Err(); err != nil {
		panic(err.Error())
	}
	for _, cycle := range cycles {
		println(cycle.String())
	}
}
