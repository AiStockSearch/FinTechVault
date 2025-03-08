// package main

// import (
//     "fmt"
//     "github.com/dana-team/go-r"
// )

// func dynamicConditionalCorrelation(arr1, arr2 []float64) float64 {
//     rScript := `
//     library(rmgarch)
//     data <- cbind(as.numeric(Arr1), as.numeric(Arr2))
//     dcc_spec <- dccspec(uspec = multispec(ugarchspec(mean.model = list(armaOrder = c(0,0)), variance.model = list(garchOrder = c(1,1))), distribution.model = "mvnorm"))
//     dcc_fit <- dccfit(dcc_spec, data = data)
//     dcc_cor <- dcc_fit$DCCcoef[1,2]
//     dcc_cor
//     `

//     r.Set("Arr1", arr1)
//     r.Set("Arr2", arr2)

//     result, err := r.Eval(rScript)
//     if err != nil {
//         fmt.Println("Ошибка:", err)
//         return 0
//     }

//     return result.(float64)
// }

// // func main() {
// //     arr1 := []float64{100, 105, 110, 115, 120}
// //     arr2 := []float64{200, 210, 220, 230, 240}

// //     correlation := dynamicConditionalCorrelation(arr1, arr2)
// //     fmt.Println("DCC Correlation:", correlation)
// // }