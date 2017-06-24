package main

import "fmt"

func main() {

	a := 43

	fmt.Println(a)  // 43
	fmt.Println(&a) // 0x20818a220

	var b = &a
	fmt.Println(b)  // 0x20818a220
	fmt.Println(*b) // 43

	*b = 42 /* gán giá trị tại vùng nhớ mà b trỏ tới*/
	// b says, "The value at this address, change it to 42"
	fmt.Println(a) // 42
	/*cách này tiện dụng
	  ta có thể truyền địa chỉ vùng nhớ thay vì 1 nhóm giá trị( truyền tham chiếu)
	  và ta có thể thay đổi giá trị tại vùng nhớ được lưu trữ đó.
	  điều này làm tăng hiệu suất
	  ta ko cần phải truyền 1 lượng lớn dữ liệu
	  mà chỉ cần truyền địa chỉ vùng nhớ
	*/

	// this is useful
	// we can pass a memory address instead of a bunch of values (we can pass a reference)
	// and then we can still change the value of whatever is stored at that memory address
	// this makes our programs more performant
	// we don't have to pass around large amounts of data
	// we only have to pass around addresses

	// everything is PASS BY VALUE in go, btw
	// when we pass a memory address, we are passing a value
}
