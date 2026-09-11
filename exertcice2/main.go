package main

import "fmt"

func main() {
	fmt.Println("=== GESTIONNAIRE DE NOTES ===")

	var nombreNotes int
	fmt.Print("Combien de notes voulez-vous saisir ? ")
	fmt.Scan(&nombreNotes)
}
