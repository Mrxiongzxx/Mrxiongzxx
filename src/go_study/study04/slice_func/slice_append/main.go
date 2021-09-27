package main

import "fmt"

func main() {
	//向 slice 尾部添加数据，返回新的 slice 对象。
	s := make([]int, 0, 5)
	fmt.Printf("%p\n", &s)
	s2 := append(s, 1)
	fmt.Printf("%p\n", &s2)
	fmt.Println(s, s2)
	//输出：
	//0x210230000
	//0x210230040
	//[] [1]


	data := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s = data[:3]
	s2 = append(s, 100, 200) // 添加多个值。
	fmt.Println(data)
	fmt.Println(s)
	fmt.Println(s2)

	//输出：
	//[0 1 2 100 200 5 6 7 8 9]
	//[0 1 2]
	//[0 1 2 100 200]

	//⼀旦超出原 slice.cap 限制，就会重新分配底层数组，即便原数组并未填满。
	data1 := [...]int{0, 1, 2, 3, 4, 10: 0}
	s = data1[:2:3]
	s = append(s, 100, 200) // ⼀次 append 两个值，超出 s.cap 限制。
	fmt.Println(s, data) // 重新分配底层数组，与原数组⽆关。
	fmt.Println(&s[0], &data[0]) // ⽐对底层数组起始指针。
	//输出：
	//[0 1 100 200] [0 1 2 3 4 0 0 0 0 0 0]
	//0x20819c180 0x20817c0c0

	//从输出结果可以看出，append 后的 s 重新分配了底层数组，并复制数据。如果只追加⼀ 个值，则不会超过 s.cap 限制，也就不会重新分配。
	//通常以 2 倍容量重新分配底层数组。在⼤批量添加数据时，建议⼀次性分配⾜够⼤的空间，
	//以减少内存分配和数据复制开销。或初始化⾜够⻓的 len 属性，改⽤索引号进⾏操作。
	//及时释放不再使⽤的 slice 对象，避免持有过期数组，造成 GC ⽆法回收。
	s = make([]int, 0, 1)
	c := cap(s)
	for i := 0; i < 50; i++ {
		s = append(s, i)
		if n := cap(s); n > c {
			fmt.Printf("cap: %d -> %d\n", c, n)
			c = n
		}
	}
	//输出:
	//cap: 1 -> 2
	//cap: 2 -> 4

}
