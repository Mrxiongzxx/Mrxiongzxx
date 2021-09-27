package main

import "fmt"

func main() {
	//可将类型分为命名和未命名两⼤类。命名类型包括 bool、int、string 等，⽽ array、
	//slice、map 等和具体元素类型、⻓度等有关，属于未命名类型。
	//具有相同声明的未命名类型被视为同⼀类型。
	//• 具有相同基类型的指针。
	//• 具有相同元素类型和⻓度的 array。
	//• 具有相同元素类型的 slice。
	//• 具有相同键值类型的 map。
	//• 具有相同元素类型和传送⽅向的 channel。
	//• 具有相同字段序列 (字段名、类型、标签、顺序) 的匿名 struct。
	//	• 签名相同 (参数和返回值，不包括参数名称) 的 function。
	//	• ⽅法集相同 (⽅法名、⽅法签名相同，和次序⽆关) 的 interface。
	//	var a struct { x int `a` }
	//	var c struct { x int `ab` }
	//	// cannot use a (type struct { x int "a" }) as type struct { x int "ab" } in assignment 不同的tag反射 所以两个结构体不同
	//	c = a
		//可⽤ type 在全局或函数内定义新类型。

		type bigint int64
		var x bigint = 100
		println(x)

		//新类型不是原类型的别名，除拥有相同数据存储结构外，它们之间没有任何关系，不会持
		//有原类型任何信息。除⾮⺫标类型是未命名类型，否则必须显式转换。
		y := 1234
		var b bigint = bigint(y) // 必须显式转换，除⾮是常量。
		var b2 int64 = int64(b)
		var s  = []int{1, 2, 3} // 未命名类型，隐式转换。
		var s2 []int = s
		fmt.Println(b2,s2)
}
