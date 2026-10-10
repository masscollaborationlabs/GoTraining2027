/**
  For License Please see LICENSE file BNGPL version 1 or later
  Copyright (C) 2026-2027 Mass Collaboration Labs 
  **/

package main

import "fmt"

func main() {
	fmt.Println("hello")

	var name string = "Jack "
	var surname = "Nicholson"

	fmt.Print(name)
	fmt.Println(surname)

	// integer

	// https://youtu.be/RM4NXAp7Fg0?list=PL1i2Llx7XoAWrIXRPi3YAJ5vZ8lh3083O Teşekkürler hocam :)

	var number int64 = 111

	fmt.Println(number)

	// float

	var y float32 = 6.4

	fmt.Println(y)

	//	fmt.Println( number + int(y))
	fmt.Println(number + int64(y)) // type conversion

	//	var isActive bool = false // bool türü

	salary := 1024 // var salary int başlangıç değeri verilmeli

	fmt.Println(salary)
}

