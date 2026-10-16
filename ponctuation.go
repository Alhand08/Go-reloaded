package main

import (
	"strings"
	"unicode"
)

func TraiterPonctuationLigne(ligne string) string {
	caracteres := []rune(ligne)
	for position := 1; position < len(caracteres); position++ {
		if EstPonctuation(caracteres[position]) {
			if position-1 >= 0 && caracteres[position-1] == ' ' {
				caracteres = SupprimerRune(caracteres, position-1)
				position -= 2
				if position < 0 {
					position = 0
				}
			}
		}
	}

	// Traiter les espaces APRÈS la ponctuation
	for position := 0; position < len(caracteres)-1; position++ {
		// Supprimer les espaces multiples après une ponctuation
		if EstPonctuation(caracteres[position]) && caracteres[position+1] == ' ' {
			for position+2 < len(caracteres) && caracteres[position+2] == ' ' {
				caracteres = SupprimerRune(caracteres, position+1)
			}
		}
		if EstPonctuation(caracteres[position]) && !EstPonctuation(caracteres[position+1]) && caracteres[position+1] != ' ' && caracteres[position+1] != '\'' && caracteres[position+1] != ')' && caracteres[position+1] != '"' {
			caracteres = AjouterEspaceApres(caracteres, position)
		}
	}

	return string(caracteres)
}

func TraiterPonctuation(texte string) string {
	lignes := strings.Split(texte, "\n")
	for numeroLigne := range lignes {
		lignes[numeroLigne] = TraiterPonctuationLigne(lignes[numeroLigne])
	}
	return strings.Join(lignes, "\n")
}

func TraiterGuillemetsSimplesLigne(ligne string) string {
	caracteres := []rune(ligne)
	guillemetOuvert := false
	positionOuverture := 0
	for position := 0; position < len(caracteres); position++ {
		// Ignorer les apostrophes au milieu des mots (ex: don't)
		if position-1 >= 0 && position+1 < len(caracteres) && caracteres[position] == '\'' && unicode.IsLetter(caracteres[position-1]) && unicode.IsLetter(caracteres[position+1]) {
			continue
		}
		if caracteres[position] == '\'' && !guillemetOuvert {
			guillemetOuvert = true
			positionOuverture = position
		} else if caracteres[position] == '\'' && guillemetOuvert {
			if positionOuverture+1 < len(caracteres) && caracteres[positionOuverture+1] == ' ' {
				caracteres = SupprimerRune(caracteres, positionOuverture+1)
				position -= 2
				if position < 0 {
					position = 0
				}
			} else if position-1 >= 0 && caracteres[position-1] == ' ' {
				caracteres = SupprimerRune(caracteres, position-1)
				position -= 2
				if position < 0 {
					position = 0
				}
			} else {
				guillemetOuvert = false
			}
		}
	}
	return string(caracteres)
}

func TraiterGuillemetsSimples(texte string) string {
	lignes := strings.Split(texte, "\n")
	for numeroLigne := range lignes {
		lignes[numeroLigne] = TraiterGuillemetsSimplesLigne(lignes[numeroLigne])
	}
	return strings.Join(lignes, "\n")
}
