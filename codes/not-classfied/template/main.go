//create: 2018/01/09 16:33:01 change: 2018/01/09 16:54:52 lijiaocn@foxmail.com
package main

import (
	"os"
	"text/template"
)

type Student struct {
	Name string
	Age  int
}

func main() {
	s := Student{
		Name: "lijiaocn",
		Age:  18,
	}

	f, err := os.OpenFile("./example.result", os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		println(err.Error())
		return
	}

	t := template.Must(template.ParseFiles("./example.tpl"))
	if err := t.Execute(f, s); err != nil {
		println(err.Error())
		return
	}

	if err := f.Close(); err != nil {
		println(err.Error())
		return
	}
}
