package main

import (
	 "fmt"
	 "math"
)

func go_if(x float64) float64  {
	if x <0 {
		return math.Sqrt(x)
	}
	return math.Sqrt(x)
}

func go_if_else(x, n, lim float64) float64 {
	// Pow(3,４)表示取三的4次方
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("来到else:%g >= %g\n", v, lim)
	}
	// 这里开始就不能使用 v 了
	return lim
}

func Sqrt(x float64) float64 {
	// 实现一个类似math.Sqrt开平方的函数
    // 根据提示　定义一个初始值　来套用公式
    z := 1.00
    //  临时变量 记录z 上次的值
    temp := 0.00

    for {
        // 计算出最新的z值
        z = z - (z*z-x)/(2*z)
        fmt.Println(z)
        if math.Abs(z-temp) < 0.000000000000001 {
            //  当值停止改变（或改变非常小）的时候退出循环
            break
        } else {
            //　赋值最后的值
            temp = z
        }
    }
    return z
}

func main(){
	// fmt.Println("传入负数的结果:",go_if(-9),"传入正数的结果:",go_if(9))
	// fmt.Println(go_if_else(3,3,10))
	Sqrt(2)
}