package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	WRONG    = "Некорректный ввод"
	MISMATCH = "Оба числа должны быть в одной системе счисления"
	NEGATIVE = "Результатом операции с Римскими числами не может быть отрицательное число или ноль"
	RANGE    = "Введенные числа не могут быть больше 10"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println(calculate(input))
}

func parseInt(str string) int {
	if isRoman(str) {
		return convertRomanToArabic(str)
	}
	v, err := strconv.Atoi(str)
	if err != nil {
		panic(WRONG)
	}
	return v
}

func convertRomanToArabic(str string) int {
	var roman = map[string]int{
		"I":    1,
		"II":   2,
		"III":  3,
		"IV":   4,
		"V":    5,
		"VI":   6,
		"VII":  7,
		"VIII": 8,
		"IX":   9,
		"X":    10,
	}
	return roman[str]
}

func isRoman(str string) bool {
	for _, c := range str {
		if c == 'L' || c == 'C' {
			panic(RANGE)
		} else if c != 'I' && c != 'V' && c != 'X' {
			return false
		}
	}
	return true
}

func calculate(input string) string {
	inputArray := strings.Fields(input)
	if len(inputArray) != 3 {
		panic(WRONG)
	}

	a := parseInt(inputArray[0])
	b := parseInt(inputArray[2])
	var calcResult int
	switch inputArray[1] {
	case "+":
		calcResult = a + b
	case "-":
		calcResult = a - b
	case "/":
		calcResult = a / b
	case "*":
		calcResult = a * b
	default:
		panic(WRONG)
	}
	if a > 10 || b > 10 {
		panic(RANGE)
	}
	isRomanResult := isRoman(inputArray[0]) && isRoman(inputArray[2])
	if isRomanResult && (a == 0 || b == 0) {
		panic(RANGE)
	}
	if !isRoman(inputArray[0]) && isRoman(inputArray[2]) || isRoman(inputArray[0]) && !isRoman(inputArray[2]) {
		panic(MISMATCH)
	}

	if isRomanResult && calcResult <= 0 {
		panic(NEGATIVE)
	}

	if isRoman(inputArray[0]) && isRoman(inputArray[2]) {
		return convertArabicToRoman(calcResult)
	}

	return strconv.Itoa(calcResult)
}

type romanNumber struct {
	value     int
	character string
}

func convertArabicToRoman(number int) string {
	numbers := []romanNumber{
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}
	var result string
	for i, num := range numbers {
		count := number / num.value
		number -= count * num.value
		for j := 0; j < count; j++ {
			result += numbers[i].character
		}
	}
	return result
}
