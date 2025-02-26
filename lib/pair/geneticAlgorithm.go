package main

import (
    "fmt"
    "math/rand"
)

func geneticAlgorithm(arr1, arr2 []float64, generations, populationSize int) float64 {
    // Простой пример GA для оптимизации коэффициента beta
    population := make([]float64, populationSize)
    for i := range population {
        population[i] = rand.Float64()
    }

    for gen := 0; gen < generations; gen++ {
        newPopulation := make([]float64, populationSize)
        for i := range newPopulation {
            parent1, parent2 := population[rand.Intn(populationSize)], population[rand.Intn(populationSize)]
            child := (parent1 + parent2) / 2 // Кроссинговер
            child += rand.NormFloat64() * 0.1 // Мутация
            newPopulation[i] = child
        }
        population = newPopulation
    }

    return population[0]
}

// func main() {
//     arr1 := []float64{100, 105, 110, 115, 120}
//     arr2 := []float64{200, 210, 220, 230, 240}

//     bestBeta := geneticAlgorithm(arr1, arr2, 100, 50)
//     fmt.Println("Best Beta:", bestBeta)
// }