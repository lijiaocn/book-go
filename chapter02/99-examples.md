---
layout: default
title: 99-examples
author: lijiaocn
createdate: 2017/12/16 23:16:46
changedate: 2017/12/18 14:31:08
categories:
tags:
keywords:
description: 

---


<!-- toc -->


### raw string

codes: 

	package main
	
	func main() {
	    raw := `Hello,
	World
	    !`
	
	    println(raw)
	}

output: 

	Hello,
	World
		!

### same string

codes: 

	package main
	
	func main() {
	
	    str1 := "日本語"                                  // UTF-8 input text
	    str2 := `日本語`                                  // UTF-8 input text as a raw literal
	    str3 := "\u65e5\u672c\u8a9e"                   // the explicit Unicode code points
	    str4 := "\U000065e5\U0000672c\U00008a9e"       // the explicit Unicode code points
	    str5 := "\xe6\x97\xa5\xe6\x9c\xac\xe8\xaa\x9e" // the explicit UTF-8 bytes
	
	    println(str1)
	    println(str2)
	    println(str3)
	    println(str4)
	    println(str5)
	}

output:

	日本語
	日本語
	日本語
	日本語
	日本語

### constant max int

codes:

	package main
	
	func main() {
	    //2^256=115792089237316195423570985008687907853269984665640564039457584007913129639936
	    i := 115792089237316195423570985008687907853269984665640564039457584007913129639936
	}

output:

	./main.go:6: constant 115792089237316195423570985008687907853269984665640564039457584007913129639936 overflows int

### unamed type

codes:

	package main
	
	type Student struct {
	    name string
	    age  int
	}
	
	func Display(s struct {
	    name string
	    age  int
	}) {
	    println(s.name)
	    println(s.age)
	}
	
	func main() {
	    alice := Student{
	        name: "alice",
	        age:  16,
	    }
	    Display(alice)
	}

output:

	alice
	16

### Method Sets

codes:

	package main
	
	type Str string
	
	func (s Str) Show() {
	    println(s)
	}
	
	func main() {
	    str := Str("Hello World!")
	    pstr := &str
	    pstr.Show()
	}

output:

	Hello World!

### String Index

codes:

	package main
	
	import (
	    "fmt"
	)
	
	func main() {
	    str := "Hello World!"
	    fmt.Printf("%c\n", str[6])
	
	    //not allow
	    //ptr := &str[6]
	
	    //not allow
	    //str[6] = 'w'
	}

output:

	W

### Create Slice from Array

codes:

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

output:

	array's length:  30
	slice1's length:  5
	slice1's capacity:  20
	10
	11
	12
	13
	14

### Create Slice by make

codes:

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

output:

	slice2， len is  10 capacity is  10
	slice3， len is  10 capacity is  20
