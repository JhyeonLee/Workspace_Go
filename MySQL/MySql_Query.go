package main
 
import (
    "database/sql"
    "fmt"
    "log"
    _ "github.com/go-sql-driver/mysql"
)

func main() {
	// sql.DB 객체 생성
	// db, err := sql.Open("mysql", "root:pwd@tcp(127.0.0.1:3306)/testdb")
	db, err := sql.Open("mysql", "root:pwd@unix(127.0.0.1:3306)/testdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 하나의 Row를 갖는 SQL 쿼리
	var name string
	err = db.QueryRow("SELECT name FROM test1 WHERE id = 1").Scan(&name)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(name)
	
	// 복수 Row를 갖는 SQL 쿼리
	var id int
	var name02 string
	rows, err := db.Query("SELECT if, name FROM test1 where id >= ?", 1)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close() // 반드시 닫는다 (지연하여 닫기)

	for rows.Next() {
		err := rows.Scan(&id, &name02)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println(name02)
}
