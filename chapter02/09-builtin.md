---
layout: default
title: 09-builtin-funcs
author: lijiaocn
createdate: 2017/12/22 00:11:56
changedate: 2017/12/22 00:24:00
categories:
tags:
keywords:
description: 

---


<!-- toc -->

# go的内置函数 

go内置了一些函数，这些函数可以直接使用。

## 关闭(Close)

	close

## 长度和容量(Length and capacit)

	Call      Argument type    Result
	
	len(s)    string type      string length in bytes
	          [n]T, *[n]T      array length (== n)
	          []T              slice length
	          map[K]T          map length (number of defined keys)
	          chan T           number of elements queued in channel buffer
	
	cap(s)    [n]T, *[n]T      array length (== n)
	          []T              slice capacity
	          chan T           channel buffer capacity

## 分配(Allocation)

	new(T)

## 创建(Making slices, maps and channels)

	Call             Type T     Result
	
	make(T, n)       slice      slice of type T with length n and capacity n
	make(T, n, m)    slice      slice of type T with length n and capacity m
	
	make(T)          map        map of type T
	make(T, n)       map        map of type T with initial space for n elements
	
	make(T)          channel    unbuffered channel of type T
	make(T, n)       channel    buffered channel of type T, buffer size n

例如：

	s := make([]int, 10, 100)       // slice with len(s) == 10, cap(s) == 100
	s := make([]int, 1e3)           // slice with len(s) == cap(s) == 1000
	s := make([]int, 1<<63)         // illegal: len(s) is not representable by a value of type int
	s := make([]int, 10, 0)         // illegal: len(s) > cap(s)
	c := make(chan int, 10)         // channel with a buffer size of 10
	m := make(map[string]int, 100)  // map with initial space for 100 elements

## 分片的追加(Appending to slices)

	append(s S, x ...T) S  // T is the element type of S

例如：

	s0 := []int{0, 0}
	s1 := append(s0, 2)                // append a single element     s1 == []int{0, 0, 2}
	s2 := append(s1, 3, 5, 7)          // append multiple elements    s2 == []int{0, 0, 2, 3, 5, 7}
	s3 := append(s2, s0...)            // append a slice              s3 == []int{0, 0, 2, 3, 5, 7, 0, 0}
	s4 := append(s3[3:6], s3[2:]...)   // append overlapping slice    s4 == []int{3, 5, 7, 2, 3, 5, 7, 0, 0}

	var t []interface{}
	t = append(t, 42, 3.1415, "foo")   //                             t == []interface{}{42, 3.1415, "foo"}

	var b []byte
	b = append(b, "bar"...)            // append string contents      b == []byte{'b', 'a', 'r' }

## 分片的复制(Copying slices)

	copy(dst, src []T) int
	copy(dst []byte, src string) int

例如：

	var a = [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	var s = make([]int, 6)
	var b = make([]byte, 5)
	n1 := copy(s, a[0:])            // n1 == 6, s == []int{0, 1, 2, 3, 4, 5}
	n2 := copy(s, s[2:])            // n2 == 4, s == []int{2, 3, 4, 5, 4, 5}
	n3 := copy(b, "Hello, World!")  // n3 == 5, b == []byte("Hello")

## 字典成员的删除(Deletion of map elements)

	delete(m, k)  // remove element m[k] from map m 

## 复数的操作(Manipulating complex numbers)

	complex(realPart, imaginaryPart floatT) complexT
	real(complexT) floatT
	imag(complexT) floatT

例如：

	var a = complex(2, -2)             // complex128
	const b = complex(1.0, -1.4)       // untyped complex constant 1 - 1.4i
	x := float32(math.Cos(math.Pi/2))  // float32
	var c64 = complex(5, -x)           // complex64
	var s uint = complex(1, 0)         // untyped complex constant 1 + 0i can be converted to uint
	_ = complex(1, 2<<s)               // illegal: 2 assumes floating-point type, cannot shift
	var rl = real(c64)                 // float32
	var im = imag(a)                   // float64
	const c = imag(b)                  // untyped constant -1.4
	_ = imag(3 << s)                   // illegal: 3 assumes complex type, cannot shift

## panic处理(Handling panics)

	panic(interface{})
	func recover() interface{}

例如：

	func protect(g func()) {
		defer func() {
			log.Println("done")  // Println executes normally even if there is a panic
			if x := recover(); x != nil {
				log.Printf("run time panic: %v", x)
			}
		}()
		log.Println("start")
		g()
	}

## 自举函数(Bootstrapping)

	Function   Behavior

	print      prints all arguments; formatting of arguments is implementation-specific
	println    like print but prints spaces between arguments and a newline at the end



