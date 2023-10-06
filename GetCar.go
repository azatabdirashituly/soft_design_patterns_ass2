package main

import "fmt"

func getCar(carModel string) (ICar, error) {
	if carModel == "Audi S5" {
		return newAudi(), nil
	}
	if carModel == "Mercedes GT" {
		return newMercedes(), nil
	}
	return nil, fmt.Errorf("Wrong model of the car")
}
