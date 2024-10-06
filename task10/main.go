package main

import "fmt"

func MakeTemperatures(slice []float64) map[int][]float64 {
	result := make(map[int][]float64)

	for _, val := range slice {
		group := (int(val) / 10) * 10
		result[group] = append(result[group], val)
	}

	return result
}

func main() {
	slice := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	result := MakeTemperatures(slice)

	fmt.Print(result)
}
