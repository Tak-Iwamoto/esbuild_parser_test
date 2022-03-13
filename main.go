package main

import (
	"fmt"
	"github.com/evanw/esbuild"
	"log"
	"os"
)

func main() {
	content, err := os.ReadFile("Test.tsx")

	if err != nil {
		log.Fatal(err)
	}

	tree, _ := esbuild.ParseTsx(string(content))
	fmt.Println(tree.ExportKeyword)
}
