// - Работает с двумя операндами и одной операцией.(+, -, *, /)
// - Работает с двумя системами счисления. Арабская и Римская СИ.
// - Одна СИ в одном выражении иначе указать на ошибку и прекратить работу
// - Результат деления является целое число, остаток отбрасывается
// - Результат работы с арабскими числами мб отрицательные числа и 0
// - Результат работы с римскими числами только положительные числа, иначе ошибка и прекратить работу

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var operation []rune = []rune{'+', '-', '*', '/'}

func isHas(array []rune, element rune) bool {
	for i := 0; i < len(array); i++ {
		if array[i] == element {
			return true
		}
	}

	return false
}

func isRome(rome string) bool {
	romes := []rune{'I', 'V', 'X', 'L', 'C', 'D', 'M'}
	for i := 0; i < len(rome); i++ {
		if !isHas(romes, rune(rome[i])) {
			return false
		}
	}
	return true
}

func arabicToRome(num int) string {
	ones := [10]string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	tens := [10]string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	hunds := [10]string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	thous := [5]string{"", "M", "MM", "MMM", "MMMM"}

	//var result string

	var t string = thous[int(num/1000)]
	var h string = hunds[int(num/100%10)]
	var ten string = tens[int(num/10%10)]
	var one string = ones[num%10]

	return t + h + ten + one
}

func romeToArabic(num string) int {
	var rome_map map[rune]int = map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	var result int

	for i := 0; i < len(num); i++ {
		if i > 0 && rome_map[rune(num[i])] > rome_map[rune(num[i-1])] {
			result += rome_map[rune(num[i])] - 2*rome_map[rune(num[i-1])]
		} else {
			result += rome_map[rune(num[i])]
		}
	}
	return result
}

func calculate(fnum, snum int, oper string) int {
	var result int
	if oper == "*" {
		result = fnum * snum
	} else if oper == "/" {
		result = fnum / snum
	} else if oper == "+" {
		result = fnum + snum
	} else if oper == "-" {
		result = fnum - snum
	}

	return result
}

func parseInput(exc string) (int, string) {
	var fnum string
	var snum string
	var oper string
	parts := false
	var error string

	//Разбиваю выражение на разные части
	for i := 0; i < len(exc); i++ {
		if isHas(operation, rune(exc[i])) {
			oper += string(exc[i])
			parts = true
		} else if parts {
			snum += string(exc[i])
		} else {
			fnum += string(exc[i])
		}
		fnum = strings.TrimSpace(fnum)
		snum = strings.TrimSpace(snum)
	}
	if !parts || snum == "" || fnum == "" {
		error := "Вывод ошибки, так как строка не является математической операцией"
		return 0, error
	} else if len(oper) > 1 {
		return 0, "Вывод ошибки, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)."
	}

	//Преобразую в int, если не получится пробую считать в римской си
	fn, ferr := strconv.Atoi(fnum)
	sn, serr := strconv.Atoi(snum)

	if ferr != nil || serr != nil {
		if isRome(fnum) && isRome(snum) {
			result := calculate(romeToArabic(fnum), romeToArabic(snum), oper)
			if result > 0 {
				return 0, arabicToRome(result)
			} else {
				return 0, "Вывод ошибки, так как в римской системе нет отрицательных чисел."
			}

		} else {
			error = "Вывод ошибки, так как используются одновременно разные системы счисления."
			return 0, error
		}
	}

	return calculate(fn, sn, oper), error
}

func main() {
	// fmt.Println(arabicToRome(1))
	// fmt.Println(romeToArabic("IV"))
	// fmt.Println(parseInput("I-II"))

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Введите значение: ")
		text, _ := reader.ReadString('\n')
		num, str := parseInput(text)
		if str != "" {
			fmt.Println(str)
		} else {
			fmt.Println(num)
		}
	}
}
