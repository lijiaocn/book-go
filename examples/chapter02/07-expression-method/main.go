//create: 2017/10/26 22:04:34 change: 2017/12/20 22:13:37 lijiaocn@foxmail.com
package main

type Str string

func (s Str) show_str() {
	println(s)
}

func (s *Str) show_strp() {
	println(*s)
}

func main() {
	var s Str = "hello"

	Str.show_str(s)
	//not allow
	//Str.show_strp(s)

	(*Str).show_strp(&s)
	//allow
	(*Str).show_str(&s)
}
