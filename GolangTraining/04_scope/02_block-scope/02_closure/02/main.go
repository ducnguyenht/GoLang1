package main

import "fmt"

var x = 0

func increment() int {
	x++
	return x
}

func main() {
	fmt.Println(increment())
	fmt.Println(increment())
}

/*
closure cho phép giới hạn vùng truy cập biến sử dụng bới nhiều hàm,
ko có closure thì 1 hoặc nhiều hàm có thể truy cập vào cùng 1 biến,
mà biến đó phải được gói trong 1 vùng truy cập
*/
