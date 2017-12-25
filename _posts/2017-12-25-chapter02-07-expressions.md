---
layout: post
title: 07-表达式语法
author: lijiaocn
createdate: 2017/12/20 16:38:41
changedate: 2017/12/25 12:44:22
categories: chapter02
tags:
keywords:
description: 

---

* auto-gen TOC:
{:toc}

## 表达式(Expressions) 与运算符(Operator)

## 表达式

表达式是用运算符和函数的描述的一个计算过程。

### 常量表达式(Constant expressions)

常量表达式在编译时执行，常量表达式中只能使用常量。

使用常量表达式时，需要特别注意未明确声明类型的常量的类型。

	const a = 2 + 3.0             // a == 5.0   (untyped floating-point constant)
	const b = 15 / 4              // b == 3     (untyped integer constant)
	const c = 15 / 4.0            // c == 3.75  (untyped floating-point constant)
	const Θ float64 = 3/2         // Θ == 1.0   (type float64, 3/2 is integer division)
	const Π float64 = 3/2.        // Π == 1.5   (type float64, 3/2. is float division)
	const d = 1 << 3.0            // d == 8     (untyped integer constant)
	const e = 1.0 << 3            // e == 8     (untyped integer constant)
	const f = int32(1) << 33      // illegal    (constant 8589934592 overflows int32)
	const g = float64(2) >> 1     // illegal    (float64(2) is a typed floating-point constant)
	const h = "foo" > "bar"       // h == true  (untyped boolean constant)
	const j = true                // j == true  (untyped boolean constant)
	const k = 'w' + 1             // k == 'x'   (untyped rune constant)
	const l = "hi"                // l == "hi"  (untyped string constant)
	const m = string(k)           // m == "x"   (type string)
	const Σ = 1 - 0.707i          //            (untyped complex constant)
	const Δ = Σ + 2.0e-4          //            (untyped complex constant)
	const Φ = iota*1i - 1/1i      //            (untyped complex constant)
	
complex是内置的函数，返回常量：

	const ic = complex(0, c)      // ic == 3.75i  (untyped complex constant)
	const iΘ = complex(0, Θ)      // iΘ == 1i     (type complex128)
	
如果常量的值超过了能够表达的范围，这个常量可以作为中间值使用：

	const Huge = 1 << 100         // Huge == 1267650600228229401496703205376  (untyped integer constant)
	const Four int8 = Huge >> 98  // Four == 4                                (type int8)

除数不能为0：

	const n = 3.14 / 0.0          // illegal: division by zero

不可以将常量转换为不匹配的类型：

	uint(-1)     // -1 cannot be represented as a uint
	int(3.14)    // 3.14 cannot be represented as an int
	int64(Huge)  // 1267650600228229401496703205376 cannot be represented as an int64
	Four * 300   // operand 300 cannot be represented as an int8 (type of Four)
	Four * 100   // product 400 cannot be represented as an int8 (type of Four)

### 选择表达式(Selector)

选择表达式的格式如下:

	x.f

其中f是选择器(selector)，类型为f，它不能是空白标记符`_`。

如果x是包名，那么选择的是包中的标记符。

f可以是x的成员、方法、匿名成员、匿名成员的方法，到达f时经过的选择次数是f的深度(depth)。

如果f是x的直接成员，深度为0，f是x的直接匿名成员的成员，深度为f在匿名成员中的深度+1。

选择表达式遵循下面的规则：

	x的类型为T或者*T，并且T不是指针和接口，x.f是T中深度最小的名为f的成员。
	x的类型为T，T是接口, x.f是x的动态类型的名为f的方法。
	如果x是指针，x.f是(*x).f的简写，两者等同

如果按照上面的规则，找不到f，编译或运行时报错。

对于下面的代码：

	type T0 struct {
		x int
	}
	
	func (*T0) M0()
	
	type T1 struct {
		y int
	}
	
	func (T1) M1()
	
	type T2 struct {
		z int
		T1
		*T0
	}
	
	func (*T2) M2()
	
	type Q *T2
	
	var t T2     // with t.T0 != nil
	var p *T2    // with p != nil and (*p).T0 != nil
	var q Q = p

可以有这么些选择方法：

	t.z          // t.z
	t.y          // t.T1.y
	t.x          // (*t.T0).x
	
	p.z          // (*p).z
	p.y          // (*p).T1.y
	p.x          // (*(*p).T0).x
	
	q.x          // (*(*q).T0).x        (*q).x is a valid field selector
	
	p.M0()       // ((*p).T0).M0()      M0 expects *T0 receiver
	p.M1()       // ((*p).T1).M1()      M1 expects T1 receiver
	p.M2()       // p.M2()              M2 expects *T2 receiver
	t.M2()       // (&t).M2()           M2 expects *T2 receiver, see section on Calls

注意q没有选择`M0()`，因为M0()的Reciver类型是`*T1`，类型Q中不能继承T1的方法。

### 方法表达式(Method expressions)

方法(method)表达式就是方法的实现语句。

	MethodExpr    = ReceiverType "." MethodName .
	ReceiverType  = TypeName | "(" "*" TypeName ")" | "(" ReceiverType ")" .

与函数的不同的是，方法是有接收者(Receiver)的，如下：

	type T struct {
	        a int
	}
	func (tv  T) Mv(a int) int         { return 0 }  // value receiver
	func (tp *T) Mp(f float32) float32 { return 1 }  // pointer receiver
	
	var t T

方法是属于类型的，类型的方法和类型的指针的方法是不同的。

类型的方法是一个将接收者作为参数传入的函数，例如在上面例子中:

	T.Mv 的类型为 func(tv T, a int) int
	T.Mp 的类型为 func(tv *T, a int) int

类型的方法可以直接通过类型名调用：

	T.Mv(t, 7)             //注意要传入接收者
	(T).Mv(t, 7)
	(*T).Mp(&t, 7)         //注意传入的是接收者是指针

类型不能调用类型指针的方法，类型指针可以调用类型的方法：

	T.Mp(&t,7)       //Mp是(*T)的方法，不允许T调用
	(*T).Mv(t,7)     //Mv是T的方法，*T可以调用

也可以把方法赋值给变量，然后通过变量调用:

	f1 := T.Mv; f1(t, 7)         //要传入接受者t
	f2 := (T).Mv; f2(t, 7)       //要传入接受者t
	
	f3 := T.Mp; f3(&t, 7)         //要传入接受者&t
	f4 := (T).Mp; f4(&t, 7)       //要传入接受者&t

也可以通过该类型的变量调用，这时候不需要传入接收者。

	t.Mv(7)
	t.Mp(7)

因为变量的方法和类型的方法是不同的，所以不需要传入接收者。

	t.Mv 的类型为 func(a int) int
	t.Mp 的类型为 func(a int) int

无论一个变量(t)是不是指针(类型为`*T`的变量），它都既可以调用类型(T)的方法，也可以调用类型指针(`*T`)的方法。go语言自身代为完成了取址和取值操作。

变量的方法也可以存放单独的变量中，然后通过变量调用：

	f := t.Mv; f(7)   // like t.Mv(7)
	f := pt.Mp; f(7)  // like pt.Mp(7)
	f := pt.Mv; f(7)  // like (*pt).Mv(7)
	f := t.Mp; f(7)   // like (&t).Mp(7)
	f := makeT().Mp   // invalid: result of makeT() is not addressable

变量的类型为接口时，用同样的方式调用方法：

	var i interface { M(int) } = myVal
	f := i.M; f(7)  // like i.M(7)

### 索引表达式(Index expressions)

索引表达式格式如下：

	a[x]

a的类型不同，表达式的运行结果不同。

	如果a不是字典，x的必须是整数，并且0<= x <len(a)
	如果a是数组，返回数组中x位置处的成员，如果x超出数组范围，程序panic
	如果a是指向数组的指针，a[x]等同于(*a)[x]
	如果a是分片(Slice)， a[x]返回x位置处的数组成员，如果x超出范围，程序panic
	如果a是字符串，返回x位置处的字符，如果x超出范围，程序panic，且a[x]不能被赋值
	如果a是字典(map)，x的类型必须是字典的key的类型，返回字典中x对应的值，和表示对应成员是否存在的布尔类型的值(bool)
	如果a是字典(map)，且a的值是nil，a[x]返回字典中成员类型的零值

### 分片表达式(Slice expressions)

分片表达式适用于字符串、数组、指向数组的指针和分片。

	a[low : high]

返回一个从零开始，长度为high-low的分片。

	a := [5]int{1, 2, 3, 4, 5}
	s := a[1:4]

得到分片s的情况如下：

	s[0] == 2
	s[1] == 3
	s[2] == 4

分片表达式中low和high省略：

	a[2:]  // same as a[2 : len(a)]
	a[:3]  // same as a[0 : 3]
	a[:]   // same as a[0 : len(a)]

如果a是指向数组的指针，a[low:high]等同于`(*a)[low:high]`。

如果a是字符串、数组、指向数组的指针，low和high的取值范围为：

	0 <= low <= high <= len(a)

如果a是分片，low和high的取值范围为：

	0 <= low <= high <= cap(a)

low和high超出范围时，引发panic。

如果a是已经声明字符串、分片，返回值也是字符串、分片。

如果a是未声明的字符串，返回一个类型为字符串的变量.

如果a是数组，返回指向这个数组的分片。

### 完整分片表达式(Full slice expressions)

完整的分片表达式还带有一个max，限定返回的分片的容量为(capacity)为`max-low`。

	a[low : high : max]

在完整的分片表达式中，只有low可以省略，默认为0。

如果a是字符串、数组、指向数组的指针，low和high的取值范围为：

	0<= low <= high <= max <= len(a)

如果a是分片，low、high和max的取值范围为：

	0<= low <= high <= max <= cap(a)

如果超出范围，引发panic。

### 类型断言表达式(Type assertions expressions)

断言表达式用来判断x是否不为nil，且它的类型是否与T匹配。

	x.(T)

如果T不是接口类型，x的类型必须是接口，判断T是否可以成为x的动态类型。

如果T是接口类型，判断x是否实现了接口T。

如果T不是接口类型，x的类型也不是接口，引发panic。

如果断言成立，表达式的值就是类型为T的x，和布尔值true；如果断言不成立，表达式的值是类型T的零值，和布尔值false。

### 调用表达式(Call expressions)

调用表达式适用于函数和方法：

	f(a1, a2, … an)

针对方法使用时，需要带有receiver:

	math.Atan2(x, y)  // function call
	var pt *Point
	pt.Scale(3.5)     // method call with receiver pt

传入值按值、按顺序传递给函数或方法的参数，返回值也是按值传递的。

如果一个函数的返回值，满足另一个参数的传入参数要求，可以写成`f(g(parameters_of_g))`，例如：

	func Split(s string, pos int) (string, string) {
	        return s[0:pos], s[pos:]
	}
	
	func Join(s, t string) string {
	        return s + t
	}
	
	if Join(Split(value, len(value)/2)) != value {
	        log.Panic("test fails")
	}

调用表达式支持变长参数，变长参数必须是最后一个，且类型前是`...`。

例如在下面的函数中：

	func Greeting(prefix string, who ...string)

如果以这种方式调用，参数who的值是nil：

	Greeting("nobody")

如果以这种方式调用，参数who的值的类型是[]string：

	Greeting("hello:", "Joe", "Anna", "Eileen")

如果以这种方式调用，参数who等于s：

	s:= []string{"James", "Jasmine"}
	Greeting("goodbye:", s...)

## 运算符(Operator)

运算符用于构成表达式。

	Expression = UnaryExpr | Expression binary_op Expression .
	UnaryExpr  = PrimaryExpr | unary_op UnaryExpr .

	binary_op  = "||" | "&&" | rel_op | add_op | mul_op .
	rel_op     = "==" | "!=" | "<" | "<=" | ">" | ">=" .
	add_op     = "+" | "-" | "|" | "^" .
	mul_op     = "*" | "/" | "%" | "<<" | ">>" | "&" | "&^" .

	unary_op   = "+" | "-" | "!" | "^" | "*" | "&" | "<-" .

运算符都是go语言内置的。

	Precedence    Operator
	    5             *  /  %  <<  >>  &  &^
	    4             +  -  |  ^
	    3             ==  !=  <  <=  >  >=
	    2             &&
	    1             ||

优先级相同的二元运算符按照先左后右的顺序结合：

	x / y * z

等同于：

	(x / y) * z

### 算数运算符(Arithmetic operators)

	+    sum                    integers, floats, complex values, strings
	-    difference             integers, floats, complex values
	*    product                integers, floats, complex values
	/    quotient               integers, floats, complex values
	%    remainder              integers
	
	&    bitwise AND            integers
	|    bitwise OR             integers
	^    bitwise XOR            integers
	&^   bit clear (AND NOT)    integers
	
	<<   left shift             integer << unsigned integer
	>>   right shift            integer >> unsigned integer

### 字符串拼接(String concatenation)

字符串可以用运算符"+"进行拼接：

	:= "hi" + string(c)
	s += " and good bye"

### 比较运算符(Comparison operators)

	==    equal
	!=    not equal
	<     less
	<=    less or equal
	>     greater
	>=    greater or equal

### 逻辑运算符(Logical operators)

	&&    conditional AND    p && q  is  "if p then q else false"
	||    conditional OR     p || q  is  "if p then true else q"
	!     NOT                !p      is  "not p"

### 地址运算符(Address operators)

	&     
	*  

### 读取运算符(Receive operator)

	v1 := <-ch
	v2 = <-ch
	f(<-ch)
	<-strobe  // wait until clock pulse and discard received value

### 类型转换(Conversions)

	Conversion = Type "(" Expression [ "," ] ")" .
