package main

import (
	"fmt"

	"github.com/junkd0g/covid/lib/continent"
)

func main() {
	arr, _ := continent.GetContinentData()

	for _, x := range arr {
		fmt.Println(x)

	}
}
