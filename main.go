package main

import (
	"fmt"

	"github.com/sethvargo/go-githubactions"
)

func main() {
	val := githubactions.GetInput("val")

	fmt.Println(val)
}
