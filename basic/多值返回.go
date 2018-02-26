package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("大吉", "狗年")
	fmt.Println(a, b)
}
