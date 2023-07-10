package main

import (
	"fmt"
	"math"
	"sort"
)

func main10() {
	// Исходная последовательность температурных колебаний
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	// Создаем map для группировки температур
	groups := make(map[int][]float64)

	// Группируем температуры по 10-градусным интервалам
	for _, temp := range temperatures {
		// Округляем температуру до ближайшего числа, кратного 10
		group := int(math.Round(temp/10.0)) * 10

		// Добавляем температуру в соответствующую группу
		groups[group] = append(groups[group], temp)
	}

	// Выводим результаты
	keys := make([]int, 0, len(groups))
	for key := range groups {
		keys = append(keys, key)
	}
	sort.Ints(keys)

	for _, key := range keys {
		fmt.Printf("%d: %v\n", key, groups[key])
	}
}
