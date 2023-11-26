package client

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

// Создаём структуру "Телнет" с полями таймаута, хоста и порта
type Telnet struct {
	timeout time.Duration
	host    string
	port    string
}

// И передаём параметры в эту структуру
func NewTelnet(timeout time.Duration, host, port string) *Telnet {
	return &Telnet{
		timeout: timeout,
		host:    host,
		port:    port,
	}
}

// Эта функция нужна для старта работы
func (t *Telnet) Start() {
	// Формируем подключение с таймаутом, беря те данные, которые сохранили в струтуре "Телнет"
	conn, err := net.DialTimeout("tcp", AddressBuilder(t.host, t.port), t.timeout)
	if err != nil {
		log.Fatal(err)
	}
	// Отложенно закрываем соединение
	defer conn.Close()

	// Начинаем бесконечный цикл работы клиента
	for {
		reader := bufio.NewReader(os.Stdin)
		cmd, err := reader.ReadString('\n')
		if err != nil {
			if errors.Is(err, io.EOF) {
				fmt.Fprintln(os.Stdout, "Закрываю соединение...")
			}
			fmt.Fprintln(os.Stderr, err)
			return
		}
		_, err = fmt.Fprint(conn, cmd)
		if err != nil {
			fmt.Fprintln(conn, "Ошибка во время отправки сообщения")
		}
		serverReader := bufio.NewReader(conn)
		text, err := serverReader.ReadString('\n')
		if err != nil {
			if errors.Is(err, io.EOF) {
				fmt.Fprintln(os.Stdout, "Закрыва соединение с сервером...")
			}
			fmt.Fprintln(os.Stderr, err)
			return
		}
		fmt.Fprint(os.Stdout, text)

	}
}

// Эта функция просто объединяет две строки в одну
func AddressBuilder(host, port string) string {
	return host + port
}
