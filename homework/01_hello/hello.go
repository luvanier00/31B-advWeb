package main

import (
	"fmt"
	"time"
)

func main() {
	var name string
	var year int
	var yearActual int
	var ageActual int

	fmt.Println("Hello, <enter your name>")
	fmt.Scanln(&name)
	fmt.Println("What year where you born <year>")
	fmt.Scanln(&year)

	//fmt.Println(name)
	//fmt.Println(year)

	yearActual = time.Now().Year()
	ageActual = yearActual - year

	fmt.Printf("Hello, %v! You are currently %v years old.", name, ageActual)

}
