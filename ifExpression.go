package main

import "fmt"

func main() {
	name := "Kurniawan"

	if name == "Eko" {
		fmt.Println("Hello Eko")

	} else if name == "Joko" {
		fmt.Println("Hello Joko")

	} else {
		fmt.Print("Hi, Boleh Kenalan ? ")
	}

	if length := len(name); length > 5 {
		fmt.Println("\nTerlalu Panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
