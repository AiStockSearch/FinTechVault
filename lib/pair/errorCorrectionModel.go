package main

import (
    "fmt"
    "gonum.org/v1/gonum/mat"
)

func errorCorrectionModel(arr1, arr2 []float64) ([]float64, error) {
    // Создаем матрицы для регрессии
    x := mat.NewDense(len(arr1)-1, 2, nil)
    y := mat.NewDense(len(arr2)-1, 1, nil)

    // Заполняем матрицу X (изменение arr1 + spread в предыдущий момент)
    for i := 1; i < len(arr1); i++ {
        x.Set(i-1, 0, arr1[i]-arr1[i-1]) // Изменение arr1
        x.Set(i-1, 1, arr2[i-1]-(arr1[i-1]*0.5)) // Spread (упрощенно)
        y.Set(i-1, 0, arr2[i]-arr2[i-1]) // Изменение arr2
    }

    // Решаем уравнение Y = X * Beta
    qr := new(mat.QR)
    qr.Factorize(x)
    beta := mat.NewVecDense(2, nil)
    qr.Solve(beta, false, y.ColView(0))

    // Возвращаем коэффициенты
    coefficients := make([]float64, 2)
    coefficients[0] = beta.AtVec(0)
    coefficients[1] = beta.AtVec(1)

    return coefficients, nil
}

// func main() {
//     arr1 := []float64{100, 105, 110, 115, 120}
//     arr2 := []float64{200, 210, 220, 230, 240}

//     coefficients, err := errorCorrectionModel(arr1, arr2)
//     if err != nil {
//         fmt.Println("Ошибка:", err)
//         return
//     }

//     fmt.Println("Коэффициенты ECM:", coefficients)
// }