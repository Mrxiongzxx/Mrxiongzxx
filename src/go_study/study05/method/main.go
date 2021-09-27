package main

import "fmt"


//⽅法总是绑定对象实例，并隐式将实例作为第⼀实参 (receiver)。
//• 只能为当前包内命名类型定义⽅法。
//• 参数 receiver 可任意命名。如⽅法中未曾使⽤，可省略参数名。
//• 参数 receiver 类型可以是 T 或 *T。基类型 T 不能是接⼝或指针。
//• 不⽀持⽅法重载，receiver 只是参数签名的组成部分。
//• 可⽤实例 value 或 pointer 调⽤全部⽅法，编译器⾃动转换。
//没有构造和析构⽅法，通常⽤简单⼯⼚模式返回对象实例。

type Queue struct {
	elements []interface{}
}
func NewQueue() *Queue { // 创建对象实例。
	return &Queue{make([]interface{}, 10)}
}
func (*Queue) Push(e interface{}) error { // 省略 receiver 参数名。
	panic("not implemented")
}
// func (Queue) Push(e int) error { // Error: method redeclared: Queue.Push
// panic("not implemented")
// }
func (self *Queue) length() int { // receiver 参数名可以是 self、this 或其他。
	return len(self.elements)
}

//⽅法不过是⼀种特殊的函数，只需将其还原，就知道 receiver T 和 *T 的差别。
type Data struct{
	x int
}
func (self Data) ValueTest() { // func ValueTest(self Data);
	fmt.Printf("Value: %p\n", &self)
}
func (self *Data) PointerTest() { // func PointerTest(self *Data);
	fmt.Printf("Pointer: %p\n", self)
}

type X struct{}
func (*X) test() {
	println("X.test")
}

func main() {
	d := Data{}
	p := &d
	fmt.Printf("Data: %p\n", p)
	d.ValueTest() // ValueTest(d)
	d.PointerTest() // PointerTest(&d)
	p.ValueTest() // ValueTest(*p)
	p.PointerTest() // PointerTest(p)
	//输出：
	//Data : 0x2101ef018
	//Value : 0x2101ef028
	//Pointer: 0x2101ef018
	//Value : 0x2101ef030
	//Pointer: 0x2101ef018

	//从 1.4 开始，不再⽀持多级指针查找⽅法成员。
	q := &X{}
	q.test()
	// Error: calling method with receiver &p (type **X) requires explicit dereference
	// (&q).test()
}


