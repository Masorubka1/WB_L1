package main

import "fmt"

type Set struct {
	data map[string]bool
}

func NewSet() *Set {
	return &Set{
		data: make(map[string]bool),
	}
}

func (s *Set) Add(item string) {
	s.data[item] = true
}

func (s *Set) Remove(item string) {
	delete(s.data, item)
}

func (s *Set) Contains(item string) bool {
	_, exists := s.data[item]
	return exists
}

func (s *Set) Size() int {
	return len(s.data)
}

func (s *Set) Clear() {
	s.data = make(map[string]bool)
}

func main12() {
	sequence := []string{"cat", "cat", "dog", "cat", "tree"}

	set := NewSet()

	// Добавление элементов из последовательности в множество
	for _, item := range sequence {
		set.Add(item)
	}

	// Вывод элементов множества
	for item := range set.data {
		fmt.Println(item)
	}
}
