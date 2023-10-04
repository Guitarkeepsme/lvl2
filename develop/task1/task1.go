package main

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	tm, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		fmt.Printf("Ошибка во время выполнения команды: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Время системы:", time.Now())
	fmt.Println("Точное время:", tm)
}
