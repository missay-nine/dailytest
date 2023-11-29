package main

import "fmt"

func SliceRise(s []int) {
	s = append(s, 0)
	for i := range s {
		s[i]++
	}

	fmt.Println(s, "函数")
}
func main() {
	s1 := []int{1, 2}
	fmt.Println(len(s1), cap(s1))
	s2 := s1
	s2 = append(s2, 3)
	fmt.Println(len(s2), cap(s2))
	SliceRise(s1)
	SliceRise(s2)
	fmt.Println(s1, s2)
}
