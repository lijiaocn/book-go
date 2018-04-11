//create: 2018/01/22 10:39:15 change: 2018/01/22 10:40:28 lijiaocn@foxmail.com
package main

import (
	"os"
)

func main() {
	if name, err := os.Hostname(); err != nil {
		println(err.Error())
	} else {
		println(name)
	}
}
