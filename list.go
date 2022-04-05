package main

//Declare Interface
type printer interface {
	printDetails()
}

type list []printer

func (l list) printList() {
	for _, item := range l {
		item.printDetails()
	}
}

//create methods for print
