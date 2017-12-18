---
layout: default
title: 05-types
author: lijiaocn
createdate: 2017/12/18 10:55:10
changedate: 2017/12/18 14:32:09
categories:
tags:
keywords:
description: 

---

* auto-gen TOC:
{:toc}

# go的类型 

类型是用来诠释如何解读指定位置中存放的数据，以及约定操作符的含义的。

## 类型命名

类型可以是命名的(named)，也可以是未命名的(unnamed)。

	Type      = TypeName | TypeLit | "(" Type ")" .
	TypeName  = identifier | QualifiedIdent .
	TypeLit   = ArrayType | StructType | PointerType | FunctionType | InterfaceType |
	            SliceType | MapType | ChannelType .

使用`type`指定了名字的类型是命名的，例如下面的类型的名字为Student

	type Student struct {
		Name string
		age int
	}

由其它类型组合成的新类型，可以不被命名，例如下面的类型是没有名字的：

	[] string
	[] int

无类型的名字，通用用于定义其它类型时:

	type Array []int

或者在函数定义中使用：

	func Display(s struct {
	    name string
	    age  int
	}) {
	    println(s.name)
	    println(s.age)
	}

## 实际类型(underlying type)

类型是可以用来定义其它类型的，例如：

	type T1 string 
	type T2 T1

这里定义了一个类型T1，然后又用T1定义了类型T2。

T1的实际类型(underlying type)是string，T2的实际类型不是T1，而是T1的实际类型string。

实际类型必须是go的内置类型或者类型的组合。

例如，string、T1、T2的实际类型是string。

	type T1 string
	type T2 T1

[]T1、T3、T4的实际类型是[]T1。

	type T3 []T1
	type T4 T3

## 内置类型(predeclared)

go语言内置了下面的类型：

	bool byte complex64 complex128 error float32 float64
	int int8 int16 int32 int64 rune string
	uint uint8 uint16 uint32 uint64 uintptr

## 类型的方法(Method sets)

类型可以有自己的方法(Method)，也就是其它语言中的函数。

一个非接口类型的方法集就所有接受者(receiver)为改类型的方法，接口类型的方法集就是接口定义中包含的方法。

需要注意的是指针类型类型（例如 * T)，它的方法集是所有接受者为所指类型(T)和指针类型( * T)的方法集。

例如下面的代码中，方法的Show的Receiver是Str，但是类型为 * Str的pstr也可以调用。

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

方法集中的方法不能重名、且必须有名字。

## 布尔(Boolean types)

布尔类型是内置的类型`bool`，它的value只能是两个内置的常量：

	true
	false

## 数值(Numeric types)

数值类型都是内置的类型，一共有以下几种。

	uint8       the set of all unsigned  8-bit integers (0 to 255)
	uint16      the set of all unsigned 16-bit integers (0 to 65535)
	uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
	uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)
	
	int8        the set of all signed  8-bit integers (-128 to 127)
	int16       the set of all signed 16-bit integers (-32768 to 32767)
	int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
	int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)
	
	float32     the set of all IEEE-754 32-bit floating-point numbers
	float64     the set of all IEEE-754 64-bit floating-point numbers
	
	complex64   the set of all complex numbers with float32 real and imaginary parts
	complex128  the set of all complex numbers with float64 real and imaginary parts
	
	byte        alias for uint8
	rune        alias for int32

另外还有三个数值类型，它们占用的空间取决于实现：

	uint     either 32 or 64 bits
	int      same size as uint
	uintptr  an unsigned integer large enough to store the uninterpreted bits of a pointer value

## 字符串(String types)

字符串是内置的类型`string`，字符串的值是连续的字节，这些字节是不可更改的。

可以通过内置函数`len`获取字符串的长度，可以用通过[i]读取字符串的第i个(从0开始)字节。

字符串的字节只能读取，不能更改，也不能取址。

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

## Array types

数组是多个相同类型的值，在go中，数组必须有长度，长度是数组类型的一部分。

	ArrayType   = "[" ArrayLength "]" ElementType .
	ArrayLength = Expression .
	ElementType = Type .

数组是单维的，可以累进成多维数组：

	[32]byte
	[2*N] struct { x, y int32 }
	[1000]*float64
	[3][5]int
	[2][2][2]float64  // same as [2]([2]([2]float64))

要注意长度是数组类型的一部分，长度不同的数组是不同的类型，例如：

	package main
	
	func main() {
	    var array1 [32]byte
	    var array2 [24]byte
	
	    array1[0] = 'a'
	    array2[0] = 'b'
	
	    //not allow
	    //array2 = array1
	}

数组成员可以用从0开始的坐标索引，长度可以用内置的函数`len`获取。

## Slice types

分片(Slice)是用来索引数组(Array)中的一段连续的成员的。

	SliceType = "[" "]" ElementType .

分片初始化后就绑定到了一个数组，多个分片可以绑定到同一个数组。

与数组不同的是，分片有长度(length)和容量(capacity)两个属性。

长度是分片所索引的数组成员的数量，可以用内置的函数`len`获取。

容量是分片能够索引的数组成员的最大数量，等于数组的长度减去分片索引的第一个数组成员在数组中位置。

例如在下面的代码中，分片slice1的长度是5，容量是20(=30-10)

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

分片可以通过两种方式创建，第一种方式就是上面的代码中使用的方式：

	    slice1 := array1[10:15]

这样创建的slice1索引的是数组的从0开始编号的第10个、第11个、第12个、第13个、第14个个成员，总计5个。

	10
	11
	12
	13
	14

>注意[10:15]是一个前闭后开的集合，即包括10，不包括15。

第二种方式是使用内置的`make`函数创建。

	make([]T, length, capacity)

使用make创建的时候，至少需要指定分片的长度，make会为分片创建一个隐藏的数组。

如果指定了capacity，数组的长度就是capacity，如果没有指定，数组的长度等于分片的长度。

例如下面的代码中slice2的长度和容量都是10，slice3的长度是10，容量是20。

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

通过make创建分片，相当与新建一个数组，然后取它的[0:length]。

	make([]int, 50, 100)

等同于：

	new([100]int)[0:50]

## Struct types

结构体(Struct)是比较复杂的类型。

## Pointer types
## Function types
## Interface types
## Map types
## Channel types

## 参考

1. [文献1][1]
2. [文献2][2]

[1]: 1.com  "文献1" 
[2]: 2.com  "文献1" 
