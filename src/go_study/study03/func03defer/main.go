package main

import (
	"os"
	"sync"
	"testing"
)

//延迟调用函数
//关键字 defer ⽤于注册延迟调⽤。这些调⽤直到 ret 前才被执⾏，通常⽤于释放资源或错
//误处理。
func test() error {
	f, err := os.Create("test.txt")
	if err != nil { return err }
	defer f.Close() // 注册调⽤，⽽不是注册函数。必须提供参数，哪怕为空。
	f.WriteString("Hello, World!")
	return nil
}

//多个 defer 注册，按 FILO 次序执⾏。哪怕函数或某个延迟调⽤发⽣错误，这些调⽤依旧
//会被执⾏。
func test2(x int) {
	defer println("a")
	defer println("b")
	defer func() {
		println(100 / x) // div0 异常未被捕获，逐步往外传递，最终终⽌进程。
	}()
	defer println("c")
}
//输出
//c
//b
//a
//延迟调⽤参数在注册时求值或复制，可⽤指针或闭包 "延迟" 读取。
func test3() {
	x, y := 10, 20
	defer func(i int) {
		println("defer:", i, y) // y 闭包引⽤

	}(x) // x 被复制
	x += 10
	y += 100
	println("x =", x, "y =", y)
}

//滥⽤ defer 可能会导致性能问题，尤其是在⼀个 "⼤循环" ⾥。
var lock sync.Mutex
func test4() {
	lock.Lock()
	lock.Unlock()
}
func testdefer() {
	lock.Lock()
	defer lock.Unlock()
}
func BenchmarkTest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		test4()
	}
}
func BenchmarkTestDefer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testdefer()
	}
}

func main() {
	test2(0)
	test3()
}

