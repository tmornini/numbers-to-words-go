package main

func portionToWords(value int) []string {
	parts := []string{}

	hundreds := value / 100

	if hundreds > 0 {
		parts = append(parts, lowWords[hundreds], "hundred")
	}

	value -= (hundreds * 100)

	tens := value / 10

	if tens > 1 {
		parts = append(parts, tensWords[tens])
		value -= tens * 10
	}

	if value > 0 {
		parts = append(parts, lowWords[value])
	}

	return parts
}
