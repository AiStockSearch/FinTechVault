package main

import (
    "fmt"
    "math"
)

func spectralAnalysis(spread []float64) []float64 {
    // Простой пример FFT
    n := len(spread)
    spectrum := make([]float64, n/2)
    for k := 0; k < n/2; k++ {
        real, imag := 0.0, 0.0
        for t := 0; t < n; t++ {
            angle := 2 * math.Pi * float64(k) * float64(t) / float64(n)
            real += spread[t] * math.Cos(angle)
            imag -= spread[t] * math.Sin(angle)
        }
        spectrum[k] = math.Sqrt(real*real + imag*imag)
    }
    return spectrum
}

// func main() {
//     spread := []float64{1.2, 1.3, 1.1, 1.0, 1.4, 1.5}

//     spectrum := spectralAnalysis(spread)
//     fmt.Println("Spectral Analysis:", spectrum)
// }