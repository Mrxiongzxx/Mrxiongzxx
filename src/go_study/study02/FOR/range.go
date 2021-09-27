package main

import "fmt"

func main()  {
	s := "abc"
	for i := range s { // 忽略 2nd value，⽀持 string/array/slice/map。
		println(s[i])
	}
	for _, c := range s { // 忽略 index。
		println(c)
	}
	for range s { // 忽略全部返回值，仅迭代。
		//...
	}
	m := map[string]int{"a": 1, "b": 2}
	for k, v := range m { // 返回 (key, value)。
		println(k, v)
	}
	//注意，range 会复制对象。
	a := [3]int{0, 1, 2}
	for i, v := range a { // index、value 都是从复制品中取出。
		if i == 0 { // 在修改前，我们先修改原数组。
			a[1], a[2] = 999, 999
			fmt.Println(a) // 确认修改有效，输出 [0, 999, 999]。
		}
		a[i] = v + 100 // 使⽤复制品中取出的 value 修改原数组。
	}
	fmt.Println(a) // 输出 [100, 101, 102]。
	test()
}

func test()  {
	//建议改⽤引⽤类型，其底层数据不会被复制。
	s := []int{1, 2, 3, 4, 5}
	for i, v := range s { // 复制 struct slice { pointer, len, cap }。
		if i == 0 {
			s = s[:3] // 对 slice 的修改，不会影响 range。
			s[2] = 100 // 对底层数据的修改。
		}
		println(i, v)
	}
}