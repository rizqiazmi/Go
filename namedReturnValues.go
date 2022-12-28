package main

import "fmt"

func getCompleteName() (firstName string, middleName string, lastName string) {
	firstName = "Eko"
	middleName = "Kurniawan"
	lastName = "Khannedy"
	return
}

func main() {
	a, b, c := getCompleteName()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
