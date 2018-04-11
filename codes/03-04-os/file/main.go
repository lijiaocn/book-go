//create: 2018/01/22 11:10:24 change: 2018/01/22 14:21:16 lijiaocn@foxmail.com
package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	for {
		file, err := os.Open("./file.txt")
		if err != nil {
			log.Fatal(err.Error())
		}

		stat, err := file.Stat()
		if err != nil {
			log.Fatal(err.Error())
		}

		if stat.Mode()&os.ModeSymlink != 0 {
			file.Close()

			realfile, err := os.Readlink("./file.txt")
			if err != nil {
				log.Fatal(err.Error())
			}

			file, err = os.Open(realfile)
			if err != nil {
				log.Fatal(err.Error())
			}

			stat, err = file.Stat()
			if err != nil {
				log.Fatal(err.Error())
			}
		}

		fmt.Println(stat.ModTime().Format(time.RFC3339))

		if err := file.Close(); err != nil {
			log.Fatal(err.Error())
		}

		time.Sleep(1 * time.Second)
	}
}
