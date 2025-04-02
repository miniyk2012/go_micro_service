package main

import (
	"bufio"
	"fmt"
	"os"

	"demo/proto/book"
	"demo/proto/sniff"

	"google.golang.org/protobuf/encoding/protodelim"
)

func main() {
	println(decodeFromFile("./part-6-0"))
	println(decodeFromFile("./part-3-0"))
}

func foo() {
	price := &book.Price{
		MarketPrice: 100,
		SalePrice:   80,
	}
	fmt.Println(price)
}

func decodeFromFile(fileName string) int {
	var sniffList []*sniff.ReqResp
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	br := bufio.NewReader(file)
	for {
		reqResp := &sniff.ReqResp{}
		err = protodelim.UnmarshalFrom(br, reqResp)
		if err != nil {
			break
		}
		sniffList = append(sniffList, reqResp)
	}
	fmt.Println(sniffList)
	return len(sniffList)
}
