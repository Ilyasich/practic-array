package sort

import (
	"fmt"
	"sort"
)

// Вывести в порядке возрастания цифры из которых состоит число
func PrintDigitsInOrder(num int) {
	digits := make([]int, 0, 10)

	for num != 0 {
		digits = append(digits, num%10)
		num /= 10
	}

	sort.Ints(digits)

	for _, digit := range digits {
		fmt.Print(digit)
	}
	fmt.Println()
}
