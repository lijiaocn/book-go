//create: 2017/12/25 14:14:15 change: 2017/12/25 14:16:27 lijiaocn@foxmail.com
package main

import (
	"flag"
)

type CmdLine struct {
	Help    bool
	Host    string
	APIPath string
	Prefix  string
	Token   string
	SkipTLS bool
}

var cmdline CmdLine

func init() {
	flag.BoolVar(&cmdline.Help, "help", false, "show usage")
	flag.StringVar(&cmdline.Host, "host", "https://127.0.0.1:6443", "kubernetes api host")
	flag.StringVar(&cmdline.APIPath, "apipath", "/", "kubernetes api path")
	flag.StringVar(&cmdline.Prefix, "prefix", "", "kubernetes api prefix")
	flag.StringVar(&cmdline.Token, "token", "", "user's bearer token")
	flag.BoolVar(&cmdline.SkipTLS, "skiptls", false, "don't verify TLS certificate")
}

func main() {
	flag.Parse()
	if cmdline.Help {
		flag.Usage()
		return
	}
	println("Help:", cmdline.Help)
	println("Host:", cmdline.Host)
	println("APIPath:", cmdline.APIPath)
	println("Prefix:", cmdline.Prefix)
	println("Token:", cmdline.Token)
	println("SkipTLS:", cmdline.SkipTLS)
}
