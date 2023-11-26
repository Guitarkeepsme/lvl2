// Реализовать утилиту wget с возможностью скачивать сайты целиком.

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

func main() {
	// 	var address string

	// 	fmt.Println("Введите адрес сайта:")
	// 	_, err := fmt.Scanf("%s", address)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	file, err := download(address)
	file, err := download("https://backendinterview.ru")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Файл с названием: %s создан\n", file.Name())
}

func download(path string) (*os.File, error) {
	if err := parseURL(path); err != nil {
		return nil, err
	}
	res, err := http.Get(path)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Соединение разорвано, код: %d \nи тело: %s\n", res.StatusCode, body)
	}
	if err != nil {
		return nil, err
	}
	var (
		index    = 0
		filename = "index_"
	)
	for {
		if _, err := os.Open(filename + strconv.Itoa(index) + ".html"); err == nil {
			index += 1
			continue
		}
		break
	}
	file, err := os.Create(filename + strconv.Itoa(index) + ".html")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	if _, err := file.Write(body); err != nil {
		return nil, err
	}
	return file, nil
}

func parseURL(rawURL string) error {
	_, err := url.Parse(rawURL)
	if err != nil {
		return err
	}
	return nil
}
