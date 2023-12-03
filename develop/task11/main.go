package main

import (
	"log"
	"net/http"

	"task11/cache"
	"task11/handlers"
)

func main() {

	// Для хранения данных используем кэш
	storage := cache.NewCache()

	// Создаём мультиплексор запросов: Он сопоставляет URL-адрес каждого входящего запроса
	// со списком зарегистрированных шаблонов и вызывает обработчик для шаблона,
	// который наиболее точно соответствует URL-адресу.
	mux := http.NewServeMux()

	// Посольку мы создали свой мультиплексор запросов, нам необходимо явно вызвать http.HandlerFunc().
	// Это заставит тип http.HandlerFunc действовать в качестве адаптера,
	// который позволяет задействовать обычные функции как HTTP-обработчики при условии,
	// что они обладают требуемой сигнатурой.
	createEventHandler := http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		// handlers.CreateEventHandler(writer, request, storage)
	})

	// Итак, мы реализуем все необходимые вызовы обработчиков: создание обработчика,
	// обновление, удаление, получение дня, недели и месяца.
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

	// И создаём соответствующие обработчики для ServeHux:

	mux.Handle("/create_event", mwLogger(createEventHandler))
	mux.Handle("/update_event", mwLogger(updateEventHandler))
	mux.Handle("/delete_event", mwLogger(deleteEventHandler))
	mux.Handle("daily_events", mwLogger(getDayEventHandler))
	mux.Handle("/weekly_events", mwLogger(getWeekEventHandler))
	mux.Handle("/monthly_events", mwLogger(getMonthEventHandler))

	// Выводим на консоль сообщение о том, что сервер успешно запущен

	log.Println("Сервер запущен по адресу localhost:8000...")

	// Печатаем ошибки
	log.Fatal(http.ListenAndServe(":8080", mux))
}

// Эта функция нужна для того, чтобы
func mwLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Запуск %s...", r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
