package main

func Resoudre(str string) string {
	str = TraiterOrdres(str)
	str = TraiterPonctuation(str)
	str = TraiterGuillemetsSimples(str)
	str = TraiterVoyelles(str)
	return str
}