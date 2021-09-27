package main

import (
	"fmt"
	"unsafe"
)

//⽀持指针类型 *T，指针的指针 **T，以及包含包名前缀的 *<package>.T。
//• 默认值 nil，没有 NULL 常量。
//• 操作符 "&" 取变量地址，"*" 透过指针访问⺫标对象。
//• 不⽀持指针运算，不⽀持 "->" 运算符，直接⽤ "." 访问⺫标成员。

func main()  {
	type data struct {
		a int
	}
	var d = data{1234}
	var p *data
	p = &d
	fmt.Printf("%p,%v\n",p,p.a)

	//不能对指针做加减法等运算。
	//x := 1234
	//p := &x
	//p++ // Error: invalid operation: p += 1 (mismatched types *int and int)

	x := 0x12345678 //*int
	//Pointer类型用于表示任意类型的指针。有4个特殊的只能用于Pointer类型的操作：
	//1) 任意类型的指针可以转换为一个Pointer类型值
	//2) 一个Pointer类型值可以转换为任意类型的指针
	//3) 一个uintptr类型值可以转换为一个Pointer类型值
	//4) 一个Pointer类型值可以转换为一个uintptr类型值
	point1 := unsafe.Pointer(&x) // *int -> Pointer int指针 转为任意类型数组
	n := (*[4]byte)(point1)           // Pointer -> *[4]byte  任意类型数组 转为 [4]byte
	for i := 0; i < len(n); i++ {
		fmt.Printf("%X ", n[i])
	}


}

//将 Pointer 转换成 uintptr，可变相实现指针运算。
func PointTranslate() {
	d := struct {
		s string
		x int
	}{"abc", 100}
	p := uintptr(unsafe.Pointer(&d)) // *struct -> Pointer -> uintptr
	p += unsafe.Offsetof(d.x) // uintptr + offset
	p2 := unsafe.Pointer(p) // uintptr -> Pointer
	px := (*int)(p2) // Pointer -> *int
	*px = 200 // d.x = 200
	fmt.Printf("%#v\n", d)
}