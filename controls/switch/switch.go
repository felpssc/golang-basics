package main

import "fmt"

func options(option int) string {

	selectedOption := ""

	switch option {
	case 1:
		selectedOption = "VOL +"
	case 2:
		selectedOption = "VOL -"
	case 3, 4:
		selectedOption = "ON/OFF"
	}

	return selectedOption
}

func main() {
	option := options(2)

	fmt.Println(option)
}
