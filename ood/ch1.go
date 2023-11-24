package main

import "fmt"

type P struct {
	name string
}

func (p P) setname(n string) string {
	p.name = n
	return n
}

func (p P) getname() string {
	return p.name
}

func main() {
	p := &P{name: "zhangsan"}
	fmt.Println(p.setname("lisi"))
	fmt.Println(p.getname())
}
