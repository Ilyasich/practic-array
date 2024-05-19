package maxsize


//Вывести размер самой длинной подпоследовательности состоящей из одинаковых элементов
func MaxSubsequenceLength(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	maxLength := 1
	currentLength := 1

	for i := 1; i < len(arr); i++ {
		if arr[i] == arr[i-1] {
			currentLength++
		} else {
			if currentLength > maxLength {
				maxLength = currentLength
			}
			currentLength = 1
		}
	}

	if currentLength > maxLength {
		maxLength = currentLength
	}

	return maxLength
}
