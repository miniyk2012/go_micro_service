package main

import (
	"fmt"
	"time"

	"protobufv3_demo/pb"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func oneOfDemo() {
	// 初始化
	sampleMessage := &pb.SampleMessage{}
	sampleMessage.TestOneof = &pb.SampleMessage_Name{Name: "hello world"}
	fmt.Println(sampleMessage.GetName(), sampleMessage.GetTestOneof())

	sampleMessage.TestOneof = &pb.SampleMessage_SubMessage_{SubMessage: &pb.SampleMessage_SubMessage{Age: 12}}
	fmt.Println(sampleMessage.GetSubMessage(), sampleMessage.GetTestOneof())
}

func basicDemo() {
	searchRequest := pb.SearchRequest{Query: "hello world", Corpus: pb.SearchRequest_LOCAL, DateOfBirth: timestamppb.New(time.Now())}
	searchRequest.Corpus = pb.SearchRequest_NEWS
	jsonBytes, err := protojson.Marshal(&searchRequest)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Printf("json format: %s\n", jsonBytes)
	reqRespBytes, err := proto.Marshal(&searchRequest)
	fmt.Printf("bytes: %v\n", reqRespBytes)

	dateOfBirth := searchRequest.GetDateOfBirth().AsTime()
	loc, _ := time.LoadLocation("Asia/Shanghai") // UTC+8 timezone
	utc8Time := dateOfBirth.In(loc)
	fmt.Println(utc8Time.Format("2006-01-02 15:04:05"))
}

func anyDemo() {
	// Create a Result message
	result := &pb.Result{
		Url:      "https://example.com",
		Title:    "Example Page",
		Snippets: []string{"snippet1", "snippet2"},
		Week:     pb.Week_MONDAY,
	}

	// Pack the Result message into an Any
	anyResult, err := anypb.New(result)
	if err != nil {
		fmt.Println("Error packing to Any:", err)
		return
	}

	// Demonstrate unpacking Any
	var unpackedResult pb.Result
	if err := anyResult.UnmarshalTo(&unpackedResult); err != nil {
		fmt.Println("Error unpacking Any:", err)
		return
	}
	fmt.Printf("Unpacked Result: %+v\n", &unpackedResult)

	// Create a SearchResponse with the Any message
	response := &pb.SearchResponse{
		Ret:     []string{"result1", "result2"},
		Corpus:  pb.SearchRequest_WEB,
		Details: []*anypb.Any{anyResult},
	}

	// Marshal to JSON
	jsonBytes, err := protojson.Marshal(response)
	if err != nil {
		fmt.Println("Error marshaling to JSON:", err)
		return
	}
	fmt.Println("Marshaled JSON:", string(jsonBytes))

}

func main() {
	//oneOfDemo()
	//anyDemo()
	basicDemo()
}
