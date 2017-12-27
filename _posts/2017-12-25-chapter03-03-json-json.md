---
layout: post
title: 03-json字符串的序列化与反序列化
author: lijiaocn
createdate: 2017/12/25 15:52:28
changedate: 2017/12/25 17:04:04
categories: chapter03
tags:
keywords:
description: 

---

* auto-gen TOC:
{:toc}

## 说明

`encoding/json`是Go的标准包，用来将变量序列化成json字符串，以及将json字符串反序列化为变量。

## 代码

{% highlight go %}
//create: 2017/12/25 15:36:05 change: 2017/12/25 16:35:35 lijiaocn@foxmail.com
package main

import (
    "encoding/json"
    "fmt"
)

type (
    Inner struct {
        Prefix string `json:"Prefix"`
    }

    Outer struct {
        Addr  string `json:"Addr"`
        Port  int    `json:"Port"`
        Inner `json:",inline"`
        //Inner
        //Inner `json:"Inner"`
    }
)

func main() {

    var err error
    var b []byte

    outer := Outer{
        Addr: "10.1.1.1",
        Port: 80,
        Inner: Inner{
            Prefix: "prefix",
        },
    }

    if b, err = json.Marshal(outer); err != nil {
        println(err)
        return
    }
    fmt.Printf("marshal result: %s\n", string(b))

    var new_outer Outer
    if err = json.Unmarshal(b, &new_outer); err != nil {
        println(err)
        return
    }
    fmt.Printf("unmarshal result: %v\n", new_outer)
}
{% endhighlight %}

## 运行

{% highlight shell %}
$ go run main.go
marshal result: {"Addr":"10.1.1.1","Port":80,"Inner":{"Prefix":"prefix"}}
unmarshal result: {10.1.1.1 80 {prefix}}
{% endhighlight %}

注意，如果Outer中的Innner没有加json tag：

{% highlight go %}
Outer struct {
    Addr string `json:"Addr"`
    Port int    `json:"Port"`
    Inner
}
{% endhighlight %}

或者json tag中没有名称：

{% highlight go %}
Outer struct {
    Addr string `json:"Addr"`
    Port int    `json:"Port"`
    Inner       `json: ",inline"`
}
{% endhighlight %}

序列化的结果中没有`Inner`字段，直接展示`Inner`的成员：

{% highlight shell %}
marshal result: {"Addr":"10.1.1.1","Port":80,"Prefix":"prefix"}
unmarshal result: {10.1.1.1 80 {prefix}}
{% endhighlight %}
