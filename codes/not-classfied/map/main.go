//create: 2018/01/24 14:14:33 change: 2018/01/25 14:21:06 lijiaocn@foxmail.com
package main

func main() {
	m := make(map[string]string, 0)
	m["1"] = "1"
	m["2"] = "2"
	println(len(*(&m)))
	delete(m, "2")
	println(len(m))
	c, _ := m["33"]
	if c == "" {
		println("is empty")
	}
}
