package main

import (
	"strings"
	"unicode"
)

func SupprimerRune(caracteres []rune, position int) []rune {
	if position < 0 || position >= len(caracteres) {
		return caracteres
	}
	return append(caracteres[:position], caracteres[position+1:]...)
}

func SupprimerChaine(mots []string, position int) []string {
	if position < 0 || position >= len(mots) {
		return mots
	}
	return append(mots[:position], mots[position+1:]...)
}

func Capitaliser(mot string) string {
	caracteres := []rune(mot)
	premiereLettreTrouvee := false
	for position := 0; position < len(caracteres); position++ {
		if unicode.IsLetter(caracteres[position]) && !premiereLettreTrouvee {
			caracteres[position] = unicode.ToUpper(caracteres[position])
			premiereLettreTrouvee = true
		} else {
			caracteres[position] = unicode.ToLower(caracteres[position])
		}
	}
	return string(caracteres)
}

func AjouterChaine(mots []string, position int, nouveauMot string) []string {
	resultat := make([]string, len(mots)+1)
	positionSource := 0
	for positionResultat := 0; positionResultat < len(resultat); positionResultat++ {
		if positionResultat == position {
			resultat[positionResultat] = nouveauMot
		} else {
			resultat[positionResultat] = mots[positionSource]
			positionSource++
		}
	}
	return resultat
}

func AjouterEspaceApres(caracteres []rune, position int) []rune {
	if position < 0 || position >= len(caracteres) {
		return caracteres
	}
	nouveauxCaracteres := []rune{}
	for numero := 0; numero <= position; numero++ {
		nouveauxCaracteres = append(nouveauxCaracteres, caracteres[numero])
	}
	nouveauxCaracteres = append(nouveauxCaracteres, ' ')
	for numero := position + 1; numero < len(caracteres); numero++ {
		nouveauxCaracteres = append(nouveauxCaracteres, caracteres[numero])
	}
	return nouveauxCaracteres
}

func EstPonctuation(caractere rune) bool {
	return caractere == '.' || caractere == ',' || caractere == '!' || caractere == '?' || caractere == ';' || caractere == ':'
}

func estVoyelle(mot string) bool {
	mot = strings.TrimLeft(mot, "'\"(  ")
	if mot == "" {
		return false
	}
	premiereLettre := rune(strings.ToLower(string(mot[0]))[0])
	return premiereLettre == 'a' || premiereLettre == 'e' || premiereLettre == 'i' || premiereLettre == 'o' || premiereLettre == 'u' || premiereLettre == 'h'
}
