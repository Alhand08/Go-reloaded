package main

import (
	"strings"
	"unicode"
)

func SupprimerRune(s []rune, i int) []rune {
	if i < 0 || i >= len(s) {
		return s
	}
	return append(s[:i], s[i+1:]...)
}

func SupprimerChaine(arr []string, i int) []string {
	if i < 0 || i >= len(arr) {
		return arr
	}
	return append(arr[:i], arr[i+1:]...)
}

func Capitaliser(str string) string {
	runes := []rune(str)
	start := false
	for i := 0; i < len(runes); i++ {
		if unicode.IsLetter(runes[i]) && !start {
			runes[i] = unicode.ToUpper(runes[i])
			start = true
		} else {
			runes[i] = unicode.ToLower(runes[i])
		}
	}
	return string(runes)
}

func AjouterChaine(s []string, i int, v string) []string {
	result := make([]string, len(s)+1)
	j := 0
	for k := 0; k < len(result); k++ {
		if k == i {
			result[k] = v
		} else {
			result[k] = s[j]
			j++
		}
	}
	return result
}

func AjouterEspaceApres(s []rune, i int) []rune {
	if i < 0 || i >= len(s) {
		return s
	}
	arr := []rune{}
	for j := 0; j <= i; j++ {
		arr = append(arr, s[j])
	}
	arr = append(arr, ' ')
	for j := i + 1; j < len(s); j++ {
		arr = append(arr, s[j])
	}
	return arr
}

func EstPonctuation(r rune) bool {
	return r == '.' || r == ',' || r == '!' || r == '?' || r == ';' || r == ':'
}

func estVoyelle(str string) bool {
	str = strings.TrimLeft(str, "'\"(  ")
	if str == "" {
		return false
	}
	v := rune(strings.ToLower(string(str[0]))[0])
	return v == 'a' || v == 'e' || v == 'i' || v == 'o' || v == 'u' || v == 'h'
}