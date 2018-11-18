# 实现命令行参数

## Flag

`flag`是Go语言的一个标准包，用来处理命令行参数。

### 代码

```go
//create: 2017/10/26 22:04:34 change: 2017/12/25 13:38:09 lijiaocn@foxmail.com
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
```

### 运行

查看使用方法：

```bash
$ ./k8s-haproxy -h
Usage of ./k8s-haproxy:
  -apipath string
        kubernetes api path (default "/")
  -help
        show usage
  -host string
        kubernetes api host (default "https://127.0.0.1:6443")
  -prefix string
        kubernetes api prefix
  -skiptls
        don't verify TLS certificate
  -token string
        user's bearer token
```

指定参数值：

```bash
$ ./k8s-haproxy -host=192.168.2.1:443 -token="this is my token"
Help: false
Host: 192.168.2.1:443
APIPath: /
Prefix:
Token: this is my token
SkipTLS: false
```
