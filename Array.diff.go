/*
Your goal in this kata is to implement a difference function, which subtracts one list from another and returns the result.

It should remove all values from list a, which are present in list b keeping their order.

array_diff({1, 2}, 2, {1}, 1, *z) == {2} (z == 1)
If a value is present in b, all of its occurrences must be removed from the other:

array_diff({1, 2, 2, 2, 3}, 5, {2}, 1, *z) == {1, 3} (z == 2)
*/

// Мое решение:
package main

func ArrayDiff(a, b []int) []int {
	// Создаем пустой срез, в который будем добавлять различающиеся элементы
	diff := []int{}

	// Проходим по всем элементам из среза 'a'
	for _, valA := range a {
		// Инициализируем переменную для отслеживания, был ли найден совпадающий элемент в 'b'
		found := false

		// Проходим по всем элементам из среза 'b'
		for _, valB := range b {
			// Если найдено совпадение элементов, помечаем 'found' как true и выходим из цикла 'for'
			if valA == valB {
				found = true
				break
			}
		}

		// Если не было найдено совпадение среди элементов 'b', добавляем 'valA' в срез 'diff'
		if !found {
			diff = append(diff, valA)
		}
	}

	// Возвращаем срез, содержащий различающиеся элементы из 'a'
	return diff
}
