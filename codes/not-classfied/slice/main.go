//create: 2018/01/15 15:47:29 change: 2018/01/18 11:12:13 lijiaocn@foxmail.com
package main

import (
	"fmt"
)

func main() {
	var s []string
	s = append(s, "hello")

	var array = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := array[:]
	s2 := array[:]
	s1[0] = 100
	fmt.Println(s2)

	array = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 = array[:]
	s2 = array[:]
	s2 = append(s2, 1000, 1001, 1002) //s2 point to an new array
	s1[0] = 100
	fmt.Println(s2)
}
