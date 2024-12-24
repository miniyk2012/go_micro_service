package main

import (
	"fmt"

	"protobufv3_demo/pb"

	"google.golang.org/protobuf/encoding/protojson"
)

func oneOfDemo() {
	// 初始化
	sampleMessage := &pb.SampleMessage{}
	sampleMessage.TestOneof = &pb.SampleMessage_Name{Name: "hello world"}
	fmt.Println(sampleMessage.GetTestOneof())

	sampleMessage.TestOneof = &pb.SampleMessage_SubMessage_{SubMessage: &pb.SampleMessage_SubMessage{Age: 12}}
	fmt.Println(sampleMessage.GetTestOneof())
}

func basicDemo() {
	searchRequest := pb.SearchRequest{Query: "hello world", Corpus: pb.SearchRequest_LOCAL}
	searchRequest.Corpus = pb.SearchRequest_NEWS
	jsonBytes, err := protojson.Marshal(&searchRequest)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(string(jsonBytes))
}

func main() {
	oneOfDemo()
}
