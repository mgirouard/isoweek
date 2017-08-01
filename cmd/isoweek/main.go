package main

import (
	"fmt"
	"time"

	"github.com/mgirouard/isoweek/cmd/isoweek/flags"
)

func main() {
	then := time.Date(flags.Y, time.Month(flags.M), flags.D, 0, 0, 0, 0, time.UTC)
	_, week := then.ISOWeek()

	fmt.Printf("%d", week)
}
