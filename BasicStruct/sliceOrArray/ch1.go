package main

import "fmt"

func main() {

	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := a[0:3]
	fmt.Println(a, b)
	b = append(b, 11, 12, 13)
	fmt.Print(len(b), cap(b), "\n")
	fmt.Println(a, b)
}

// 输出结果是：

// [1 2 3 4 5 6 7 8 9 10] [1 2 3]
// 6 10
// [1 2 3 11 12 13 7 8 9 10] [1 2 3 11 12 13]

// 解释如下：

// - a 是一个长度为 10 的数组，b 是 a 的一个切片，表示 a 的前三个元素。
// - append 函数可以向切片追加元素，如果切片的容量足够，就直接在原数组上修改，如果不够，就会创建一个新的数组，并复制原数组的元素，然后返回新的切片¹³。
// - b 的初始长度和容量都是 3，当向 b 追加 11，12，13 时，b 的长度变为 6，但容量还是 10，因为 a 的容量是 10，所以 b 还可以在 a 的空间上扩展。
// - 追加元素后，b 的元素变为 [1 2 3 11 12 13]，同时也修改了 a 的前六个元素，所以 a 的元素变为 [1 2 3 11 12 13 7 8 9 10]。
