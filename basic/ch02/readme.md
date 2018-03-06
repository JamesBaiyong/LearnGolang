#### 控制语句

###### for循环

1. Go 只有一种循环结构： `for` 循环,Go 的 for 语句后面没有小括号，大括号 `{ }` 则是必须的。

2. 基本的 `for` 循环由三部分组成，它们用分号隔开：

   初始化语句：在第一次迭代前执行

   条件表达式：在每次迭代前求值

   后置语句：在每次迭代的结尾执行

3. 如果省略循环条件，该循环就不会结束，因此无限循环可以写得很紧凑。

```go
package main

func main() {
	for {
	}
}
```

###### if 语句

1. Go 的 `if` 语句与 `for` 循环类似，表达式外无需小括号 `( )` ，而大括号 `{ }` 则是必须的。

2. 同 `for` 一样， `if` 语句可以在条件表达式前执行一个简单的语句。

   ```go
   func pow(x, n, lim float64) float64 {
   	if v := math.Pow(x, n); v < lim {
   		return v
   	}
   	return lim
   }
   ```


###### switch语句

1. `switch` 是编写一连串 `if - else` 语句的简便方法。它运行第一个值等于条件表达式的 case 语句。且Go 只运行选定的 case，而非之后所有的 case。 实际上，Go 自动提供了在这些语言中每个 case 后面所需的 `break` 语句。
2. 除非以 `fallthrough` 语句结束，否则分支会自动终止。 Go 的另一点重要的不同在于 switch 的 case 无需为常量，且取值不必为整数。
3. 无条件的switch,这种形式能将一长串 "if-then-else" 写得更加清晰。

```go
package main

import "fmt"

func main(){
  switch parm := 1
  parm {
    case 1:
    	fmt.Println(123)
    case 2:
    	fmt.Println(321)
    default:
    	fmt.Println('h31l0')
  }
}
```

###### defer语句

1. defer 语句会将函数推迟到外层函数返回之后执行,推迟调用的函数其参数会立即求值，但直到外层函数返回前该函数都不会被调用。
2. defer后的函数或语句调用会被压入一个栈中。 当外层函数返回时，被推迟的函数会按照后进先出的顺序调用。也就是说，一个函数中，defer先近后出。

```go
package main

import (
	"fmt"
)

func defer_test() {
	fmt.Println("World")
}

func main() {
	defer fmt.Println("done")
	fmt.Println("Hello")
	fmt.Println("倒计时...")
	for i:=0;i<=10; i++{
		defer fmt.Println(i)
	}
	fmt.Println("go...")
	defer_test()
}

/*out:
Hello
倒计时...
go...
World
10
9
8
7
6
5
4
3
2
1
0
done
*/
```

