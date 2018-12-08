package main

import (
	"fmt"
)

func main() {
	s := "abc"
	fmt.Println(s[0] == '\x61', s[1] == 'b', s[2] == 0x63) //使用索引号访问字符 (byte)

	s1 := `a
b\r\n\x00
c`
	fmt.Println(s1) //使用 "`" 定义不做转义处理的原始字符串，支持跨行

	a := "hello" +
		",world!"
	fmt.Println(a)
	//	b := "hello"
	//	+",world!"
	//	fmt.Println(b) //连接跨行字符串时，"+" 必须在上一行末尾，否则导致编译错误

	c := "hello, world!"
	c1 := c[:5]
	c2 := c[7:]
	c3 := c[1:5]
	fmt.Println(c1, c2, c3) //支持用两个索引号返回子串。子串依然指向原字节数组，仅修改了指针和长度属性

	fmt.Printf("%T\n", 'a')

	//要修改字符串，可先将其转换成 []rune 或 []byte，完成后再转换为 string。无论哪种转换，都会重新分配内存，并复制字节数组
	d := "abcd"
	bs := []byte(d)
	bs[1] = 'B'
	fmt.Println(string(bs))

	e := "电脑"
	us := []rune(e)
	us[1] = '话'
	fmt.Println(string(us))

	//用for循环遍历字符串时,也有byte和rune两种方式
	f := "abc汉字"
	for i := 0; i < len(f); i++ {
		fmt.Printf("%c,", f[i])
	}
	fmt.Println()
	for _, g := range f {
		fmt.Printf("%c,", g)
	}

}
