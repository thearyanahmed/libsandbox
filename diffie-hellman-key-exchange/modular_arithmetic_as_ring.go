package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {
	// Example values for g and a
	g := 5
	a := 3

	// Different values of N to demonstrate modular arithmetic as a "ring"
	moduli := []int{5, 7, 11, 13, 17, 19, 23, 29,} 

	// Calculate g^a using math.Pow for each modulus N
	for _, N := range moduli {
		// Calculate g^a using math.Pow
		result := math.Pow(float64(g), float64(a))

		// Convert the result to big.Int for proper modulus calculation
		resultInt := new(big.Int)
		resultInt.SetInt64(int64(result))

		// Calculate resultInt % N
		modResult := new(big.Int).Mod(resultInt, big.NewInt(int64(N)))

		// Print the result for the given modulus N
		fmt.Printf("g^a mod N = %v^%v mod %v = %d\n", g, a, N, modResult)
	}
}

// output
// g^a mod N = 5^3 mod 13 = 8
// g^a mod N = 5^3 mod 17 = 6
// g^a mod N = 5^3 mod 19 = 11
// g^a mod N = 5^3 mod 23 = 10
// g^a mod N = 5^3 mod 29 = 9
// g^a mod N = 5^3 mod 30 = 5
