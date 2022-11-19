package main

type Nota float64

func (n Nota) Entre(inicio, fim float64) bool {
	return float64(n) >= inicio && float64(n) <= fim
}

func NotaParaConceito(n Nota) string {
	if n.Entre(9, 10) {
		return "A"
	} else if n.Entre(7, 8.99) {
		return "B"
	} else if n.Entre(5, 6.99) {
		return "C"
	} else if n.Entre(3, 4.99) {
		return "D"
	} else {
		return "E"
	}
}

func main() {
	var nota = Nota(9.3)
	println(nota.Entre(7, 10))
	println(nota.Entre(0, 6.9))
}
