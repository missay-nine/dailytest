package Map

func main() {
	var mp map[int]int
	mp = make(map[int]int)
	for i := 0; i < 100; i++ {
		mp[i] = i
	}

}
