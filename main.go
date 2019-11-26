package main

import (
	"fmt"
	"os"
	"time"
	"strconv"
	"github.com/geremachek/escape"
	fu "./funcs"
)

const help = `Usage: warlock [OPTION]...
A simple plain english clock

--help, -h: Display this information
-u: Use the U.S. standard AM/PM time
-s: Put spacing around the text
-b: Use bold text`

var (
	us bool
	space bool
	bold bool

	hour int
	xm string
	esc int = 0
)

func main() {
	argument := os.Args[1:]
	t := time.Now()

	hour = t.Hour()

	if len(argument) == 0 {
		fmt.Println(fu.ConvertHour(hour) + " " + fu.ConvertMinute(t.Minute()))
	} else {
		if argument[0] == "--help" || argument[0] == "-h" {
			fmt.Println(help)
		} else if rune(string(argument[0])[0]) == '-' {
			for _, elem := range string(argument[0])[1:] {
				if rune(elem) == 'u' {
					us = true
				} else if rune(elem) == 's' {
					space = true
				} else if rune(elem) == 'b' {
					bold = true
				}
			}

			if us {
				i, _ := strconv.Atoi(t.Format("3"))
				hour = i
				xm = t.Format(" AM")
			}

			if bold {
				esc = 1
			}

			if space {
				fmt.Print("\n " + escape.Vint(esc) + fu.ConvertHour(hour) + " " + fu.ConvertMinute(t.Minute()) + xm + escape.Vint(0) + "\n\n")
			} else {
				fmt.Println(escape.Vint(esc) + fu.ConvertHour(hour) + " " + fu.ConvertMinute(t.Minute()) + xm + escape.Vint(0))
			}

		} else {
			fmt.Println(escape.Vint(31, 1) + "Error: '" + argument[0] + "' is not a valid option" + escape.Vint(0))
		}
	}
}
