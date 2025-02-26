package main

import (
    "fmt"
    "math"
)

func meanReversionWithBands(spread []float64, multiplier float64) ([]float64, []float64) {
    mean := 0.0
    for _, val := range spread {
        mean += val
    }
    mean /= float64(len(spread))

    var sumSquared float64
    for _, val := range spread {
        sumSquared += math.Pow(val-mean, 2)
    }
    stdDev := math.Sqrt(sumSquared / float64(len(spread)))

    upperBand := make([]float64, len(spread))
    lowerBand := make([]float64, len(spread))
    for i := range spread {
        upperBand[i] = mean + multiplier*stdDev
        lowerBand[i] = mean - multiplier*stdDev
    }

    return upperBand, lowerBand
}

// func main() {
//     spread := []float64{1.2, 1.3, 1.1, 1.0, 1.4, 1.5}

//     upper, lower := meanReversionWithBands(spread, 2)
//     fmt.Println("Upper Band:", upper)
//     fmt.Println("Lower Band:", lower)
// }