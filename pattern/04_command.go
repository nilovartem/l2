package pattern

import (
	"fmt"
	"os/exec"
)

/*
	Реализовать паттерн «комманда».

Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.

	https://en.wikipedia.org/wiki/Command_pattern
*/
var (
	history []string
)

type Command interface {
	Execute(...string)
	Name() string
}

type MakeDirectory struct {
}

func (m MakeDirectory) Execute(args ...string) {
	cmd := exec.Command("mkdir", args[0])
	_, _ = cmd.Output()
	history = append(history, m.Name())
}
func (m MakeDirectory) Name() string {
	return "mkdir"
}

type Touch struct {
}

func (t Touch) Execute(args ...string) {
	cmd := exec.Command("touch", args[0])
	_, _ = cmd.Output()
	history = append(history, t.Name())
}
func (t Touch) Name() string {
	return "touch"
}

type History struct {
}

func (h History) Execute(args ...string) {
	fmt.Println("History")
	for _, entry := range history {
		fmt.Println(entry)
	}
	history = append(history, h.Name())
}
func (h History) Name() string {
	return "history"
}

type Executor struct {
}

func (e *Executor) Execute(command Command, args ...string) {
	command.Execute(args...)
}

// Пример использования
func RunCommand() {
	executor := Executor{}
	executor.Execute(MakeDirectory{}, "/Users/artem/Documents/tests/folder")
	executor.Execute(Touch{}, "/Users/artem/Documents/tests/file.empty")
	executor.Execute(History{})

}
