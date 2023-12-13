// Поскольку оригинальный wget скачивает файл, добавил сооветствующую реализацию.

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

var bytesToMegaBytes = 1048576.0

type PassThru struct {
	io.Reader
	curr  int64
	total float64
}

func (pt *PassThru) Read(p []byte) (int, error) {
	n, err := pt.Reader.Read(p)
	pt.curr += int64(n)

	if err == nil || (err == io.EOF && n > 0) {
		printProgress(float64(pt.curr), pt.total)
	}

	return n, err
}

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s url outfile\n", os.Args[0])
		os.Exit(1)
	}

	resp, err := http.Get(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	out, err := os.Create(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	src := &PassThru{Reader: resp.Body, total: float64(resp.ContentLength)}

	size, err := io.Copy(out, src)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Printf("\nФайл загружен. (%.1f MB)\n", float64(size)/bytesToMegaBytes)
}

func printProgress(curr, total float64) {
	width := 40.0
	output := ""
	threshold := (curr / total) * float64(width)
	for i := 0.0; i < width; i++ {
		if i < threshold {
			output += "="
		} else {
			output += " "
		}
	}

	fmt.Printf("\r[%s] %.1f of %.1fMB", output, curr/bytesToMegaBytes, total/bytesToMegaBytes)
}
