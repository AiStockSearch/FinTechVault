package main

import (
    "fmt"
    "math"
)

func bayesianInference(arr1, arr2 []float64) float64 {
    // Простой пример Bayesian Inference для коэффициента корреляции
    mean := 0.0
    for i := range arr1 {
        mean += arr1[i] * arr2[i]
    }
    mean /= float64(len(arr1))

    var sumSquared float64
    for i := range arr1 {
        sumSquared += math.Pow(arr1[i]*arr2[i]-mean, 2)
    }
    stdDev := math.Sqrt(sumSquared / float64(len(arr1)))

    return mean / stdDev
}

// func main() {
//     arr1 := []float64{100, 105, 110, 115, 120}
//     arr2 := []float64{200, 210, 220, 230, 240}

//     result := bayesianInference(arr1, arr2)
//     fmt.Println("Bayesian Inference Result:", result)
// }