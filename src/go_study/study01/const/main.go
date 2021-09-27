package main

import "unsafe"

//常量值必须是编译期可确定的数字、字符串、布尔值。
const x, y int = 1, 2 // 多常量初始化
const s = "Hello, World!" // 类型推断
const ( // 常量组
	a, b = 10, 100
	c bool = false
)
//在常量组中，如不提供类型和初始化值，那么视作与上⼀常量相同。
const (
	str = "abc"
	str2// str2 = "abc"
)
//常量值还可以是 len、cap、unsafe.Sizeof 等编译期可确定结果的函数返回值。
const (
	d = "abc"
	e = len(d)
	f = unsafe.Sizeof(e)
)
//如果常量类型⾜以存储初始化值，那么不会引发溢出错误。
const (
	g byte = 100 // int to byte
	h int = 1e20 // float64 to int, overflows
)
func main() {
	const x = "xxx" // 未使⽤局部常量不会引发编译错误。
}
