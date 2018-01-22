//create: 2018/01/09 14:18:55 change: 2018/01/09 16:27:34 lijiaocn@foxmail.com
package main

import (
	"strconv"
)

func main() {
	a := 10
	//string(a) not work
	b := "hello " + string(a)
	print(b)

	c := "hello " + strconv.Itoa(a)
	print(c)

}
