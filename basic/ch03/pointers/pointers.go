package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i         // 指向 i
	fmt.Println(*p) // 通过指针读取读取i的值
	*p = 21         // 通过指针设置i的值
	fmt.Println(i)  // 查看新赋值

	p = &j         // 指向 j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}
