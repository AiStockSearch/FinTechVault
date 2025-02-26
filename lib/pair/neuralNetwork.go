package main

import (
    "fmt"
    "github.com/ziutek/go-deep"
)

func neuralNetwork(spread []float64) ([]float64, error) {
    // Создаем нейронную сеть
    net := deep.NewNet(
        deep.Layer{Neurons: 1, Activation: deep.Linear}, // Входной слой
        deep.Layer{Neurons: 5, Activation: deep.ReLU},   // Скрытый слой
        deep.Layer{Neurons: 1, Activation: deep.Linear}, // Выходной слой
    )

    // Подготавливаем данные для обучения
    inputs := make([][]float64, len(spread)-1)
    outputs := make([][]float64, len(spread)-1)
    for i := 1; i < len(spread); i++ {
        inputs[i-1] = []float64{spread[i-1]}
        outputs[i-1] = []float64{spread[i]}
    }

    // Обучаем сеть
    err := net.Train(inputs, outputs, deep.MSE, 0.01, 100)
    if err != nil {
        return nil, err
    }

    // Предсказываем следующее значение spread
    prediction, _ := net.Forward([]float64{spread[len(spread)-1]})
    return prediction, nil
}

// func main() {
//     spread := []float64{1.2, 1.3, 1.1, 1.0, 1.4, 1.5}

//     prediction, err := neuralNetwork(spread)
//     if err != nil {
//         fmt.Println("Ошибка:", err)
//         return
//     }

//     fmt.Println("Предсказанное значение spread:", prediction)
// }