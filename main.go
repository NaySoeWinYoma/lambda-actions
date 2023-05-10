package main

import (
	"fmt"
	"log"
	"os"

	"github.com/sethvargo/go-githubactions"
)

func main() {
	val := githubactions.GetInput("aws-key")

	if checkFolderExists("functions") {
		fmt.Println("Ok")
	}

	fmt.Println(val)
}

func checkFolderExists(folder string) bool {
	files, err := os.ReadDir(folder)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name(), file.IsDir())
	}

	return true
}
