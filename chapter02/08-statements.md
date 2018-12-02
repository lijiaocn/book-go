---
layout: default
title: 08-statements
author: lijiaocn
createdate: 2017/12/21 23:11:24
changedate: 2017/12/22 00:11:36
categories:
tags:
keywords:
description: 

---


<!-- toc -->

# go的状态语句 

	Statement =
	    Declaration | LabeledStmt | SimpleStmt |
	    GoStmt | ReturnStmt | BreakStmt | ContinueStmt | GotoStmt |
	    FallthroughStmt | Block | IfStmt | SwitchStmt | SelectStmt | ForStmt |
	    DeferStmt .
	
	SimpleStmt = EmptyStmt | ExpressionStmt | SendStmt | IncDecStmt | Assignment | ShortVarDecl .

## 终止语句(Terminating statements)

终止语句是指下的情况：

	return
	goto
	调用内置函数panic(interface{})
	if语句以及else语句中语句的结束
	for语句语句的结束
	switch语句语句
	select
	labeled 

## 空白语句(Empty statements)

空白语句不做任何事情：

	EmptyStmt = .

## 标记语句(Labeled statements)

标记语句可以是goto、break、continue的目标。

	LabeledStmt = Label ":" Statement .
	Label       = identifier .

## 表达式语句(Expression statements)

除了下面的内置函数，其它的函数、方法和接收操作符都可以用于表达式语句。

	append cap complex imag len make new real
	unsafe.Alignof unsafe.Offsetof unsafe.Sizeof

## 发送语句(Send statements)

发送语句是专用于向通道(channel)发送数据的。

	SendStmt = Channel "<-" Expression .
	Channel  = Expression .

## 递增递减语句(IncDec statements)

	IncDecStmt = Expression ( "++" | "--" ) .

## 赋值语句(Assignments)

	Assignment = ExpressionList assign_op ExpressionList .
	assign_op = [ add_op | mul_op ] "=" .

## if语句(If statements)

	IfStmt = "if" [ SimpleStmt ";" ] Expression Block [ "else" ( IfStmt | Block ) ] .

## switch语句(Switch statements)

switch语句分为以表达式为依据，和以类型为依据两种形式。

	SwitchStmt = ExprSwitchStmt | TypeSwitchStmt .

使用表达式作为分支依据：
	
	ExprSwitchStmt = "switch" [ SimpleStmt ";" ] [ Expression ] "{" { ExprCaseClause } "}" .
	ExprCaseClause = ExprSwitchCase ":" StatementList .
	ExprSwitchCase = "case" ExpressionList | "default" .

例如:

	switch tag {
	default: s3()
	case 0, 1, 2, 3: s1()
	case 4, 5, 6, 7: s2()
	}
	
	switch x := f(); {  // missing switch expression means "true"
	case x < 0: return -x
	default: return x
	}
	
	switch {
	case x < y: f1()
	case x < z: f2()
	case x == 4: f3()
	}

使用类型为依据：

	TypeSwitchStmt  = "switch" [ SimpleStmt ";" ] TypeSwitchGuard "{" { TypeCaseClause } "}" .
	TypeSwitchGuard = [ identifier ":=" ] PrimaryExpr "." "(" "type" ")" .
	TypeCaseClause  = TypeSwitchCase ":" StatementList .
	TypeSwitchCase  = "case" TypeList | "default" .
	TypeList        = Type { "," Type } .

例如：

	switch i := x.(type) {
	case nil:
	    printString("x is nil")                // type of i is type of x (interface{})
	case int:
	    printInt(i)                            // type of i is int
	case float64:
	    printFloat64(i)                        // type of i is float64
	case func(int) float64:
	    printFunction(i)                       // type of i is func(int) float64
	case bool, string:
	    printString("type is bool or string")  // type of i is type of x (interface{})
	default:
	    printString("don't know the type")     // type of i is type of x (interface{})
	}

## for语句(For statements)

for的语句循环条件有三种。

	ForStmt = "for" [ Condition | ForClause | RangeClause ] Block .
	Condition = Expression .

简略条件判断：

	for a < b {
	    a *= 2
	}

完整条件判断：

	ForClause = [ InitStmt ] ";" [ Condition ] ";" [ PostStmt ] .
	InitStmt = SimpleStmt .
	PostStmt = SimpleStmt .

例如:

	for i := 0; i < 10; i++ {
	    f(i)
	}
	
	for cond { S() }    is the same as    for ; cond ; { S() }
	for      { S() }    is the same as    for true     { S() }

range判断：

	RangeClause = [ ExpressionList "=" | IdentifierList ":=" ] "range" Expression .

需要特别注意的是Expression不同时，`range Expression`的返回值不同。

	Range expression                          1st value          2nd value
	
	array or slice  a  [n]E, *[n]E, or []E    index    i  int    a[i]       E
	string          s  string type            index    i  int    see below  rune
	map             m  map[K]V                key      k  K      m[k]       V
	channel         c  chan E, <-chan E       element  e  E

## go语句(Go statements)

	Stmt = "go" Expression .

## select语句(Select statements)

select语句用于执行当前可以执行的语句。

	SelectStmt = "select" "{" { CommClause } "}" .
	CommClause = CommCase ":" StatementList .
	CommCase   = "case" ( SendStmt | RecvStmt ) | "default" .
	RecvStmt   = [ ExpressionList "=" | IdentifierList ":=" ] RecvExpr .
	RecvExpr   = Expression .

如果有多个语句当前都可以执行，需要特别注意这些语句的执行顺序。

	1. 通道(channel)相关的语句如果同时进入可执行状态，只执行在源码中位置靠前的语句
	2. 如果多个语句可以执行，随机选择一个执行。
	3. 如果所有语句都不能执行，那么执行default语句，如果没有default语句，进入等待状态

例如:

	var a []int
	var c, c1, c2, c3, c4 chan int
	var i1, i2 int
	select {
	case i1 = <-c1:
	    print("received ", i1, " from c1\n")
	case c2 <- i2:
	    print("sent ", i2, " to c2\n")
	case i3, ok := (<-c3):  // same as: i3, ok := <-c3
	    if ok {
	        print("received ", i3, " from c3\n")
	    } else {
	        print("c3 is closed\n")
	    }
	case a[f()] = <-c4:
	    // same as:
	    // case t := <-c4
	    //    a[f()] = t
	default:
	    print("no communication\n")
	}

	for {  // send random sequence of bits to c
	    select {
	    case c <- 0:  // note: no statement, no fallthrough, no folding of cases
	    case c <- 1:
	    }
	}

	select {}  // block forever

## 返回语句(Return statements)

	ReturnStmt = "return" [ ExpressionList ] .

例如:

	func simpleF() int {
	    return 2
	}

支持多值返回:

	func complexF1() (re float64, im float64) {
	    return -7.0, -4.0
	}

可以直接将表达式的结果返回：

	func complexF2() (re float64, im float64) {
	    return complexF1()
	}

还可以命名返回：

	func complexF3() (re float64, im float64) {
	    re = 7.0
	    im = 4.0
	    return
	}

	func (devnull) Write(p []byte) (n int, _ error) {
	    n = len(p)
	    return
	}

命名的返回的时候，不同有同名的其它变量：

	func f(n int) (res int, err error) {
	    if _, err := f(n-1); err != nil {
	        return  // invalid return statement: err is shadowed
	    }
	    return
	}

## break语句(Break statement)

	BreakStmt = "break" [ Label ] .

例如:

	OuterLoop:
	    for i = 0; i < n; i++ {
	        for j = 0; j < m; j++ {
	            switch a[i][j] {
	            case nil:
	                state = Error
	                break OuterLoop
	            case item:
	                state = Found
	                break OuterLoop
	            }
	        }
	    }

## continue语句(Continue statements)

	ContinueStmt = "continue" [ Label ] .

例如:

	RowLoop:
		for y, row := range rows {
			for x, data := range row {
				if data == endOfRow {
					continue RowLoop
				}
				row[x] = data + bias(x, y)
			}
		}

## goto语句(Goto statements)

	GotoStmt = "goto" Label .

使用goto的时候要特别注意，不要在goto与Label直接存在变量的声明。

例如下面的做法符合语法要求，但是容易造成混乱，在`L:`之后的位置使用i，会报错：

		goto L  // BAD
		v := 3
	L:

goto只能跳转到所在区块中的标记位置。

例如下面的做法是不符合语法的，L1是另一个区块中的标记。

	if n%2 == 1 {
		goto L1
	}
	for n > 0 {
		f()
		n--
	L1:
		f()
		n--
	}

## fallthrough语句(Fallthrough statements)

fallthrouch用于switch语句中，表示紧邻的下一个语句需要被执行。

	FallthroughStmt = "fallthrough" .

## defer语句(Defer statements)

defer表示跟随的语句需要在函数执行结束的时候执行。

	DeferStmt = "defer" Expression .

例如：

	lock(l)
	defer unlock(l)  // unlocking happens before surrounding function returns

	// prints 3 2 1 0 before surrounding function returns
	for i := 0; i <= 3; i++ {
		defer fmt.Print(i)
	}

	// f returns 1
	func f() (result int) {
		defer func() {
			result++
		}()
		return 0
	}
