package main

func main() {
	//很特别的写法：
	//• 可省略条件表达式括号。
	//• ⽀持初始化语句，可定义代码块局部变量。
	//• 代码块左⼤括号必须在条件表达式尾部。
	x := 0
	// if x > 10 // Error: missing condition in if statement
	// {
	// }
	if n := "abc"; x > 0 { // 初始化语句未必就是定义变量，⽐如 println("init") 也是可以的。
		println(n[2])
	} else if x < 0 { // 注意 else if 和 else 左⼤括号位置。
		println(n[1])
	} else {
		println(n[0])
	}
	//不⽀持三元操作符 "a > b ? a : b"。
}
