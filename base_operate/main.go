package main

import (
	"fmt"

	"./test"
)

// 测试golang的一些基本操作
func main() {
	fmt.Println("test script staring")
	fmt.Println("")

	test.Test_array()
	test.Test_slice()

	fmt.Println("")
	fmt.Println("test script ending")
}
