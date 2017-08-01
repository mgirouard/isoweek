package flags

import (
	"flag"
	"time"
)

var Y, M, D int

func init() {
	now := time.Now()

	flag.IntVar(&Y, "year", now.Year(), "Specify the year")
	flag.IntVar(&M, "month", int(now.Month()), "Specify the month")
	flag.IntVar(&D, "day", now.Day(), "Specify the day")

	flag.Parse()
}
