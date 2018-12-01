//create: 2017/10/26 22:04:34 change: 2017/12/18 11:42:13 lijiaocn@foxmail.com
package main

type Student struct {
	name string
	age  int
}

func Display(s struct {
	name string
	age  int
}) {
	println(s.name)
	println(s.age)
}

func main() {
	alice := Student{
		name: "alice",
		age:  16,
	}
	Display(alice)
}
