package main

import (
	"fmt"
)

func main() {
	const LENGTH int = 10	// 显示类型定义
	const WIDTH = 5			// 隐式类型定义
	var area int
	const a, b, c = 1, false, "Hello"

	area = LENGTH * WIDTH
	fmt.Println("Area of rectangle:", area)

	// println() 是一个​​内建函数（built-in function）
	println()	// 打印空行
	println(a,b,c)
	/*
	​​内建函数​​：Go 语言提供了一系列无需导入包即可直接使用的内置函数，println 就是其中之一（类似的还有 print、len、cap 等）。
	​​作用​​：println(a, b, c) 会将变量 a、b、c 的值打印到标准错误输出（stderr），并自动添加空格分隔和换行符。
	​​与 fmt 包的区别​​：
	fmt.Println() 来自标准库，功能更强大（支持格式化、多类型等），输出到 stdout。
	println() 是简单的内置函数，主要用于调试，输出到 stderr，不保证未来版本兼容性（官方建议生产代码用 fmt 包）。
	*/


	// 常量用于枚举
	const (
		Unknown = 0
		Male    = 1
		Female  = 2
	)
	println(Unknown, Male, Female)

	

}