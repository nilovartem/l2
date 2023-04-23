package pattern

import (
	"errors"
	"fmt"
	"strings"
)

/*
	Реализовать паттерн «строитель».

Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.

	https://en.wikipedia.org/wiki/Builder_pattern
*/

// Строим Select запрос + отдельно храним его составляющие
type Query struct {
	Table   string
	Columns []string
	Sql     string
}

// Абстрактный строитель Select запроса
type ISelectBuilder interface {
	SelectColumns()
	SelectFromTable()
	Build() (Query, error)
}

// Определенный строитель
type SelectBuilder struct {
	table   string
	columns []string
}

var (
	errBuilder = errors.New("not enough arguments, SelectBuilder can't build object")
)

func (selectBuilder *SelectBuilder) SelectFromTable(table string) *SelectBuilder {
	selectBuilder.table = table
	return selectBuilder
}
func (selectBuilder *SelectBuilder) SelectColumns(columns ...string) *SelectBuilder {
	selectBuilder.columns = columns
	return selectBuilder
}
func (selectBuilder *SelectBuilder) Build() (*Query, error) {
	if selectBuilder.table == "" {
		return nil, errBuilder
	}
	if len(selectBuilder.columns) == 0 {
		selectBuilder.columns = []string{"*"}
	}
	return &Query{
		Table:   selectBuilder.table,
		Columns: selectBuilder.columns,
		//Построение запроса сделано как пример, очень запутанный код
		Sql: strings.Join([]string{"SELECT", strings.Join(selectBuilder.columns, ","), "FROM", selectBuilder.table}, " "),
	}, nil
}

/*
До внедрения паттерна приходилось писать SQL запрос вручную / использовать длинный конструктор.
Теперь можно вызвать методы построения по цепочке и получить тот же запрос, но без длинного конструктора
*/
//Пример использования
func RunBuilder() {
	selectBuilder := SelectBuilder{}
	query, err := selectBuilder.
		SelectFromTable("Products").
		SelectColumns("ID", "Name", "Color", "Type").
		Build()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(query.Table)
	fmt.Println(query.Columns)
	fmt.Println(query.Sql)
}
