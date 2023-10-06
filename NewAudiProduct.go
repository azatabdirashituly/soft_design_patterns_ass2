package main

type AudiS5ultra struct {
	newAudi ICar
}

func (a *AudiS5ultra) getPrice() float64 {
	price := a.newAudi.GetPrice()
	return price + 4
}
