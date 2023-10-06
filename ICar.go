package main

type ICar interface {
	SetModel(string)
	SetColor(string)
	SetPrice(float64)
	GetModel() string
	GetColor() string
	GetPrice() float64
}
