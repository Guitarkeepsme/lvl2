Что выведет программа? Объяснить вывод программы.

package main

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	for n := range ch {
		println(n)
	}
}

Вывод: fatal error: all goroutines are asleep - deadlock!

Объяснение: причина в том, что мы не закрыли канал, и продолжили читать из пустого канала.
Следует не забывать закрывать канал.

Для исправления ошибки необходимо после 10 строки добавить close(ch)