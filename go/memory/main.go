package main

import "fmt"

func Round(n, a uintptr) uintptr {
	return (n + a - 1) &^ (a - 1)
}

func test() *int {
	x := new(int)
	*x = 0xAABB
	return x
}

func main() {
	fmt.Println(*test())
}

/*
1.
go build -gcflags "-l" -o test test.go 关闭内联优化
go tool objdump -s "main\.test" test
*
test.go:6		0x1092fe8		e8437ef7ff		CALL runtime.newobject(SB)  堆上分配
*
没有内联时，需要在两个栈帧间传递对象，因此在堆上分配 而不是返回一个失效栈帧里的数据

2.
go build -o test test.go  使用默认参数
go tool objdump -s "main\.main" test
内联优化后的代码没有调用 newobject 在堆上分配内存
当内联后，实际上就成了 main 栈帧内的局部变量， 无需去堆上操作。

*/
