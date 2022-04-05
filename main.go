package main

func main() {
	//Declare objects
	mineCraft := game{title: "Minecraft", price: 5}
	worldofWarcraft := game{title: "World of Warcraft", price: 19}
	eliteDangerous := game{title: "Elite Dangerous", price: 54}
	candleInTheTomb := books{title: "Candle in the tomb", price: 20}
	barneyAndFriends := books{title: "Barney and Friends", price: 10}
	razerBtEarpiece := ComAccessories{title: "Razer BT earpiece", price: 159}
	razerKeyboard := ComAccessories{title: "Razer keyboard", price: 110}
	logiTechMouse := ComAccessories{title: "Logitech Mouse", price: 80}
	//Insert codes here

	var store list
	store = append(store, &mineCraft, &worldofWarcraft, &eliteDangerous, &candleInTheTomb, &barneyAndFriends, &razerBtEarpiece, &razerKeyboard, &logiTechMouse)

	store.printList()

}
