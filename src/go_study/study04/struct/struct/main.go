package main

import "fmt"

//值类型，赋值和传参会复制全部内容。可⽤ "_" 定义补位字段，⽀持指向⾃⾝类型的指针成员。
type Node struct {
	_ int
	id int
	data *byte
	next *Node
}
func main() {
	n1 := Node{
		id: 1,
		data: nil,
	}
	n2 := Node{
		id: 2,
		data: nil,
		next: &n1,
	}

	//顺序初始化必须包含全部字段，否则会出错。
	type User struct {
		name string
		age int
	}
	u1 := User{"Tom", 20}
	//u2 := User{"Tom"} // Error: too few values in struct initializer
	fmt.Println(n1,n2,u1)

	//⽀持匿名结构，可⽤作结构成员或定义变量。
	type File struct {
		name string
		size int
		attr struct {
			perm int
			owner int
		}
	}
	f := File{
		name: "test.txt",
		size: 1025,
		// attr: {0755, 1}, // Error: missing type in composite literal
	}
	f.attr.owner = 1
	f.attr.perm = 0755
	var attr = struct {
		perm int
		owner int
	}{2, 0755}
	f.attr = attr

	//⽀持 "=="、"!=" 相等操作符，可⽤作 map 键类型。
	type User1 struct {
		id int
		name string
	}
	m1 := map[User1]int{
		User1{1, "Tom"}: 100,
	}

	//可定义字段标签，⽤反射读取。标签是类型的组成部分。
	var u2 struct { name string "username" }
	var u3 struct { name string }
	fmt.Println(m1,u2,u3)
	//u3 = u2 // Error: cannot use u2 (type struct { name string "username" }) as

	// type struct { name string } in assignment
	//空结构 "节省" 内存，⽐如⽤来实现 set 数据结构，或者实现没有 "状态" 只有⽅法的 "静态类"。
	var null struct{}
	set := make(map[string]struct{})
	set["a"] = null

}
