package main

import (
    "fmt"
)

func crossCorrelation(arr1, arr2 []float64, lag int) float64 {
    if len(arr1) != len(arr2) {
        return 0
    }

    n := len(arr1)
    if n < lag {
        return 0
    }

    sum := 0.0
    for i := lag; i < n; i++ {
        sum += arr1[i-lag] * arr2[i]
    }

    return sum / float64(n-lag)
}

// func main() {
//     arr1 := []float64{100, 105, 110, 115, 120}
//     arr2 := []float64{200, 210, 220, 230, 240}

//     lag := 1
//     result := crossCorrelation(arr1, arr2, lag)
//     fmt.Printf("Cross-Correlation (lag=%d): %.2f\n", lag, result)
// }