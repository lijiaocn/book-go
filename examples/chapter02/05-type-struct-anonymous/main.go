//create: 2017/10/26 22:04:34 change: 2017/12/19 20:25:52 lijiaocn@foxmail.com
package main

import (
	"fmt"
)

type A struct {
	A1 string
	A2 string
}

type B struct {
	A
	B1 string
	B2 string
}

func main() {
	b := B{
		A: A{
			A1: "a1",
			A2: "a2",
		},
		B1: "b1",
		B2: "b2",
	}
	fmt.Println(b.A)
	fmt.Println(b.A.A1)
	fmt.Println(b.A1)
}
