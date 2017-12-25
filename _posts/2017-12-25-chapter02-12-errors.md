---
layout: post
title: 12-错误处理
author: lijiaocn
createdate: 2017/12/22 00:33:14
changedate: 2017/12/25 12:44:57
categories: chapter02
tags:
keywords:
description: 

---

* auto-gen TOC:
{:toc}

## go的错误处理

## 返回的错误值

	type error interface {
		Error() string
	}

## panic的传入参数类型

	package runtime
	
	type Error interface {
		error
		// and perhaps other methods
	}



