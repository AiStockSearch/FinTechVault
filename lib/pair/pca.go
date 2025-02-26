package main

import (
    "fmt"
    "gonum.org/v1/gonum/mat"
)

func pca(arr1, arr2 []float64) (*mat.Dense, error) {
    // Создаем матрицу данных
    data := mat.NewDense(len(arr1), 2, nil)
    for i := range arr1 {
        data.Set(i, 0, arr1[i])
        data.Set(i, 1, arr2[i])
    }

    // Вычисляем ковариационную матрицу
    cov := mat.Dense{}
    cov.Mul(data.T(), data)

    // Вычисляем собственные значения и векторы
    ev := mat.Eigen(cov.T())
    vectors := ev.Vectors

    return vectors, nil
}

// func main() {
//     arr1 := []float64{100, 105, 110, 115, 120}
//     arr2 := []float64{200, 210, 220, 230, 240}

//     vectors, err := pca(arr1, arr2)
//     if err != nil {
//         fmt.Println("Ошибка:", err)
//         return
//     }

//     fmt.Println("Principal Components:")
//     vectors.Formatted(&vectors)
// }