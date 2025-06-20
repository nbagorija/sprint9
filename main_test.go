package main

import (
	"testing"
)

// TestGenerateRandomElements тестирует функцию генерации случайных чисел.
func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name     string
		size     int
		expected int
	}{
		{
			name:     "Обычный случай",
			size:     10,
			expected: 10,
		},
		{
			name:     "Нулевой размер",
			size:     0,
			expected: 0,
		},
		{
			name:     "Отрицательный размер",
			size:     -1,
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := generateRandomElements(tt.size)
			if len(result) != tt.expected {
				t.Errorf("generateRandomElements(%d) вернул слайс длиной %d, ожидалось %d",
					tt.size, len(result), tt.expected)
			}
			if tt.size > 0 {
				for _, num := range result {
					if num < 0 {
						t.Errorf("generateRandomElements(%d) вернул отрицательное число %d",
							tt.size, num)
					}
				}
			}
		})
	}
}

// TestMaximum тестирует функцию поиска максимума в один поток.
func TestMaximum(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "Обычный случай",
			input:    []int{1, 5, 3, 8, 2},
			expected: 8,
		},
		{
			name:     "Пустой слайс",
			input:    []int{},
			expected: 0,
		},
		{
			name:     "Один элемент",
			input:    []int{42},
			expected: 42,
		},
		{
			name:     "Все элементы одинаковые",
			input:    []int{5, 5, 5, 5},
			expected: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maximum(tt.input)
			if result != tt.expected {
				t.Errorf("maximum(%v) = %d, ожидалось %d", tt.input, result, tt.expected)
			}
		})
	}
}

// TestMaxChunks тестирует функцию поиска максимума с использованием горутин.
func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "Обычный случай",
			input:    []int{1, 5, 3, 8, 2, 7, 4, 6},
			expected: 8,
		},
		{
			name:     "Пустой слайс",
			input:    []int{},
			expected: 0,
		},
		{
			name:     "Один элемент",
			input:    []int{42},
			expected: 42,
		},
		{
			name:     "Все элементы одинаковые",
			input:    []int{5, 5, 5, 5, 5, 5, 5, 5},
			expected: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maxChunks(tt.input)
			if result != tt.expected {
				t.Errorf("maxChunks(%v) = %d, ожидалось %d", tt.input, result, tt.expected)
			}
		})
	}
}
