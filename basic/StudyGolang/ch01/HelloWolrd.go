package main //首行这个是必须的。所有的 Go 文件以 package <something>开头，对于独立运行的执行文件必须是 package main 。

import "fmt" //这是说需要将 fmt 包加入 main。不是 main 的其他包都被称为库。

func main(){ //package main 必须首先出现，紧跟着是 import。在 Go 中，package 总是首先出现，然后是 import，然后是其他所有内容。
			//当 Go 程序在执行的时候，首先调用的函数是 main.main() ，这是从 C 中继承而来。这里定义了这个函数。
	fmt.Printf("Hello,World,or καληµ´ϵρα κ ó σµϵ; or こんにちは;世界! ") //调用了来自于 fmt 包的函数打印字符串到屏幕。字符串由包裹，
																	//并且可以包含非 ASCII 的字符。这里使用了希腊文和日文。
	fmt.Print("Hello,World,or καληµ´ ϵρα κ ó σµϵ; or こんにちは;世界! ")
}
