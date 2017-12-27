//create: 2017/12/25 14:24:03 change: 2017/12/25 15:15:33 lijiaocn@foxmail.com
package main

import (
	"flag"
	"github.com/golang/glog"
)

func main() {
	flag.Parse()
	glog.Infof("%s\n", "This is a info")
	glog.Warningf("%s\n", "This is a warning")
	glog.Errorf("%s\n", "This is a error")
	glog.V(1).Infof("%s\n", "This is a v1 log")
	glog.V(2).Infof("%s\n", "This is a v2 log")
	//this call will make the progress exit.
	//glog.Fatalf("%s\n", "This is a fatal")
	//this call will make the progress exit.
	//glog.Exitf("%s\n", "This is a exit")
}
