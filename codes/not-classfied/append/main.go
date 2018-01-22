//create: 2018/01/09 14:01:09 change: 2018/01/09 14:18:47 lijiaocn@foxmail.com
package main

type Student struct {
	Name string
}

func main() {
	slice := []Student{}
	s := Student{
		Name: "s1",
	}
	slice = append(slice, s)
	s.Name = "s2"
	for _, s := range slice {
		print(s.Name)
	}
}
