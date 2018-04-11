//create: 2018/01/23 13:49:31 change: 2018/01/23 13:52:17 lijiaocn@foxmail.com
package main

import (
	"strings"
)

func main() {
	strs := strings.Split("a,b,c,d", ",")
	println(len(strs))
	for _, str := range strs {
		println(str)
	}

	strs = strings.Split("a", ",")
	println(len(strs))
	for _, str := range strs {
		println(str)
	}
}
