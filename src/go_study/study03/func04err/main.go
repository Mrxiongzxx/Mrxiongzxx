package main

import (
	"errors"
	"fmt"
)

//没有结构化异常，使⽤ panic 抛出错误，recover 捕获错误。
func test() {
	defer func() {
		if err := recover(); err != nil {
			println(err.(string)) // 将 interface{} 转型为具体类型。
		}
	}()
	panic("panic error!")
}

//由于 panic、recover 参数类型为 interface{}，因此可抛出任何类型对象。
//func panic(v interface{})
//func recover() interface{}
//延迟调⽤中引发的错误，可被后续延迟调⽤捕获，但仅最后⼀个错误可被捕获。
func test1() {
	defer func() {
		fmt.Println(recover())
	}()
	defer func() {
		panic("defer panic")
	}()
	panic("test panic")
}

//捕获函数 recover 只有在延迟调⽤内直接调⽤才会终⽌错误，否则总是返回 nil。任何未 捕获的错误都会沿调⽤堆栈向外传递。
func test2() {
	defer recover() // ⽆效！
	defer fmt.Println(recover()) // ⽆效！
	defer func() {
		func() {
			println("defer inner")
			recover() // ⽆效！
		}()

	}()
	panic("test panic")
}

//使⽤延迟匿名函数或下⾯这样都是有效的。
func except() {
	recover()
}
func test3() {
	defer except()
	panic("test panic")
}


//如果需要保护代码⽚段，可将代码块重构成匿名函数，如此可确保后续代码被执⾏。
func test4(x, y int) {
	var z int
	func() {
		defer func() {
			if recover() != nil { z = 0 }
		}()
		z = x / y
		return
	}()
	println("x / y =", z)
}
//除⽤ panic 引发中断性错误外，还可返回 error 类型错误对象来表⽰函数调⽤状态。
type error interface {
	Error() string
}
//标准库 errors.New 和 fmt.Errorf 函数⽤于创建实现 error 接⼝的错误对象。通过判断
//错误对象实例来确定具体错误类型。
var ErrDivByZero = errors.New("division by zero")
func div(x, y int) (int, error) {
	if y == 0 { return 0, ErrDivByZero }
	return x / y, nil
}

//如何区别使⽤ panic 和 error 两种⽅式？惯例是：导致关键流程出现不可修复性错误的 使⽤ panic，其他使⽤ error。

func main() {
	test()
	test1()
	//test2()
	test3()
	test4(1,0)
	switch z, err := div(10, 0); err {
	case nil:
		println(z)
	case ErrDivByZero:
		panic(err)
	}
}
