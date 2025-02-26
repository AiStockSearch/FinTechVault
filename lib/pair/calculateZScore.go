package main

import (
    "fmt"
    "math"
)

func calculateZScore(spread []float64) []float64 {
    // Рассчитываем среднее и стандартное отклонение
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

    // Рассчитываем Z-score
    zScores := make([]float64, len(spread))
    for i, val := range spread {
        zScores[i] = (val - mean) / stdDev
    }

    return zScores
}

// func main() {
//     arr1 := []float64{100, 105, 110, 115, 120}
//     arr2 := []float64{200, 210, 220, 230, 240}

//     // Рассчитываем spread
//     spread := make([]float64, len(arr1))
//     for i := range arr1 {
//         spread[i] = arr2[i] - arr1[i]
//     }

//     // Рассчитываем Z-score
//     zScores := calculateZScore(spread)
//     fmt.Println("Z-scores:", zScores)
// }