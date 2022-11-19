package main

type Sport interface {
	TurnTurboOn()
}

type Luxury interface {
	TurnOnAutoPilot()
}

type SportLuxury interface {
	Sport
	Luxury
}

type BMW7 struct{}

func (b *BMW7) TurnTurboOn() {
	println("Turning turbo on...")
}

func (b *BMW7) TurnOnAutoPilot() {
	println("Turning auto pilot on...")
}

func main() {
	var b SportLuxury = &BMW7{}
	b.TurnTurboOn()
	b.TurnOnAutoPilot()
}
