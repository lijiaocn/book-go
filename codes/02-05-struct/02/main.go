//create: 2017/12/28 18:44:09 change: 2017/12/28 18:47:09 lijiaocn@foxmail.com

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
	A1 string
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
		A1: "i'm B's A1",
	}
	fmt.Println(b.A)
	fmt.Println(b.A.A1)
	fmt.Println(b.A1)
}
