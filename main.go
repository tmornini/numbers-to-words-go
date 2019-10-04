package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		value, err := strconv.Atoi(scanner.Text())

		if err != nil {
			panic(err)
		}

		fmt.Println(
			integerToWords(value),
		)
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}
}

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

func portionToWords(value int) []string {
	parts := []string{}

	hundreds := value / 100

	if hundreds > 0 {
		parts = append(parts, onesWords[hundreds], "hundred")
	}

	value -= (hundreds * 100)

	tens := value / 10

	if tens > 1 {
		parts = append(parts, tensWords[tens])
		value -= tens * 10
	}

	if value > 0 {
		parts = append(parts, onesWords[value])
	}

	return parts
}

func toThePowerOf(i, j int) int {
	return int(math.Pow(float64(i), float64(j)))
}

var onesWords = []string{
	"zero",
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
	"ten",
	"eleven",
	"twelve",
	"thirteen",
	"fourteen",
	"fifteen",
	"sixteen",
	"seventeen",
	"eighteen",
	"nineteen",
}

var tensWords = []string{
	"",
	"",
	"twenty",
	"thirty",
	"forty",
	"fifty",
	"sixty",
	"seventy",
	"eighty",
	"ninety",
}

var longScaleWords = []string{
	"",
	"",
	"",
	"thousand",
	"",
	"",
	"million",
	"",
	"",
	"billion",
	"",
	"",
	"trillion",
	"",
	"",
	"quadrillion",
}
