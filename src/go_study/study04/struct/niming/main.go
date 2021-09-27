package main

import "fmt"

//匿名字段不过是⼀种语法糖，从根本上说，就是⼀个与成员类型同名 (不含包名) 的字段。
//被匿名嵌⼊的可以是任何类型，当然也包括指针。

func main()  {
	type User struct {
		name string
	}
	type Manager struct {
		User
		title string
	}

	m := Manager{
		User: User{"Tom"}, // 匿名字段的显式字段名，和类型名相同。
		title: "Adminstrator",
	}
	fmt.Println(m)

}

func test()  {
	//可以像普通字段那样访问匿名字段成员，编译器从外向内逐级查找所有层次的匿名字段，
	//直到发现⺫标或出错。
	type Resource struct {
		id int
	}
	type User struct {
		Resource
		name string
	}
	type Manager struct {
		User
		title string
	}
	var m Manager
	m.id = 1 //找到Resource里的id赋值
	m.name = "Jack" //找到User里面的name字段赋值
	m.title = "Administrator"

}

func test2()  {
	//外层同名字段会遮蔽嵌⼊字段成员，相同层次的同名字段也会让编译器⽆所适从。解决⽅
	//法是使⽤显式字段名。
	type Resource struct {
		id int
		name string
	}
	type Classify struct {
		id int
	}
	type User struct {
		Resource // Resource.id 与 Classify.id 处于同⼀层次。
		Classify
		name string // 遮蔽 Resource.name。
	}
	u := User{
		Resource{1, "people"},
		Classify{100},
		"Jack",
	}
	println(u.name) // User.name: Jack
	println(u.Resource.name) // people
	// println(u.id) // Error: ambiguous selector u.id  不指明要那个匿名字段的id 就会报错
	println(u.Classify.id) // 100
}

func test3(){
	//不能同时嵌⼊某⼀类型和其指针类型，因为它们名字相同。
	type Resource struct {
		id int
	}
	type User struct {
		*Resource
		// Resource // Error: duplicate field Resource
		name string
	}
	u := User{
		&Resource{1},
		"Administrator",
	}
	println(u.id)
	println(u.Resource.id)
}