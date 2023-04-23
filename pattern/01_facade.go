package pattern

import (
	"fmt"
	"net/smtp"
)

/*
	Реализовать паттерн «фасад».

Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.

	https://en.wikipedia.org/wiki/Facade_pattern
*/
var (
	Host = "smtp.gmail.com"
	Port = "587"
)

type SMTPConfig struct {
	host string
	port string
}

func (config *SMTPConfig) Address() string {
	return config.host + ":" + config.port
}

type Message struct {
	subject string
	body    string
}

func (message *Message) Content() []byte {
	return []byte(message.subject + "\n" + message.body)
}

// Фасад отправки сообщения
type MailFacade struct {
	from    string
	to      string
	subject string
	body    string
	key     string
}

/*
Причина использования фасада
Без применения паттерна пользователю программы приходилось:
1. Настраивать адрес, объединяя хост и порт
2. Проводить авторизацию
3. Преобразовывать сообщение в нужный формат
Теперь пользователю доступен только метод Send,
применив который он отправит сообщение с меньшими усилиями и
не зная о сложной внутренней логике
*/
func (mail *MailFacade) Send() (err error) {
	smtpConfig := SMTPConfig{host: Host, port: Port}
	auth := smtp.PlainAuth("", mail.from, mail.key, smtpConfig.host)
	message := Message{subject: mail.subject, body: mail.body}
	err = smtp.SendMail(smtpConfig.Address(), auth, mail.from, []string{mail.to}, message.Content())
	return
}

// Пример использования
func RunFacade() {
	mailFacade := MailFacade{
		from:    "gmail Отправителя",
		to:      "email Получателя",
		subject: "Тема",
		body:    "Письмо",
		key:     "Ключ отправителя"}
	err := mailFacade.Send()
	if err != nil {
		fmt.Println("Произошла ошибка при отправке сообщения")
	}
}
