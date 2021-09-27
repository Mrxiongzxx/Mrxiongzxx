package main

import "fmt"

func main()  {
	//⾯向对象三⼤特征⾥，Go 仅⽀持封装，尽管匿名字段的内存布局和⾏为类似继承。没有
	//class 关键字，没有继承、多态等等。

	type User struct {
		id int
		name string
	}
	type Manager struct {
		User
		title string
	}
	m := Manager{User{1, "Tom"}, "Administrator"}
	// var u User = m // Error: cannot use m (type Manager) as type User in assignment
	// 没有继承，⾃然也不会有多态。
	var u User = m.User // 同类型拷⻉。
	fmt.Println(u)
	//内存布局和 C struct 相同，没有任何附加的 object 信息。
	//	|<-------- User:24 ------->|<-- title:16 -->|
	//	+--------+-----------+------------+ +---------------+
	//	m | 1 | string | string | | Administrator | [n]byte
	//	+--------+-----------+------------+ +---------------+
	//	| | |
	//	| +--->>>------------------->>>--------+
	//	|
	//	+--->>>----------------------------------->>>-----+
	//	|
	//	+--->>>----------------------------------->>>-+ |
	//	| | |
	//	+--------+-----------+ +---------+
	//	u | 1 | string | | Tom | [n]byte
	//	+--------+-----------+ +---------+
	//	|<- id:8 -->|<- name:16 -->|
	//	可⽤ unsafe 包相关函数输出内存地址信息。
	//	m : 0x2102271b0, size: 40, align: 8
	//	m.id : 0x2102271b0, offset: 0
	//	m.name : 0x2102271b8, offset: 8
	//	m.title: 0x2102271c8, offset: 24
}
