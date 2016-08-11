package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(calculateEntropy("Hello,playground"))
}

func calculateEntropy(input string) float64 {
	characterCounts := make(map[byte]int, 255)
	for char := range input {
		fmt.Println(input[char])
		characterCounts[input[char]] += 1
		//fmt.Println(characterCounts[char])

	}
	fmt.Println(characterCounts)
	characterFrequencies := make(map[byte]float64, 255)
	for key := range characterCounts {
		if characterCounts[key] > 0 {
			fmt.Println(characterCounts[key], len(input))
			characterFrequencies[key] = float64(characterCounts[key]) / float64(len(input))
			fmt.Println(characterFrequencies[key])
		}
	}

	entropyMeasure := float64(0)
	for key := range characterFrequencies {
		entropyValue := characterFrequencies[key]
		fmt.Println(entropyValue * math.Ln2 * entropyValue)
		entropyMeasure += float64(entropyValue * math.Ln2 * entropyValue)
	}
	entropyMeasure = entropyMeasure * -1

	return entropyMeasure
}
