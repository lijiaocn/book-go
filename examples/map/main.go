//create: 2017/10/26 22:04:34 change: 2017/12/21 16:18:13 lijiaocn@foxmail.com
package main

func main() {
	m := make(map[string]int)
	m["a"] = 1
	m["b"] = 2
	m["c"] = 3
	m["a"] = 4
	println(m["a"])

	for _, v := range m {
		println(v)
	}
}
