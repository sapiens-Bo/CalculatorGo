package main

import "fmt"

func arabicToRome(num int) string {
	ones := []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	tens := []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	hunds := []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	thous := []string{"", "M", "MM", "MMM", "MMMM"}

	//var result string

	var t string = thous[int(num/1000)]
	var h string = hunds[int(num/100%10)]
	var ten string = tens[int(num/10%10)]
	var one string = ones[num%10]

	return t + h + ten + one
}

func main() {
	fmt.Println(arabicToRome(1))
}
