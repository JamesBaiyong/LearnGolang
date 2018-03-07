package main

import "fmt"

func slices() {
	//普通切片
	array := [6] int {1,2,3,4,5,6}
	//包括第一个数，但不包括第４个数
	var s [] int = array[1:4]
	fmt.Println(s)
}

func slices_pointers() {
	names := [5] string {"Jake","Tom","James","Messi","Buffon"}
	a := names[1:4]
	b := names[0:1]
	fmt.Println(a,b,names)
	//修改切片元素，会修改底层的数组对应的元素
	b[0] = "BOOM"
	fmt.Println(a,b,names)

}

func slices_literals() {
	//无长度切片
	q := []int{1,2,3,4,5,6}
	fmt.Println(q)

	//结构体构建的无长度切片
	s := [] struct{
		i int
		b bool
	}{
		{1,true},
		{2,false},
		{3,true},
		{4,false},
	}
	fmt.Println(s)
}

func slices_bounds(){
	//切片默认行为
	s := [] int {1,2,3,4,5,6}
	var a = s[1:4]
	var b = s[:4]
	var c = s[1:]
	fmt.Println(a,b,c)
}

func slices_len_cap(){
	//切片的长度和容量
	s := [] int {1,2,3,4,5,6}
	fmt.Printf("len=%d,cap=%d,%v\n",len(s),cap(s),s)

	s = s[:0]
	fmt.Printf("len=%d,cap=%d,%v\n",len(s),cap(s),s)

	s = s[:4]
	fmt.Printf("len=%d,cap=%d,%v\n",len(s),cap(s),s)

	s = s[2:]
	fmt.Printf("len=%d,cap=%d,%v\n",len(s),cap(s),s)
}

func slices_nil(){
	//nil切片
	var s [] int 
	fmt.Printf("len=%d,cap=%d,%v\n",len(s),cap(s),s)
	if s == nil {
		fmt.Println("nil！")
	}
}

func slices_append(){
	//向切片追加元素
	var s [] int
	fmt.Printf("len=%d,cap=%d,%v\n",len(s),cap(s),s)

	s = append(s,0)
	fmt.Printf("len=%d,cap=%d,%v\n",len(s),cap(s),s)

	s = append(s,1,2,3)
	fmt.Printf("len=%d,cap=%d,%v\n",len(s),cap(s),s)
}

func main() {
	slices()
	slices_pointers()
	slices_literals()
	slices_bounds()
	slices_len_cap()
	slices_nil()
	slices_append()
}
