package main

import "strings"

func TraiterVoyelles(str string) string {
	lines := strings.Split(str, "\n")
	for i := 0; i < len(lines); i++ {
		arr := strings.Split(lines[i], " ")
		for j := 0; j < len(arr)-1; j++ {
			if (arr[j] == "a" || arr[j] == "A") && estVoyelle(arr[j+1]) {
				arr[j] = arr[j] + "n"
			}
		}
		lines[i] = strings.Join(arr, " ")
	}
	return strings.Join(lines, "\n")
}