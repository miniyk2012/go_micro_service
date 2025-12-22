package main

import (
	"fmt"
	"lesson41/privateConfig"
)

func main() {
	cf := privateConfig.NewConfigWithOptions(privateConfig.WithA("a"), privateConfig.WithB("B"), privateConfig.WithC(2))
	fmt.Printf("%+v\n", *cf)

	cf2 := privateConfig.NewConfigWithFuncOptions(privateConfig.WithD(50))
	fmt.Printf("%+v\n", *cf2)
}
