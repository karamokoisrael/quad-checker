package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {

	//   0, 1, 2, 3, 4
	// values := [5]int{1, 2, 3, 4, 5}

	// fmt.Println(values[4])

	files, err := ioutil.ReadDir("./")

	if err != nil {
		log.Fatal(err)
	}

	nombreTotalDeFichiers := len(files)
	for i := 0; i < nombreTotalDeFichiers; i++ {
		indiceDuFichierActuel := i
		fichierActuel := files[indiceDuFichierActuel]
		nomDuFichierActuel := fichierActuel.Name()
		fmt.Println(nomDuFichierActuel)
	}

	// for i, file := range files {
	// 	name := file.Name()
	// 	fmt.Println(i)
	// 	fmt.Println(name)
	// }
}
