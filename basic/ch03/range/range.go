package main

import "fmt"

var r = [] int {1,2,4,8,16,32,64,128}

func main(){
	for i, v := range r {
		fmt.Printf("2**%d = %d\n",i,v)
	}

	pow := make([]int, 10)
	for i := range pow {
		fmt.Println(i)
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}

	range_more()
}
/*
"<<"运算符表示，将"<<"左边的数变成二进制数，往左边移动"<<"右边的位数
*/

type Foo struct {
    bar string
}

func range_more() {
	list := []Foo{
        {"A"},
        {"B"},
        {"C"},
    }

	list2 := make([]*Foo, len(list))
	fmt.Println(list2)
    for i, value := range list {
		fmt.Println(value)
		// list2[i] = &list[i] // {A} {B} {C}
		list2[i] = &value // &{C} &{C} &{C}
    }

    fmt.Println(list[0], list[1], list[2])
    fmt.Println(list2[0], list2[1], list2[2])
}
