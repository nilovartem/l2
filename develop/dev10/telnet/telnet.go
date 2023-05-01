package telnet

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func send(conn net.Conn, signalChan chan<- os.Signal, connChan chan<- error) {
	for {
		// Читаем из stdin
		reader := bufio.NewReader(os.Stdin)
		payload, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				signalChan <- syscall.Signal(syscall.SIGQUIT)
				return
			}
			connChan <- err
		}
		// Пишем в сокет
		fmt.Fprintln(conn, payload)
	}
}
func receive(conn net.Conn, connChan chan<- error) {
	for {
		// Читаем ответ от сервера
		reader := bufio.NewReader(conn)
		response, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				connChan <- fmt.Errorf("connection closed by foreign host")
				return
			}
			connChan <- err
		}
		// Выводим что прочитали
		fmt.Print(response)
	}
}
func Connect(timeout uint, host, port string) error {
	address := net.JoinHostPort(host, port)
	conn, err := net.DialTimeout("tcp", address, time.Duration(timeout*uint(time.Second)))
	if err != nil {
		fmt.Printf("%v", err)
		return err
	}
	defer conn.Close()
	fmt.Printf("Connected to %v\n", address)
	// syscall chan
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGTERM, syscall.SIGINT)

	// connection error chan
	connChan := make(chan error, 1)
	go send(conn, signalChan, connChan)
	go receive(conn, connChan)

	select {
	case <-signalChan:
		conn.Close()
	case err = <-connChan:
		if err != nil {
			return err
		}
	}
	return nil

}
