//create: 2017/10/26 22:04:34 change: 2017/12/18 14:12:23 lijiaocn@foxmail.com
package main

func main() {
	var array1 [30]int
	for i := 0; i < len(array1); i++ {
		array1[i] = i
	}

	slice1 := array1[10:15]

	println("array's length: ", len(array1))
	println("slice1's length: ", len(slice1))
	println("slice1's capacity: ", cap(slice1))

	for i := 0; i < len(slice1); i++ {
		println(slice1[i])
	}
}
