//create: 2017/10/26 22:04:34 change: 2017/12/18 13:41:24 lijiaocn@foxmail.com
package main

import (
	"fmt"
)

func main() {
	str := "Hello World!"
	fmt.Printf("%c\n", str[6])

	//not allow
	//ptr := &str[6]

	//not allow
	//str[6] = 'w'
}
