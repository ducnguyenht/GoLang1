package main

import "fmt"

func main() {

	a := 43

	fmt.Println(a)  // 43
	fmt.Println(&a) // 0xc0420401d0

	var b = &a
	fmt.Println(b)  // 0xc0420401d0
	fmt.Println(&b) //0xc042058020
	fmt.Println(*b) // 43
	/*
	   b là con trỏ dạng int trỏ tới địa chỉ vùng nhớ mà giá trị tại đó được lưu trữ dạng int
	   để lấy giá trị được lưu trữ tại vùng nhớ đó thì sử dụng : *b
	   dấu * là toán tử
	*/

	// b is an int pointer;
	// b points to the memory address where an int is stored
	// to see the value in that memory address, add a * in front of b
	// this is known as dereferencing
	// the * is an operator in this case
}
