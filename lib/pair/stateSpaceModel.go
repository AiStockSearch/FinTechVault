package main

import (
    "fmt"
    "gonum.org/v1/gonum/mat"
)

func stateSpaceModel(arr1, arr2 []float64) (*mat.Dense, error) {
    // Простой пример State Space Model
    data := mat.NewDense(len(arr1), 2, nil)
    for i := range arr1 {
        data.Set(i, 0, arr1[i])
        data.Set(i, 1, arr2[i])
    }

    // Вычисляем скрытое состояние (упрощенно)
    state := mat.NewDense(len(arr1), 1, nil)
    for i := range arr1 {
        state.Set(i, 0, (arr1[i]+arr2[i])/2)
    }

    return state, nil
}

// func main() {
//     arr1 := []float64{100, 105, 110, 115, 120}
//     arr2 := []float64{200, 210, 220, 230, 240}

//     state, err := stateSpaceModel(arr1, arr2)
//     if err != nil {
//         fmt.Println("Ошибка:", err)
//         return
//     }

//     fmt.Println("State Space Model States:")
//     fmt.Println(state)
// }