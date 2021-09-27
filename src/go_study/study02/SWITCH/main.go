package main

func main()  {
	//分⽀表达式可以是任意类型，不限于常量。可省略 break，默认⾃动终⽌。
	x := []int{1, 2, 3}
	i := 2
	switch i {
	case x[1]:
		println("a")
	case 1, 3:
		println("b")
	default:
		println("c")
	}
}

func test()  {
	//如需要继续下⼀分⽀，可使⽤ fallthrough，但不再判断条件。
	x := 10
	switch x {
	case 10:
		println("a")
		fallthrough
	case 0:
		println("b")
	}
}

func test1()  {
	//省略条件表达式，可当 if...else if...else 使⽤。
	var x = [3]int{}
	switch {
	case x[1] > 0:
		println("a")
	case x[1] < 0:
		println("b")
	default:
		println("c")
	}
	switch i := x[2]; { // 带初始化语句
	case i > 0:
		println("a")
	case i < 0:
		println("b")
	default:
		println("c")
	}
}