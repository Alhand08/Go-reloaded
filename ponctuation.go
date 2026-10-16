package main

import (
	"strings"
	"unicode"
)

func TraiterPonctuationLigne(str string) string {
	arr := []rune(str)

	for i := 1; i < len(arr); i++ {
		if EstPonctuation(arr[i]) {
			if i-1 >= 0 && arr[i-1] == ' ' {
				arr = SupprimerRune(arr, i-1)
				i -= 2
				if i < 0 {
					i = 0
				}
			}
		}
	}

	for i := 0; i < len(arr)-1; i++ {
		if EstPonctuation(arr[i]) && arr[i+1] == ' ' {
			for i+2 < len(arr) && arr[i+2] == ' ' {
				arr = SupprimerRune(arr, i+1)
			}
		}
		// Ajouter un espace après la ponctuation si nécessaire
		if EstPonctuation(arr[i]) && !EstPonctuation(arr[i+1]) && arr[i+1] != ' ' && arr[i+1] != '\'' && arr[i+1] != ')' && arr[i+1] != '"' {
			arr = AjouterEspaceApres(arr, i)
		}
	}

	return string(arr)
}

func TraiterPonctuation(str string) string {
	lines := strings.Split(str, "\n")
	for i := range lines {
		lines[i] = TraiterPonctuationLigne(lines[i])
	}
	return strings.Join(lines, "\n")
}

func TraiterGuillemetsSimplesLigne(str string) string {
	r := []rune(str)
	start := false
	li := 0
	for i := 0; i < len(r); i++ {
		// Ignorer les apostrophes au milieu des mots (ex: don't)
		if i-1 >= 0 && i+1 < len(r) && r[i] == '\'' && unicode.IsLetter(r[i-1]) && unicode.IsLetter(r[i+1]) {
			continue
		}
		if r[i] == '\'' && !start {
			start = true
			li = i
		} else if r[i] == '\'' && start {
			if li+1 < len(r) && r[li+1] == ' ' {
				r = SupprimerRune(r, li+1)
				i -= 2
				if i < 0 {
					i = 0
				}
			} else if i-1 >= 0 && r[i-1] == ' ' {
				r = SupprimerRune(r, i-1)
				i -= 2
				if i < 0 {
					i = 0
				}
			} else {
				start = false
			}
		}
	}
	return string(r)
}

func TraiterGuillemetsSimples(str string) string {
	lines := strings.Split(str, "\n")
	for i := range lines {
		lines[i] = TraiterGuillemetsSimplesLigne(lines[i])
	}
	return strings.Join(lines, "\n")
}