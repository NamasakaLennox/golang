package main

import "fmt"

func main() {
	var age int
	var name, gender string

	age = 10
	name = "Lennox"
	gender = "Male"

	account := true
	lastName := "Namasaka"

	fmt.Println(age, name, gender, account, lastName)
	printname()

}

func printname() {
	firstName := "Mike"
	fmt.Println(firstName)
}
