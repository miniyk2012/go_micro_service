package main

import (
	"fmt"
	api "lesson11/proto"
	"reflect"

	"github.com/iancoleman/strcase"
	fieldmask_utils "github.com/mennanov/fieldmask-utils"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func oneofDemo() {
	req := api.NoticeReaderRequest{
		Msg:       "Hello, Reader!",
		NoticeWay: &api.NoticeReaderRequest_Email{Email: "yk_ecust_2007@163.com"},
	}
	switch req.NoticeWay.(type) {
	case *api.NoticeReaderRequest_Email:
		fmt.Println("Notice via Email:", req.GetEmail())
	case *api.NoticeReaderRequest_Phone:
		fmt.Println("Notice via Phone:", req.GetPhone())
	default:
		fmt.Println("Unknown notice way")
	}
	fmt.Printf("email: %v, phone: %v\n", req.GetEmail(), req.GetPhone())
}

func wrapperDemo() {
	book1 := api.Book{
		Title:  "Go Programming",
		Price1: proto.Int64(100),
		Price2: &wrapperspb.Int64Value{Value: 200},
		Price3: 100,
	}
	book2 := api.Book{
		Title: "Go Programming",
	}
	if book2.GetPrice2() == nil {
		fmt.Println("book2 with no price")
	}
	fmt.Printf("book1 price: %+v\n", book1.GetPrice2().GetValue())
	fmt.Printf("book1.price1 is set: %t\n", book1.Price1 != nil)
	fmt.Printf("book2.price1 is not set: %t\n", book2.Price1 == nil)
}

// fieldMaskDemo 使用field_mask实现部分更新实例
func fieldMaskDemo() {
	// client
	paths := []string{"title", "info.b"}
	req := api.UpdateBookRequest{
		Op: "yangkai",
		// 我只想改title和info.b字段
		Book: &api.Book{
			Title:  "New Title",
			Author: "yangkai.04", // 不在paths中, server端可以不接收
			Info:   &api.Book_Info{B: "New B"},
		},
		UpdateMask: &fieldmaskpb.FieldMask{Paths: paths},
	}
	fmt.Printf("originBook: %+v\n", req.Book)
	// server, 需要依赖三方库
	mask, _ := fieldmask_utils.MaskFromProtoFieldMask(req.UpdateMask, strcase.ToCamel)
	var bookDst = make(map[string]interface{})
	var updateBook = &api.Book{}
	fieldmask_utils.StructToMap(mask, req.Book, bookDst)
	fmt.Printf("bookDst: %+v\n", bookDst)
	fieldmask_utils.StructToStruct(mask, req.Book, updateBook)
	fmt.Printf("updateBook: %+v\n", updateBook)
	fmt.Printf("book.info is %v", reflect.TypeOf(bookDst["Info"]))
	if infoMap, ok := bookDst["Info"].(map[string]any); ok {
		fmt.Printf(", info.b = %v\n", infoMap["B"])
	}
}

func main() {
	//oneofDemo()
	//wrapperDemo()
	fieldMaskDemo()
}
