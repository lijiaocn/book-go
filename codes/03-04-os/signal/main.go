//create: 2017/12/27 17:34:37 change: 2017/12/28 10:31:34 lijiaocn@foxmail.com
package main

import (
	"os"
	"os/signal"
	"syscall"
)

func main() {
	println(syscall.Getpid())

	receive := make(chan os.Signal, 10)

	signal.Ignore(syscall.SIGUSR1)
	signal.Notify(receive, syscall.SIGUSR2)

	for {
		select {
		case s := <-receive:
			switch s {
			case syscall.SIGUSR1:
				println("receive SIGUSR1")
			case syscall.SIGUSR2:
				println("receive SIGUSR2")
			}
		}
	}
}
