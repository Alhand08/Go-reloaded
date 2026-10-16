package main

func Resoudre(texte string) string {
	texte = TraiterOrdres(texte)
	texte = TraiterPonctuation(texte)
	texte = TraiterGuillemetsSimples(texte)
	texte = TraiterVoyelles(texte)
	return texte
}
