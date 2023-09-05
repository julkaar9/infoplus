package utils

import (
	"fmt"
	"math"
)

func HumanReadable(num float64) string {
	suffix := []string{"", "K", "M", "B", "T", "Q"}
	magnitude := 0
	if math.Abs(num) >= 10000 {
		for math.Abs(num) >= 1000 {
			magnitude += 1
			num /= 1000
		}
	}
	return fmt.Sprintf("%.f%s", num, suffix[magnitude])
}
