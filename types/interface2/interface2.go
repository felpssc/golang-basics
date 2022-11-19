package main

type Sport interface {
	TurnTurboOn()
}

type Ferrari struct {
	Model       string
	IsTurboOn   bool
	ActualSpeed int
}

func (f *Ferrari) TurnTurboOn() {
	f.IsTurboOn = true
}

func main() {
	car1 := Ferrari{"F40", false, 0}
	car1.TurnTurboOn()

	var sport1 Sport = &Ferrari{
		Model:       "F40",
		IsTurboOn:   false,
		ActualSpeed: 0,
	}

	sport1.TurnTurboOn()
}
