package main

import "fmt"

//关键字 iota 定义常量组中从 0 开始按⾏计数的⾃增枚举值。
const (
	Sunday = iota // 0
	Monday // 1，通常省略后续⾏表达式。
	Tuesday // 2
	Wednesday // 3
	Thursday // 4
	Friday // 5
	Saturday // 6
)
const (
	_ = iota // iota = 0
	KB int64 = 1 << (10 * iota) // iota = 1
	MB // 与 KB 表达式相同，但 iota = 2
	GB
	TB
)
//在同⼀常量组中，可以提供多个 iota，它们各⾃增⻓。
const (
	A, B = iota, iota << 10 // 0, 0 << 10
	C, D // 1, 1 << 10
)
//如果 iota ⾃增被打断，须显式恢复。
const (
	E = iota // 0
	F // 1
	G = "c" // c
	H // c，与上⼀⾏相同。
	I = iota // 4，显式恢复。注意计数包含了 G、H 两⾏。
	J // 5
)
//可通过⾃定义类型来实现枚举类型限制。
type Color int
const (
	Black Color = iota
	Red
	Blue
)
func test(c Color) {}

func main() {
	c := Black
	test(c)
	x := 1
	fmt.Println(x)
	//test(x) // Error: cannot use x (type int) as type Color in function argument
	test(1) // 常量会被编译器⾃动转换
}