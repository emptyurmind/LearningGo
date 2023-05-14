package main

import "fmt"

func main() {
	// 函数值不仅仅是一串代码，还记录了状态
	// Go使用闭包（closures）技术实现函数值
	f := squares()
	fmt.Println(f()) // "1"
	fmt.Println(f()) // "4"
	fmt.Println(f()) // "9"
	fmt.Println(f()) // "16"
	// 变量的生命周期不由它的作用域决定：squares返回后，变量x仍然隐式的存在于f中。
}

func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}
func negative(n int) int   { return -n }
func product(m, n int) int { return m * n }
