// package main

// import (
//     "fmt"
//     "github.com/montanaflynn/go-wavelet"
// )

// func waveletTransform(spread []float64) ([]float64, error) {
//     // Выполняем wavelet transform
//     result, err := wavelet.Transform(spread, wavelet.Haar, 1)
//     if err != nil {
//         return nil, err
//     }

//     return result.Approximation, nil
// }

// // func main() {
// //     spread := []float64{1.2, 1.3, 1.1, 1.0, 1.4, 1.5}

// //     waveletResult, err := waveletTransform(spread)
// //     if err != nil {
// //         fmt.Println("Ошибка:", err)
// //         return
// //     }

// //     fmt.Println("Wavelet Result:", waveletResult)
// // }