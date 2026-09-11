package main

import "fmt"

func main() {
	fmt.Println("=== GESTIONNAIRE DE NOTES ===")

	var nombreNotes int
	fmt.Print("Combien de notes voulez-vous saisir ? ")
	fmt.Scan(&nombreNotes)

	notes := make([]int, nombreNotes)
	for i := range notes {
		fmt.Printf("Note %d : ", i+1)
		fmt.Scan(&notes[i])
	}
}
