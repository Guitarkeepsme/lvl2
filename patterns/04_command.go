package patterns

import "fmt"

// Интерфейс команды
type Command interface {
	Execute()
}

// Реализация команды "Приветствие"
type GreetingCommand struct {
	name string
}

func (g *GreetingCommand) Execute() {
	fmt.Printf("Привет, %s!\n", g.name)
}

// Реализация команды "Прощание"
type FarewellCommand struct {
	name string
}

func (f *FarewellCommand) Execute() {
	fmt.Printf("До свидания, %s!\n", f.name)
}

// Реализация командного исполнителя
type Invoker struct {
	command Command
}

func (i *Invoker) SetCommand(command Command) {
	i.command = command
}

func (i *Invoker) ExecuteCommand() {
	i.command.Execute()
}

func CommandPattern() {
	// Создание команд и командного исполнителя
	greetingCmd := &GreetingCommand{name: "Вася"}
	farewellCmd := &FarewellCommand{name: "Петя"}
	invoker := &Invoker{}

	// Установка и исполнение команд
	invoker.SetCommand(greetingCmd)
	invoker.ExecuteCommand()

	invoker.SetCommand(farewellCmd)
	invoker.ExecuteCommand()
}
