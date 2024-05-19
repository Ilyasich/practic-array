package mix

import (
	
	"math/rand"
	"time"
)

// функция для перемешивания массива
func ShuffleArray(arr []int) []int {
	rand.NewSource(time.Now().UnixNano()) // инициализация генератора случайных чисел
	for i := range arr {
		j := rand.Intn(i + 1)
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}