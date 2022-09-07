package main

import (
    "database/sql"
    "fmt"
    _ "github.com/denisenkom/go-mssqldb"
    "log"
)

func main() {
	// sql.DB 객체 생성
	db, err := sql.Open("mssql", "server=(local);user id=sa;password=pws;database=pubs")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 하나의 Row를 갖는 SQL 쿼리 : QueryRow()
	var lname, fname string
	err = db.QueryRow("SELECT au_lname, au_fname FROM authors WHERE AU_ID='172-32-1176'").Scan(&lname, &fname)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fname, lname)

	// 복수 Row를 갖는 SQL 쿼리 : Query()
	rows, err := db.Query("SELECT au_lname, au_fname FROM authors WHERE au_lname=?", "Ringer")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close() //반드시 닫는다 (지연하여 닫기)

	for rows.Next() {
		err := rows.Scan(&lname, &fname)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(fname, lname)
	}
}