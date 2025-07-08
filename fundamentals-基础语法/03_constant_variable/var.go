package main

import (
	"fmt"
)

// 因式分解关键字的写法一般用于声明全局变量
var (
	g1 int
	g2 bool
	g3 string
)

func main() {
	// Declare a variable of type string
	var a string = "Zin"
	fmt.Println("a =", a)

	var b, c int = 1, 2		// 同时声明多个变量并赋值
	fmt.Println("b =", b, "c =", c)

	// 声明变量时不初始化，默认值为零值
	// 数值类型为 0 
	// 布尔类型为 false
	// 字符串类型为 ""
	// 引用类型为 nil
	var i int
    var f float64
    var b2 bool
    var s string

    fmt.Printf("%v %v %v %q\n", i, f, b2, s)

	// 自动推断类型
	x := 42
	fmt.Println("x =", x)

	// 不带声明格式的只能在函数体中1出现
	z1, z2 := 3.14, "Hello"
	fmt.Println("z1 =", z1, "z2 =", z2)

}