package main

import "fmt"

func main()  {
	//初始化复合对象，必须使⽤类型标签，且左⼤括号必须在类型尾部。
	// var a struct { x int } = { 100 } // syntax error
	// var b []int = { 1, 2, 3 } // syntax error
	// c := struct {x int; y string} // syntax error: unexpected semicolon or newline
	// {
	// }
	var a = struct{ x int }{100}
	var b = []int{1, 2, 3}
	//初始化值以 "," 分隔。可以分多⾏，但最后⼀⾏必须以 "," 或 "}" 结尾。
	fmt.Println(a,b)
}

func test(){
	//a := []int{
	//
	//	1,
	//	2 // Error: need trailing comma before newline in composite literal
	//}
	b := []int{
		1,
		2, // ok
	}
	c := []int{
		1,
		2} // ok
	fmt.Println(b,c)
}