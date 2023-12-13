package patterns

import "fmt"

// Интерфейс обработчика
type Handler interface {
	SetNext(handler Handler)
	HandleRequest(request string)
}

// Реализация обработчика "Логин"
type LoginHandler struct {
	next Handler
}

func (l *LoginHandler) SetNext(handler Handler) {
	l.next = handler
}

func (l *LoginHandler) HandleRequest(request string) {
	if request == "login" {
		fmt.Println("Обработчик Логин: обрабатываю запрос")
	} else if l.next != nil {
		fmt.Println("Обработчик Логин: передаю запрос дальше")
		l.next.HandleRequest(request)
	} else {
		fmt.Println("Обработчик Логин: не могу обработать запрос")
	}
}

// Реализация обработчика "Регистрация"
type RegistrationHandler struct {
	next Handler
}

func (r *RegistrationHandler) SetNext(handler Handler) {
	r.next = handler
}

func (r *RegistrationHandler) HandleRequest(request string) {
	if request == "registration" {
		fmt.Println("Обработчик Регистрация: обрабатываю запрос")
	} else if r.next != nil {
		fmt.Println("Обработчик Регистрация: передаю запрос дальше")
		r.next.HandleRequest(request)
	} else {
		fmt.Println("Обработчик Регистрация: не могу обработать запрос")
	}
}

// Реализация обработчика "Платеж"
type PaymentHandler struct {
	next Handler
}

func (p *PaymentHandler) SetNext(handler Handler) {
	p.next = handler
}

func (p *PaymentHandler) HandleRequest(request string) {
	if request == "payment" {
		fmt.Println("Обработчик Платеж: обрабатываю запрос")
	} else if p.next != nil {
		fmt.Println("Обработчик Платеж: передаю запрос дальше")
		p.next.HandleRequest(request)
	} else {
		fmt.Println("Обработчик Платеж: не могу обработать запрос")
	}
}

func ChainPattern() {
	// Создание обработчиков
	loginHandler := &LoginHandler{}
	registrationHandler := &RegistrationHandler{}
	paymentHandler := &PaymentHandler{}

	// Установка цепочки обработчиков
	loginHandler.SetNext(registrationHandler)
	registrationHandler.SetNext(paymentHandler)

	// Обработка запросов
	loginHandler.HandleRequest("login")
	loginHandler.HandleRequest("registration")
	loginHandler.HandleRequest("payment")
	loginHandler.HandleRequest("invalid")
}
