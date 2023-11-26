/*
Реализовать простейший telnet-клиент.

Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123


Требования:
Программа должна подключаться к указанному хосту (ip или доменное имя + порт) по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные, полученные и[з] сокета, должны выводиться в STDOUT.
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s)
При нажатии Ctrl+D программа должна закрывать сокет и завершаться.
Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему серверу программа должна завершаться через timeout.

*/

package main

import (
	"flag"
	"os"
	"time"

	"github.com/guitarkeepsme/task10.go/internal/client"
	"github.com/guitarkeepsme/task10.go/internal/server"
)

var (
	timeout time.Duration
	host    string = "localhost"
	port    string = ":8080"
)

func main() {
	// Функция для инициализации флага
	initFlags()

	go server.RunServer(host + port)

	client := client.NewTelnet(timeout, host, port)
	client.Start()

}

func initFlags() {
	// В данном случае у нас всего один флаг
	flag.DurationVar(&timeout, "timeout", time.Second*10, "таймаут запуска сервера")
	flag.Parse()

	// Если имеем всего один аргумент, это значит, что нам передали адрес хоста
	if len(os.Args) == 1 {
		host = flag.Arg(1)
	}

	// Если аргументов два, то второй -- номер порта
	if len(os.Args) == 2 {
		port = ":" + flag.Arg(2)
	}
}
