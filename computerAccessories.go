package main

import "fmt"

//Declare Computer Accessoies Struct
type ComAccessories struct {
	title string
	price int
}

//Create print method
func (c *ComAccessories) printDetails() {
	fmt.Printf("Item name: %v - Price: $%v\n", c.title, c.price)
}
