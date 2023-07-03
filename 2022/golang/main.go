package main

import (
	"fmt"
	"jshoemaker/helpers"
)

func main() {

	fmt.Println("Helper 1:")

	joined := helpers.JoinStr("poopy", "face")
	fmt.Println(joined)

	fmt.Println(helpers.Double(2))
}
