package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {

	files, err := ioutil.ReadDir("./")

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		name := file.Name()
		if strings.Contains(name, "quad") {
			fmt.Println(name)
		}
	}
}
