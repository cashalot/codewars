/*
Write a function that takes an integer as input, and returns the number of bits that are equal to one in the binary representation of that number. You can guarantee that input is non-negative.

Example: The binary representation of 1234 is 10011010010, so the function should return 5 in this case
*/

// мое решение:
package main

func CountBits(n uint) int {
	// Инициализируем переменную для подсчета битов
	count := 0

	// Пока число 'n' больше нуля
	for n > 0 {
		// Прибавляем к 'count' значение последнего бита числа 'n'
		count += int(n & 1)

		// Сдвигаем биты числа 'n' вправо на одну позицию
		n >>= 1
	}

	// Возвращаем общее количество установленных битов в числе 'n'
	return count
}
