package main

import (
    "database/sql"
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

	// 트랜잭션 시작
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback() // 중간에 에러시 롤백

	// INSERT 문 실행
	_, err = db.Exec("INSERT discounts(discounttype,discount) VALUES(?, ?)", "Test1", 12)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("INSERT discounts(discounttype,discount VALUES(?, ?)", "Test2", 11)
	if err != nil {
		log.Fatal(err)
	}

	// 트랜잭션 커밋
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
}