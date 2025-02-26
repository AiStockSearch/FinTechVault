package main

import (
    "fmt"
)

func grangerCausality(arr1, arr2 []float64) bool {
    // Простой пример: если arr1 предсказывает arr2, то считаем, что есть причинно-следственная связь
    correlation := 0.0
    for i := 1; i < len(arr1); i++ {
        correlation += arr1[i-1] * arr2[i]
    }
    return correlation > 0
}

// func main() {
//     arr1 := []float64{100, 105, 110, 115, 120}
//     arr2 := []float64{200, 210, 220, 230, 240}

//     result := grangerCausality(arr1, arr2)
//     fmt.Println("Granger Causality:", result)
// }