package main

import (
    "fmt"
    "github.com/dana-team/go-r"
)

func vectorAutoregression(arr1, arr2 []float64) ([]float64, error) {
    rScript := `
    library(vars)
    data <- cbind(as.numeric(Arr1), as.numeric(Arr2))
    var_model <- VAR(data, p=2)
    predictions <- predict(var_model)$fcst<button class="citation-flag" data-index="1">[, "mean"]
    predictions
    `

    r.Set("Arr1", arr1)
    r.Set("Arr2", arr2)

    result, err := r.Eval(rScript)
    if err != nil {
        return nil, err
    }

    return result.([]float64), nil
}

// func main() {
//     arr1 := []float64{100, 105, 110, 115, 120}
//     arr2 := []float64{200, 210, 220, 230, 240}

//     predictions, err := vectorAutoregression(arr1, arr2)
//     if err != nil {
//         fmt.Println("Ошибка:", err)
//         return
//     }

//     fmt.Println("VAR Predictions:", predictions)
// }