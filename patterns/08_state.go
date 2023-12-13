package patterns

import "fmt"

// Интерфейс состояния
type State interface {
	Handle(context *Context)
}

// Реализация состояния "Включено"
type OnState struct{}

func (o *OnState) Handle(context *Context) {
	fmt.Println("Текущее состояние: Включено")
	context.SetState(&OffState{})
}

// Реализация состояния "Выключено"
type OffState struct{}

func (o *OffState) Handle(context *Context) {
	fmt.Println("Текущее состояние: Выключено")
	context.SetState(&OnState{})
}

// Контекст
type Context struct {
	state State
}

func (c *Context) SetState(state State) {
	c.state = state
}

func (c *Context) Request() {
	c.state.Handle(c)
}

func StatePattern() {
	// Создание контекста
	context := &Context{}

	// Изначальное состояние - Выключено
	context.SetState(&OffState{})

	// Выполнение запросов
	context.Request()
	context.Request()
	context.Request()
}
