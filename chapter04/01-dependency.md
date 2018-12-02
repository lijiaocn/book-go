<!-- toc -->
# 依赖代码管理

这里记录用于Go项目的依赖管理工具。

## Dep

[dep][3]是一个比较新的依赖管理工具，它计划成为Go的标准工具(截止2017-12-26 20:50:43)。

### 安装

可以下面的方法安装：

```bash
go get -u github.com/golang/dep/cmd/dep
```

在Mac上还可以用brew安装：

```bash
brew install dep
brew upgrade dep
```

### 使用

在项目的根目录中执行`dep init`完成初始化设置。

```bash
$ dep init
$ tree
.
├── Gopkg.lock
├── Gopkg.toml
└── vendor

1 directory, 2 files
```

`Gopkg.toml`和`Gopkg.lock`用来记录项目的依赖包，如果项目中没有代码，这两个文件是空的。

使用`dep ensure -add`添加依赖包（项目目录中需要有源码）：

```bash
dep ensure -add  k8s.io/client-go
```

这时依赖包的依赖包还没有安装到项目本地，需要执行`dep ensure`进行安装。

代码中引用的其它`未通过dep添加`的依赖包，也需要执行`dep ensure`进行安装。

如果要更改依赖的版本，编辑`Gopkg.toml`之后，执行`dep ensure`。

`dep ensure`将所有的依赖包信息更新到Gopkg.lock文件中，并将所有的依赖包下载到vendor目录中。

### 发布

可以只将`Gopkg.toml`和`Gopkg.lock`包含到项目代码库中。之后在项目根目录中执行`dep ensure`下载依赖包：

```bash
dep ensure
```

## Glide

[glide][2]是一个使用比较多的依赖管理工具。

### 安装

在linux上可以用下面的方法安装：

```bash
curl https://glide.sh/get | sh
```

Mac上还可以用brew安装：

```bash
brew install glide
```

### 使用

#### glide create初始化

在项目的根目录中执行`glide create`完成初始化设置。

```bash
$ glide create
[INFO]	Generating a YAML configuration file and guessing the dependencies
[INFO]	Attempting to import from other package managers (use --skip-import to skip)
[INFO]	Scanning code to look for dependencies
[INFO]	Writing configuration file (glide.yaml)
[INFO]	Would you like Glide to help you find ways to improve your glide.yaml configuration?
[INFO]	If you want to revisit this step you can use the config-wizard command at any time.
[INFO]	Yes (Y) or No (N)?
N
[INFO]	You can now edit the glide.yaml file. Consider:
[INFO]	--> Using versions and ranges. See https://glide.sh/docs/versions/
[INFO]	--> Adding additional metadata. See https://glide.sh/docs/glide.yaml/
[INFO]	--> Running the config-wizard command to improve the versions in your configuration
```

如果项目中已经有代码，会将已有代码的依赖写入到glide.yaml中。

之后直接使用`glide up --quick`将已有代码的依赖收集到vendor目录中。

#### glide get 添加依赖包

使用glide安装依赖包的时候，会自动提示版本。例如用glide的命令引入`k8s.io/client-go`时：

```bash
$ glide get k8s.io/client-go
[INFO]	Preparing to install 1 package.
[INFO]	Attempting to get package k8s.io/client-go
[INFO]	--> Gathering release information for k8s.io/client-go
[INFO]	The package k8s.io/client-go appears to have Semantic Version releases (http://semver.org).
[INFO]	The latest release is v6.0.0. You are currently not using a release. Would you like
[INFO]	to use this release? Yes (Y) or No (N)
Y
[INFO]	The package k8s.io/client-go appears to use semantic versions (http://semver.org).
[INFO]	Would you like to track the latest minor or patch releases (major.minor.patch)?
[INFO]	The choices are:
[INFO]	 - Tracking minor version releases would use '>= 6.0.0, < 7.0.0' ('^6.0.0')
[INFO]	 - Tracking patch version releases would use '>= 6.0.0, < 6.1.0' ('~6.0.0')
[INFO]	 - Skip using ranges
[INFO]	For more information on Glide versions and ranges see https://glide.sh/docs/versions
[INFO]	Minor (M), Patch (P), or Skip Ranges (S)?
P
[INFO]	--> Adding k8s.io/client-go to your configuration with the version ~6.0.0
[INFO]	Downloading dependencies. Please wait...
[INFO]	--> Fetching updates for k8s.io/client-go
[INFO]	Resolving imports
[INFO]	Downloading dependencies. Please wait...
[INFO]	--> Detected semantic version. Setting version for k8s.io/client-go to v6.0.0
[INFO]	Exporting resolved dependencies...
[INFO]	--> Exporting k8s.io/client-go
[INFO]	Replacing existing vendor dependencies
```

安装完成之后目录结构如下：

```bash
$ tree -L 2
.
├── glide.lock
├── glide.yaml
└── vendor
    └── k8s.io
```

`glide get`在glide.yaml和glide.lock中记录了依赖的client-go的版本。

这时，依赖包已经安装到了项目本地，但是依赖包的依赖包还需要通过执行`glide up`进行安装。

#### glide up 更新vendor目录

项目的源代码中直接引入的`不是通过glide命令安装`的包，也需要执行`glide up`安装。

`glide up`会在glide.lock中记录依赖包的依赖包和源码引用的依赖包，并将这些依赖包安装到项目本地。

>⚠️  glide up安装的依赖包是直接从github或者其它网址下载的，既不会使用、也不会更改$GOPATH中的源码。

>⚠️  没有指定版本的依赖包，glide up每次执行的时候，都会去获取安装最新的代码。

>⚠️  glide up默认递归搜集依赖，可以用glide up --quick

```bash
$ glide up
[INFO]	Downloading dependencies. Please wait...
[INFO]	--> Fetching updates for k8s.io/client-go
[INFO]	--> Detected semantic version. Setting version for k8s.io/client-go to v5.0.1
[INFO]	Resolving imports
[INFO]	--> Fetching github.com/lijiaocn/GoPkgs
[INFO]	Found Godeps.json file in /Users/lijiao/.glide/cache/src/https-k8s.io-apimachinery
[INFO]	--> Parsing Godeps metadata...
...
[INFO]	--> Exporting k8s.io/kube-openapi
[INFO]	--> Exporting k8s.io/api
[INFO]	Replacing existing vendor dependencies
[INFO]	Project relies on 31 dependencies.

$ head glide.lock
hash: 1002f0b1fae48b0c9e90737a6071f892d98b6bd9016d55f91cca24d25672e4cb
updated: 2017-12-27T10:44:30.038822286+08:00
imports:
- name: github.com/davecgh/go-spew
  version: 782f4967f2dc4564575ca782fe2d04090b5faca8
  subpackages:
  - spew
- name: github.com/emicklei/go-restful
  version: ff4f55a206334ef123e4f79bbf348980da81ca46
  subpackages:
```

如果要更改依赖包的版本，修改`glide.yaml`文件的version后，也执行`glide up`。

```bash
$ cat glide.yaml
package: github.com/lijiaocn/handbook-go/codes/04-01-version/glide
import:
- package: k8s.io/client-go
  version: ~6.0.0
$ glide up -v
```

### 管理

查看依赖包：

	glide list

删除依赖包：

	glide rm

设置镜像站：

	glide mirror

### 发布

发布的时候将`glide.lock`和`glide.yaml`包含到项目源码库中。

获取到项目代码后，项目的根目录中执行`glide install`下载项目的依赖包。

```bash
$ glide install
[INFO]	Downloading dependencies. Please wait...
[INFO]	--> Found desired version locally github.com/davecgh/go-spew 782f4967f2dc4564575ca782fe2d04090b5faca8!
[INFO]	--> Found desired version locally github.com/emicklei/go-restful ff4f55a206334ef123e4f79bbf348980da81ca46!
[INFO]	--> Found desired version locally github.com/emicklei/go-restful-swagger12 dcef7f55730566d41eae5db10e7d6981829720f6!
...
[INFO]	--> Exporting gopkg.in/inf.v0
[INFO]	--> Exporting k8s.io/client-go
[INFO]	--> Exporting k8s.io/kube-openapi
[INFO]	Replacing existing vendor dependencies
```

## Godep

Go语言早期没有提供依赖包管理的功能，[godep][2]是一个比较简单的第三方依赖管理工具。

### 安装

可以用下面的命令安装godep：

	go get github.com/tools/godep

### 下载

godep不会自动拉取代码，使用前，先要确认中`$GOPATH`中的源码的版本是正确的。

例如，如果要使用`k8s.io/client-go`的v6.0.0版本的代码。

先下载源码：

	go get k8s.io/client-go/...

然后将下载的client-go源码切换到v6.0.0：

	cd $GOPATH/src/k8s.io/client-go
	git checkout v6.0.0

因为client-go自身也使用godep进行了依赖管理，所以还需要在client-go中执行：

	godep restore ./...

>执行restore的目的下载client-go的依赖包。

这时可以直接在项目中引用目标版本的代码，但依赖包的文件还是到$GOPATH中读取。

如果另一个项目使用了同名依赖包的另一个版本，两者就会冲突。所以还需要将项目的依赖信息、依赖文件保存到项目本地。

### 保存

通过以下命令，将项目引用的第三方代码以及版本信息保存在本地:

	godep save          #保存当前目录下的go文件(不遍历子目录)引用的第三方代码
	godep save ./...    #保存当前目录以及子目录下的go文件引用的第三方代码

在Go1.5之前，godep将版本信息和第三方代码保存的项目的Godeps目录下。因此在go1.5之前，需要通过godep调用go命令，才会使用保存在项目本地的依赖包。例如：

	godep go build

在Go1.5以后，godep将版本信息保存在godeps目录中,将依赖包保存在项目本地的vendor目录中，可以直接使用go命令。

### 发布

可以将项目本地的依赖包，即Godep目录和vendor目录提交到项目的代码库中。

也可以只将Godep目录提交到项目的代码库中。这样下载了项目代码后，需要用restore下载依赖包：

	go restore

restore会命令将$GOPATH中的源码更新为`Godeps/Godeps.json`中指定的版本，因此最好还是使用第一种发布方式。

例如Godeps.json中的指定metadata的版本为`3b1ae45394a234c385be014e9a488f2bb6eef821`：

```json
{
    "ImportPath": "k8s.io/client-go",
    "GoVersion": "go1.9",
    "GodepVersion": "v79",
    "Packages": [
        "./..."
    ],
    "Deps": [
        {
            "ImportPath": "cloud.google.com/go/compute/metadata",
            "Rev": "3b1ae45394a234c385be014e9a488f2bb6eef821"
        },
    ...(省略)...
```

执行restore命令之后，`$GOPATH/cloud.google.com/go/compute/metadata`中的版本将为变成指定的版本：

```bash
$ cd $GOPATH/cloud.google.com/go/compute/metadata
$ git log |head
commit 3b1ae45394a234c385be014e9a488f2bb6eef821
Author: Sarah Adams <shadams@google.com>
Date:   Thu Sep 8 15:39:53 2016 -0700

    datastore: adds support for loading entity values

    Adds support for conditionally loading either flattened fields or
    entity values into nested structs on read.

    Issue #75
```

### 设置代理

godep是用go语言开发，go语言开发的程序都支持环境变量`http_proxy`。

```
$ http_proxy=127.0.0.1:53100 godep restore ./...
```

godep在加载源码时，会使用git等版本管理工具，所以可能还需要给这些版本管理设置代理。

git可以在~/.gitconfig中设置:

```ini
[https]
proxy = 127.0.0.1:53100
sslverify = false
[http]
proxy = 127.0.0.1:53100
```

## Vendor

vendor是1.5中的一个试验特性，在1.6版本中被正式引入。编译过程中，会先引用vendor目录中的代码。

对于同样的代码main.go:

```go
package main

import (
    "github.com/lijiaocn/GoPkgs/version"
)

func main() {
    version.Show()
}
```

没有vendor之前，项目vendor_test目录结构如下:

```bash
▾ vendor_test/
  ▾ main/
      main.go
```

main.go中引用的是`$GOPATH/.../version`中的文件。

使用vendor之后，目录结构如下：

```bash
▾ vendor_test/
  ▸ Godeps/
  ▾ main/
      main.go
  ▾ vendor/
    ▾ github.com/
      ▾ lijiaocn/
        ▾ GoPkgs/
          ▸ version/
            LICENSE
```

main.go中引用的是本地的`vendor/.../version`中的文件。

不需要对main.go做任何修改。

## 参考

1. [godep][1]
2. [glide][2]
3. [dep][3]

[1]: https://github.com/tools/godep "godep"
[2]: https://github.com/Masterminds/glide "glide"
[3]: https://github.com/golang/dep "dep"
