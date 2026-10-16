package main

import "strings"

func TraiterVoyelles(texte string) string {
	lignes := strings.Split(texte, "\n")
	for numeroLigne := 0; numeroLigne < len(lignes); numeroLigne++ {
		mots := strings.Split(lignes[numeroLigne], " ")
		for numeroMot := 0; numeroMot < len(mots)-1; numeroMot++ {
			if (mots[numeroMot] == "a" || mots[numeroMot] == "A") && estVoyelle(mots[numeroMot+1]) {
				mots[numeroMot] = mots[numeroMot] + "n"
			}
		}
		lignes[numeroLigne] = strings.Join(mots, " ")
	}
	return strings.Join(lignes, "\n")
}
