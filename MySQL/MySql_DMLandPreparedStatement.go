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

	// INSERT 문 실행
	result, err := db.Exec("INSERT INTO test1 VALUES (?, ?)", 11, "Jack")
	if err != nil {
		log.Fatal(err)
	}
	
	// sql.Result.RowsAffected() 체크
	n, err := result.RowsAffected()
	if n == 1 {
		fmt.Println("1 row inserted.")
	}

	// db 서버에 placeholder 를 가진 sql문을 미리 준비시키기
	// Prepared Statement 생성
	stmt, err := db.Prepare("UPDATE test1 SET name=? WHERE id=?")
	checkError(err)
	defer stmt.Close()

	// Prepared Statement 실행
	_, err = stmt.Exec("Tom", 1) // Placerholder 파라미터 순서대로 전달
	checkError(err)
	_, err = stmt.Exec("Jack", 2)
	checkError(err)
	_, err = stmt.Exec("SHawm", 3)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}