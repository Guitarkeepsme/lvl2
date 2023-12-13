package main

import (
	"log"
	"net/http"

	"task11/cache"
	"task11/handlers"
)

func main() {

	storage := cache.NewCache()

	mux := http.NewServeMux()

	createEventHandler := http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		handlers.CreateEventHandler(writer, request, storage)
	})

	updateEventHandler := http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		handlers.UpdateEventHandler(writer, request, storage)
	})

	deleteEventHandler := http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		handlers.DeleteEventHandler(writer, request, storage)
	})

	getDayEventHandler := http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		handlers.GetDayEventHandler(writer, request, storage)
	})

	getWeekEventHandler := http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		handlers.GetWeekEventHandler(writer, request, storage)
	})

	getMonthEventHandler := http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		handlers.GetMonthEventHandler(writer, request, storage)
	})

	// mux.HandleFunc("/", homePage)
	mux.Handle("/create_event", mwLogger(createEventHandler))
	mux.Handle("/update_event", mwLogger(updateEventHandler))
	mux.Handle("/delete_event", mwLogger(deleteEventHandler))
	mux.Handle("/daily_events", mwLogger(getDayEventHandler))
	mux.Handle("/weekly_events", mwLogger(getWeekEventHandler))
	mux.Handle("/monthly_events", mwLogger(getMonthEventHandler))

	// Выводим на консоль сообщение о том, что сервер запущен

	log.Println("Сервер запущен по адресу localhost:8080...")

	// Печатаем ошибки
	log.Fatal(http.ListenAndServe(":8080", mux))
}

// Функция для стартовой страницы
// func homePage(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("Добро пожаловать в утилиту календаря!"))
// }

// Для того, чтобы наглядно запускать каждый обработчик, создаём специальную функцию
func mwLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Запуск %s...", r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
