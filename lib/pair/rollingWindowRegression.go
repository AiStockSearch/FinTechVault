package main

import (
    "fmt"
    "gonum.org/v1/gonum/mat"
)

func rollingWindowRegression(arr1, arr2 []float64, windowSize int) ([]float64, []float64) {
    alphaValues := make([]float64, 0)
    betaValues := make([]float64, 0)

    for i := windowSize; i < len(arr1); i++ {
        x := mat.NewDense(windowSize, 2, nil)
        y := mat.NewDense(windowSize, 1, nil)

        for j := 0; j < windowSize; j++ {
            x.Set(j, 0, 1) // Константа
            x.Set(j, 1, arr1[i-windowSize+j])
            y.Set(j, 0, arr2[i-windowSize+j])
        }

        qr := new(mat.QR)
        qr.Factorize(x)
        beta := mat.NewVecDense(2, nil)
        qr.Solve(beta, false, y.ColView(0))

        alphaValues = append(alphaValues, beta.AtVec(0))
        betaValues = append(betaValues, beta.AtVec(1))
    }

    return alphaValues, betaValues
}

// func main() {
//     arr1 := []float64{100, 105, 110, 115, 120, 125, 130}
//     arr2 := []float64{200, 210, 220, 230, 240, 250, 260}

//     windowSize := 3
//     alpha, beta := rollingWindowRegression(arr1, arr2, windowSize)

//     fmt.Println("Alpha Values:", alpha)
//     fmt.Println("Beta Values:", beta)
// }