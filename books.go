package main

import "fmt"

//Declare books Struct
type books struct {
	title string
	price int
}

func (b *books) printDetails() {
	fmt.Printf("Book Title: %v - Price: $%v\n", b.title, b.price)
}

//Create print method
