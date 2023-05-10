package main

import (
	"bytes"
	"embed"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"text/template"

	"github.com/sethvargo/go-githubactions"
)

type LambdaFunction struct {
	FunctionName string
	RoleName     string
	RoleArm      string
	EnvKeys      map[string]string
	SourceFile   string
}

//go:embed templates/*
var fs embed.FS

func createFunction(lambdaFunction LambdaFunction) {
	functionFolder := githubactions.GetInput("functionFolder")

	fmt.Println(functionFolder)

	tmpl, err := template.New("lambda.tf.stuff").ParseFS(fs, "templates/*")
	if err != nil {
		panic(err)
	}

	var tpl bytes.Buffer

	if err := tmpl.Execute(&tpl, lambdaFunction); err != nil {
		fmt.Println(err)
	}

	result := tpl.String()

	infraPath := ".infra"

	if _, err := os.Stat(infraPath); os.IsNotExist(err) {
		if err := os.Mkdir(infraPath, 0755); err != nil {
			log.Fatal(err)

			return
		}
	}

	if err := os.WriteFile(fmt.Sprintf(
		"./%v/%v.tf",
		infraPath,
		lambdaFunction.FunctionName),
		[]byte(result), 0644); err != nil {

		log.Fatal(err)
	}

	// data, err := os.ReadFile("./templates/main.tf.stuff")

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// if err := os.WriteFile(fmt.Sprintf(
	// 	"./%v/%v.tf",
	// 	infraPath,
	// 	"main.tf"),
	// 	data, 0644); err != nil {

	// 	fmt.Println(err)
	// }

	fmt.Print("Success")
}

func main2() {
	lambdaFunctions := []LambdaFunction{}

	// define the path to the functions folder
	functionsPath := "functions"

	// define the target directory name
	targetDirName := ".infra"

	// use filepath.Walk to traverse the functions folder and its subdirectories
	err := filepath.Walk(functionsPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// if the current path is the top-level functions directory, skip it
		if path == functionsPath {
			return nil
		}

		if info.IsDir() {
			// create the target directory with the desired name
			targetDirPath := filepath.Join(targetDirName, info.Name())

			if err := os.MkdirAll(targetDirPath, os.ModePerm); err != nil {
				return err
			}

			lambdaFunctions = append(lambdaFunctions, LambdaFunction{
				FunctionName: info.Name(),
				RoleName:     fmt.Sprintf("lambda_%v_role_3", info.Name()),
				RoleArm:      "aws:arn",
				SourceFile:   fmt.Sprintf("%v/main.go", info.Name()),
				EnvKeys:      map[string]string{},
			})
		}

		return nil
	})

	if err != nil {
		fmt.Println(err)
	}

	if len(lambdaFunctions) > 0 {
		for _, f := range lambdaFunctions {
			createFunction(f)
		}
	}
}

func main() {
	fmt.Println("HI Actions")

	err := filepath.Walk("functions", func(path string, info os.FileInfo, err error) error {

		fmt.Println(info.Name())

		return nil
	})

	if err != nil {
		fmt.Println(err)
	}

}
