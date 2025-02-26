package main

import (
    "fmt"
    "math/rand"
)

func monteCarloSimulation(spread []float64, simulations int) ([]float64, error) {
    // Расчет среднего и стандартного отклонения
    mean := 0.0
    for _, val := range spread {
        mean += val
    }
    mean /= float64(len(spread))

    var sumSquared float64
    for _, val := range spread {
        sumSquared += (val - mean) * (val - mean)
    }
    stdDev := math.Sqrt(sumSquared / float64(len(spread)))

    // Генерация сценариев
    scenarios := make([]float64, simulations)
    for i := range scenarios {
        scenarios[i] = rand.NormFloat64()*stdDev + mean
    }

    return scenarios, nil
}

// func main() {
//     spread := []float64{1.2, 1.3, 1.1, 1.0, 1.4, 1.5}

//     scenarios, err := monteCarloSimulation(spread, 10)
//     if err != nil {
//         fmt.Println("Ошибка:", err)
//         return
//     }

//     fmt.Println("Scenarios:", scenarios)
// }