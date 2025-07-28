import "fmt"

func int2word(num int) string {

	unit1 := [...]string{"", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen"}
	unit2 := [...]string{"", "", "Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"}
	unit3 := [...]string{"Quadrillion,", "Trillion,", "Billion,", "Million,", "Thousand,", "hundred"}
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
				ret += int2word(d) + " " + unit3[i] + " " + int2word(r)
				break
			}
			i++
		}
	}

	return ret
}

func main() {
	fmt.Println(int2word(99999))
}
