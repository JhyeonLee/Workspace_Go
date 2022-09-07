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

	// INSERT 문 실행 : eXEC()
	result, err := db.Exec("INSER INTO discounts(discounttye,discount) VALUES(?, ?)", "Test", 10)
	if err != nil {
		log.Fatal(err)
	}

	n, err := result.RowsAffected() // sql.Result.RowAffected() 체크
	if n == 1 {
		fmt.Println("Successfully inserted.")
	}

	// Prepared Statement 생성
	stmt, err := db.Prepare("UPDATE discounts Set discount=?2 WHERE discounttye=?1")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	// Prepared Statement 실행
	_, err = stmt.Exec("Test", 15) // Placeholder 피림;터 ?1, ?2 전달
	if err != nil {
		log.Fatal(err)
	}
}