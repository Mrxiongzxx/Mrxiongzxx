package main

import "fmt"

//需要说明，slice 并不是数组或数组指针。它通过内部指针和相关属性引⽤数组⽚段，以
//实现变⻓⽅案。
//runtime.h
//struct Slice   切片实质也是个结构体
//{ // must not move anything
//byte* array; // actual data
//uintgo len; // number of elements
//uintgo cap; // allocated number of elements
//}


//• 引⽤类型。但⾃⾝是结构体，值拷⻉传递。
//• 属性 len 表⽰可⽤元素数量，读写操作不能超过该限制。
//• 属性 cap 表⽰最⼤扩张容量，不能超出数组限制。
//• 如果 slice == nil，那么 len、cap 结果都等于 0。

func main()  {

	data := [...]int{0, 1, 2, 3, 4, 5, 6}
	slice := data[1:4:5] // [low : high : max]

	for i := range slice {
		fmt.Println(slice[i])
	}

	//创建表达式使⽤的是元素索引号，⽽⾮数量。
	//data := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	//expression slice              len cap 	comment
	//------------+----------------------+------+-------+---------------------
	//data[:6:8] [0 1 2 3 4 5]   	6 	8 		省略 low.
	//data[5:] [5 6 7 8 9]         	5 	5 		省略 high、max。
	//data[:3] [0 1 2]             	3 	10 		省略 low、max。
	//data[:] [0 1 2 3 4 5 6 7 8 9] 10 	10 		全部省略。

	//读写操作实际⺫标是底层数组，只需注意索引号的差别。
	//切片引用类型  s的改变也会改变data里的值
	s := data[2:4]
	s[0] += 100
	s[1] += 200
	fmt.Println(s)
	fmt.Println(data)

	//可直接创建 slice 对象，⾃动分配底层数组。
	s1 := []int{0, 1, 2, 3, 8: 100} // 通过初始化表达式构造，可使⽤索引号。
	fmt.Println(s1, len(s1), cap(s1))
	s2 := make([]int, 6, 8) // 使⽤ make 创建，指定 len 和 cap 值。
	fmt.Println(s2, len(s2), cap(s2))
	s3 := make([]int, 6) // 省略 cap，相当于 cap = len。
	fmt.Println(s3, len(s3), cap(s3))

	//使⽤ make 动态创建 slice，避免了数组必须⽤常量做⻓度的⿇烦。还可⽤指针直接访问
	//底层数组，退化成普通数组操作。
	s4 := []int{0, 1, 2, 3}
	p := &s4[2] // *int, 获取底层数组元素指针。
	*p += 100
	fmt.Println(s4)
	//	输出：
	//[0 1 102 3]

	//⾄于 [][]T，是指元素类型为 []T 。
	data2 := [][]int{
		[]int{1, 2, 3},
		[]int{100, 200},
		[]int{11, 22, 33, 44},
	}
	fmt.Println(data2)

	//可直接修改 struct array/slice 成员。
	d := [5]struct {
		x int
	}{}
	str := d[1:4]

	d[1].x = 10
	str[2].x = 20

	fmt.Println(d,str)
	fmt.Printf("%p, %p\n", &d, &d[0])

}

func test() {
	//4.2.1 reslice
	//所谓 reslice，是基于已有 slice 创建新 slice 对象，以便在 cap 允许范围内调整属性。
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := s[2:5] // [2 3 4]
	s2 := s1[2:6:7] // [4 5 6 7]
	//s3 := s2[3:6] // Error
	fmt.Println(s,s1,s2)
	//+---+---+---+---+---+---+---+---+---+---+
	//	data | 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 |
	//	+---+---+---+---+---+---+---+---+---+---+
	//	0 2 5
	//+---+---+---+---+---+---+---+---+
	//	s1 | 2 | 3 | 4 | | | | | | len = 3, cap = 8
	//+---+---+---+---+---+---+---+---+
	//	0 2 6 7
	//+---+---+---+---+---+
	//	s2 | 4 | 5 | 6 | 7 | | len = 4, cap = 5
	//+---+---+---+---+---+
	//	0 3 4 5
	//+---+---+---+
	//	s3 | 7 | 8 | X | error: slice bounds out of range
	//+---+---+---+
}

func test1()  {
	//新对象依旧指向原底层数组。
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := s[2:5] // [2 3 4]
	s1[2] = 100
	s2 := s1[2:6] // [100 5 6 7]
	s2[3] = 200
	fmt.Println(s)
	//	输出：
	//[0 1 2 3 100 5 6 200 8 9]
}