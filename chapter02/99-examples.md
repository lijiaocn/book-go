---
layout: default
title: 99-examples
author: lijiaocn
createdate: 2017/12/16 23:16:46
changedate: 2017/12/17 01:10:52
categories:
tags:
keywords:
description: 

---

* auto-gen TOC:
{:toc}


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
