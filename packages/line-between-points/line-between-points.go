package main

import "math"

// Iniciando com a letra maiúscula é PÚBLICO (visível dentro e fora do pacote)!
// Iniciando com a letra minúscula é PRIVADO (visível apenas dentro do pacote)!

// Por exemplo...
// Ponto - gerará algo público
// ponto ou _Ponto - gerará algo privado

// Point represents an coordinate in a 2D plane
type Point struct {
	X float64
	Y float64
}

func catetos(p1, p2 Point) (cx, cy float64) {
	cx = math.Abs(p1.X - p2.X)
	cy = math.Abs(p1.Y - p2.Y)

	return
}

// Distance is a function to calculate the distance between two points
func Distance(p1, p2 Point) float64 {
	cx, cy := catetos(p1, p2)

	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}
