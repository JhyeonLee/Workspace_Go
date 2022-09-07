package main

import (
    "encoding/xml"
    "fmt"
)

// Member
type Member struct {
	Name string
	Age int
	Active bool
}

func main() {
	mem := Member{"Alex", 10, true}

	xmlBytes, err := xml.Marshal(mem)
	if err != nil {
		panic(err)
	}

	xmlString := string(xmlBytes)
	fmt.Println(xmlString)

	
	// 테스트용 데이터
	xmlBytes02, _ := xml.Marshal(Member{"Tim", 1, true})

	// XML 디코딩
	var memDecoded Member
	err = xml.Unmarshal(xmlBytes02, &memDecoded)
	if err != nil {
		panic(err)
	}

	// mem02 구조체 필드 엑세스
	fmt.Println(memDecoded.Name, memDecoded.Age, memDecoded.Active)
	
}