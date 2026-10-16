Un outil simple de complétion, d'édition et d'autocorrection de texte écrit en Go, en utilisant uniquement la bibliothèque standard.

Fonctionnalités:
Hex → Décimal : convertit en décimal le nombre hexadécimal placé avant (hex).
Bin → Décimal : convertit en décimal le nombre binaire placé avant (bin).
Majuscules / Minuscules / Capitalisation : gère (up), (low), (cap) ainsi que les modificateurs multi-mots comme (up, 2).
Mise en forme de la ponctuation : espace correctement la ponctuation . , ! ? : ; et gère les groupes de ponctuation comme ... ou !?.
Apostrophes : met en forme 'mot' ou 'plusieurs mots' correctement.
Correction grammaticale : remplace a par an devant une voyelle ou un « h ».


Utilisation:
bash
./go-reloaded sample.txt result.txt

ou directement avec Go :

dans le terminal 

go run . sample.txt result.txt
sample.txt – fichier d'entrée contenant le texte à modifier
result.txt – fichier de sortie contenant le texte transformé
Exemple

Entrée (sample.txt) :

it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom,...

Exécution :

dans le terminal: 

./go-reloaded sample.txt result.txt

Sortie (sample.txt → result.txt) :

It was the best of times, it was the worst of TIMES,...

Structure du projet:


Go-Reloaded/
├── main.go              # Point d'entrée
├── go.mod
├── resoudre.go          # Enchaîne toutes les transformations
├── cap.go               # Drapeaux (hex), (bin), (up), (low), (cap)
├── ponctuation.go       # Ponctuation et guillemets simples
├── voyelles.go          # Correction a → an
├── fonction-base.go     # Fonctions utilitaires
└── main/
    ├── sample.txt       # Exemple d'entrée
    └── result.txt       # Exemple de sortie