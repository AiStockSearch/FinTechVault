package main

import (
    "fmt"
    "math"
)

func copulaBasedPairTrading(arr1, arr2 []float64) float64 {
    // Простая корреляция как пример Copula
    mean1, mean2 := 0.0, 0.0
    for i := range arr1 {
        mean1 += arr1[i]
        mean2 += arr2[i]
    }
    mean1 /= float64(len(arr1))
    mean2 /= float64(len(arr2))

    var cov, std1, std2 float64
    for i := range arr1 {
        cov += (arr1[i] - mean1) * (arr2[i] - mean2)
        std1 += math.Pow(arr1[i]-mean1, 2)
        std2 += math.Pow(arr2[i]-mean2, 2)
    }
    cov /= float64(len(arr1))
    std1 = math.Sqrt(std1 / float64(len(arr1)))
    std2 = math.Sqrt(std2 / float64(len(arr2)))

    return cov / (std1 * std2)
}

// func main() {
//     arr1 := []float64{100, 105, 110, 115, 120}
//     arr2 := []float64{200, 210, 220, 230, 240}

//     correlation := copulaBasedPairTrading(arr1, arr2)
//     fmt.Println("Copula Correlation:", correlation)
// }