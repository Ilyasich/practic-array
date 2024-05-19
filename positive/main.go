package positive

//Количество положительных элементов массива
func CountPositiveNumbers(arr []int) int {
    count := 0
    for _, value := range arr {
        if value > 0 {
            count++
        }
    }
    return count
}

/* В этом коде:
1. Функция `countPositiveNumbers` принимает на вход массив целых чисел `arr`.
2. В цикле `for range` проходим по каждому элементу массива.
3. Если элемент `value` больше нуля (`value > 0`), увеличиваем счетчик `count`.
4. Функция возвращает количество положительных элементов.
5. В функции `main` создается массив и вызывается функция `countPositiveNumbers` для подсчета положительных элементов. Результат выводится на экран.  */
