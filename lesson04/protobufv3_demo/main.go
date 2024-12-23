package main

import (
	"fmt"
	"protobufv3_demo/pb"

	"google.golang.org/protobuf/encoding/protojson"
)


func main()  {
	searchRequest := pb.SearchRequest{Query: "hello world", Corpus: pb.SearchRequest_LOCAL}
	searchRequest.Corpus = pb.SearchRequest_NEWS
	jsonBytes, err := protojson.Marshal(&searchRequest)
	if err!= nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(string(jsonBytes))
}