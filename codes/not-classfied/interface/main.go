//create: 2018/01/17 13:38:21 change: 2018/01/17 13:39:56 lijiaocn@foxmail.com

package main

type People interface {
	Hi()
}

func main() {
	var p People
	println(p)
}
