package main

import (
	"strconv"
	"strings"
)

// Separe le drapeau et la ponctuation collee a la fin (ex: "(hex)," -> "(hex)", ",")
func ExtrairePonctuation(s string) (string, string) {
	punct := ""
	for len(s) > 0 {
		lastChar := s[len(s)-1]
		if lastChar == '.' || lastChar == ',' || lastChar == '!' || lastChar == '?' || lastChar == ';' || lastChar == ':' {
			punct = string(lastChar) + punct
			s = s[:len(s)-1]
		} else {
			break
		}
	}
	return s, punct
}

func EstDrapeauAvecNombre(s string) (bool, string, int) {
	paramKeys := []string{"up", "low", "cap"}

	if strings.HasPrefix(s, "(") && strings.HasSuffix(s, ")") {
		content := strings.TrimSuffix(strings.TrimPrefix(s, "("), ")")
		parts := strings.SplitN(content, ",", 2)

		if len(parts) == 2 {
			flagName := strings.TrimSpace(parts[0])
			numberStr := strings.TrimSpace(parts[1])

			for _, key := range paramKeys {
				if flagName == key {
					if nbr, err := strconv.Atoi(numberStr); err == nil {
						return true, key, nbr
					}
				}
			}
		}
	}

	return false, "", -1
}

func EstDrapeauSimple(str string) bool {
	return str == "(hex)" || str == "(bin)" || str == "(up)" || str == "(low)" || str == "(cap)"
}

func estMotValide(s string) bool {
	if strings.TrimSpace(s) == "" {
		return false
	}
	for _, r := range s {
		if r != '.' && r != ',' && r != '!' && r != '?' && r != ';' && r != ':' {
			return true
		}
	}
	return false
}

func appliquerOrdres(arr []string) []string {

	for i := 0; i < len(arr)-1; i++ {
		if strings.HasPrefix(arr[i], "(") && strings.HasSuffix(arr[i], ",") {
			p2, _ := ExtrairePonctuation(arr[i+1])
			if strings.HasSuffix(p2, ")") {
				arr[i] = arr[i] + " " + arr[i+1]
				arr = SupprimerChaine(arr, i+1)
			}
		}
	}


	counter := 0
	for len(arr) > 0 && counter < len(arr) {
		token, _ := ExtrairePonctuation(arr[counter])
		iFv, _, _ := EstDrapeauAvecNombre(token)

		if !(EstDrapeauSimple(token) || iFv) && arr[counter] != "" {
			break
		}

		if arr[counter] == "" {
			counter++
			continue
		}
		arr = SupprimerChaine(arr, counter)
	}

	for i := 1; i < len(arr); i++ {
		token, trailingPunct := ExtrairePonctuation(arr[i])

		if EstDrapeauSimple(token) {
			target := i - 1
			for target >= 0 && !estMotValide(arr[target]) {
				target--
			}
			if target >= 0 {
				switch token {
				case "(hex)":
					if nbr, err := strconv.ParseInt(arr[target], 16, 64); err == nil {
						arr[target] = strconv.Itoa(int(nbr))
					}
				case "(bin)":
					if nbr, err := strconv.ParseInt(arr[target], 2, 64); err == nil {
						arr[target] = strconv.Itoa(int(nbr))
					}
				case "(up)":
					arr[target] = strings.ToUpper(arr[target])
				case "(low)":
					arr[target] = strings.ToLower(arr[target])
				case "(cap)":
					arr[target] = Capitaliser(arr[target])
				}
				if trailingPunct != "" {
					arr[target] = arr[target] + trailingPunct
				}
			}
			arr = SupprimerChaine(arr, i)
			i--
		} else if bl, key, nbr := EstDrapeauAvecNombre(token); bl {
			if nbr > 0 {
				modifies := 0
				lastTarget := -1
				for j := i - 1; j >= 0 && modifies < nbr; j-- {
					if estMotValide(arr[j]) {
						switch key {
						case "up":
							arr[j] = strings.ToUpper(arr[j])
						case "low":
							arr[j] = strings.ToLower(arr[j])
						case "cap":
							arr[j] = Capitaliser(arr[j])
						}
						if lastTarget == -1 {
							lastTarget = j
						}
						modifies++
					}
				}
				if lastTarget != -1 && trailingPunct != "" {
					arr[lastTarget] = arr[lastTarget] + trailingPunct
				}
			}
			arr = SupprimerChaine(arr, i)
			i--
		}
	}

	return arr
}

func TraiterOrdres(str string) string {
	lines := strings.Split(str, "\n")
	finalLines := []string{}
	for _, line := range lines {
		words := strings.Split(line, " ")
		words = appliquerOrdres(words)
		finalLines = append(finalLines, strings.Join(words, " "))
	}

	return strings.Join(finalLines, "\n")
}