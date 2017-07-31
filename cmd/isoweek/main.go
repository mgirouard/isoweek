package main

import (
	"fmt"
	"time"
)

func main() {
	_, week := time.Now().ISOWeek()
	fmt.Printf("%d", week)
}
