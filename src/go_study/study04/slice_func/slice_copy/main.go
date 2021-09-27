package main

import "fmt"

func main()  {
	//函数 copy 在两个 slice 间复制数据，复制⻓度以 len ⼩的为准。两个 slice 可指向同⼀
	//底层数组，允许元素区间重叠。
	data := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := data[8:]
	s2 := data[:5]
	copy(s2, s) // dst:s2, src:s
	fmt.Println(s2)
	fmt.Println(data)
	//输出：
	//[8 9 2 3 4]
	//[8 9 2 3 4 5 6 7 8 9]
	//应及时将所需数据 copy 到较⼩的 slice，以便释放超⼤号底层数组内存。
}
