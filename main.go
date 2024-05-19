package main

import (
	"fmt"
	"practic-array/days"
	"practic-array/dell"
	"practic-array/index"
	"practic-array/maxsize"
	"practic-array/merger"
	"practic-array/mix"
	"practic-array/positive"
	"practic-array/sort"
)

func main() {
	//пример исользования функции поиска положительных элементов
	array := []int{-2, 3, 0, 5, -1, 6, 7}
	result := positive.CountPositiveNumbers(array)
	fmt.Printf("Количество положительных элементов в массиве: %d\n", result)

	// Пример использования функции для вывода индексов элементов в заданном диапазоне
	numbers := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	index.PrintIndicesInRange(numbers, 30, 70)

	//пример использования функции возрастания цифры из которых состоит число
	sort.PrintDigitsInOrder(1234567890)

	//пример использования функции размера самой длинной подпоследовательности состоящей из одинаковых элементов
	arr := []int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4}
	fmt.Println(maxsize.MaxSubsequenceLength(arr))

	//пример использования функции удаления повторяющиеся элементы из массива
	nums := []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5}
	fmt.Println(dell.RemoveDuplicates(nums))

	//пример использования функции слияние двух упорядоченных массивов
	arr1 := []int{1, 3, 5, 7, 9}
	arr2 := []int{2, 4, 6, 8, 10}
	merged := merger.MergeSortedArrays(arr1, arr2)
	fmt.Println(merged)

	// Пример использования функции високосный год
	day := 24
	month := 3
	fmt.Println("Номер дня с начала года:", days.DayOfYear(day, month))

	// Пример использования функции для перемешивания массива
	number := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	shuffled := mix.ShuffleArray(number)
	fmt.Println(shuffled)

}
