---
layout: post
title: 04-变量语法
author: lijiaocn
createdate: 2017/12/18 10:43:19
changedate: 2017/12/25 12:44:07
categories: chapter02
tags:
keywords:
description: 

---

* auto-gen TOC:
{:toc}

## go的变量 

变量记录的是一个位置，这个位置中存放的值是可以变化的。

变量是有类型的，变量的类型规定了如何解读指定位置中存放的值。

## 静态类型和动态类型

变量的类型分为静态类型和动态类型。

声明变量时，指定的类型是变量的静态类型。

如果变量静态类型是接口(interface type)，它还会有一个动态类型，动态类型就是被赋予的值的类型。

	var x interface{}  // x is nil and has static type interface{}
	x = 42             // x has value 42 and dynamic type int
	x = v              // x has value (*T)(nil) and dynamic type *T

如果声明变量时没有设置变量的值，它的值就是对应类型的零值（zero value)。

## 参考

1. [文献1][1]
2. [文献2][2]

[1]: 1.com  "文献1" 
[2]: 2.com  "文献1" 
