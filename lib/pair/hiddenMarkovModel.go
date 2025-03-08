// package main

// import (
//     "fmt"
//     "github.com/jbrukh/hmm"
// )

// func hiddenMarkovModel(spread []float64) ([]int, error) {
//     // Преобразуем spread в дискретные состояния
//     discreteSpread := make([]int, len(spread))
//     for i, val := range spread {
//         if val > 0 {
//             discreteSpread[i] = 1 // Бычье состояние
//         } else {
//             discreteSpread[i] = 0 // Медвежье состояние
//         }
//     }

//     // Создаем HMM модель
//     model := hmm.New(hmm.Config{
//         States:     2, // Два состояния: бычье и медвежье
//         Observables: 2,
//     })

//     // Обучаем модель
//     err := model.Train(discreteSpread)
//     if err != nil {
//         return nil, err
//     }

//     // Предсказываем скрытые состояния
//     states, err := model.Predict(discreteSpread)
//     if err != nil {
//         return nil, err
//     }

//     return states, nil
// }

// // func main() {
// //     spread := []float64{1.2, 1.3, 1.1, 1.0, 1.4, 1.5}

// //     states, err := hiddenMarkovModel(spread)
// //     if err != nil {
// //         fmt.Println("Ошибка:", err)
// //         return
// //     }

// //     fmt.Println("Скрытые состояния:", states)
// // }