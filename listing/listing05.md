package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}

Вывод: error 

Объяснение: проблема, схожая с тем, что было в вопросе 3. 
Поле data интерфейса error nil, а поле с метаданными не пустое, 
следовательно, нужно сравнивать поле data c nil, а не весь тип.