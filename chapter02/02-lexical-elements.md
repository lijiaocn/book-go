---
layout: default
title: 02-lexical-elements
author: lijiaocn
createdate: 2017/12/16 17:58:54
changedate: 2017/12/16 23:22:48
categories:
tags:
keywords:
description: 

---

* auto-gen TOC:
{:toc}

# go语言的词法规则

英语有单词，汉语有词语，编程语言也有自己的词汇。

很多工作多年的程序员都没有从语言的层面了解自己正在使用的程序语言。这不应当被批评谴责，`程序语言设计`是一个大牛云集的小众领域，而程序员是使用程序语言的芸芸大众。

绝大多数程序员的任务是使用程序语言开发出软件，而不是设计程序语言，正如厨师的任务是做出美味佳肴，而不是制作锅碗瓢盆和生产原材料。

但，了解一下还是有好处的，能够从更高的层面、更抽象的看问题，这很重要。

go语言使用的编码是utf-8，用go语言写的源代码是utf-8编码的文本文件。

## 构成要素：字符(Characters)、字母(Letters)、数字(Digits)

英语单词的构成要素是字母，汉语词语的构成要素是字，go语言词汇的构成要素是字符。

### 字符(Characters)

go语言用utf-8编码，它完全是由下面的字符(Characters)组成的。

	newline        = /* the Unicode code point U+000A */ .
	unicode_char   = /* an arbitrary Unicode code point except newline */ .
	unicode_letter = /* a Unicode code point classified as "Letter" */ .
	unicode_digit  = /* a Unicode code point classified as "Number, decimal digit" */ .

### 字母(Letters)、数字(Digits)

go语言将部分utf-8字符称为字母(letter)与数字(digit)。

	letter        = unicode_letter | "_" .
	decimal_digit = "0" … "9" .
	octal_digit   = "0" … "7" .
	hex_digit     = "0" … "9" | "A" … "F" | "a" … "f" .

## 注释(Comments)

注释有`单行注释`(line comments)和`通用注释`(general comments)两种形式。

单行注释以`//`开始，到行尾结束。

通用注释以`/*`开始，到`*/`结束。

注释不能位于字符(rune)、字符串(string literal)和另一个注释当中。

## 词汇(Tokens)

Tokens是一个统称，它由标识符、关键字、运算符、分隔符、整数、浮点数、虚数、字符、字符串组成。

简而言之，在一个go文件中，去掉空格(spaces, U+0020)、TAB(horizontal tabs, U+0009)、回车(carriage returns, U+000D)、换行(newlines, U+000A)、分号(Semicolons)之后，剩下肉眼可见的部分就是Tokens。

Tokens是编译原理中一个常用的术语。编译器在进行词法分析的时候，会连续的读取源码文件中的内容，它从第一个非空白的符号开始记录，遇到下一个空白的符号后，将已经记录的内容作为一个Token。

## 分号(Semicolons)

很多编程语言都用“;”作为结束符号，标记一行代码的结束。go语言也用分号做结束符，但是在源码中可以不写出分号，go可以自主推断出是否结束。

当一行代码的最后一个Token是下面的类型时，go会自动在行尾补上分号：

	标识符(identifier)
	整数、浮点数、虚数、字(rune)、字符串
	关键字：break、continue、fallthrough、return
	运算符和分隔符：++、--、)、]、}

## 标识符(Identifiers)

标识符是用来索引程序中的实体(entities)的。

譬如说变量、常量、函数等，这些都是程序中的entity，它们需要有一个名字，这个名字就是它们的标识符。

当我们说变量A的时候，其实是在指标识符A关联的数值。

go的标识符语法格式如下：

	identifier = letter { letter | unicode_digit } .

即，由字母和数字组成，但必须以字母开头，且不能是关键字。

## 关键字(Keywords)

关键字是go语言保留的一些单词，它们都是由特定功能的，不用用来做标识符。

关键字的数量是有限的，下面是go的全部关键字：

	break        default      func         interface    select
	case         defer        go           map          struct
	chan         else         goto         package      switch
	const        fallthrough  if           range        type
	continue     for          import       return       var

## 运算符和分隔符(Operators and Delimiters)

运算符和分隔符是一类有特殊的意义的非字母符号。

它们的数量也是有限的，下面是go的全部运算符和分隔符：

	+    &     +=    &=     &&    ==    !=    (    )
	-    |     -=    |=     ||    <     <=    [    ]
	*    ^     \*=    ^=     <-    >     >=    {    }
	/    <<    /=    <<=    ++    =     :=    ,    ;
	%    >>    %=    >>=    --    !     ...   .    :
	     &^          &^=

## 整数(Integer literals)

整数就是数学意义上的整数，在go中有十进制、八进制、十六进制三种表示方式。

	int_lit     = decimal_lit | octal_lit | hex_lit .
	decimal_lit = ( "1" … "9" ) { decimal_digit } .
	octal_lit   = "0" { octal_digit } .
	hex_lit     = "0" ( "x" | "X" ) hex_digit { hex_digit } .

注意，没有2进制的表示方式。

注意，`decimal_digit`前面章节中给出的数字，后面再遇到前面已经定义的词法时，不再提示。

在十六进制表示方式中，大写字母与小写字母的含义是相同的。

	42            //十进制
	0600          //八进制，以0开头
	0xBadFace     //十六进制，以0x开头，忽略大小写

## 浮点数(Floating-point literals)

浮点数就是数学上的浮点数，带有小数点的数，go支持用科学计数表示浮点数。

	float_lit = decimals "." [ decimals ] [ exponent ] |
	            decimals exponent |
	            "." decimals [ exponent ] .
	decimals  = decimal_digit { decimal_digit } .
	exponent  = ( "e" | "E" ) [ "+" | "-" ] decimals .

浮点数可以有以下几种样式：

	0.
	72.40
	072.40          //== 72.40
	2.71828
	1.e+0
	6.67428e-11
	1E6
	.25
	.12345E+5

注意在浮点数中，全是十进制，没有八进制和十六进制，`0720.40`等于`720.40`。

## 虚数(Imaginary literals)

虚数是复数的组成部分，在样式上，它就是在整数或者浮点数后面加上“i”。

	imaginary_lit = (decimals | float_lit) "i" .

虚数也只能用十进制表示。

	0i
	011i  // == 11i
	0.i
	2.71828i
	1.e+0i
	6.67428e-11i
	1E6i
	.25i
	.12345E+5i

## 符号(Rune literals)

在go语言中，Rune literal就是一个utf-8字符的值。

go语言使用utf8编码，utf8是一种变长的编码，它使用1～4个字节表示一个符号。

这样的符号用C语言中的char来指示，明显是不合适的，因为char要求空间必须是1个byte。

当要对这样的字符进行处理的时候，需要有一个称呼来指示它，这个称呼就是`Rune`。

Rune，是一个单一意义的符号，它占用的空间是不固定的。

go语言中，rune的语法有一些复杂：

	rune_lit         = "'" ( unicode_value | byte_value ) "'" .
	unicode_value    = unicode_char | little_u_value | big_u_value | escaped_char .
	byte_value       = octal_byte_value | hex_byte_value .
	octal_byte_value = `\` octal_digit octal_digit octal_digit .
	hex_byte_value   = `\` "x" hex_digit hex_digit .
	little_u_value   = `\` "u" hex_digit hex_digit hex_digit hex_digit .
	big_u_value      = `\` "U" hex_digit hex_digit hex_digit hex_digit
	                           hex_digit hex_digit hex_digit hex_digit .
	escaped_char     = `\` ( "a" | "b" | "f" | "n" | "r" | "t" | "v" | `\` | "'" | `"` ) .

首先，它是用单引号包裹的。

单引号中包裹的可以是byte值，也可以是unicode编码值。

byte_value有八进制和十六进制两种表达方式，八进制以`\`开始，后面跟随三个数字，十六进度以`\x`开始，后面跟随两个十六进制数字。

unidecode编码有四种形式。

第一种是单字符，第二种是以`\u`开头后面跟随4个十六进制数字，第三种是以`\U`开头后面跟随8个十六进制数字。

第四种是以`\`开头的转义字符，转义字符的数量是有限的，只有下面这些：

	\a   U+0007 alert or bell
	\b   U+0008 backspace
	\f   U+000C form feed
	\n   U+000A line feed or newline
	\r   U+000D carriage return
	\t   U+0009 horizontal tab
	\v   U+000b vertical tab
	\\   U+005c backslash
	\'   U+0027 single quote  (valid escape only within rune literals)
	\"   U+0022 double quote  (valid escape only within string literals)

注意，用unicode_value表示时，unicode_value必须是有效的，符合utf-8编码规范。

godoc中给出的rune literal：

	'a'
	'ä'
	'本'
	'\t'
	'\000'
	'\007'
	'\377'
	'\x07'
	'\xff'
	'\u12e4'
	'\U00101234'
	'\''         // rune literal containing single quote character
	'aa'         // illegal: too many characters
	'\xa'        // illegal: too few hexadecimal digits
	'\0'         // illegal: too few octal digits
	'\uDFFF'     // illegal: surrogate half
	'\U00110000' // illegal: invalid Unicode code point

## 字符串(String literals)

字符串有两种形式，原始型(raw string literals)和解释型(interpreted string literals)。

	string_lit             = raw_string_lit | interpreted_string_lit .
	raw_string_lit         = "`" { unicode_char | newline } "`" .
	interpreted_string_lit = `"` { unicode_value | byte_value } `"` .

原始型字符串用反引号包裹，反引号中的内容都是字符串的一部分，反斜杠就是反斜杠，还包括看不到换行回车等。

	`\n
	\n`                  // same as "\\n\n\\n"

简而言之，原始型字符串，就是它看起来的样子。

解释型字符串用双引号包裹，可以使用反斜杠进行转义。

	"Hello, world!\n"
	"日本語"
	"\u65e5本\U00008a9e"
	"\xff\u00FF"
	"\uD800"             // illegal: surrogate half
	"\U00110000"         // illegal: invalid Unicode code point

注意与rune类似，unicode_value必须有效的，符合utf-8规范的。

解释型字符串可以用多种形式描述相同的内容，这个特点，有时候是特别有用的。

下面的五个解释型字符串，样式不同，但内容完全一致：

	"日本語"                                 // UTF-8 input text
	`日本語`                                 // UTF-8 input text as a raw literal
	"\u65e5\u672c\u8a9e"                    // the explicit Unicode code points
	"\U000065e5\U0000672c\U00008a9e"        // the explicit Unicode code points
	"\xe6\x97\xa5\xe6\x9c\xac\xe8\xaa\x9e"  // the explicit UTF-8 bytes

## 参考

1. [go Lexical elements][1]

[1]: http://127.0.0.1:6060/ref/spec#Lexical_elements  "go Lexical elements" 
