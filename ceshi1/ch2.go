// 测试时间生成是否正确

package main

import (
	"dubo/randomapi"
)

func main() {
	//a := randomapi.Randomtime2([]int{27, 3, 47, 44, 30})
	var n, gold int = 16, 27500

	a := randomapi.Randomtime(n)
	println(a)
	b, c := randomapi.Randomapi2(n, gold, a)
	println(b, c)
	if c < 14 {
		println("小")
	} else {
		println("大")
	}
	if c%2 == 0 {
		println("双")
	} else {
		println("单")
	}

}
