package main

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

/*
=== Базовая задача ===

Создать программу печатающую точное время с использованием NTP библиотеки.Инициализировать как go module.
Использовать библиотеку https://github.com/beevik/ntp.
Написать программу печатающую текущее время / точное время с использованием этой библиотеки.

Программа должна быть оформлена с использованием как go module.
Программа должна корректно обрабатывать ошибки библиотеки: распечатывать их в STDERR и возвращать ненулевой код выхода в OS.
Программа должна проходить проверки go vet и golint.
*/
const (
	SERVER = "0.beevik-ntp.pool.ntp.org"
)

func Time(server string) (t time.Time, err error) {
	response, err := ntp.Query(server)
	if err != nil {
		return t, err
	}
	//Получим смещение "правильного" времени сервера от "неправильного"
	//локального и добавим разницу к локальному
	t = time.Now().Add(response.ClockOffset)
	return t, err
}
func main() {
	time, err := Time(SERVER)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Got error: \n%v", err)
		os.Exit(1)
	}
	fmt.Println(time)
}
