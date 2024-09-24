package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var data []float64

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			return
		}

		num, err := strconv.ParseFloat(strings.TrimSpace(line), 64)
		if err != nil {
			log.Println("Error Parsing input")
			return
		}
		data = append(data, num)
	}
}
