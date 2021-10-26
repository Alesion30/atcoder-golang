package main

import (
	"fmt"
	"strings"
)

func main() {
	var S string

	fmt.Scan(&S)

	if strings.HasSuffix(S, "er") {
		fmt.Println("er")
	} else {
		fmt.Println("ist")
	}
}
