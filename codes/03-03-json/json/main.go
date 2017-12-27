//create: 2017/12/25 15:36:05 change: 2017/12/25 16:38:03 lijiaocn@foxmail.com
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
		Addr string `json:"Addr"`
		Port int    `json:"Port"`
		//Inner `json:",inline"`
		Inner
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
