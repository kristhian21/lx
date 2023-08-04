package main

import (
	"fmt"
	"log"
	"os"
	"time"

	utils "github.com/kristhian21/lx/src/internal"
)

func main() {
	files, err := os.ReadDir(".")
	
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fileInfo, _ := file.Info()
		
		size := utils.ByteToMB(fileInfo.Size())
		name := fileInfo.Name()
		modTime := fileInfo.ModTime().Format(time.DateTime)

		var icon string

		if fileInfo.IsDir() {
			icon = string('\uf4d3') + " "
		} else {
		 	icon = string('\uf4a5') + " "
		}
		
		fmt.Println(icon, size, modTime, name)
	}
}