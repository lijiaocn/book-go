//create: 2017/10/26 22:04:34 change: 2017/12/16 23:40:30 lijiaocn@foxmail.com
package main

func main() {

	str1 := "日本語"                                  // UTF-8 input text
	str2 := `日本語`                                  // UTF-8 input text as a raw literal
	str3 := "\u65e5\u672c\u8a9e"                   // the explicit Unicode code points
	str4 := "\U000065e5\U0000672c\U00008a9e"       // the explicit Unicode code points
	str5 := "\xe6\x97\xa5\xe6\x9c\xac\xe8\xaa\x9e" // the explicit UTF-8 bytes

	println(str1)
	println(str2)
	println(str3)
	println(str4)
	println(str5)
}
