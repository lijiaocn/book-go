//create: 2017/10/26 22:04:34 change: 2017/12/18 14:20:14 lijiaocn@foxmail.com
package main

func main() {
	//not allow
	//slice1 := make([]int)
	//println("slice1， len is ", len(slice1), "capacity is ", cap(slice1))

	slice2 := make([]int, 10)
	println("slice2， len is ", len(slice2), "capacity is ", cap(slice2))

	slice3 := make([]int, 10, 20)
	println("slice3， len is ", len(slice3), "capacity is ", cap(slice3))
}
