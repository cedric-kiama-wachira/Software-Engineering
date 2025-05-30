package main

import "fmt"

const fullnames string = "Cedric K. Wachira" // Constant declaration and assigning a fixed value.
var birthyear int = 1982 // Global variable declared and will be accessible to other blocks of code

func main(){
	var age int = 2025 - 1982 // Local variable within the main function 
	fmt.Println(age) // Using another function inside main to get my age.
}
