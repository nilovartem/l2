package pattern

import "fmt"

/*
	Реализовать паттерн «стратегия».

Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.

	https://en.wikipedia.org/wiki/Strategy_pattern
*/
type Sorter interface {
	Sort([]int) []int
}
type QuickSort struct{}

func (q QuickSort) Sort(data []int) []int {
	if len(data) < 2 {
		return data
	} else {
		pivot := data[0]
		var less []int
		for _, value := range data[1:] {
			if value <= pivot {
				less = append(less, value)
			}
		}
		var greater []int
		for _, value := range data[1:] {
			if value > pivot {
				greater = append(greater, value)
			}
		}
		var middle []int
		middle = append(middle, pivot)
		data = append(q.Sort((less)), append(middle, q.Sort(greater)...)...)
		return data
	}
}

type BubbleSort struct{}

func (s BubbleSort) Sort(data []int) []int {
	var isDone = false
	for !isDone {
		isDone = true
		var i = 0
		for i < len(data)-1 {
			if data[i] > data[i+1] {
				data[i], data[i+1] = data[i+1], data[i]
				isDone = false
			}
			i++
		}
	}
	return data
}

type Selector struct {
	sorter Sorter
	data   []int
}

func (selector *Selector) Selector(data []int, sorter Sorter) {
	selector.data = data
	selector.sorter = sorter
}
func (selector *Selector) PerformSort() []int {
	return selector.sorter.Sort(selector.data)
}

// Пример использования паттерна
func RunStrategy() {
	selector := Selector{sorter: BubbleSort{}, data: []int{5, 4, 1, 3, 7, 2, 10}}
	fmt.Println("Bubble Sort")
	fmt.Println(selector.PerformSort())
	selector = Selector{sorter: QuickSort{}, data: []int{5, 4, 1, 3, 7, 2, 10}}
	fmt.Println("Quick Sort")
	fmt.Println(selector.PerformSort())
}
