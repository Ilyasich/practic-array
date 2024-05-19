package index

import "fmt"

// Функция для вывода индексов элементов в заданном диапазоне
func PrintIndicesInRange(arr []int, low int, high int) {
    for i, num := range arr {
        if num >= low && num <= high {
            fmt.Println("Индекс:", i, "Значение:", num)
        }
    }
}