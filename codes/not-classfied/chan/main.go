//create: 2018/01/15 18:38:54 change: 2018/01/18 15:57:33 lijiaocn@foxmail.com
package main

import (
	"time"
)

func main() {
	var c chan int
	c = make(chan int, 2) //must create by make

	go func(c chan int) {
		c <- 1
		c <- 2
		c <- 3
		c <- 4
		c <- 5
		for {
			time.Sleep(10)
		}
	}(c)

	for {
		println("channel's length: ", len(c))
		select {
		case i := <-c:
			println("recv: ", i)
		}
	}
}
