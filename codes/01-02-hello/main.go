//create: 2017/10/26 22:04:34 change: 2018/01/13 18:45:41 lijiaocn@foxmail.com
package main

import (
	"fmt"
)

type People struct {
	Name string
	Age  int
}

func (self *People) SayHello() {
	fmt.Printf("I'm %s, %d years old, i'm glad to meet you!\n", self.Name, self.Age)
}

func main() {
	xiaoming := People{
		Name: "xiaoming",
	}
	xiaoming.Age = 10
	xiaoming.SayHello()
}
