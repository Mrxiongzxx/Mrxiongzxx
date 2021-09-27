package main


//不要期望编译器能理解你的想法，在初始化语句中计算出全部结果是个好主意。
func length(s string) int {
	println("call length.")
	return len(s)
}

func main()  {
	//⽀持三种循环⽅式，包括类 while 语法。
	s := "abc"
	for i, n := 0, len(s); i < n; i++ { // 常⻅的 for 循环，⽀持初始化语句。
		println(s[i])
	}
	n := len(s)
	for n > 0 { // 替代 while (n > 0) {}
		println(s[n]) // 替代 for (; n > 0;) {}
		n--
	}
	for { // 替代 while (true) {}
		println(s) // 替代 for (;;) {}
	}

	str := "abcd"
	for i, n := 0, length(str); i < n; i++ { // 避免多次调⽤ length 函数。
		println(i, str[i])
	}
}

