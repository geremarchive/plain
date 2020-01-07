package main

import (
	"fmt"
	"os"
	"time"
	"strconv"
	"github.com/geremachek/escape"
	flag "github.com/spf13/pflag"
	fu "warlock/funcs"
)

const help = `Usage: warlock [OPTION]...
A simple plain english clock

--help, -h: Display this information
--us, -u: Use the U.S. standard AM/PM time
--spaced, -s: Put spacing around the text
--bold, -b: Use bold text`

var (
	hour int
	xm string
	esc string = ""
	esce string = ""
)

func main() {
	t := time.Now()

	hour = t.Hour()

	if len(os.Args) == 1 {
		convmin, err := fu.ConvertMinute(t.Minute())

		if err != nil {
			panic(err)
		}

		convhour, err := fu.ConvertHour(hour)

		if err != nil {
			panic(err)
		}

		fmt.Println(convhour + " " + convmin)
	} else if os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Println(help)
	} else {
		us := flag.BoolP("us", "u", false, "Use the U.S. standard AM/PM time")
		space := flag.BoolP("spaced", "s", false, "Put spacing around the text")
		bold := flag.BoolP("bold", "b", false, "Use bold text")

		flag.Parse()

		if *us {
			i, err := strconv.Atoi(t.Format("3"))

			if err != nil {
				panic(err)
			}

			hour = i
			xm = t.Format(" PM")
		}

		if *bold {
			esc = escape.Vint(1)
			esce = escape.Vint(0)
		}

		convmin, err := fu.ConvertMinute(t.Minute())

		if err != nil {
			panic(err)
		}

		convhour, err := fu.ConvertHour(hour)

		if err != nil {
			panic(err)
		}

		if *space {
			fmt.Print("\n " + esc + convhour + " " + convmin + xm + esce + "\n\n")
		} else {
			fmt.Println(esc + convhour + " " + convmin + xm + esce)
		}
	}
}
