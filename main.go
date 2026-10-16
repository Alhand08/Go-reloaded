package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	arguments := os.Args[1:]
	if len(arguments) != 2 {
		fmt.Println("Invalid Number of arguments expected 2.")
		return
	}
	cheminEntree := arguments[0]
	cheminSortie := arguments[1]
	if !strings.HasSuffix(cheminSortie, ".txt") || !strings.HasSuffix(cheminEntree, ".txt") {
		fmt.Println("Files Extentions need to be .txt")
		return
	}
	fichierEntree, erreurOuverture := os.Open(cheminEntree)
	if erreurOuverture != nil {
		fmt.Println("Error in opening file :", erreurOuverture)
		return
	}
	defer fichierEntree.Close()
	contenu, erreurLecture := io.ReadAll(fichierEntree)
	if erreurLecture != nil {
		fmt.Println("Error in read the file :", erreurLecture)
		return
	}
	texte := string(contenu)
	texteCorrige := Resoudre(texte)
	fichierSortie, erreurCreation := os.Create(cheminSortie)
	if erreurCreation != nil {
		fmt.Println("Error in creating file :", erreurCreation)
		return
	}
	defer fichierSortie.Close()
	_, erreurEcriture := fichierSortie.WriteString(texteCorrige)

	if erreurEcriture != nil {
		fmt.Println("Error in Writing in the file :", erreurEcriture)
		return
	}
}
