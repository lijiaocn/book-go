---
layout: post
title: 03-常量语法
author: lijiaocn
createdate: 2017/12/16 23:42:57
changedate: 2017/12/25 12:44:01
categories: chapter02
tags:
keywords:
description: 

---

* auto-gen TOC:
{:toc}

## go的常量

常量意如起名，是不会变化的量，量就是值。常量是程序运行时，系统中某个位置里的数值。

## 常量分类

常量分为以下几类:

	布尔，boolean
	符号，rune
	整数，integer
	浮点数，floating-point
	复数，complex
	字符串，string

其中符号(rune)、整数(intrger)、浮点数(floating-point)和复数(complex)型常量，又被称为数值常量(numeric constants)。

## 常量表示

常量的值有下面几种表达方式：

	符号，rune
	整数，integer
	浮点数，floating-point
	虚数，imaginary
	字符串，string
	指向常量的标记符，identifier denoting a constant
	常量表达式，constant expression
	结果为常量的类型转化， a conversion with a result that is a constant
	内置函数的返回结果
	内置的常量true和false
	内置的常量标识符iota

数值型常量的值与所显示值一致，不会出现`溢出`，IEEE-754中的“-0”(negative zero)、“无穷大”(infinity)、“非数”(not-a-number)没有对应的常量。

## 常量的类型

常量可以是显示声明了类型的(typed)，也可以是未声明类型的(untyped)。

未声明类型的常量会依据它的值获得一个默认的类型：

	value_type       default_type
	------------------------------
	boolean          bool
	rune             rune
	integer          int
	floating-point   float64
	complex          complex128
	string           string

例如：

	i := 3
	j := 3.0

"3"是一个untyped的常量，因为3是一个整数，它的默认类型就是int。

"3.0"是一个浮点数，它的默认类型是float64。

## 数值型常量的范围

可以在代码中写出任意大小的数值，但是代码中写出数未必能被编译器支持。

编译器支持的最大数值是有上限的，在代码中可以写入的数字确实无限的。

go的编译器做了以下承诺：

	至少支持256个比特长度的整数
	至少支持256个比特长度的小数
	如果整数数值超过支持的范围，编译器报错
	如果浮点数和复数溢出，编译器报错
	如果浮点数和复数超出了支持的精度，使用最接近的数值

例如下面的代码编译时会报错：

	package main
	
	func main() {
		i := 115792089237316195423570985008687907853269984665640564039457584007913129639936
	}

Error:

	./main.go:6: constant 115792089237316195423570985008687907853269984665640564039457584007913129639936 overflows int

>2^256=115792089237316195423570985008687907853269984665640564039457584007913129639936

## 参考

1. [go Constants][1]

[1]: http://127.0.0.1:6060/ref/spec#Constants  "go Constants" 
