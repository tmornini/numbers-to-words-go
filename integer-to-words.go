package main

import (
	"strings"
)

func integerToWords(value int) string {
	if value == 0 {
		return "zero"
	}

	parts := []string{}

	for power := 15; power >= 0; power -= 3 {
		portion := value / toThePowerOf(10, power)
		if portion > 0 {
			parts = append(parts, portionToWords(portion)...)
			parts = append(parts, longScaleWords[power])
		}
		value -= portion * toThePowerOf(10, power)
	}

	return strings.Join(parts, " ")
}
