package main

import (
	"fmt"
)

func main(){
	// 1. boolean
	var b1 bool = true
	var b2 bool = false
	var type_bool = "b1=%t, b2=%t\n" // format verbs(bool): %t
	fmt.Printf(type_bool, b1, b2)
	
	// 2. 常见数字类型
	var n1 int = 64		// integer
	var n2 float32 = 2.7
	var n3 float64 = 3.141592653589793
	var type_number = "n1=%d, n2=%f, n3=%f\n" // %f 默认精度为6
	fmt.Printf(type_number, n1, n2, n3)
	
	// 3. complex复数
	var c1 complex64 = complex(3, 4) // 函数创建复数 3 + 4i
	c2 := complex(1.7, -2.3) // 1.7 - 2.3i (自动推断为complex128)
	c3 := 5 + 6i	// 直接使用字面量，自动推断complex128
	var c4 complex64 = 7 - 8i // 7-8i
	var c5 complex128
	c5 = 9 + 10i	// 先声明后赋值
	
	fmt.Printf("c1: %v\n", c1)
	fmt.Printf("c2: %+v\n", c2)	// 在结构体等类型中会显示字段名
	fmt.Printf("c3: %#v\n", c3)	// Go 语伐表示
	fmt.Printf("c4 实部 %f, c4 虚部 %f\n", real(c4), imag(c4))
	fmt.Printf("c5: %.2f + %.2fi\n", real(c5), imag(c5))

	// 4. 字符串
	var s1 string = "Hello, World!"
	var s2 string = `Hello,
World!` // 使用反引号可以创建多行字符串
	var s3 string = "你好，世界！"
	var type_string = "s1=%s\ns2=%s\ns3=%s\n" // %s 用于字符串
	fmt.Printf(type_string, s1, s2, s3)
	// 注意：字符串是不可变的，不能直接修改字符串中的字符
	// 5. 字符
	var c rune = 'A' // rune 是 int32 的别名，表示 Unicode 字符
	var type_char = "c=%c\n" // %c 用于字符
	fmt.Printf(type_char, c)
	// 6. 派生类型
	// 包括 指针、数组、切片、映射（map）、结构体（struct）、通道（channel）等
}
