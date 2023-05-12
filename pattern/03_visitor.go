package pattern

import "fmt"

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/

type Human interface {
	accept(visitor)
}
type Client struct {
	name  string
	goods []string
}
type Employee struct {
	name   string
	salary float64
}

func (c Client) accept(v visitor) {
	v.clientInfo(c)
}

func (e Employee) accept(v visitor) {
	v.employeeInfo(e)
}

type visitor interface {
	clientInfo(Client)
	employeeInfo(Employee)
}

type desc struct{}

func (desc) clientInfo(c Client) {
	fmt.Println("Информация о клиенте:")
	fmt.Printf("Имя: %v\nТовары: %v\n", c.name, c.goods)
}

func (desc) employeeInfo(e Employee) {
	fmt.Println("Информация о сотруднике:")
	fmt.Printf("Имя: %v\nЗарплата: %v", e.name, e.salary)
}

func RunVisitor() {
	humans := []Human{
		Client{name: "Артем", goods: []string{"Рыба", "Мясо", "Овощи"}},
		Employee{name: "Иван", salary: 15000},
	}
	for _, h := range humans {
		h.accept(desc{})
	}
}
