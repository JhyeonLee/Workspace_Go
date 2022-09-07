package main

import (
	"io"
	"io/ioutil"
	"os"
)

func main() {
	//입력파일 열기
	fi, err := os.Open("./1.txt")
	if err != nil {
		panic(err)
	}
	defer fi.Close()

	//출력파일 생성
	fo, err := os.Create("./2.txt")
	if err != nil {
		panic(err)
	}
	defer fo.Close()

	buff := make([]byte, 1024)

	// 루프
	for {
		//읽기
		cnt, err := fi.Read(buff)
		if err != nil && err != io.EOF {
			panic(err)
		}

		// 끝이면 루프 종료
		if cnt == 0 {
			break
		}

		// 쓰기
		_, err = fo.Write(buff[:cnt])
		if err != nil {
			panic(err)
		}
	}

	// 파일 읽기
	bytes, err := ioutil.ReadFile("./3.txt")
	if err != nil {
		panic(err)
	}
	// 파일 쓰기
	err = ioutil.WriteFile("./4.txt", bytes, 0)
	if err != nil {
		panic(err)
	}
}
