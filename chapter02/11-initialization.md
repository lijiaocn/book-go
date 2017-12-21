---
layout: default
title: 11-initialization
author: lijiaocn
createdate: 2017/12/22 00:28:06
changedate: 2017/12/22 00:32:39
categories:
tags:
keywords:
description: 

---

* auto-gen TOC:
{:toc}

# go程序的初始化

## 变量初始化为对应类型的零值

例如：

	type T struct { i int; f float64; next *T }
	t := new(T)

初始化值为：

	t.i == 0
	t.f == 0.0
	t.next == nil

## 包的初始化

下面的代码中，变量的初始化顺序为：d、b、c、a。

	var (
		a = c + b
		b = f()
		c = f()
		d = 3
	)

	func f() int {
		d++
		return d
	}

init函数：

	init() { … }

## 开始执行

	func main() { … }
