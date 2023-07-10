package main

import "fmt"

func Intersection(set1, set2 []int) []int {
	// Создаем map для хранения уникальных элементов первого множества
	set1Map := make(map[int]bool)
	for _, num := range set1 {
		set1Map[num] = true
	}

	// Создаем временный map для хранения пересечения множеств
	intersectionMap := make(map[int]bool)

	// Проверяем каждый элемент второго множества на наличие в первом множестве
	// и добавляем его во временный map, если он найден
	for _, num := range set2 {
		if set1Map[num] {
			intersectionMap[num] = true
		}
	}

	// Создаем слайс для результата и копируем элементы из временного map в слайс
	intersection := make([]int, 0, len(intersectionMap))
	for num := range intersectionMap {
		intersection = append(intersection, num)
	}

	return intersection
}

func main11() {
	// Примеры пересечения множеств
	set1 := []int{1, 2, 3, 4, 5}
	set2 := []int{4, 5, 6, 7, 8}
	result := Intersection(set1, set2)
	fmt.Println(result) // Output: [4 5]

	set3 := []int{10, 20, 30, 40}
	set4 := []int{30, 40, 50, 60}
	result = Intersection(set3, set4)
	fmt.Println(result) // Output: [30 40]

	set5 := []int{1, 2, 3}
	set6 := []int{4, 5, 6}
	result = Intersection(set5, set6)
	fmt.Println(result) // Output: []
}
