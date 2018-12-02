---
layout: default
title: 12-errors
author: lijiaocn
createdate: 2017/12/22 00:33:14
changedate: 2017/12/22 00:38:17
categories:
tags:
keywords:
description: 

---


<!-- toc -->

# go的错误处理

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



