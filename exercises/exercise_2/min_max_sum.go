package minmaxsum

import (
	"fmt"
	"math/big"
	"sort"
)

var Input = []int32{256741038, 623958417, 467905213, 714532089, 938071625}

// Execute : Dado uma lista de inteiros contendo 5 items, calculamos a soma minima e maxima de 4 itens da lista
func Execute(arr []int32) {
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
	var and12 = new(big.Int).Add(big.NewInt(int64(arr[2])), big.NewInt(int64(arr[1])))
	var and123 = new(big.Int).Add(and12, big.NewInt(int64(arr[3])))

	var max = new(big.Int).Add(and123, big.NewInt(int64(arr[4])))
	var min = new(big.Int).Add(and123, big.NewInt(int64(arr[0])))
	fmt.Println(min.Int64(), max.Int64())
}
