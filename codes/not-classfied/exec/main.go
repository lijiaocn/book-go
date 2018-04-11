//create: 2018/01/17 15:12:26 change: 2018/01/29 19:36:08 lijiaocn@foxmail.com
package main

import (
	"fmt"
	"os/exec"
)

func main() {
	/*
		if bytes, err := exec.Command("./run.sh", "hello").Output(); err != nil {
			println(err.Error())
		} else {
			fmt.Printf("%s\n", bytes)
		}
	*/
	if err := exec.Command("./run.sh", "hello").Run(); err != nil {
		println(err.Error())
	} else {
		fmt.Printf("success\n")
	}
}
