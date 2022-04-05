package main

import "fmt"

//Declare Games Struct
type game struct {
	title string
	price int
}

//Create print method
func (g *game) printDetails() {
	fmt.Printf("Game Title: %v - Price: $%v\n", g.title, g.price)
}
