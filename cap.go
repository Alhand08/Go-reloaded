package main

import (
	"strconv"
	"strings"
)

func ExtrairePonctuation(texte string) (string, string) {
	ponctuationFinale := ""
	for len(texte) > 0 {
		dernierCaractere := texte[len(texte)-1]
		if dernierCaractere == '.' || dernierCaractere == ',' || dernierCaractere == '!' || dernierCaractere == '?' || dernierCaractere == ';' || dernierCaractere == ':' {
			ponctuationFinale = string(dernierCaractere) + ponctuationFinale
			texte = texte[:len(texte)-1]
		} else {
			break
		}
	}
	return texte, ponctuationFinale
}

func EstDrapeauAvecNombre(texte string) (bool, string, int) {
	nomsAcceptes := []string{"up", "low", "cap"}

	if strings.HasPrefix(texte, "(") && strings.HasSuffix(texte, ")") {
		contenu := strings.TrimSuffix(strings.TrimPrefix(texte, "("), ")")
		morceaux := strings.SplitN(contenu, ",", 2)

		if len(morceaux) == 2 {
			nomDrapeau := strings.TrimSpace(morceaux[0])
			nombreTexte := strings.TrimSpace(morceaux[1])

			for _, nomAccepte := range nomsAcceptes {
				if nomDrapeau == nomAccepte {
					if nombre, erreur := strconv.Atoi(nombreTexte); erreur == nil {
						return true, nomAccepte, nombre
					}
				}
			}
		}
	}

	return false, "", -1
}

func EstDrapeauSimple(texte string) bool {
	return texte == "(hex)" || texte == "(bin)" || texte == "(up)" || texte == "(low)" || texte == "(cap)"
}

func estMotValide(mot string) bool {
	if strings.TrimSpace(mot) == "" {
		return false
	}
	for _, caractere := range mot {
		if caractere != '.' && caractere != ',' && caractere != '!' && caractere != '?' && caractere != ';' && caractere != ':' {
			return true
		}
	}
	return false
}

func appliquerOrdres(mots []string) []string {
	for position := 0; position < len(mots)-1; position++ {
		if strings.HasPrefix(mots[position], "(") && strings.HasSuffix(mots[position], ",") {
			motSuivantNettoye, _ := ExtrairePonctuation(mots[position+1])
			if strings.HasSuffix(motSuivantNettoye, ")") {
				mots[position] = mots[position] + " " + mots[position+1]
				mots = SupprimerChaine(mots, position+1)
			}
		}
	}

	positionDebut := 0
	for len(mots) > 0 && positionDebut < len(mots) {
		motNettoye, _ := ExtrairePonctuation(mots[positionDebut])
		estDrapeauNombre, _, _ := EstDrapeauAvecNombre(motNettoye)

		if !(EstDrapeauSimple(motNettoye) || estDrapeauNombre) && mots[positionDebut] != "" {
			break
		}

		if mots[positionDebut] == "" {
			positionDebut++
			continue
		}
		mots = SupprimerChaine(mots, positionDebut)
	}

	for position := 1; position < len(mots); position++ {
		motNettoye, ponctuationFinale := ExtrairePonctuation(mots[position])

		if EstDrapeauSimple(motNettoye) {
			positionCible := position - 1
			for positionCible >= 0 && !estMotValide(mots[positionCible]) {
				positionCible--
			}
			if positionCible >= 0 {
				switch motNettoye {
				case "(hex)":
					if valeur, erreur := strconv.ParseInt(mots[positionCible], 16, 64); erreur == nil {
						mots[positionCible] = strconv.Itoa(int(valeur))
					}
				case "(bin)":
					if valeur, erreur := strconv.ParseInt(mots[positionCible], 2, 64); erreur == nil {
						mots[positionCible] = strconv.Itoa(int(valeur))
					}
				case "(up)":
					mots[positionCible] = strings.ToUpper(mots[positionCible])
				case "(low)":
					mots[positionCible] = strings.ToLower(mots[positionCible])
				case "(cap)":
					mots[positionCible] = Capitaliser(mots[positionCible])
				}
				if ponctuationFinale != "" {
					mots[positionCible] = mots[positionCible] + ponctuationFinale
				}
			}
			mots = SupprimerChaine(mots, position)
			position--
		} else if estDrapeauNombre, nomDrapeau, nombre := EstDrapeauAvecNombre(motNettoye); estDrapeauNombre {
			if nombre > 0 {
				motsModifies := 0
				positionDernierMotModifie := -1
				for positionMot := position - 1; positionMot >= 0 && motsModifies < nombre; positionMot-- {
					if estMotValide(mots[positionMot]) {
						switch nomDrapeau {
						case "up":
							mots[positionMot] = strings.ToUpper(mots[positionMot])
						case "low":
							mots[positionMot] = strings.ToLower(mots[positionMot])
						case "cap":
							mots[positionMot] = Capitaliser(mots[positionMot])
						}
						if positionDernierMotModifie == -1 {
							positionDernierMotModifie = positionMot
						}
						motsModifies++
					}
				}
				if positionDernierMotModifie != -1 && ponctuationFinale != "" {
					mots[positionDernierMotModifie] = mots[positionDernierMotModifie] + ponctuationFinale
				}
			}
			mots = SupprimerChaine(mots, position)
			position--
		}
	}

	return mots
}

func TraiterOrdres(texte string) string {
	lignes := strings.Split(texte, "\n")
	lignesCorrigees := []string{}
	for _, ligne := range lignes {
		mots := strings.Split(ligne, " ")
		mots = appliquerOrdres(mots)
		lignesCorrigees = append(lignesCorrigees, strings.Join(mots, " "))
	}

	return strings.Join(lignesCorrigees, "\n")
}
