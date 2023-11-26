package server

import (
	"net/http"
	"time"
)

// Функция для запуска сервера
func RunServer(address string) {
	// С помощью встроенной библиотеки httр инициализируем сервер:
	// адресс получаем из функции main
	server := http.Server{
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
		Addr:         address,
	}
	// Слушаем сервер
	server.ListenAndServe()
	// Через 30 секунд сервер приостановит работу
	timeout := time.After(time.Second * 30)

	// Для отслеживания этого создаём канал
	done := make(chan bool)

	// Запускаем горутины, которые слушают канал
	go func(done chan<- bool) {
		<-timeout
		// По прошествии три секунды передаём флаг о закрытии канала
		done <- true
		// И закрываем сервер
		server.Close()
	}(done)
}
