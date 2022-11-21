package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)


func determineQuad(quadName ){
		
}

func main() {

	
	files, err := ioutil.ReadDir("./")
	
	if err != nil {
		log.Fatal(err)
	}

	// totalFiles := len(files)
	
	var quads [32]string

	for i, file := range files {
		name := file.Name()
		if strings.Contains(name, "quad") && name != "quadchecker" {
			quads[]
			fmt.Println(name)
		}
	}

}
