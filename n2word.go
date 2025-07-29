package main

import (
	"fmt"
	"strings"
)

type Number struct {
}

func NewNumber() *Number {
	return &Number{}
}
func (obj *Number) Int2Word(num int) string {
	// trim character
	return strings.Trim(strings.Trim(strings.ReplaceAll(obj.toWord(num), "  ", " "), ", "), "-")
}

func (obj *Number) toWord(num int) string {

	unit1 := [...]string{"", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen"}
	unit2 := [...]string{"", "", "Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"}
	unit3 := [...]string{"Quadrillion", "Trillion", "Billion", "Million", "Thousand", "hundred"}
	unit3_val := [...]int{1000000000000000, 1000000000000, 1000000000, 1000000, 1000, 100}

	ret := ""
	if num < 0 {
		return ""
	} else if num < 20 {
		ret = unit1[num]
	} else if num < 100 {
		ret += unit2[num/10] + " " + unit1[num%10]
	} else {
		i := 0
		unit3_len := len(unit3_val)
		for i < unit3_len {
			d := num / unit3_val[i]
			r := num % unit3_val[i]
			if d > 0 {
				ret += obj.toWord(d) + " " + unit3[i] + " " + obj.toWord(r)
				break
			}
			i++
		}
	}

	return ret

}

func main() {
	n := NewNumber()
	fmt.Println(n.Int2Word(143300))
	//fmt.Println(n.Int2Word(30))
}
