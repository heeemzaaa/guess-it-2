package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	g "guess/guess-functions"
)

func main() {
	if len(os.Args) > 1 {
		fmt.Println("Usage: go run main.go")
		return
	}
	scanner := bufio.NewScanner(os.Stdin)
	slice := []float64{}
	for scanner.Scan() {
		Strnbr := scanner.Text()
		number, err := strconv.ParseFloat(Strnbr, 64)
		if err != nil {
			fmt.Println("Error Parsing :", err)
			continue
		}
		slice = append(slice, number)
		if len(slice) > 1 {
			min, max := g.Guess_it(slice)
			fmt.Printf("%.0f %.0f\n", min, max)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
	}
}
