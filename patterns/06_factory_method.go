package patterns

import "fmt"

// Представим, что мы продаём клиентам три вида компьютеров:
// сервера, пк и ноутбуки. Сначала создадим соответствующие типы
const (
	ServerType           = "server"
	PersonalComputerType = "pc"
	NotebookType         = "notebook"
)

// Для реализации фабричного метода нам потребуется
// так называемый "божественный" интерфейс
type Computer interface {
	GetType() string
	PrintDetails()
}

// Теперь создаём структуры для наших типов и наделяем эти структуры
// методами таким образом, чтобы они выполняли контракт интерфейса

// СТРУКТУРА СЕРВЕРА //
// Пусть у сервера будет три характеристики: его тип,
// количество ядер и объём памяти
type Server struct {
	Type   string
	Core   int
	Memory int
}

// Теперь реализуем выполнение контракта
func (s Server) GetType() string {
	return s.Type
}

func (s Server) PrintDetails() {
	fmt.Printf("%s Core:[%d], Memory: [%d]\n", s.Type, s.Core, s.Memory)
}

// Создаём новый объект сервера с конкретными характеристиками
func NewServer() Computer {
	return Server{
		Type:   ServerType,
		Core:   16,
		Memory: 512,
	}
}

// СТРУКТУРА ПЕРСОНАЛЬНОГО КОМПЬЮТЕРА //

// Для персонального компьютера добавим поле с проверкой наличия монитора
type PersonalComputer struct {
	Type    string
	Core    int
	Memory  int
	Monitor bool
}

// Далее всё как в предыдущем типе
func (pc PersonalComputer) GetType() string {
	return pc.Type
}

func (pc PersonalComputer) PrintDetails() {
	fmt.Printf("%s Core:[%d], Memory: [%d], Monitor: [%v]\n", pc.Type, pc.Core, pc.Memory, pc.Monitor)
}

func NewPersonalComputer() Computer {
	return PersonalComputer{
		Type:    PersonalComputerType,
		Core:    8,
		Memory:  128,
		Monitor: true,
	}
}

// СТРУКТУРА НОУТБУКА //

// Для ноутбука добавляем характеристику веса
type Notebook struct {
	Type   string
	Core   int
	Memory int
	Weight float32
}

func (nb Notebook) GetType() string {
	return nb.Type
}

func (nb Notebook) PrintDetails() {
	fmt.Printf("%s Core:[%d], Memory: [%d], Weight: [%f]\n", nb.Type, nb.Core, nb.Memory, nb.Weight)
}

func NewNotebook() Computer {
	return Notebook{
		Type:   NotebookType,
		Core:   8,
		Memory: 256,
		Weight: 2.71,
	}
}

// Теперь инициализируем фабричный метод:
// по дефолту мы поймаем несуществующий тип компьютера и скажем об этом
func NewType(typeName string) Computer {
	switch typeName {
	default:
		fmt.Printf("%s -- это несуществующий тип объекта\n\n", typeName)
		return nil
	case ServerType:
		return NewServer()
	case PersonalComputerType:
		return NewPersonalComputer()
	case NotebookType:
		return NewNotebook()
	}
}

// Пример реализации главной функции: мы берём все существующие объекты
// и выдаём их характеристики
func Factory() {
	var types = []string{ServerType, PersonalComputerType, NotebookType, "Mobile phone"}
	for _, item := range types {
		computer := NewType(item)
		if computer == nil {
			continue
		}
		computer.PrintDetails()
	}
	fmt.Println("Фабричный метод реализован.")
}
