---
layout: default
title: 01-purpose
author: lijiaocn
createdate: 2017/12/16 17:27:21
changedate: 2017/12/16 23:36:24
categories:
tags:
keywords:
description: 

---

* auto-gen TOC:
{:toc}

# 设计目标

go语言的设计目标是`通用`的`系统`编程语言。

>Go is a general-purpose language designed with systems programming in mind. It is strongly typed and garbage-collected and has explicit support for concurrent programming. Programs are constructed from packages, whose properties allow efficient management of dependencies. The existing implementations use a traditional compile/link model to generate executable binaries. 

通用，意味着可以用go语言做很多事情，不受领域的限制。可以用它来写后台程序、应用程序，也可以用来做数据处理、分析决策。与通用相对的是`专用`，例如matlab也是一门编程语言，但它只能用来做数据处理。相比之下go语言可以做的事情丰富多了，但go可以做不等于它做的更好，譬如要做数据统计处理，还是用matlab、R等语言合适。

系统，是指go语言是面向操作系统的，使用go开发的程序直接在操作系统上运行，可以直接调用操作系统的接口。C、C++都是系统语言，Java不是。用Java开发的程序是运行在JVM上的，运行在操作系统上的JVM代为调用操作系统的接口。同理，HTML、Javascript、Excel中的宏语言等也不是系统编程语言。( [System programming language][1])

## 参考

1. [System programming language][1]

[1]: https://en.wikipedia.org/wiki/System_programming_language "System programming language" 
