---
layout: default
title: 07-expressions
author: lijiaocn
createdate: 2017/12/20 16:38:41
changedate: 2017/12/20 20:12:31
categories:
tags:
keywords:
description: 

---

* auto-gen TOC:
{:toc}

# 表达式(Expressions) 与操作符(Operator)

表达式是用操作符和函数的描述的一个计算过程。

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

### Method expressions

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

## 表达式中的操作符(Operator)

操作符用于构成表达式。

	Expression = UnaryExpr | Expression binary_op Expression .
	UnaryExpr  = PrimaryExpr | unary_op UnaryExpr .

	binary_op  = "||" | "&&" | rel_op | add_op | mul_op .
	rel_op     = "==" | "!=" | "<" | "<=" | ">" | ">=" .
	add_op     = "+" | "-" | "|" | "^" .
	mul_op     = "*" | "/" | "%" | "<<" | ">>" | "&" | "&^" .

	unary_op   = "+" | "-" | "!" | "^" | "*" | "&" | "<-" .

操作符都是go语言内置的。

	Precedence    Operator
	    5             *  /  %  <<  >>  &  &^
	    4             +  -  |  ^
	    3             ==  !=  <  <=  >  >=
	    2             &&
	    1             ||

优先级相同的二元操作符按照先左后右的顺序结合：

	x / y * z

等同于：

	(x / y) * z

### Arithmetic operators

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

字符串可以用操作符"+"进行拼接：

	:= "hi" + string(c)
	s += " and good bye"

### Comparison operators

	==    equal
	!=    not equal
	<     less
	<=    less or equal
	>     greater
	>=    greater or equal

### Logical operators

	&&    conditional AND    p && q  is  "if p then q else false"
	||    conditional OR     p || q  is  "if p then true else q"
	!     NOT                !p      is  "not p"

### Address operators

	&     
	*  


### Receive operator

	v1 := <-ch
	v2 = <-ch
	f(<-ch)
	<-strobe  // wait until clock pulse and discard received value

### Conversions

