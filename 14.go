package main

import (
	"fmt"
	"reflect"
)

func main14() {
	var val interface{} = "Hello"

	// Определение типа переменной
	switch val.(type) {
	case int:
		fmt.Println("Тип: int")
	case string:
		fmt.Println("Тип: string")
	case bool:
		fmt.Println("Тип: bool")
	case chan int:
		fmt.Println("Тип: chan int")
	default:
		fmt.Println("Неизвестный тип")
	}

	// Получение типа переменной с использованием рефлексии
	valueType := reflect.TypeOf(val)
	fmt.Println("Тип переменной:", valueType)
}
