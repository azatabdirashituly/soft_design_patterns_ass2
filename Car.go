package main

type Car struct {
	model string
	color string
	price float64
}

func (c *Car) SetModel(model string) {
	c.model = model
}
func (c *Car) GetModel() string {
	return c.model
}

func (c *Car) SetColor(color string) {
	c.color = color
}
func (c *Car) GetColor() string {
	return c.color
}

func (c *Car) SetPrice(price float64) {
	c.price = price
}
func (c *Car) GetPrice() float64 {
	return c.price
}
