package main

import (
    "fmt"
)

func kalmanFilter(arr1, arr2 []float64) []float64 {
    // Упрощенный Kalman Filter
    var stateMean float64 = 0
    var uncertainty float64 = 1

    kalmanMeans := make([]float64, len(arr1))
    for i := range arr1 {
        prediction := stateMean
        residual := arr2[i] - prediction

        gain := uncertainty / (uncertainty + 1)
        stateMean += gain * residual
        uncertainty *= (1 - gain)

        kalmanMeans[i] = stateMean
    }

    return kalmanMeans
}

// func main() {
//     arr1 := []float64{100, 105, 110, 115, 120}
//     arr2 := []float64{200, 210, 220, 230, 240}

//     kalmanMeans := kalmanFilter(arr1, arr2)
//     fmt.Println("Kalman Means:", kalmanMeans)
// }