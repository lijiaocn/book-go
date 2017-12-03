# 安装go

## 下载软件包

到golang官方网站的[下载页][1]下载软件包，golang的网站`https://golang.org`需要翻墙访问。

如果你不能翻墙，可以到golang中国的[下载页][2]中下载，golang中国的网址是`https://www.golangtc.com`。

	go1.9.2.linux-386.tar.gz     Archive   Linux   x86     88MB   574b2c4b1a248e58ef7d1f825beda15429610a2316d9cbd3096d8d3fa8c0bc1a
	go1.9.2.linux-amd64.tar.gz   Archive   Linux   x86-64  99MB   de874549d9a8d8d8062be05808509c09a88a248e77ec14eb77453530829ac02b

这里以linux版本为例，下载了安装包之后首先做一下校验（这是一个良好的习惯！）。

	$ sha256sum go1.9.2.linux-amd64.tar.gz
	de874549d9a8d8d8062be05808509c09a88a248e77ec14eb77453530829ac02b  go1.9.2.linux-amd64.tar.gz

得到的校验码与网站上公布的校验码核对无误之后，才可以放心的使用。

> 为什么要对下载的文件做校验？
> 互联网其实充满了各种秘密的，我们从网络上下载下来的文件很有可能是被掉过包的。
> 很多人都有做这种事情的动力，譬如黑客、运营商、捣蛋分子。
> 使用一个被掉过包的文件是非常危险的，因为我们无法预知它被做了什么手脚。
> 因此，一定要养成从官网上下载文件的习惯，而且下载完成之后要做校验。

[1]: https://golang.org/dl/ "golang.org download"
[2]: https://www.golangtc.com/download "golangtc.com download"
