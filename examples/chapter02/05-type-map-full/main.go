//create: 2017/10/26 22:04:34 change: 2017/12/19 22:02:21 lijiaocn@foxmail.com
package main

import (
	"fmt"
)

func main() {
	m := make(map[int]int, 10)
	for i := 0; i < 10; i++ {
		m[i] = i
	}
	println(len(m))
	fmt.Println(m)
	m[11] = 11
	println(len(m))
	fmt.Println(m)
}
