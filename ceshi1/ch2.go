// 测试时间生成是否正确

package main

import (
	"dubo/randomapi"
)

func main() {

	var n, gold int = 9, 61500
	a := randomapi.Randomtime2([]int{14, 3, 1, 31, 29, 1, 47, 13, 3})
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
	// 打印分割线
	println("====================================")

	a = randomapi.Randomtime(n)
	println(a)
	b, c = randomapi.Randomapi2(n, gold, a)
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
