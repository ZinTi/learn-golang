// run: go run hello.go 
// or: go build hello.go

// 每个go应用程序都包含一个名为main的包
package main

// 导入fmt包，以使用格式化IO函数
import "fmt"

// 一般为第一个执行的函数（如果有init()函数则会先执行该函数）
func main() {
	/* 
	多行注释，即块注释，与其他大部分语言一致
	*/

	fmt.Println("Hello, World!") // print line 自动添加换行符，注意首字母大写
	fmt.Print("Hello, Zin!\n") // 或者时使用转义换行符
	fmt.Println("Hello, " + "GitHub!") // 字符串连接可以使用+实现

	/*
	标识符（常量、变量、类型、函数名、结构字段等）以第一个字母大写开头，可以被外部包代码所时使用，类似于OOP语言中的public，若以小写字母开头，对包外不可见，但在整个包内部可见可用，相当于OOP语言中的 protected
	*/
}
// 注意：符号 { 不能像 C/C++ 一样单独放一行，而是像 Java 风格一样在一行尾部，但是强制要求，原因是 go 编译器自动在语句尾部添加分号 ;
// 当一行写多条语句时，可以手动添加分号，但是可读性不好，不推荐
// 标识符规则和其他语言基本一致（A~Z a~z 0~9 _ ），首字母不能是数字

