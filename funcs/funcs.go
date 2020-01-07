package funcs

import (
	"strconv"
	"strings"
)

func Convert0to19(num int) (out string) {
	switch num {
		case 0: out = "clock"
		case 1: out = "One"
		case 2: out = "Two"
		case 3: out = "Three"
		case 4: out = "Four"
		case 5: out = "Five"
		case 6: out = "Six"
		case 7: out = "Seven"
		case 8: out = "Eight"
		case 9: out = "Nine"
		case 10: out = "Ten"
		case 11: out = "Eleven"
		case 12: out = "Twelve"
		case 13: out = "Thirteen"
		case 14: out = "Fourteen"
		case 15: out = "Fifteen"
		case 16: out = "Sixteen"
		case 17: out = "Seventeen"
		case 18: out = "Eighteen"
		case 19: out = "Nineteen"

	}

	return
}

func Convert10s(num int) (out string, err error) {
	var prefix string
	strt := strconv.Itoa(num)

	switch rune(strt[0]) {
		case '2': prefix = "twenty"
		case '3': prefix = "thirty"
		case '4': prefix = "forty"
		case '5': prefix = "fifty"
	}

	if rune(strt[1]) == '0' {
		out = prefix
	} else {
		var i int
		i, err = strconv.Atoi(string(strt[1]))
		out = prefix + "-" + strings.ToLower(Convert0to19(i))
	}

	return

}

func ConvertHour(hour int) (out string, err error) {
	if hour < 20 {
		out = Convert0to19(hour)
	} else if hour < 25 {
		out, err = Convert10s(hour)
	}

	return
}

func ConvertMinute(minute int) (out string, err error) {
	if minute < 10 {
		out = "o'" + strings.ToLower(Convert0to19(minute))
	} else if minute < 20 {
		out = strings.ToLower(Convert0to19(minute))
	} else if minute < 60 {
		out, err = Convert10s(minute)
	}

	return
}
