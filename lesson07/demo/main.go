package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"

	"demo/proto/book"
	"demo/proto/sniff"

	"google.golang.org/protobuf/encoding/protodelim"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	if decodeFromFile("./part-6-0") != 1 {
		panic("decodeFromFile() != 1")
	}
	if decodeFromFile("./part-3-0") != 2 {
		panic("decodeFromFile() != 2")
	}
	if writeBatchThenDecode() != 10 {
		panic("writeBatchThenDecode() != 10")
	}
}

func foo() {
	price := &book.Price{
		MarketPrice: 100,
		SalePrice:   80,
	}
	fmt.Println(price)
}

func printJson(sniffList []*sniff.ReqResp) {
	for _, reqResp := range sniffList {
		v, _ := protojson.Marshal(reqResp)
		fmt.Printf("%s\n", v)
	}

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
	printJson(sniffList)
	println()
	return len(sniffList)
}

func writeBatchThenDecode() int {
	buf := &bytes.Buffer{}
	for i := 0; i < 10; i++ {
		reqResp := &sniff.ReqResp{}
		reqResp.Id = fmt.Sprintf("taskId-%d", i)
		reqResp.Request = []byte("abc")
		if n, err := protodelim.MarshalTo(buf, reqResp); err != nil {
			panic(fmt.Sprintf("protodelim.MarshalTo(_, %v) = %d, %v", reqResp, n, err))
		}
	}

	var sniffList []*sniff.ReqResp
	for {
		reqResp := &sniff.ReqResp{}
		err := protodelim.UnmarshalFrom(buf, reqResp)
		if err != nil {
			break
		}
		sniffList = append(sniffList, reqResp)
	}
	printJson(sniffList)
	println()
	return len(sniffList)
}
