package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"learn.zone01kisumu.ke/git/hiombima/guess-it-2/calculation"
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

		slope, intercept := calculation.LinearRegression(data)
		nextX := float64(len(data))

		predictedValue := slope*nextX + intercept

		lowerLimit := predictedValue - 2*calculation.StandardDeviation(data)
		upperLimit := predictedValue + 2*calculation.StandardDeviation(data)

		fmt.Printf("%.0f %0.f\n", lowerLimit, upperLimit)

	}

	if err := scanner.Err(); err != nil {
		log.Printf("Error reading input: %s", err)
	}
}
