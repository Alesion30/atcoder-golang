package main

import (
	"fmt"
	"strings"
)

func main() {
	var S string

	fmt.Scan(&S)

	if strings.Contains(S, "er") {
		fmt.Println("er")
	} else {
		fmt.Println("ist")
	}
}
