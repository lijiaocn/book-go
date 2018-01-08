//create: 2018/01/08 14:27:41 change: 2018/01/08 14:42:40 lijiaocn@foxmail.com
package main

type Student struct {
	Name   string
	Gender string
}

func (s *Student) IsBoy() bool {
	if s.Gender == "Male" {
		return true
	}
	return false
}
