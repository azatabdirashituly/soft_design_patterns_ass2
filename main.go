package main

import "fmt"

func main() {
	audi, _ := getCar("Audi S5")
	merc, _ := getCar("Mercedes GT")
	PrintDetails(audi)
	PrintDetails(merc)
	newAudi := &AudiS5ultra{
		newAudi: audi,
	}
	fmt.Printf("Price of new Model of Audi S5 which is Audi S5 ultra is: %v", newAudi.getPrice())
}

func PrintDetails(c ICar) {
	fmt.Printf("Model: %v, Color: %v, Price: %vM dollars \n", c.GetModel(), c.GetColor(), c.GetPrice())
}
