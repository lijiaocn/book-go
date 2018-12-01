//create: 2017/10/26 22:04:34 change: 2017/12/19 21:07:52 lijiaocn@foxmail.com
package main

type A struct {
	A1 string
}

func (a A) method() {
	a.A1 = "method set a1"
}

func (a *A) pointer_method() {
	a.A1 = "pointer method set a1"
}

type B struct {
	*A
	B1 string
	B2 string
}

func main() {
	b := B{
		A: &A{
			A1: "a1",
		},
		B1: "b1",
		B2: "b2",
	}

	b.method()
	println(b.A1)

	b.pointer_method()
	println(b.A1)

	pb := &b

	pb.method()
	println(b.A1)

	pb.pointer_method()
	println(b.A1)
}
