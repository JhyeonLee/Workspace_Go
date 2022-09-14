package main 

import (
    "bufio"
    "encoding/csv"
    "os"
)

func main() {
	// 파일 생성
	file, err := os.Create("./output.csv")
	if err != nil {
		panic(err)
	}

	// csv writer 생성
	wr := csv.NewWriter(bufio.NewWriter(file))

	// csv 내용 읽기
	wr.Write([]string{"A", "0.25"})
	wr.Write([]string{"B", "55.70"})
	wr.Flush()
}