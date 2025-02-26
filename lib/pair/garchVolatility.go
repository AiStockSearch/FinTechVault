package main

import (
    "fmt"
    "math"
)

func garchVolatility(spread []float64) []float64 {
    // Простая GARCH (1,1) модель
    var omega float64 = 0.01
    var alpha float64 = 0.07
    var beta float64 = 0.92

    volatility := make([]float64, len(spread))
    volatility[0] = math.Abs(spread[0])

    for i := 1; i < len(spread); i++ {
        volatility[i] = math.Sqrt(omega + alpha*math.Pow(spread[i-1], 2) + beta*math.Pow(volatility[i-1], 2))
    }

    return volatility
}

// func main() {
//     spread := []float64{1.2, 1.3, 1.1, 1.0, 1.4, 1.5}

//     volatility := garchVolatility(spread)
//     fmt.Println("GARCH Volatility:", volatility)
// }