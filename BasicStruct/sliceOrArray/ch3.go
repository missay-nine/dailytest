package main

import "fmt"

func main() {

	// 探索go slice 扩容规则
	// s := make([]int, 0)
	// oldCap := cap(s)
	// for i := 0; i < 2048; i++ {
	// 	s = append(s, i)
	// 	newCap := cap(s)
	// 	if newCap != oldCap {
	// 		fmt.Printf("[%d->%4d]cap=%-4d | after append %-4d cap=%-4d\n", 0, i-1, oldCap, i, newCap)
	// 		oldCap = newCap
	// 	}
	// }

	s := []int{}
	s = append(s, []int{1, 2, 3, 4, 5}...)
	fmt.Println(cap(s)) // 结果为6
	//内存分配各个size如下：{0, 8, 16, 24, 32, 48, 64, 80, 96, 112...}
}
