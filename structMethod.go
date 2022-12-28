package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (a Customer) sayHuuu() {
	fmt.Println("Huuuuuu from", a.Name)
}

func (customer Customer) sayHi(name string) {
	fmt.Println("Hello", name, "My Name is", customer.Name)
}

func main() {
	var eko Customer
	eko.Name = "Eko"
	eko.Address = "Indonesia"
	eko.Age = 38

	eko.sayHi("Joko")
	eko.sayHuuu()
}
