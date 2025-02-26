package main

import (
    "fmt"
    "gonum.org/v1/gonum/mat"
)

func linearRegression(arr1, arr2 []float64) (float64, float64, []float64) {
    // Создаем матрицы для линейной регрессии
    x := mat.NewDense(len(arr1), 2, nil)
    y := mat.NewDense(len(arr2), 1, arr2)

    // Заполняем матрицу X (константа + arr1)
    for i := range arr1 {
        x.Set(i, 0, 1) // Константа
        x.Set(i, 1, arr1[i])
    }

    // Решаем уравнение Y = X * Beta
    qr := new(mat.QR)
    qr.Factorize(x)
    beta := mat.NewVecDense(2, nil)
    qr.Solve(beta, false, y.ColView(0))

    // Извлекаем коэффициенты
    alpha := beta.AtVec(0)
    betaCoeff := beta.AtVec(1)

    // Рассчитываем spread
    spread := make([]float64, len(arr1))
    for i := range arr1 {
        spread[i] = arr2[i] - (alpha + betaCoeff*arr1[i])
    }

    return alpha, betaCoeff, spread
}

// func main() {
//     arr1 := []float64{100, 105, 110, 115, 120}
//     arr2 := []float64{200, 210, 220, 230, 240}

//     alpha, beta, spread := linearRegression(arr1, arr2)
//     fmt.Printf("Alpha: %.2f, Beta: %.2f\n", alpha, beta)
//     fmt.Println("Spread:", spread)
// }