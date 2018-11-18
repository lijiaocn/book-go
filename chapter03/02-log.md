# 打印日志

## glog

`glog`原本是google使用C++实现的一个日志模块([glog c++][2])，go语言出来以后，又实现了一个Go版本的[glog][1]。

### 代码

glog依赖标准包flag，在调用glog的方法之前，需要调用`flag.Parse()`。

```go
//create: 2017/12/25 14:24:03 change: 2017/12/25 15:08:49 lijiaocn@foxmail.com
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
```

### 运行

glog支持以下参数。

```bash
$ ./glog -h
Usage of ./glog:
  -alsologtostderr
        log to standard error as well as files
  -log_backtrace_at value
        when logging hits line file:N, emit a stack trace
  -log_dir string
        If non-empty, write log files in this directory
  -logtostderr
        log to standard error instead of files
  -stderrthreshold value
        logs at or above this threshold go to stderr
  -v value
        log level for V logs
  -vmodule value
        comma-separated list of pattern=N settings for file-filtered logging
```

glog将日志以文件的形式写入`-log_dir`指定的目录，使用`-log_dir`时只有ERROR日志会打印到标准输出，

```bash
$ ./glog -log_dir=./log
E1225 15:19:10.514843   83246 main.go:13] This is a error
```

生成的日志文件：

```bash
lijiaos-MacBook-Pro:glog lijiao$ ls log/
glog.ERROR
glog.INFO
glog.WARNING
glog.lijiaos-MacBook-Pro.lijiao.log.ERROR.20171225-151858.83190
glog.lijiaos-MacBook-Pro.lijiao.log.ERROR.20171225-151910.83246
glog.lijiaos-MacBook-Pro.lijiao.log.INFO.20171225-151858.83190
glog.lijiaos-MacBook-Pro.lijiao.log.INFO.20171225-151910.83246
glog.lijiaos-MacBook-Pro.lijiao.log.WARNING.20171225-151858.83190
glog.lijiaos-MacBook-Pro.lijiao.log.WARNING.20171225-151910.83246
```

参数`-logtostderr`将所有日志打印到标准输出，不保存到文件中。

```bash
$ ./glog  -logtostderr
I1225 15:22:11.770424   83995 main.go:11] This is a info
W1225 15:22:11.770510   83995 main.go:12] This is a warning
E1225 15:22:11.770513   83995 main.go:13] This is a error
```

参数`-v`指定日志的级别，只打印V(x)中，x小于等于-v指定的数值的日志。

```bash
$ ./glog  -logtostderr -v 1
I1225 15:28:03.517261   85366 main.go:11] This is a info
W1225 15:28:03.517335   85366 main.go:12] This is a warning
E1225 15:28:03.517339   85366 main.go:13] This is a error
I1225 15:28:03.517342   85366 main.go:14] This is a v1 log
```

### 建议

必须要记录的日志使用`glog.InfoX`、`glog.WarningX`、`glog.ErrorX`等函数打印。

用于开发调试的日志使用`glov.V().InfoX`打印。

### 参考

1. [glog golang][1]
2. [glog c++][2]

[1]: https://github.com/golang/glog  "glog golang" 
[2]: https://github.com/google/glog "glog c++"
