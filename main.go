package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func IsRomanNumerals(s string) bool {
	var RomanСharacters string = "IVXLCDM"
	for _, char := range s {
		if !strings.ContainsRune(RomanСharacters, char) {
			return false
		}
	}
	return true
}

func ConvertRomanToArabNumberls(s string) int {
	var num = map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	result := 0
	flag := 0
	for i := len(s) - 1; i >= 0; i-- {
		if num[s[i]] >= flag {
			result += num[s[i]]
		} else {
			result -= num[s[i]]
		}
		flag = num[s[i]]
	}

	return result
}

func ValidateOperand(operand int) error {
	if operand < 1 || operand > 10 {
		return errors.New("вывод ошибки, так как операнд находится вне диапазона от 1 до 10")
	}

	return nil
}

func ValidateOperation(operation string) error {
	if operation != "+" && operation != "-" && operation != "*" && operation != "/" {
		return errors.New("вывод ошибки, так как оператор не является допустимым")
	}
	return nil
}

func Calculate(operand1 int, operand2 int, operation string) int {
	result := 0
	switch operation {
	case "+":
		result = operand1 + operand2
	case "-":
		result = operand1 - operand2
	case "*":
		result = operand1 * operand2
	case "/":
		result = operand1 / operand2
	}
	return result
}

func ConvertArabToRomanNumerals(arab int) string {
	c := [...]byte{'I', 'X', 'C', 'M', 'V', 'L', 'D'}
	s := strconv.Itoa(arab)
	l := len(s)
	var result string
	for i := 0; i < l; i++ {
		n := int(s[i]) - 48
		switch n {
		case 0:
			continue
		case 1:
			result += string(c[l-i-1])
		case 2:
			result += string(c[l-i-1]) + string(c[l-i-1])
		case 3:
			result += string(c[l-i-1]) + string(c[l-i-1]) + string(c[l-i-1])
		case 4:
			result += string(c[l-i-1]) + string(c[l-i+3])
		case 5:
			result += string(c[l-i+3])
		case 6:
			result += string(c[l-i+3]) + string(c[l-i-1])
		case 7:
			result += string(c[l-i+3]) + string(c[l-i-1]) + string(c[l-i-1])
		case 8:
			result += string(c[l-i+3]) + string(c[l-i-1]) + string(c[l-i-1]) + string(c[l-i-1])
		case 9:
			result += string(c[l-i-1]) + string(c[l-i])
		}
	}
	return result
}

func main() {
	// Get input
	var err error
	fmt.Print("Введите пример: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	s := strings.Split(text, " ")

	if len(s) != 3 {
		panic(errors.New("вывод ошибки, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)"))
	}
	operand1, operation, operand2 := s[0], s[1], s[2]

	if IsRomanNumerals(operand1) != IsRomanNumerals(operand2) {
		panic(errors.New("вывод ошибки, так как используются одновременно разные системы счисления"))
	}

	// convert to int
	var numberOperand1, numberOperand2 int
	flagIsRomanNum := IsRomanNumerals(operand1)
	if flagIsRomanNum {
		numberOperand1 = ConvertRomanToArabNumberls(operand1)
		numberOperand2 = ConvertRomanToArabNumberls(operand2)
	} else {
		numberOperand1, err = strconv.Atoi(operand1)
		if err != nil {
			panic(err)
		}

		numberOperand2, err = strconv.Atoi(operand2)
		if err != nil {
			panic(err)
		}
	}

	err = ValidateOperand(numberOperand1)
	if err != nil {
		panic(err)
	}

	err = ValidateOperand(numberOperand2)
	if err != nil {
		panic(err)
	}

	err = ValidateOperation(operation)
	if err != nil {
		panic(err)
	}

	result := Calculate(numberOperand1, numberOperand2, operation)

	if flagIsRomanNum {
		if result <= 0 {
			panic(errors.New("вывод ошибки, так как в римской системе нет отрицательных чисел и нуля"))
		}
		arabResult := ConvertArabToRomanNumerals(result)
		fmt.Println(arabResult)
	} else {
		fmt.Println(result)
	}
}
