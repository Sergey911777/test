package main

import (
	"fmt"
	"time"
)

func main() {
	gde := " В Санкт-Петербурге"

	vremayNow := time.Now()
	fmt.Println(vremayNow.Format("02-01-2006 15:04:05"), gde)
}
