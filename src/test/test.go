package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(getMonth())
}

func getMonth() int {
	loc, _ := time.LoadLocation("America/Bogota")
	t := time.Now().In(loc)
	return int(t.Month())
}
