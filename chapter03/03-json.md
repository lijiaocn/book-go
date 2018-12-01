# json字符串的反序列化与变量序列化成json字符串

## encoding/json

`encoding/json`是Go的标准包，用来将变量序列化成json字符串，以及将json字符串反序列化为变量。

### 代码

```go
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
```

### 运行

```bash
$ go run main.go
marshal result: {"Addr":"10.1.1.1","Port":80,"Inner":{"Prefix":"prefix"}}
unmarshal result: {10.1.1.1 80 {prefix}}
```

注意，如果Outer中的Innner没有加json tag：

```go
Outer struct {
    Addr string `json:"Addr"`
    Port int    `json:"Port"`
    Inner
}
```

或者json tag中没有名称：

```go
Outer struct {
    Addr string `json:"Addr"`
    Port int    `json:"Port"`
    Inner       `json: ",inline"`
}
```

序列化的结果中没有`Inner`字段，直接展示`Inner`的成员：

```json
marshal result: {"Addr":"10.1.1.1","Port":80,"Prefix":"prefix"}
unmarshal result: {10.1.1.1 80 {prefix}}
```
