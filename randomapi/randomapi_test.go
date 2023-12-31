package randomapi

import "testing"

// func TestRandapi(t *testing.T) {
// 	//测试 Randapi 函数是否能正确返回一个三位数的字符串和一个 0-27 的整数
// 	for i := 0; i < 10; i++ {
// 		str, num := Randapi()
// 		if len(str) != 3 {
// 			t.Errorf("Randapi() returned wrong string length: %s", str)
// 		}
// 		if num < 0 || num > 27 {
// 			t.Errorf("Randapi() returned wrong number: %d", num)
// 		}
// 		t.Logf("Randapi() returned: %s, %d", str, num)
// 	}

// }

func TestRandomdata(t *testing.T) {
	for i := 0; i < 1000; i++ {
		ordNum, goldsum, _ := randomdata()

		// 检查订单数是否在4到50之间
		if ordNum < 4 || ordNum > 50 {
			t.Errorf("订单数不在范围内，期望在4-50之间，得到: %d", ordNum)
		}

		// 检查金币数是否是100的整数倍
		if goldsum%100 != 0 {
			t.Errorf("金币数不是100的整数倍，得到: %d", goldsum)
		}
	}
}

// Randapi的基准测试

// func BenchmarkRandapi(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		Randapi()
// 	}
// }
