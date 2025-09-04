package main

import (
	"fmt"
)

func main() {
	// 需要提前声明，不能直接用于接收返回值
	var ret2 int
	_, ret2 = Func1(3, 2) // 调用函数并接收返回值，抛弃第一个返回值
	fmt.Println("a * b =", ret2)	// 等号后面不用空格，逗号会在输出中额外添加空格

	// 或者使用短声明
	ret3, _ := Func1(3, 2) // 调用函数并接收返回值，抛弃第二个返回值
	fmt.Println("a + b =", ret3)

	//  Go 语言中你必须使用所有被声明的变量，但有时你并不需要使用从一个函数得到的所有返回值。
}

// 函数参数声明不需要 var 关键字
// 函数返回值声明不需要 var 关键字
func Func1(a, b int) (int, int) {
	return a + b, a * b
}