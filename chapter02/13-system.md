---
layout: default
title: 13-system
author: lijiaocn
createdate: 2017/12/22 00:38:28
changedate: 2017/12/22 00:53:30
categories:
tags:
keywords:
description: 

---


<!-- toc -->

# 操作系统相关

## unsafe 

	package unsafe

	type ArbitraryType int  // shorthand for an arbitrary Go type; it is not a real type
	type Pointer *ArbitraryType

	func Alignof(variable ArbitraryType) uintptr
	func Offsetof(selector ArbitraryType) uintptr
	func Sizeof(variable ArbitraryType) uintptr

## Size and alignment guarantees 

	type                                 size in bytes
	
	byte, uint8, int8                     1
	uint16, int16                         2
	uint32, int32, float32                4
	uint64, int64, float64, complex64     8

# 参考

. [文献1][1]
. [文献2][2]

1]: 1.com  "文献1" 
2]: 2.com  "文献1" 
	complex128                           16
