// +++ умеет в +, -, *, /.
// +++ арабские и латинские ок!
// +++ 1-10 и I-X обрабатывает, латинский ноль обрабатывает
// +++ работает только с целыми числами!
// +++ работает только с числами одной группы (3 + I == panic)
// +++ выход всегда строка в соотвтетсивии с группой чисел
// +++ обрабатывает непредусмотренные операции
// +++ выводит только целые числа при делении
// +++ арабские: выводят отрицательные, ноль и положительные значения
// +++ латинские: выводят только значения равные единице и больше

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func isValidRoman(str string) bool {
	switch str {
	case "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X":
		return true
	default:
		return false
	}
}
func isValidArabic(str string) bool {
	switch str {
	case "1", "2", "3", "4", "5", "6", "7", "8", "9", "10":
		return true
	default:
		return false
	}
}

func stringRomanToArabic(str string) string {
	romanNumbers := map[string]string{
		"I":    "1",
		"II":   "2",
		"III":  "3",
		"IV":   "4",
		"V":    "5",
		"VI":   "6",
		"VII":  "7",
		"VIII": "8",
		"IX":   "9",
		"X":    "10",
	}
	return romanNumbers[str]
}
func arabicToRoman(num int) string {
	// Проверяем, что число находится в допустимом диапазоне
	if num < 1 || num > 3999 {
		panic("значение <= 0")
	}

	// Срез структур, содержащий соответствия арабских чисел и римских символов
	romanNumerals := []struct {
		value  int
		symbol string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
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

	result := ""
	// Проходим по каждому элементу в срезе римских чисел
	for _, rn := range romanNumerals {
		// Пока число больше или равно текущему арабскому значению
		for num >= rn.value {
			result += rn.symbol // Добавляем римский символ к результату
			num -= rn.value     // Уменьшаем число на арабское значение
		}
	}
	return result // Возвращаем конечный результат
}

func main() {
	var expression string //входное выражение
	var expressionSlice []string
	validMathChars := []string{"+", "-", "/", "*"}
	fmt.Scan(&expression)
	expression = strings.ToUpper(expression)
	for i, vmc := range validMathChars {
		if strings.Contains(expression, vmc) {
			expressionSlice = strings.Split(expression, vmc)
			//Валидация латинских цифр
			if isValidRoman(expressionSlice[0]) && isValidRoman(expressionSlice[1]) {
				a, _ := strconv.Atoi(stringRomanToArabic(expressionSlice[0]))
				b, _ := strconv.Atoi(stringRomanToArabic(expressionSlice[1]))
				var res int
				switch vmc { // арифметика
				case "+":
					res = a + b
				case "-":
					res = a - b
				case "/":
					res = a / b
				case "*":
					res = a * b

				}
				fmt.Println(arabicToRoman(res))
				break
				// арифметика выведенная в строку как латинские цифры
				//валидация арабских цифр
			} else if isValidArabic(expressionSlice[0]) && isValidArabic(expressionSlice[1]) {
				a, _ := strconv.Atoi(expressionSlice[0])
				b, _ := strconv.Atoi(expressionSlice[1])
				var res int
				switch vmc {
				case "+":
					res = a + b
				case "-":
					res = a - b
				case "/":
					res = a / b
				case "*":
					res = a * b
				}
				fmt.Println(strconv.Itoa(res))
				break // арифметика выведенная в строку
			} else {
				panic("ошибка ввода")
			}
		}
		//проверка на непредусмотренную операцию
		if i == 3 && strings.Count(expression, vmc) == 0 {
			panic("нет такой операции")
		}
	}
}
