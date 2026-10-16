# Go-Reloaded

Petit outil en ligne de commande, écrit en Go (bibliothèque standard uniquement). Il lit un fichier texte, corrige ou transforme certains éléments, puis écrit le résultat dans un second fichier.

## Utilisation

```bash
go run . entree.txt sortie.txt
```

Le premier argument est le fichier à lire, le second est le fichier à créer.

## Transformations appliquées

| Élément dans le texte | Effet |
|---|---|
| `(hex)` | le mot précédent (nombre hexadécimal) est converti en décimal |
| `(bin)` | le mot précédent (nombre binaire) est converti en décimal |
| `(up)`, `(low)`, `(cap)` | le mot précédent passe en majuscules, en minuscules ou prend une majuscule initiale |
| `(up, n)`, `(low, n)`, `(cap, n)` | idem, sur les `n` mots précédents |
| `. , ! ? : ;` | collés au mot précédent, suivis d'un espace ; les groupes comme `...` ou `!?` restent ensemble |
| `' '` | les apostrophes sont collées au texte qu'elles entourent |
| `a` | devient `an` devant une voyelle ou un `h` |

## Organisation du code

- `main.go` : lecture des arguments, lecture et écriture des fichiers
- `resoudre.go` : enchaîne les étapes de transformation
- `cap.go` : gestion des drapeaux `(hex)`, `(bin)`, `(up)`, `(low)`, `(cap)`
- `ponctuation.go` : ponctuation et guillemets simples
- `voyelles.go` : correction `a` → `an`
- `fonction-base.go` : fonctions utilitaires
- `main/` : fichiers d'exemple (`sample.txt`, `result.txt`)

Le texte est traité ligne par ligne, dans cet ordre : drapeaux, ponctuation, guillemets, puis `a` / `an`.

## Limites connues

- Les nombres décimaux (`3.14`) et les heures (`10:30`) sont traités comme de la ponctuation et reçoivent un espace.
- Aucun test unitaire n'est encore écrit.
