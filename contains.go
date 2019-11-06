package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hell you")

	var command = "walk outsid"

	var exit = strings.Contains(command, "walk")

	fmt.Println("hell", exit)

}
