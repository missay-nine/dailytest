package randomapi

import "testing"

//func TestRandapi(t *testing.T) {
//	//测试 Randapi 函数是否能正确返回一个三位数的字符串和一个 0-27 的整数
//	for i := 0; i < 10; i++ {
//		str, num := Randapi()
//		if len(str) != 3 {
//			t.Errorf("Randapi() returned wrong string length: %s", str)
//		}
//		if num < 0 || num > 27 {
//			t.Errorf("Randapi() returned wrong number: %d", num)
//		}
//		t.Logf("Randapi() returned: %s, %d", str, num)
//	}
//
//}

// Randapi的基准测试

func BenchmarkRandapi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Randapi()
	}
}
