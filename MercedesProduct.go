package main

type Mercedes struct {
	Car
}

func newMercedes() ICar {
	return &Mercedes{
		Car: Car{
			model: "Mercedes GT",
			color: "Red",
			price: 7.0,
		},
	}
}
