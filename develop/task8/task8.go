/* Необходимо реализовать свой собственный UNIX-шелл-утилиту с поддержкой ряда простейших команд:


- cd <args> - смена директории (в качестве аргумента могут быть то-то и то)
- pwd - показать путь до текущего каталога
- echo <args> - вывод аргумента в STDOUT
- kill <args> - "убить" процесс, переданный в качесте аргумента (пример: такой-то пример)
- ps - выводит общую информацию по запущенным процессам в формате *такой-то формат*




Также требуется поддерживать функционал fork/exec-команд


Дополнительно необходимо поддерживать конвейер на пайпах (linux pipes, пример cmd1 | cmd2 | .... | cmdN).


*Шелл — это обычная консольная программа, которая, будучи запущенной, в интерактивном сеансе выводит некое приглашение
в STDOUT и ожидает ввода пользователя через STDIN. Дождавшись ввода, обрабатывает команду согласно своей логике
и при необходимости выводит результат на экран. Интерактивный сеанс поддерживается до тех пор, пока не будет введена команда выхода (например \quit).
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
)

func main() {
	for {
		fmt.Print("> ")

		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		// fmt.Println(input)

		args := strings.Split(input, " ")

		if len(args) == 0 {
			continue
		}

		switch args[0] {
		case "cd":

			if len(args) < 2 {
				fmt.Println("Напишите название директории")
				continue
			}
			err := os.Chdir(args[1])
			if err != nil {
				fmt.Println("Ошибка:", err)
			}
		case "pwd":
			dir, err := os.Getwd()
			if err != nil {
				fmt.Println("Ошибка:", err)
			} else {
				fmt.Println(dir)
			}
		case "echo":
			fmt.Println(strings.Join(args[1:], " "))
		case "kill":
			if len(args) < 2 {
				fmt.Println("Напишите, какую программу вы хотите завершить")
				continue
			}
			pid, errInt := strconv.ParseInt(args[1], 32, 0)
			if errInt != nil {
				fmt.Println("Некорректный процесс")
				continue
			}
			process, err := os.FindProcess(int(pid))
			if err != nil {
				fmt.Println("Ошибка:", err)
			} else {
				err := process.Signal(syscall.SIGTERM)
				if err != nil {
					fmt.Println("Ошибка:", err)
				}
			}
		case "ps":
			output, err := exec.Command("ps", "aux").Output()
			if err != nil {
				fmt.Println("Ошибка:", err)
			} else {
				fmt.Println(string(output))
			}
		default:
			fmt.Println("Неизвестная команда:", args[0])
			continue
		}
	}
}
