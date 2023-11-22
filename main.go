package main

import (
	"dubo/randomapi"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// 调用randomapi.go中的函数
	// 测试一百次
	// 统计值为小和的次数

	var n int = 4800
	var small, big, odd, even int
	f, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	for i := 1; i <= n; i++ {
		var flag, flag2 string
		a, b := randomapi.Randapi()
		if b < 14 {
			small++
			flag = "小"
		} else {
			big++
			flag = "大"
		}
		if b%2 == 0 {
			flag2 = "双"
			even++
		} else {
			flag2 = "单"
			odd++
		}
		// 在本地目录下生成一个log.txt文件，将结果写入
		// 将结果写入文件 格式为 第i 次结果为 a b
		if _, err = f.WriteString("第" + strconv.Itoa(i) + "次结果为" + a + " " + strconv.Itoa(b) + " " + flag + " " + flag2 + "\n"); err != nil {
			panic(err)
		}

	}
	// 控制台输出结果，并计算频率,并写入到log.txt中
	fmt.Printf("小的次数为%d,出现频率为:%f\n", small, float64(small)/float64(n))
	fmt.Printf("大的次数为%d,出现频率为:%f\n", big, float64(big)/float64(n))
	fmt.Printf("单的次数为%d,出现频率为:%f\n", odd, float64(odd)/float64(n))
	fmt.Printf("双的次数为%d,出现频率为:%f\n", even, float64(even)/float64(n))

	if _, err = f.WriteString("小的次数为" + strconv.Itoa(small) + ",出现频率为:" + strconv.FormatFloat(float64(small)/float64(n), 'f', 6, 64) + "\n"); err != nil {
		panic(err)
	}
	if _, err = f.WriteString("大的次数为" + strconv.Itoa(big) + ",出现频率为:" + strconv.FormatFloat(float64(big)/float64(n), 'f', 6, 64) + "\n"); err != nil {
		panic(err)
	}
	if _, err = f.WriteString("单的次数为" + strconv.Itoa(odd) + ",出现频率为:" + strconv.FormatFloat(float64(odd)/float64(n), 'f', 6, 64) + "\n"); err != nil {
		panic(err)
	}
	if _, err = f.WriteString("双的次数为" + strconv.Itoa(even) + ",出现频率为:" + strconv.FormatFloat(float64(even)/float64(n), 'f', 6, 64) + "\n"); err != nil {
		panic(err)
	}
	// 打上分割线
	if _, err = f.WriteString("--------------------------------------------------\n"); err != nil {
		panic(err)
	}
	// a, b := randomapi.Randapi()
	// fmt.Println(a, b)

}
