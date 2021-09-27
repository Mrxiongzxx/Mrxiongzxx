package main

import "fmt"

//不⽀持 嵌套 (nested)、重载 (overload) 和 默认参数 (default parameter)。
//• ⽆需声明原型。
//• ⽀持不定⻓变参。
//• ⽀持多返回值。
//• ⽀持命名返回参数。
//• ⽀持匿名函数和闭包。
//使⽤关键字 func 定义函数，左⼤括号依旧不能另起⼀⾏


func test(x, y int, s string) (int, string) { // 类型相同的相邻参数可合并。
	n := x + y // 多返回值必须⽤括号。
	return n, fmt.Sprintf(s, n)
}

//函数是第⼀类对象，可作为参数传递。建议将复杂签名定义为函数类型，以便于阅读。
func test2(fn func() int) int {
	return fn()
}
type FormatFunc func(s string, x, y int) string // 定义函数类型。
func format(fn FormatFunc, s string, x, y int) string {
	return fn(s, x, y)
}

//变参 本质上就是 slice。只能有⼀个，且必须是最后⼀个
func test3(s string,n ... int) string{
	 var x = 0
	 for i,_ := range n {
		 x += i
	 }
	 return fmt.Sprintf(s,x)
}


func test4() (int,int){
	return 1,2
}

func add(x, y int) int {
	return x + y
}

func sum(n ...int) int {
	var x int
	for _, i := range n {
		x += i
	}
	return x
}

//命名返回参数可看做与形参类似的局部变量，最后由 return 隐式返回。
func add2(x, y int) (z int) {
	z = x + y
	return  //不加参数 默认返回z
}

//命名返回参数可被同名局部变量遮蔽，此时需要显式返回。
func add3(x, y int) (z int) {
	{ // 不能在⼀个级别，引发 "z redeclared in this block" 错误。
		var z = x + y
		// return // Error: z is shadowed during return
		return z // 必须显式返回。
	}
}

//命名返回参数允许 defer 延迟调⽤通过闭包读取和修改。
func add4(x,y int) (z int){
	defer func() {
		z += 100
	}()
	z = x + y
	return
}

//显式 return 返回前，会先修改命名返回参数。
func add5(x, y int) (z int) {
	defer func() {
		println(z) // 输出: 203
	}()
	z = x + y
	return z + 200 // 执⾏顺序: (z = z + 200) -> (call defer) -> (ret)
}


func  main()  {
	test(1,2,"abc")

	s1 := test2(func() int { return 100 }) // 直接将匿名函数当参数。
	s2 := format(func(s string, x, y int) string {
		return fmt.Sprintf(s, x, y)
	}, "%d, %d", 10, 20)
	println(s1, s2)

	println(test3("sum: %d", 1, 2, 3))

	//使⽤ slice 对象做变参时，必须展开
	s := []int{1, 2, 3}
	println(test3("sum: %d", s...))

	// s4 := make([]int, 2)
	// s4 = test() // Error: multiple-value test() in single-value context
	//不能⽤容器对象接收多返回值。只能⽤多个变量，或 "_" 忽略
	x, _ := test4()
	println(x)

	//多返回值可直接作为其他函数调⽤实参。
	println(add(test4()))
	println(sum(test4()))

	println(add4(1, 2)) // 输出: 103

	println(add5(1, 2)) // 输出: 203
}