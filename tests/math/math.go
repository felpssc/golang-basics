package math

import (
	"fmt"
	"strconv"
)

// Media is responsible for calculate... well, you know what... :)
func Media(numbers ...float64) float64 {
	total := 0.0
	for _, num := range numbers {
		total += num
	}

	media := total / float64(len(numbers))

	roundedMedia, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", media), 64)

	return roundedMedia
}
