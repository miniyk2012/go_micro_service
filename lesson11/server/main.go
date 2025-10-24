package main

import (
	"fmt"
	"lesson08/proto"
)

func main() {
	req := proto.NoticeReaderRequest{
		Msg:       "Hello, Reader!",
		NoticeWay: &proto.NoticeReaderRequest_Email{Email: "yk_ecust_2007@163.com"},
	}
	switch req.NoticeWay.(type) {
	case *proto.NoticeReaderRequest_Email:
		fmt.Println("Notice via Email:", req.GetEmail())
	case *proto.NoticeReaderRequest_Phone:
		fmt.Println("Notice via Phone:", req.GetPhone())
	default:
		fmt.Println("Unknown notice way")
	}
	fmt.Printf("email: %v, phone: %v\n", req.GetEmail(), req.GetPhone())
}
