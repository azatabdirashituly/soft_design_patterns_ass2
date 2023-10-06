package main

type Audi struct {
	Car
}

func newAudi() ICar {
	return &Audi{
		Car: Car{
			model: "Audi S5",
			color: "Black",
			price: 3.5,
		},
	}
}
