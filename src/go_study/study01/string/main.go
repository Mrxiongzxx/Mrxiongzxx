package main

import "fmt"

//字符串是不可变值类型，内部⽤指针指向 UTF-8 字节数组。
//• 默认值是空字符串 ""。
//• ⽤索引号访问某字节，如 s[i]。
//• 不能⽤序号获取字节元素指针，&s[i] ⾮法。
//• 不可变类型，⽆法修改字节数组。
//• 字节数组尾部不包含 NULL。

func main() {
	//使⽤索引号访问字符 (byte)。
	s := "abc"
	println(s[0] == '\x61', s[1] == 'b', s[2] == 0x63)

	//使⽤ "`" 定义不做转义处理的原始字符串，⽀持跨⾏。
	str := `a
		b\r\n\x00
		c`
	println(str)

	//连接跨⾏字符串时，"+" 必须在上⼀⾏末尾，否则导致编译错误。
	str1 := "Hello, " +
		"World!"
	fmt.Println(str1)
	//str2 := "Hello, "
	//+ "World!" // Error: invalid operation: + untyped string

	//⽀持⽤两个索引号返回⼦串。⼦串依然指向原字节数组，仅修改了指针和⻓度属性。
	s0 := "Hello, World!"
	s1 := s0[:5] // Hello  前五个
	s2 := s0[7:] // World!  从第七位开始
	s3 := s0[1:5] // ello 从第一个开始算
	fmt.Println(s0,s1,s2,s3)

	//单引号字符常量表⽰ Unicode Code Point，⽀持 \uFFFF、\U7FFFFFFF、\xFF 格式。
	//对应 rune 类型，UCS-4。 rune 是int32别名
	fmt.Printf("%T\n", 'a')
	var c1, c2 rune = '\u6211', '们'
	println(c1 == '我', string(c2) == "\xe4\xbb\xac")


	//要修改字符串，可先将其转换成 []rune 或 []byte，完成后再转换为 string。⽆论哪种转
	//换，都会重新分配内存，并复制字节数组。

	s4 := "abcd"
	bs := []byte(s4)
	bs[1] = 'B'
	println(string(bs))
	u := "电脑"
	us := []rune(u)
	us[1] = '话'
	println(string(us))

	str_print()
}

//循环输出string
//⽤ for 循环遍历字符串时，也有 byte 和 rune 两种⽅式
func str_print() {
	s := "abc汉字"
	for i := 0; i < len(s); i++ { // byte
		fmt.Printf("%c,", s[i])
	}
	fmt.Println()
	for _, r := range s { // rune
		fmt.Printf("%c,", r)
	}
}