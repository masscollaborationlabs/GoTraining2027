/**
Copyright (c) 2026-2027 Mass Collaboration Labs

This program is free software: you can redistribute it and/or modify 
it under the terms of the Broad Netwide General Public License as 
published by the Project Owner under Version 2.0 of this License, 
or (at your option) any later version officially released by the Project Owner 
that explicitly and legally permits such upgrade path within its terms.

This program is distributed in the hope that it will be useful, 
but WITHOUT ANY WARRANTY; without even the implied warranty of 
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the 
Broad Netwide General Public License for more details.

You should have received a copy of the Broad Netwide General Public License
along with this program. If not, see the root directory of this repository 
or the official project legal notices.

Consistent with Section 5 of BNGPL v2.0, if this software is conveyed 
as a binary or deployed via a cloud/SaaS interface, the end-user 
must be provided access to this copyright notice via the application's 
legal GUI, installation directory text file, or primary service webpage.
**/

package main

import "fmt"

func main() {
	fmt.Println("hello")

	var name string = "Jack "
	var surname = "Nicholson"
	var school string = "Los Angeles Lakers"
	var house string
	
	fmt.Print(name)
	fmt.Println(surname)
	fmt.Println(school)
	fmt.Println(house)
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

