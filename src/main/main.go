package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	files, err := os.ReadDir(".")
	
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fileInfo, _ := file.Info()
		var typeIcon string

		if fileInfo.IsDir() {
			typeIcon = string('\uf4d3') + " "
		} else {
			typeIcon = string('\uf4a5') + " "
		}
		
		fmt.Println(typeIcon, fileInfo.Size(), fileInfo.Name())
	}
}