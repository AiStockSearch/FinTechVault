// package main

// import (
//     "fmt"
//     "github.com/dana-team/go-r"
// )

// func johansenTest(arr1, arr2 []float64) bool {
//     rScript := `
//     library(urca)
//     data <- cbind(as.numeric(Arr1), as.numeric(Arr2))
//     jtest <- ca.jo(data, type="eigen", ecdet="none", K=2)
//     p_value <- summary(jtest@teststat[2, 1]) < summary(jtest@cval[2, 1])
//     p_value
//     `

//     r.Set("Arr1", arr1)
//     r.Set("Arr2", arr2)

//     result, err := r.Eval(rScript)
//     if err != nil {
//         fmt.Println("Ошибка выполнения R-скрипта:", err)
//         return false
//     }

//     return result.(bool)
// }

// // func main() {
// //     arr1 := []float64{100, 105, 110, 115, 120}
// //     arr2 := []float64{200, 210, 220, 230, 240}

// //     isCointegrated := johansenTest(arr1, arr2)
// //     fmt.Println("Are the assets cointegrated?", isCointegrated)
// // }