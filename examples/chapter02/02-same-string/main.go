//create: 2017/10/26 22:04:34 change: 2017/12/16 23:24:48 lijiaocn@foxmail.com
package main

func main() {

	raw1 := "日本語"                                  // UTF-8 input text
	raw2 := `日本語`                                  // UTF-8 input text as a raw literal
	raw3 := "\u65e5\u672c\u8a9e"                   // the explicit Unicode code points
	raw4 := "\U000065e5\U0000672c\U00008a9e"       // the explicit Unicode code points
	raw5 := "\xe6\x97\xa5\xe6\x9c\xac\xe8\xaa\x9e" // the explicit UTF-8 bytes

	println(raw1)
	println(raw2)
	println(raw3)
	println(raw4)
	println(raw5)
}
