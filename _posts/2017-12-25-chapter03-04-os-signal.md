---
layout: post
title: 04-捕获系统信号
author: lijiaocn
createdate: 2017/12/27 17:19:02
changedate: 2017/12/28 15:10:48
categories: chapter03
tags:
keywords:
description: 

---

* auto-gen TOC:
{:toc}

## 说明

`os/signal`是Go的一个标准库，用来接收信号量。

信号是通过操作系统传递给进程的，例如命令`kill -9 1000`，就是向进程号为1000的进程发出了信号`9`。

## 发送信号

可以使用`kill`命令向任意一个进程发送信号。有三种发送形式。

使用`-s`指定信号名称，例如：

	kill -s SIGINT  进程号

使用`-信号名称`，例如：

	kill -SIGINT  进程号

使用`-信号数值`，例如：

	kill -2  进程号

使用`kill -l`可以查看所有信号的名称以及数值。

{% highlight shell %}
$ kill -l
 1) SIGHUP	 2) SIGINT	 3) SIGQUIT	 4) SIGILL
 5) SIGTRAP	 6) SIGABRT	 7) SIGEMT	 8) SIGFPE
 9) SIGKILL	10) SIGBUS	11) SIGSEGV	12) SIGSYS
13) SIGPIPE	14) SIGALRM	15) SIGTERM	16) SIGURG
17) SIGSTOP	18) SIGTSTP	19) SIGCONT	20) SIGCHLD
21) SIGTTIN	22) SIGTTOU	23) SIGIO	24) SIGXCPU
25) SIGXFSZ	26) SIGVTALRM	27) SIGPROF	28) SIGWINCH
29) SIGINFO	30) SIGUSR1	31) SIGUSR2
{% endhighlight %}

## 信号的种类

在Mac/Linux上，可以通过`man signal`查看所有的信号量，以及这些信号量的默认处理动作。

{% highlight shell %}
No    Name         Default Action       Description
1     SIGHUP       terminate process    terminal line hangup
2     SIGINT       terminate process    interrupt program
3     SIGQUIT      create core image    quit program
4     SIGILL       create core image    illegal instruction
5     SIGTRAP      create core image    trace trap
6     SIGABRT      create core image    abort program (formerly SIGIOT)
7     SIGEMT       create core image    emulate instruction executed
8     SIGFPE       create core image    floating-point exception
9     SIGKILL      terminate process    kill program
10    SIGBUS       create core image    bus error
11    SIGSEGV      create core image    segmentation violation
12    SIGSYS       create core image    non-existent system call invoked
13    SIGPIPE      terminate process    write on a pipe with no reader
14    SIGALRM      terminate process    real-time timer expired
15    SIGTERM      terminate process    software termination signal
16    SIGURG       discard signal       urgent condition present on socket
17    SIGSTOP      stop process         stop (cannot be caught or ignored)
18    SIGTSTP      stop process         stop signal generated from keyboard
19    SIGCONT      discard signal       continue after stop
20    SIGCHLD      discard signal       child status has changed
21    SIGTTIN      stop process         background read attempted from control terminal
22    SIGTTOU      stop process         background write attempted to control terminal
23    SIGIO        discard signal       I/O is possible on a descriptor (see fcntl(2))
24    SIGXCPU      terminate process    cpu time limit exceeded (see setrlimit(2))
25    SIGXFSZ      terminate process    file size limit exceeded (see setrlimit(2))
26    SIGVTALRM    terminate process    virtual time alarm (see setitimer(2))
27    SIGPROF      terminate process    profiling timer alarm (see setitimer(2))
28    SIGWINCH     discard signal       Window size change
29    SIGINFO      discard signal       status request from keyboard
30    SIGUSR1      terminate process    User defined signal 1
31    SIGUSR2      terminate process    User defined signal 2
{% endhighlight %}

## Go的信号定义

每个信号量都有一个唯一的整数值和名字，Go在标准包`syscall`中定义了这些信号量：

{% highlight go %}
const (
	SIGABRT   = Signal(0x6)
	SIGALRM   = Signal(0xe)
	SIGBUS    = Signal(0xa)
	SIGCHLD   = Signal(0x14)
	SIGCONT   = Signal(0x13)
	SIGEMT    = Signal(0x7)
	SIGFPE    = Signal(0x8)
	SIGHUP    = Signal(0x1)
	SIGILL    = Signal(0x4)
	SIGINFO   = Signal(0x1d)
	SIGINT    = Signal(0x2)
	SIGIO     = Signal(0x17)
	SIGIOT    = Signal(0x6)
	SIGKILL   = Signal(0x9)
	SIGPIPE   = Signal(0xd)
	SIGPROF   = Signal(0x1b)
	SIGQUIT   = Signal(0x3)
	SIGSEGV   = Signal(0xb)
	SIGSTOP   = Signal(0x11)
	SIGSYS    = Signal(0xc)
	SIGTERM   = Signal(0xf)
	SIGTRAP   = Signal(0x5)
	SIGTSTP   = Signal(0x12)
	SIGTTIN   = Signal(0x15)
	SIGTTOU   = Signal(0x16)
	SIGURG    = Signal(0x10)
	SIGUSR1   = Signal(0x1e)
	SIGUSR2   = Signal(0x1f)
	SIGVTALRM = Signal(0x1a)
	SIGWINCH  = Signal(0x1c)
	SIGXCPU   = Signal(0x18)
	SIGXFSZ   = Signal(0x19)
)
{% endhighlight %}

## 示例

代码：

{% highlight go %}
//create: 2017/12/27 17:34:37 change: 2017/12/28 10:31:34 lijiaocn@foxmail.com
package main

import (
	"os"
	"os/signal"
	"syscall"
)

func main() {
	println(syscall.Getpid())

	receive := make(chan os.Signal, 10)

	signal.Ignore(syscall.SIGUSR1)
	signal.Notify(receive, syscall.SIGUSR2)

	for {
		select {
		case s := <-receive:
			switch s {
			case syscall.SIGUSR1:
				println("receive SIGUSR1")
			case syscall.SIGUSR2:
				println("receive SIGUSR2")
			}
		}
	}
}
{% endhighlight %}

运行:

	$ ./signal
	18199

用kill向进程18199发送SIGUSR1和SIGUSR2信号:

	$ kill -SIGUSR1 18199
	$ kill -SIGUSR2 18199

SIGUSR1被忽略，SIGUSR2收到：

	$ ./signal
	18199
	receive SIGUSR2

## 参考

1. [godoc: syscall#SIGINT][1]

[1]: http://golang.org/pkg/syscall/#SIGINT "godoc: syscall#SIGINT"
