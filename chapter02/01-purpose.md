---
layout: default
title: 01-purpose
author: lijiaocn
createdate: 2017/12/16 17:27:21
changedate: 2017/12/16 17:50:32
categories:
tags:
keywords:
description: 

---

* auto-gen TOC:
{:toc}

# 设计目标与编码

## 设计目标

>Go is a general-purpose language designed with systems programming in mind. It is strongly typed and garbage-collected and has explicit support for concurrent programming. Programs are constructed from packages, whose properties allow efficient management of dependencies. The existing implementations use a traditional compile/link model to generate executable binaries. 

go语言的设计目标是`通用`的`系统`编程语言。

通用，意味着可以用go语言做很多事情，不受领域的限制。与通用相对的是`专用`，例如matlab也是一门编程语言，但它只能用来做数据处理。go语言可以做的事情就丰富多了，即可以用它来写后台程序、应用程序，也可以用来做数据处理、分析决策。go可以做的事情多，不等于它都可以做的更好，譬如要做数据统计处理，还是用matlab、R等语言合适。

系统，是指go语言是面向操作系统的，使用go开发的程序是直接在操作系统上运行，可以调用操作系统的接口。C、C++都是系统语言。Java不是。用Java开发的程序是运行在JVM上的，JVM运行在操作系统上，代为调用操作系统的接口。同理，HTML、Javascript、Excel中的宏语言等也不是系统编程语言。( [System programming language][1])

## 编码方案

go的源代码使用UTF-8编码。

## 参考

1. [System programming language][1]

[1]: https://en.wikipedia.org/wiki/System_programming_language "System programming language" 
