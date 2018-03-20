###### 方法	

1. 方法就是一类带特殊的 **接收者** 参数的函数。

   ```go
   func (v Vertex) Abs() float64 {
   	return math.Sqrt(v.X*v.X + v.Y*v.Y)
   }
   ```

   这段例子是有一个接收者为v，名为Abs的方法，是一个结构体类型方法。

2. 方法只是个带接收者参数的函数。

```go
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
```

​	此时Abs就是一个普通的函数了

###### 指针接收者

1. 指针接收者的方法可以修改接收者指向的值，由于方法经常需要修改它的接收者，指针接收者比值接收者更常用。

以下就是个指针接收者的方法，并且修改了传入的值

```go
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}
```

若使用值接收者(接受者前面无`*`号)，那么 方法会对原始 结构体的 值的副本进行操作。（对于函数的其它参数也是如此。）

###### 指针与函数

1. 非指针函数也和非指针接受者方法类似，无法改变结构体本身的值。

###### 方法与指针重定向

1. 带指针参数的函数必须接受一个指针

   ```go
   var v Vertex
   ScaleFunc(v, 5)  // 编译错误！
   ScaleFunc(&v, 5) // OK
   ```

2. 以指针为接收者的方法被调用时，接收者既能为值又能为指针：

```go
var v Vertex
v.Scale(5)  // OK
p := &v
p.Scale(10) // OK
```

3. 接受一个值作为参数的函数必须接受一个指定类型的值：

   ```go
   var v Vertex
   fmt.Println(AbsFunc(v))  // OK
   fmt.Println(AbsFunc(&v)) // 编译错误！
   ```

4. 以值为接收者的方法被调用时，接收者既能为值又能为指针：

   ```go
   var v Vertex
   fmt.Println(v.Abs()) // OK
   p := &v
   fmt.Println(p.Abs()) // OK
   ```