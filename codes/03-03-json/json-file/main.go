//create: 2018/01/22 10:14:35 change: 2018/01/22 10:18:26 lijiaocn@foxmail.com
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Student struct {
	Name string
	Age  int
}

func main() {
	file, err := os.Open("./test.json")
	if err != nil {
		log.Fatal(err)
	}

	var s Student
	if err := json.NewDecoder(file).Decode(&s); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(s)
	}

}
