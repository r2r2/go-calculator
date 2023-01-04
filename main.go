package main

import (
	"SimpleCalculator/calculator"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const defaultLength = 20

func main() {
	greetings()
	doCalculation()
}

func doCalculation() {
	for {
		fmt.Print("enter digits...: ")
		a, b, operator, err := readFromInput()
		if string(a) == "q" {
			fmt.Println(err)
			return
		}
		if err != nil {
			printErr(err)
			continue
		}

		af, bf, err := convertToFloat64(a, b)
		if err != nil {
			printErr(err)
			continue
		}

		result, err := calculator.SimpleCalculator(af, bf, operator)
		if err != nil {
			printErr(err)
			continue
		}

		c := offset(a, b)
		fmt.Printf("\u001BM\u001B[%dC =\u001B[3;38;5;121m %.2f\n", c, result)
		fmt.Print("\u001B[0m") // reset all modes (styles and colors)
	}
}

func readFromInput() (a, b, op string, err error) {
	text, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		return "", "", "", err
	}
	text = strings.TrimSpace(text)
	str := strings.Split(text, " ")

	if str[0] == "q" {
		return str[0], "", "", fmt.Errorf("\n\texiting calculator\n\tBye 👋!")
	}

	if len(str) != 3 {
		return "", "", "", fmt.Errorf("enter digits in format <2 + 2>")
	}

	return str[0], str[2], str[1], nil
}

func convertToFloat64(a, b string) (sf, bf float64, err error) {

	if sf, err = strconv.ParseFloat(a, 64); err != nil {
		return sf, 0, fmt.Errorf("not a valid number: %v", a)
	}

	if bf, err = strconv.ParseFloat(b, 64); err != nil {
		return 0, bf, fmt.Errorf("not a valid number: %v", b)
	}
	return sf, bf, nil
}

func offset(a, b string) int {
	return defaultLength + len(a) + len(b)
}

func greetings() {
	fmt.Println("\n\tAvailable operators: '+' '-' '*' '/'")
	fmt.Println("\tenter digits in format <2 + 2>")
	fmt.Println("\tto quit type: 'q'")
	fmt.Println()
}

func printErr(err error) {
	fmt.Printf("\u001B[3;31m%v\u001B[0m\n", err)
}
