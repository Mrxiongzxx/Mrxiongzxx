package main

import "fmt"

func main()  {
	//go 静态类型语言  go语言定义的值一定要使用  不然报错
	//定义变量方法
	//方法一
	var x int
	x = 123
	//方法二
	var f float32 = 1.6
	var s = "abc"
	//方法三
	a := 1
	//可以一次定义多个变量
	var b,c int
	var (
		d float32
		e string
	)
	 g,h:= 1.6,"abc"
	fmt.Println(a,b,c,d,e,f,g,h,s,x)

	//多变量赋值时，先计算所有相关值，然后再从左到右依次赋值。
	data, i := [3]int{0, 1, 2}, 0
	i, data[i] = 2, 100 // (i = 0) -> (i = 2), (data[0] = 100)


	//可以用 _ 占位
	var array = [3]int{0,1,2}
	for _,v := range array{
		fmt.Println(v)
	}

	j := []int{0, 0, 0} // 提供初始化表达式。
	j[1] = 10
	k := make([]int, 3) // makeslice
	k[1] = 10
	//l := new([]int)
	//l[1] = 10 // Error: invalid operation: c[1] (index of type *[]int) 没有分配内存会报错

	//类型转换
	//不⽀持隐式类型转换，即便是从窄向宽转换也不⾏。
	var m byte = 100
	// var n int = b // Error: cannot use b (type byte) as type int in assignment
	var n int = int(m) // 显式转换
	fmt.Println(n)


	//使⽤括号避免优先级错误。
	//*Point(p) // 相当于 *(Point(p))
	//(*Point)(p)
	//<-chan int(c) // 相当于 <-(chan int(c))
	//(<-chan int)(c)

	//同样不能将其他类型当 bool 值使⽤。
	//o := 100
	//if o { // Error: non-bool a (type int) used as if condition
	//	println("true")
	//}
}

