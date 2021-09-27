package main

import "fmt"

//闭包复制的是原对象指针，这就很容易解释延迟引⽤现象。
func test() func() {
	x := 100
	fmt.Printf("x (%p) = %d\n", &x, x)
	return func() {
		fmt.Printf("x (%p) = %d\n", &x, x)
	}
}

//匿名函数可赋值给变量，做为结构字段，或者在 channel ⾥传送。
func main()  {
	// --- function variable ---
	fn := func() { println("Hello, World!") }
	fn()  //会输出 "Hello, World!"
	// --- function collection ---
	fns := [](func(x int) int){
		func(x int) int { return x + 1 },
		func(x int) int { return x + 2 },
	}
	//对比int切片  fns := []int{}
	//fns[0](100) => func(x int) int { return x + 1 }, x传100进去 得出101
	println(fns[0](100))
	// --- function as field ---
	d := struct {
		fn func() string
	}{
		fn: func() string { return "Hello, World!" },
	}
	println(d.fn())
	// --- channel of function ---
	fc := make(chan func() string, 2)
	fc <- func() string { return "Hello, World!" }
	println((<-fc)())

	f := test()
	f()
	//输出：
	//x (0xc00000a098) = 100
	//x (0xc00000a098) = 100

	//在汇编层⾯，test 实际返回的是 FuncVal 对象，其中包含了匿名函数地址、闭包对象指
	//针。当调⽤匿名函数时，只需以某个寄存器传递该对象即可。
	//FuncVal { func_address, closure_var_pointer ...
}



