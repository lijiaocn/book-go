//create: 2017/10/26 22:04:34 change: 2017/12/19 21:30:07 lijiaocn@foxmail.com
package main

func main() {
	i := 8
	pi := &i
	ppi := &pi

	println(*ppi, pi)
	println(*pi, i)
}
