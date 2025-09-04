// hello2.go
package main

import (
	"fmt"
) // 可以使用圆括号包含所有要导入的库，注意不是花括号，而是圆括号

func main(){
	var name = "ZhangSan"
	var age = 20
	var url = "Name=%s&Age=%d"
	// Sprintf 格式化字符串并赋值给新字符串
	var target_url = fmt.Sprintf(url, name, age)
	fmt.Println(target_url)

	var name2 = "Jack"
	var age2 = 23
	var url2 = "Name=%s&Age=%d"
	// Printf 格式化字符串并写入标准输出
	fmt.Printf(url2, name2, age2)

	// var x int
	// const Pi float64 = 3.14159265358979323846
	// fmt.Println(x)
	// fmt.Println(Pi)


}
