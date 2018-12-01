//create: 2017/10/26 22:04:34 change: 2017/12/18 12:38:05 lijiaocn@foxmail.com
package main

type Str string

func (s Str) Show() {
	println(s)
}

func main() {
	str := Str("Hello World!")
	pstr := &str
	pstr.Show()
}
