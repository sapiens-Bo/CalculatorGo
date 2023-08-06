package main

import "fmt"

var operation [4]string = [4]string{"+", "-", "*", "/"}

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
		// "I" : 1,
		// "V" : 5,
		// "X" : 10,
		// "L" : 50,
		// "C" : 100,
		// "D" : 500,
		// "M" : 1000,
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

func main() {
	fmt.Println(arabicToRome(1))
	fmt.Println(romeToArabic("IV"))
}
