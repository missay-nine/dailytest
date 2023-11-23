package main

import "fmt"

func A1(a int) {
	fmt.Println(a)
}

// func A() {
// 	a, b := 1, 2
// 	defer A1(a)
// 	a = a + b
// 	fmt.Println(a, b)
// }
func A() {
	a, b := 1, 2
	defer func(b int) {
		a = a + b
		fmt.Println(a, b)
	}(b)
	a = a + b
	fmt.Println(a, b)
}

func f() (x int) {
	defer func() {
		x++ // 修改返回值x
	}()
	return 1 // 返回值是x，不是0
}

func main() {
	//A()
	fmt.Println(f())
}
