// 设计一种猜数字大小的 算法思路是这样的 选取五个订单的时间值，比如说 有一个订单是2023-11-20 09:20:28:058  那么这一个订单的时间值就是28058
// 然后还有个订单数，这里是5 ，下单的总金币是32700  ，那么加密方式就是 SHA256(订单数 + 总金币+时间值) 得到一个hash 值，然后将64 位hash 取中间60位，分成五段。
// 第一段转为十进制，后面依次类推  将五段十进制数求和，得到一个数 ，根据第一个数字 ，比如说第一个数字是6 从第6个数 截取三位，得到一个数，这个数就是最终的结果。，然后0-13为小 14-27为大。

package randomapi

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"math/big"
	"strconv"
)

func Randapi() (string, int) {
	// type pair struct {
	// 	x string
	// 	y int
	// } // 一个存对应字符串 一个存对应和
	//var log []pair
	ordNumInt, goldsum, timesum := randomdata()
	// ordNumInt, goldsum, timesum := 14, 12200, 378423 //自己测试用的
	ans := jiami(ordNumInt, goldsum, timesum)
	//	fmt.Println(ans)

	// 然后分成五段，从下标为2 取 每段取12位，将对应十位变成十进制
	// 每段的和
	var duanSum int64 = 0
	for i := 0; i < 5; i++ {
		ans1 := ans[2+12*i : 14+12*i]
		// 将该字符串转为十进制
		ans2, _ := strconv.ParseInt(ans1, 16, 64)
		duanSum += ans2
	}

	//fmt.Println(duanSum)
	//取duanSum 的第一个数字
	duanSumString := strconv.FormatInt(duanSum, 10)
	firstDigit := duanSumString[0]
	//将firstDigit 从byte 转为int 类型
	firstDigitInt := int(firstDigit - '0')

	// 然后根据第一个数字，从第firstDigit 个数截取三位 如果第一位是0  则从第10位开始截取duansum三位 // 不可能是第10为 因为一个数不可能有前导0
	var result string = duanSumString[firstDigitInt-1 : firstDigitInt+2]

	// 然后取result 每一位的数字
	var resultInt int = 0
	for i := 0; i < 3; i++ {
		resultInt += int(result[i] - '0')
	}
	//log = append(log, pair{ans, resultInt})
	return result, resultInt
}

func Randomapi2(ordernum, goldsum int, timesum int) (string, int) {
	// type pair struct {
	// 	x string
	// 	y int
	// } // 一个存对应字符串 一个存对应和
	//var log []pair
	// timesum := randomtime(ordernum)
	//fmt.Println(timesum)
	ans := jiami(ordernum, goldsum, timesum)
	//	fmt.Println(ans)

	// 然后分成五段，从下标为2 取 每段取12位，将对应十位变成十进制
	// 每段的和
	var duanSum int64 = 0
	for i := 0; i < 5; i++ {
		ans1 := ans[2+12*i : 14+12*i]
		// 将该字符串转为十进制
		ans2, _ := strconv.ParseInt(ans1, 16, 64)
		duanSum += ans2

	}

	//fmt.Println(duanSum)
	//取duanSum 的第一个数字
	duanSumString := strconv.FormatInt(duanSum, 10)
	firstDigit := duanSumString[0]
	//将firstDigit 从byte 转为int 类型
	firstDigitInt := int(firstDigit - '0')

	// 然后根据第一个数字，从第firstDigit 个数截取三位 如果第一位是0  则从第10位开始截取duansum三位 // 不可能是第10为 因为一个数不可能有前导0
	var result string = duanSumString[firstDigitInt-1 : firstDigitInt+2]

	// 然后取result 每一位的数字
	var resultInt int = 0
	for i := 0; i < 3; i++ {
		resultInt += int(result[i] - '0')
	}
	//log = append(log, pair{ans, resultInt})
	return result, resultInt
}

func jiami(ordNumInt int, goldsum int, timesum int) string {
	// 用sha256 算法 将订单数和金币数和时间值之和加密
	sha256 := sha256.New()
	total := ordNumInt + goldsum + timesum
	totalString := strconv.Itoa(total)
	//fmt.Println(total)
	sha256.Write([]byte(totalString))
	sha256sum := sha256.Sum(nil)
	// 将加密后的值转换为字符串,得到一个64位的字符串
	ans := hex.EncodeToString(sha256sum)
	return ans
}

func randomdata() (int, int, int) {
	//随机生成一个订单数 4-50包括50

	goldsum := 0
	timesum := 0
	//订单数是4-50
	ordNum, err := rand.Int(rand.Reader, big.NewInt(47))
	if err != nil {
		panic(err)
	}
	// 将ordNum 转换为int类型
	ordNumInt := int(ordNum.Int64()) + 4
	// 根据订单数随机生成对应数量的时间值 值的类型是 0-59999 ,用一个切片接受
	for i := 0; i < ordNumInt; i++ {
		// 随机生成金币数 存入tmp 范围是 100-200000  是100的整数倍 rand 随机生成0-1999  +1就是 1-2000 *100 就是对应 100 -200000
		tmp, err := rand.Int(rand.Reader, big.NewInt(2000))
		if err != nil {
			panic(err)
		}
		// 将tmp 转换为int 类型
		tmpInt := (int(tmp.Int64()) + 1) * 100
		// 将金币数存入goldsum
		goldsum += tmpInt
		// 随机生成时间值
		tmp2, _ := rand.Int(rand.Reader, big.NewInt(60000))
		tmp2Int := int(tmp2.Int64())
		// 将时间值求和
		timesum += tmp2Int
	}
	return ordNumInt, goldsum, timesum
}

func Randomtime(ordernum int) int {
	var timesum int = 0
	for i := 0; i < ordernum; i++ {
		// 随机生成时间值
		tmp2, _ := rand.Int(rand.Reader, big.NewInt(60000))
		tmp2Int := int(tmp2.Int64())
		// 将时间值求和
		timesum += tmp2Int

	}
	return timesum
}

// 给定一个数组 里面包括秒数，为数组中的每个秒数随机生成0-999的毫秒值，然后秒数和毫秒拼接在一起，得到一个新的数组，
// 比如说秒数是32 生成的毫秒数为333 那么这个数就是32333 然后将数组所有时间值求和，得到一个时间值
func Randomtime2(a []int) int {
	n := len(a)
	var timesum int = 0
	for i := 0; i < n; i++ {
		// 随机生成时间值
		tmp2, _ := rand.Int(rand.Reader, big.NewInt(1000))
		tmp2Int := int(tmp2.Int64())
		// 将时间值求和
		timesum += a[i]*1000 + tmp2Int
	}
	return timesum

}
