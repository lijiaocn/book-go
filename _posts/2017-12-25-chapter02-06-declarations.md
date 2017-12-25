---
layout: post
title: 06-声明语法
author: lijiaocn
createdate: 2017/12/20 11:12:59
changedate: 2017/12/25 12:44:17
categories: chapter02
tags:
keywords:
description: 

---

* auto-gen TOC:
{:toc}

## 代码区块、声明、声明作用域名

声明的影响是有范围的，它的影响范围叫做作用域，作用域对应的是代码区块。

## 代码区块(blocks) 

	Block = "{" StatementList "}" .
	StatementList = { Statement ";" } .

在go中有这样几种代码区块：

	所有的代码组成一个终极区块(universe block)
	隶属于同一个package的代码，组成对应的包区块(package block)
	同一个文件中的代码，组成一个文件区块(file block)
	if、for、switch语句包裹的代码，组成了独立的隐式区块(implicit block)
	switch、select的条件(clause)语句中的代码，组成了独立的隐式区块
	"{"和"}"包裹的代码，组成一个隐式区块

## 声明的作用域(Declarations and scope)

声明就是设置标记符(identifier)的过程，实现标记符以下内容的绑定：

	constant，常量
	type，类型
	variable，变量
	function，函数
	label，标记
	package，包

声明的语法格式：

	Declaration   = ConstDecl | TypeDecl | VarDecl .
	TopLevelDecl  = Declaration | FunctionDecl | MethodDecl .

在包区块中，`init`只能用于声明init函数。

声明的效果是限定在区块中的。

	go内置的标记符在终极区块(universe block)中有效
	在函数之外声明的常量、类型、变量、函数在包区块(package block)中有效，注意不包括方法(method)。
	通过import导入的包(package)的名字，在文件区块(file block)中有效
	声明的方法的接收者(receiver)、函数参数、函数返回值，在函数的代码区块中有效
	在函数的代码区块中声明的常量、变量、类型，在声明位置和所在的最内层区块的末尾之间有效

代码区块是可以嵌套的，内层代码区块中的声明在内存代码区块中覆盖外层代码区块中的声明。

## 标记作用域(Label scopes)

标记(Label)的作用域与其它的标记符不同，它被用于`break`、`continue`、`goto`。

标记一旦声明，必须使用，否则编译报错。

标记在函数内声明，它在整个函数区块以及函数区块的嵌入区块中有效，并且可以与其它标识符同名，

## 内置的标记符

go内置了空白标记符(Blank identifier)和预声明的标记服(Predeclared identifiers)。

空白标记符就是一个下划线`_`，表示对应的目标不被声明。

预声明的标记符有以下这些:

	Types:
	    bool byte complex64 complex128 error float32 float64
	    int int8 int16 int32 int64 rune string
	    uint uint8 uint16 uint32 uint64 uintptr
	
	Constants:
	    true false iota
	
	Zero value:
	    nil
	
	Functions:
	    append cap close complex copy delete imag len
	    make new panic print println real recover

## 标识符的导出(Exported identifiers)

可以将包区块(package block)中的满足条件的标记符导出到其它包区块中。

	标记符必须以大写字母开头(Unicode upper case letter)
	标记符是在包区块中声明的，或者是结构的成语名（filed name)、方法名(method name)

不符合这两点的标记符不能导出。

## 标记符的唯一性(Uniqueness of identifiers)

不同名的两个标记符是不同的，在不同的包中同名的两个标记符也不同的。

## 常量的声明(Constant declarations)

常量的声明是将常量、常量表达式绑定到指定的标记符，之后可用标记符读取常量。

	ConstDecl      = "const" ( ConstSpec | "(" { ConstSpec ";" } ")" ) .
	ConstSpec      = IdentifierList [ [ Type ] "=" ExpressionList ] .
	
	IdentifierList = identifier { "," identifier } .
	ExpressionList = Expression { "," Expression } .

有以下几种声明样式：

	const Pi float64 = 3.14159265358979323846
	const zero = 0.0         // untyped floating-point constant
	const (
	    size int64 = 1024
	    eof        = -1  // untyped integer constant
	)
	const a, b, c = 3, 4, "foo"  // a = 3, b = 4, c = "foo", untyped integer and string constants
	const u, v float32 = 0, 3    // u = 0.0, v = 3.0

在使用小括号样式时，如果后续声明的常量表达式是相同的，那么可以省略这些常量表达样式。

下面的声明:

	const (
	    Sunday = iota
	    Monday
	    Tuesday
	)

等同于：

	const (
	    Sunday = iota
	    Monday = iota
	    Tuesday = iota
	)

go内置常量iota是一个特殊的常量表达式，它在`const`关键字之后第一次出现是value是0，在后续的每次声明中，value增加1，直到遇到下一个const后，重新归零。

	const ( // iota is reset to 0
	    c0 = iota  // c0 == 0
	    c1 = iota  // c1 == 1
	    c2 = iota  // c2 == 2
	)
	
	const ( // iota is reset to 0
	    a = 1 << iota  // a == 1
	    b = 1 << iota  // b == 2
	    c = 3          // c == 3  (iota is not used but still incremented)
	    d = 1 << iota  // d == 8
	)
	
	const ( // iota is reset to 0
	    u         = iota * 42  // u == 0     (untyped integer constant)
	    v float64 = iota * 42  // v == 42.0  (float64 constant)
	    w         = iota * 42  // w == 84    (untyped integer constant)
	)
	
	const x = iota  // x == 0  (iota has been reset)
	const y = iota  // y == 0  (iota has been reset)

注意，在同一个声明中出现的多个itoa的value是相同的：

	const (
	    bit0, mask0 = 1 << iota, 1<<iota - 1  // bit0 == 1, mask0 == 0
	    bit1, mask1                           // bit1 == 2, mask1 == 1
	    _, _                                  // skips iota == 2
	    bit3, mask3                           // bit3 == 8, mask3 == 7
	)

## 类型声明

类型用关键字type进行声明。

	TypeDecl     = "type" ( TypeSpec | "(" { TypeSpec ";" } ")" ) .
	TypeSpec     = identifier Type .

类型声明也有几种样式:

	type IntArray [16]int
	
	type (
	    Point struct{ x, y float64 }
	    Polar Point
	)
	
	type TreeNode struct {
	    left, right *TreeNode
	    value *Comparable
	}
	
	type Block interface {
	    BlockSize() int
	    Encrypt(src, dst []byte)
	    Decrypt(src, dst []byte)
	}

需要注意，为一个类型声明了另一个标记符之后，这个标记符对应的类型不会得到被声明的类型的方法。

	// A Mutex is a data type with two methods, Lock and Unlock.
	type Mutex struct         { /* Mutex fields */ }
	func (m *Mutex) Lock()    { /* Lock implementation */ }
	func (m *Mutex) Unlock()  { /* Unlock implementation */ }
	
	// NewMutex has the same composition as Mutex but its method set is empty.
	type NewMutex Mutex
	
	// The method set of the base type of PtrMutex remains unchanged,
	// but the method set of PtrMutex is empty.
	type PtrMutex *Mutex

## 变量的声明(Variable declarations)

	VarDecl     = "var" ( VarSpec | "(" { VarSpec ";" } ")" ) .
	VarSpec     = IdentifierList ( Type [ "=" ExpressionList ] | "=" ExpressionList ) .

变量有以下几种声明格式:

	var i int
	var U, V, W float64
	var k = 0
	var x, y float32 = -1, -2
	var (
	    i       int
	    u, v, s = 2.0, 3.0, "bar"
	)
	var re, im = complexSqrt(-1)
	var _, found = entries[name]  // map lookup; only interested in "found"

如果声明时没有赋值，变量的值为对应的类型的零值(zero value)。

如果声明时没有指定类型，变量的类型根据赋的值推导出来：

	var d = math.Sin(0.5)  // d is float64
	var i = 42             // i is int
	var t, ok = x.(T)      // t is T, ok is bool

特别注意，如果没有指定类型，赋值时不能使用nil：

	var n = nil            // illegal

## 变量的简单声明(Short variable declarations)

变量还可以使用简短的方式声明。

	ShortVarDecl = IdentifierList ":=" ExpressionList .

例如：

	i, j := 0, 10
	f := func() int { return 7 }
	ch := make(chan int)
	r, w := os.Pipe(fd)  // os.Pipe() returns two values
	_, y, _ := coord(p)  // coord() returns three values; only interested in y coordinate

使用简短方式时必须注意，":="右边必须有新的标记符：

	field1, offset := nextField(str, 0)
	field2, offset := nextField(str, offset)  // redeclares offset
	a, a := 1, 2                              // illegal: double declaration of a or no new variable 
	                                          // if a was declared elsewhere

简单方式比较适合在"if"、"for"、"switch"语句声明只会在本区块中使用的变量。

## 函数的声明(Function declarations)

函数使用关键字`func`声明。

	FunctionDecl = "func" FunctionName ( Function | Signature ) .
	FunctionName = identifier .
	Function     = Signature FunctionBody .
	FunctionBody = Block .

如果函数类型中有返回值，函数声明中必须在每个路径的最后进行return。

	func IndexRune(s string, r rune) int {
	    for i, c := range s {
	        if c == r {
	            return i
	        }
	    }
	    // invalid: missing return statement
	}

可以声明一个不是用go实现的函数，在声明中省略函数体即可。

	func flushICache(begin, end uintptr)  // implemented externally

## 方法的声明(Method declarations)

方法也用关键字`func`声明，但是格式不同，比函数声明多了一个Receiver。

	MethodDecl   = "func" Receiver MethodName ( Function | Signature ) .
	Receiver     = Parameters .

Receiver的类型是`T`或者`*T`，T的类型不能是指针和接口，并且必须是在同一个包中定义的。

Receiver可以设置标记符，标记符在方法的区块中有效，且不能与方法中的其它标记符重名。

	func (p *Point) Length() float64 {
	   return math.Sqrt(p.x * p.x + p.y * p.y)
	}
	
	func (p *Point) Scale(factor float64) {
	   p.x *= factor
	   p.y *= factor
	}

方法的类型是函数，例如上面声明的方法Scale，它的类型是：

	func(p *Point, factor float64)
