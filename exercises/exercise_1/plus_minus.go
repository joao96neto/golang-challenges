package plusminus

import (
	"fmt"
	"math"
	"math/big"
)

var Input = []int32{-4, 3, -9, 0, 4, 1}
var positive = func(a int32) bool {
	if a > 0 {
		return true
	}
	return false
}
var negative = func(a int32) bool {
	if a < 0 {
		return true
	}
	return false
}
var zero = func(a int32) bool {
	if a == 0 {
		return true
	}
	return false
}

// Execute desc: Dada uma lista de inteiros, calcula se a quantidade de valores, positivos, negativos e neutros,
// depois exibe a porcentagem de cada um comparado a quantidade total.
func Execute(arr []int32) {
	resultPositive := getQuantityByFunc(arr, positive)
	resultNegative := getQuantityByFunc(arr, negative)
	resultZero := getQuantityByFunc(arr, zero)
	totalQuantityFloat := big.NewFloat(float64(len(arr)))

	calculatePercentage(totalQuantityFloat, big.NewFloat(resultNegative))

	fmt.Println(resultPositive, resultNegative, resultZero)
}

func getQuantityByFunc(arr []int32, fn func(int322 int32) bool) float64 {
	quantity := 0
	for _, val := range arr {
		if fn(val) {
			quantity++
		}
	}
	return float64(quantity)
}
func calculatePercentage(total, quantity *big.Float) {
	ratio := math.Pow(10, float64(6))
	result, _ := new(big.Float).Quo(quantity, total).Float64()
	fmt.Println(math.Round(result*ratio) / ratio)
}
