package main

import (
	"fmt"
	"runtime"
	"time"
)

func os_know() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS
	os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("deflult : %s.",os)
	}
}

func when_saturday(){
	fmt.Println("when is Saturday??")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Just Today.")
	case today + 1:
		fmt.Println("one day ago.")
	case today + 2:
		fmt.Println("two day ago.")
	default:
		fmt.Println("too far awy.")
	}
}

func no_condition(){
	t := time.Now()
	fmt.Println(t)
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning :)")
	case t.Hour() < 17:
		fmt.Println("Good afternoon :)")
	default:
		fmt.Println("Good evening :)")
	}
}

func main() {
	os_know()
	fmt.Println()
	when_saturday()
	no_condition()
}
