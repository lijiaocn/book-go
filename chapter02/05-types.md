---
layout: default
title: 05-types
author: lijiaocn
createdate: 2017/12/18 10:55:10
changedate: 2017/12/20 11:11:51
categories:
tags:
keywords:
description: 

---

* auto-gen TOC:
{:toc}

# go的类型 

类型是用来诠释如何解读指定位置中存放的数据，以及约定操作符的含义的。

## 类型的属性

### 内置类型(predeclared)

go语言内置了下面的类型：

	bool byte complex64 complex128 error float32 float64
	int int8 int16 int32 int64 rune string
	uint uint8 uint16 uint32 uint64 uintptr

### 命名类型(named)

类型可以是命名的(named)，也可以是未命名的(unnamed)。

	Type      = TypeName | TypeLit | "(" Type ")" .
	TypeName  = identifier | QualifiedIdent .
	TypeLit   = ArrayType | StructType | PointerType | FunctionType | InterfaceType |
	            SliceType | MapType | ChannelType .

使用`type`指定了名字的类型是命名的，例如下面的类型的名字为Student

	type Student struct {
		Name string
		age int
	}

由其它类型组合成的新类型，可以不被命名，例如下面的类型是没有名字的：

	[] string
	[] int

无类型的名字，通用用于定义其它类型:

	type Array []int

或者在函数定义中使用：

	func Display(s struct {
	    name string
	    age  int
	}) {
	    println(s.name)
	    println(s.age)
	}


### 实际类型(underlying type)

类型是可以用来定义其它类型的，例如：

	type T1 string 
	type T2 T1

这里定义了一个类型T1，然后又用T1定义了类型T2。

T1的实际类型(underlying type)是string，T2的实际类型不是T1，而是T1的实际类型string。

实际类型必须是go的内置类型或者类型的组合。

例如，string、T1、T2的实际类型是string。

	type T1 string
	type T2 T1

[]T1、T3、T4的实际类型是[]T1。

	type T3 []T1
	type T4 T3

### 类型的方法(method sets)

类型可以有自己的方法(Method)，也就是其它语言中的函数。

一个非接口类型的方法集就所有接受者(receiver)为改类型的方法，接口类型的方法集就是接口定义中包含的方法。

需要注意的是指针类型类型（例如 * T)，它的方法集是所有接受者为所指类型(T)和指针类型( * T)的方法集。

例如下面的代码中，方法的Show的Receiver是Str，但是类型为 * Str的pstr也可以调用。

	package main
	
	type Str string
	
	func (s Str) Show() {
		println(s)
	}
	
	func main() {
		str := Str("Hello World!")
		pstr := &str
		pstr.Show()
	}

方法集中的方法不能重名、且必须有名字。

### 类型的等同性(identical)

命名语句不同的两个命名类型，是不等同的。例如下面的T1和T2，虽然实际类型都是string，但它们是两个类型。

	type T1 string
	type T2 string

命名类型与未命名类型是不等同的，例如下面的T1与[]string是两个类型。

	type T1 []string
	[]string

命名语句和定义语句`完全相同`的两个命名类型是才等同的，例如下面的T1。

	type T1 string
	type T1 string

定义语句`完全相同`的两个未命名类型才是等同的，例如下面的[]string。

	[5]string
	[5]string

在编程时，同一个类型只会定义一次。

在代码中定义`两个`等同的类型其是做不到的，因为它们如果等同，那么其实就是一个。例如下面的代码。

	package main
	
	type T string
	type T string
	
	func main() {
	}

编译时会报错。

	./main.go:5: T redeclared in this block
	    previous declaration at ./main.go:4

两个类型等同是一个用来解释类型不等同的规则，即如果不符合遵守等同的规则，那么就是不等同的。

对于未命名类型需要特别注意，只要不满足下面的条件，那么就是两个不同的类型：

	两个数组类型要等同，不仅数组中成员的类型要相同，而且数组的长度也要相同。
	两个分片类型等同，只需要分片的成员类型相同。
	两个结构体等同，结构体成员的顺序、名称、标签(tag)都必须相同。
	两个指针类型，所指向的类型相同。
	两个函数类型，要有相同的数量的参数和返回值，参数和返回值的类型要相同，参数名和返回值的名字可以不同。
	两个接口类型，要有相同的方法，方法的顺序可以不同。
	两个字典类型，key的类型和value的类型必须相同。
	两个通道(channel)类型，通道的方向和传递的类型必须相同。

例如下面两个函数类型符合上面的条件，所以是相同的：

	func(x int, y float64) *[]string
	func(int, float64) (result *[]string)

### 类型的赋值(Assignability)

一个值(value)只有在满足下面的条件时，才可以被赋给对应的类型的变量(variable)。

	值的类型与变量的类型相同
	值的类型与变量的实际类型相同，且其中一个的类型是未命名的类型
	变量的类型是一个接口，值实现了接口中方法
	值是一个双向的通道(channel)，变量类型也是通道，传递的数据类型相同，并且其中一个的类型是未命名的。
	值是内置的数值nil，变量的类型是指针(pointer)、函数(function)、分片(slice)、字典(map)、通道(channel)、接口(interface)
	值是一个符合变量的类型要求的常量。

## go支持的类型

### 布尔(Boolean types)

布尔类型是内置的类型`bool`，它的value只能是两个内置的常量：

	true
	false

### 数值(Numeric types)

数值类型都是内置的类型，一共有以下几种。

	uint8       the set of all unsigned  8-bit integers (0 to 255)
	uint16      the set of all unsigned 16-bit integers (0 to 65535)
	uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
	uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)
	
	int8        the set of all signed  8-bit integers (-128 to 127)
	int16       the set of all signed 16-bit integers (-32768 to 32767)
	int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
	int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)
	
	float32     the set of all IEEE-754 32-bit floating-point numbers
	float64     the set of all IEEE-754 64-bit floating-point numbers
	
	complex64   the set of all complex numbers with float32 real and imaginary parts
	complex128  the set of all complex numbers with float64 real and imaginary parts
	
	byte        alias for uint8
	rune        alias for int32

另外还有三个数值类型，它们占用的空间取决于实现：

	uint     either 32 or 64 bits
	int      same size as uint
	uintptr  an unsigned integer large enough to store the uninterpreted bits of a pointer value

### 字符串(String types)

字符串是内置的类型`string`，字符串的值是连续的字节，这些字节是不可更改的。

可以通过内置函数`len`获取字符串的长度，可以用通过[i]读取字符串的第i个(从0开始)字节。

字符串的字节只能读取，不能更改，也不能取址。

	package main
	
	import (
	    "fmt"
	)
	
	func main() {
	    str := "Hello World!"
	    fmt.Printf("%c\n", str[6])
	
	    //not allow
	    //ptr := &str[6]
	
	    //not allow
	    //str[6] = 'w'
	}

### 数组(Array types)

数组是多个相同类型的值，在go中，数组必须有长度，长度是数组类型的一部分。

	ArrayType   = "[" ArrayLength "]" ElementType .
	ArrayLength = Expression .
	ElementType = Type .

数组是单维的，可以累进成多维数组：

	[32]byte
	[2*N] struct { x, y int32 }
	[1000]*float64
	[3][5]int
	[2][2][2]float64  // same as [2]([2]([2]float64))

要注意长度是数组类型的一部分，长度不同的数组是不同的类型，例如：

	package main
	
	func main() {
	    var array1 [32]byte
	    var array2 [24]byte
	
	    array1[0] = 'a'
	    array2[0] = 'b'
	
	    //not allow
	    //array2 = array1
	}

数组成员可以用从0开始的坐标索引，长度可以用内置的函数`len`获取。

### 分片(Slice types)

分片(Slice)是用来索引数组(Array)中的一段连续的成员的。

	SliceType = "[" "]" ElementType .

分片初始化后就绑定到了一个数组，多个分片可以绑定到同一个数组。

与数组不同的是，分片有长度(length)和容量(capacity)两个属性。

长度是分片所索引的数组成员的数量，可以用内置的函数`len`获取。

容量是分片能够索引的数组成员的最大数量，等于数组的长度减去分片索引的第一个数组成员在数组中位置。

例如在下面的代码中，分片slice1的长度是5，容量是20(=30-10)

	package main
	
	func main() {
	    var array1 [30]int
	    for i := 0; i < len(array1); i++ {
	        array1[i] = i
	    }
	
	    slice1 := array1[10:15]
	
	    println("array's length: ", len(array1))
	    println("slice1's length: ", len(slice1))
	    println("slice1's capacity: ", cap(slice1))
	
	    for i := 0; i < len(slice1); i++ {
	        println(slice1[i])
	    }
	}

分片可以通过两种方式创建，第一种方式就是上面的代码中使用的方式：

	    slice1 := array1[10:15]

这样创建的slice1索引的是数组的从0开始编号的第10个、第11个、第12个、第13个、第14个个成员，总计5个。

	10
	11
	12
	13
	14

>注意[10:15]是一个前闭后开的集合，即包括10，不包括15。

第二种方式是使用内置的`make`函数创建。

	make([]T, length, capacity)

使用make创建的时候，至少需要指定分片的长度，make会为分片创建一个隐藏的数组。

如果指定了capacity，数组的长度就是capacity，如果没有指定，数组的长度等于分片的长度。

例如下面的代码中slice2的长度和容量都是10，slice3的长度是10，容量是20。

	package main
	
	func main() {
	    //not allow
	    //slice1 := make([]int)
	    //println("slice1， len is ", len(slice1), "capacity is ", cap(slice1))
	
	    slice2 := make([]int, 10)
	    println("slice2， len is ", len(slice2), "capacity is ", cap(slice2))
	
	    slice3 := make([]int, 10, 20)
	    println("slice3， len is ", len(slice3), "capacity is ", cap(slice3))
	}

通过make创建分片，相当与新建一个数组，然后取它的[0:length]。

	make([]int, 50, 100)

等同于：

	new([100]int)[0:50]

### 结构体(Struct types)

结构体(Struct)是比较复杂的类型，它是由多个命名变量组成，这些变量每个都有名字和类型，被成为"结构体成员(field)"。

	StructType     = "struct" "{" { FieldDecl ";" } "}" .
	FieldDecl      = (IdentifierList Type | AnonymousField) [ Tag ] .
	AnonymousField = [ "*" ] TypeName .
	Tag            = string_lit .

go语言的struct用法与C语言中的不同，C语言中是“struct 结构体名{ 结构体成员...}”，go语言中没有中间的结构体名。如果要给go的结构体命名，需要使用关键type：

	type 结构体名 struct{
		结构体成员
	}

结构体成员的名称可以显示声明(IdentifierList Type)，也可以隐式声明(AnonymousField)。

隐式声明明是不为变量设置明确的标识符时，变量的名字默认为类型的名字。例如：

	struct {
	    T1        // field name is T1
	    *T2       // field name is T2
	    P.T3      // field name is T3
	    *P.T4     // field name is T4
	    x, y int  // field names are x and y
	}


go语言中的隐式声明的成员，其实有一点C++中的继承的意思，例如在下面的定义中，结构体B可以直接使用它的隐式成员A的结构体成员：

	package main
	
	import (
	    "fmt"
	)
	
	type A struct {
	    A1 string
	    A2 string
	}
	
	type B struct {
	    A
	    B1 string
	    B2 string
	}
	
	func main() {
	    b := B{
	        A: A{
	            A1: "a1",
	            A2: "a2",
	        },
	        B1: "b1",
	        B2: "b2",
	    }
	    fmt.Println(b.A)
	    fmt.Println(b.A.A1)
	    fmt.Println(b.A1)
	}

在上面的代码中，结构体B没有显示声明为A1的成员，因此`b.A1`索引的是它的隐式成员A的成员。

如果结构体B有一个名为A1的显示成员，那么只能通过`b.A.A1`的方式索引到A的成员A1，`b.A`索引的将是B的显示成员A1。

例如下面代码中，最后一行打印的是`b1's a1`。

	package main
	
	import (
	    "fmt"
	)
	
	type A struct {
	    A1 string
	    A2 string
	}
	
	type B struct {
	    A
	    A1 string
	    B1 string
	    B2 string
	}
	
	func main() {
	    b := B{
	        A: A{
	            A1: "a1",
	            A2: "a2",
	        },
	        A1: "b's a1",
	        B1: "b1",
	        B2: "b2",
	    }
	    fmt.Println(b.A)
	    fmt.Println(b.A.A1)
	    fmt.Println(b.A1)
	}

同一个结构体内的成员不能重名，在使用隐式声明的时候要特别注意，因为一个类型与它的指针类型，在被隐式声明的时候，会得到相同的变量名。例如下面的结构体的三个成员的名字都是`T`，这是不允许的。

	struct {
	    T     // conflicts with anonymous field *T and *P.T
	    *T    // conflicts with anonymous field T and *P.T
	    *P.T  // conflicts with anonymous field T and *T
	}

隐式声明的`T`和隐式声明的`*T`的区别之一是这个隐式声明的变量的存放位置。另外go语言声称：

	If S contains an anonymous field T, the method sets of S and *S both include promoted methods with receiver T. The method set of *S also includes promoted methods with receiver *T.
	
	If S contains an anonymous field *T, the method sets of S and *S both include promoted methods with receiver T or *T.

但是试验却发现，下面的两段代码执行的效果是相同的。

代码一，隐式成员为`A`：

	package main
	
	type A struct {
	    A1 string
	}
	
	type B struct {
	    A
	    B1 string
	    B2 string
	}
	
	func main() {
	    b := B{
	        A: A{
	            A1: "a1",
	        },
	        B1: "b1",
	        B2: "b2",
	    }
	
	    b.method()
	    println(b.A1)
	
	    b.pointer_method()
	    println(b.A1)
	
	    pb := &b
	
	    pb.method()
	    println(b.A1)
	
	    pb.pointer_method()
	    println(b.A1)
	}

代码二，隐式成员为`*A`：

	package main
	
	type A struct {
	    A1 string
	}
	
	func (a A) method() {
	    a.A1 = "method set a1"
	}
	
	func (a *A) pointer_method() {
	    a.A1 = "pointer method set a1"
	}
	
	type B struct {
	    *A
	    B1 string
	    B2 string
	}
	
	func main() {
	    b := B{
	        A: &A{
	            A1: "a1",
	        },
	        B1: "b1",
	        B2: "b2",
	    }
	
	    b.method()
	    println(b.A1)
	
	    b.pointer_method()
	    println(b.A1)
	
	    pb := &b
	
	    pb.method()
	    println(b.A1)
	
	    pb.pointer_method()
	    println(b.A1)
	}

go语言中可以在每个结构体成员后面跟随一个标签(tag)，标签用来注明成员的属性。标签可以是解释型字符串，也可以是原始型字符串。

	Tag            = string_lit .
	string_lit     = raw_string_lit | interpreted_string_lit .

另外，在结构体中还可以添加只起到填充(padding)作用的成员：

	// A struct with 6 fields.
	struct {
	    x, y int
	    u float32
	    _ float32  // padding
	    A *[]int
	    F func()
	}

### 指针(Pointer types)

指针类型比较简单：

	PointerType = "*" BaseType .
	BaseType    = Type .

支持多重指针：

	package main
	
	func main() {
	    i := 8
	    pi := &i
	    ppi := &pi
	
	    println(*ppi, pi)
	    println(*pi, i)
	}

### 函数(Function types)

go语言的函数的声明格式与其它语言也有所不同。

	FunctionType   = "func" Signature .
	Signature      = Parameters [ Result ] .
	Result         = Parameters | Type .
	Parameters     = "(" [ ParameterList [ "," ] ] ")" .
	ParameterList  = ParameterDecl { "," ParameterDecl } .
	ParameterDecl  = [ IdentifierList ] [ "..." ] Type .

可以由以下几种样式的函数：

	func()
	func(x int) int
	func(a, _ int, z float32) bool
	func(a, b int, z float32) (bool)
	func(prefix string, values ...int)
	func(a, b int, z float64, opt ...interface{}) (success bool)
	func(int, int, float64) (float64, *[]int)
	func(n int) func(p *T)

最显著的不同是，参数的类型是在参数名之后的，如果两个参数类型相同且位置相临，可以省略前一个参数的类型，例如：

	func(a, b int, z float32) (bool)

函数的最后一个参数可以是变长参数(variadic)，可以对应0个到多个输入参数：

	func(prefix string, values ...int)

函数可以有多个返回值：

	func(int, int, float64) (float64, *[]int)

也可以返回函数：

	func(n int) func(p *T)

注意，这里给出的是函数类型，函数类型不等于函数的声明与实现，函数的声明与实现在后面章节中。

### 接口(Interface types)

接口类型的格式如下：

	InterfaceType      = "interface" "{" { MethodSpec ";" } "}" .
	MethodSpec         = MethodName Signature | InterfaceTypeName .
	MethodName         = identifier .
	InterfaceTypeName  = TypeName .

例如：

	interface {
	    Read(b Buffer) bool
	    Write(b Buffer) bool
	    Close()
	}

接口的成员是方法(method)，一个类型只要实现一个接口中的所有方法的类型，可以作为类型为该接口的变量的的动态类型。

例如下面的T就实现了上面的接口：

	func (p T) Read(b Buffer) bool { return … }
	func (p T) Write(b Buffer) bool { return … }
	func (p T) Close() { … }

一个类型可以实现多个接口的方法，也可以是空的，不包含任何的方法：

	interface{}

接口可以包含其它的接口，但是不能包含它自身，或者通过其它接口形成了重复包含：

	// illegal: Bad cannot embed itself
	type Bad interface {
	    Bad
	}
	
	// illegal: Bad1 cannot embed itself using Bad2
	type Bad1 interface {
	    Bad2
	}
	type Bad2 interface {
	    Bad1
	}

### 字典(Map types)

go语言原生支持字典(map)。

	MapType     = "map" "[" KeyType "]" ElementType .
	KeyType     = Type .

Key的类型不能是函数(function)、字典(map)、分片(slice)

如果Key的类型是接口，可以作为该接口变量的动态类型的类型必须是可比较的，否则会panic。

字典中的成员数量成为字典的长度(length)，可以通过内置函数len()获取。

字典的成员可以通过赋值操作增加，用Key作为index读取。

如果要删除字典中的成员，需要使用内置的delete()函数。

map需要使用内置函数make创建:

	make(map[string]int)
	make(map[string]int, 100)

创建时指定length意思是，预先分配出可以容纳这么多成员的空间，而不是只能容纳这么多。

map的长度不受创建时指定的length的限制，可以无限增加成员。

	package main
	
	import (
	    "fmt"
	)
	
	func main() {
	    m := make(map[int]int, 10)
	    for i := 0; i < 10; i++ {
	        m[i] = i
	    }
	    println(len(m))
	    fmt.Println(m)
	    m[11] = 11
	    println(len(m))
	    fmt.Println(m)
	}

### 通道(Channel types)

通道是用来在并发编程中传递value的。

	ChannelType = ( "chan" | "chan" "<-" | "<-" "chan" ) ElementType .

它可以是可读、可写、既可读又可写的，例如：

	chan T          // can be used to send and receive values of type T
	chan<- float64  // can only be used to send float64s
	<-chan int      // can only be used to receive ints

<-是靠左临近的，通道类型本身也开始被传递：

	chan<- chan int    // same as chan<- (chan int)
	chan<- <-chan int  // same as chan<- (<-chan int)
	<-chan <-chan int  // same as <-chan (<-chan int)
	chan (<-chan int)

通道类型的变量必须用内置的make函数创建：

	make(chan int, 100)

第二参数是指定通道中可以缓存的成员的数量，如果没有第二个参数或者第二个参数为0，那么该通道是不做缓存的，必须等对方接收或者写入完成后，才可以完成写入或接收。

通道需要由写入方使用内置的close函数关闭，接收方收取了最后一个数据后，再从通道中试图读取的时候，会立即返回失败。

例如，如果通道c被关闭，且通道中没有数据了，下面的语句将会立即返回，且ok是false。

	x, ok := <-c

通道是并发安全的，使用内置函数len读取通道中缓存的数据个数，或者用cap读取通道容量，不需要考虑并发的影响。

另外通道中的数据遵循先入先出的规则。
