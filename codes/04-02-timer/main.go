//create: 2017/12/28 20:46:39 change: 2017/12/28 21:31:32 lijiaocn@foxmail.com
package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		select {
		case c := <-time.After(10 * time.Second):
			fmt.Println(c)
		}
	}
}
