package main

import (
	"dubo/randomapi"
	"fmt"
)

func main() {

	a, b := randomapi.Randomapi2(10, 228800)
	fmt.Println(a, b)
}
