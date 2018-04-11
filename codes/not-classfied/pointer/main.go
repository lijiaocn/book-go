//create: 2018/01/23 10:42:04 change: 2018/01/23 13:49:12 lijiaocn@foxmail.com
package main

import (
	"fmt"
)

type Student struct {
	Name string
	Age  int
}

func main() {
	var s Student
	s.Name = "lijiaocn"
	s.Age = 10

	p := &s
	c := *p

	s.Age = 11

	fmt.Printf("%s\n", p)
	fmt.Printf("%s\n", c)
}
