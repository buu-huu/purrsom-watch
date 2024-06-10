package utility

import (
	"math"
)

func Entropy(data []byte) float64 {
	freq := make(map[byte]int)
	for _, b := range data {
		freq[b]++
	}

	entropy := 0.0
	total := float64(len(data))
	for _, count := range freq {
		probability := float64(count) / total
		entropy -= probability * math.Log2(probability)
	}
	return entropy
}
