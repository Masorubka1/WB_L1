/*К каким негативным последствиям может привести данный фрагмент кода, и как это исправить? Приведите корректный пример реализации.


var justString string
func someFunc() {
  v := createHugeString(1 << 10)
  justString = v[:100]
}

func main() {
  someFunc()
}

Проблема - слишком большие строки едят много памяти, которые gc потом будет удалять(это долго по времени)
*/

package main

import (
	"fmt"
	"strings"
)

var justString string

func someFunc() {
	var builder strings.Builder
	builder.Grow(1 << 10) // Устанавливаем начальный размер строки

	// Добавляем содержимое в builder
	builder.WriteString(createHugeString(1 << 10))

	// Присваиваем только первые 100 символов переменной justString
	justString = builder.String()[:100]
}

func main15() {
	someFunc()
	fmt.Println(justString)
}

func createHugeString(size int) string {
	// Здесь должна быть логика создания огромной строки
	// Возвращаем пустую строку для примера
	return ""
}
