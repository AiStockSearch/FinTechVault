package main

import (
    "fmt"
    "math"
)

func adfTest(series []float64) bool {
    // Простой ADF тест (упрощенный вариант)
    mean := 0.0
    for _, val := range series {
        mean += val
    }
    mean /= float64(len(series))

    var sumSquared float64
    for _, val := range series {
        sumSquared += math.Pow(val-mean, 2)
    }
    stdDev := math.Sqrt(sumSquared / float64(len(series)))

    // Если стандартное отклонение близко к нулю, считаем стационарным
    if stdDev < 0.01 {
        return true
    }
    return false
}

func cointegrationTest(arr1, arr2 []float64) bool {
    // Рассчитываем spread
    spread := make([]float64, len(arr1))
    for i := range arr1 {
        spread[i] = arr2[i] - arr1[i]
    }

    // Проверяем стационарность spread
    return adfTest(spread)
}

// func main() {
//     arr1 := []float64{100, 105, 110, 115, 120}
//     arr2 := []float64{200, 210, 220, 230, 240}

//     isCointegrated := cointegrationTest(arr1, arr2)
//     fmt.Println("Are the assets cointegrated?", isCointegrated)
// }