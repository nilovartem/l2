package pattern

import "fmt"

/*
	Реализовать паттерн «состояние».

Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.

	https://en.wikipedia.org/wiki/State_pattern
*/
type OrderState interface {
	next(*Order)
	cancel(*Order)
}

type Order struct {
	state OrderState
	block bool
}

func (o *Order) next() {
	o.state.next(o)
	if o.block {
		fmt.Println("Невозможно выполнить операцию, заказ был отменен")
	}
}
func (o *Order) cancel() {
	o.state.cancel(o)
	o.block = true
}
func (o *Order) setState(state OrderState) {
	o.state = state
}

// Состояния
type Delivery struct{}
type Ready struct{}
type Canceled struct{}

func (Delivery) next(order *Order) {
	fmt.Println("Заказ доставлен на пункт выдачи")
	order.setState(Ready{})
}
func (Delivery) cancel(order *Order) {
	order.setState(Canceled{})
}
func (Ready) next(order *Order) {
	fmt.Println("Заказ ожидает на пункте выдачи")
}
func (Ready) cancel(order *Order) {
	order.setState(Canceled{})
}
func (Canceled) next(order *Order) {
	fmt.Println("Штраф за отмену заказа")
}
func (Canceled) cancel(order *Order) {
}
func RunState() {
	order := Order{Delivery{}, false}
	order.next()
	order.cancel()
	order.next()
}
