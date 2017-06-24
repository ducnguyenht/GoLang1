package main

import "fmt"

func wrapper() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func main() {
	increment := wrapper()
	fmt.Println(increment())
	fmt.Println(increment())
}

/*
closure cho phép giới hạn vùng truy cập biến sử dụng bới nhiều hàm,
ko có closure thì 1 hoặc nhiều hàm có thể truy cập vào cùng 1 biến,
mà biến đó phải được gói trong 1 vùng truy cập
closure helps us limit the scope of variables used by multiple functions
without closure, for two or more funcs to have access to the same variable,
that variable would need to be package scope
*/
