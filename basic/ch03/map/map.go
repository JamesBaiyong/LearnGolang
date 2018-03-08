package main

import "fmt"

type Veritx struct{
	Lat,Long float64
}

//类似字典
var m =  map[string] Veritx{
	"A":Veritx{33.3,22.2,},
	"B":Veritx{22.2,33.3,},

}

//顶级类型只是一个类型名，你可以在文法的元素中省略它。
var m1 = map[string] Veritx{
	"A":{33.3,22.3},
	"B":{222.2,333.3},
}

func main(){
	fmt.Println(m)
	fmt.Println(m1)
	fmt.Println(m["A"])
	modifiy_map()
}

func modifiy_map(){
	var m = make(map [string] int)
	m ["yes"] = 34
	fmt.Println("The value:",m["yes"])

	m["yes"] = 121
	sf := m["yes"]
	fmt.Println("The value:",m["yes"])
	fmt.Println(sf)

	delete(m,"yes") //删除值
	fmt.Println("The value:",m["yes"])

	v, k := m["yes"]
	fmt.Println("The value:",v,"是否存在key:",k)

}