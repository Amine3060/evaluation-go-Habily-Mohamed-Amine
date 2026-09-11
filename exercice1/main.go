package main

import "fmt"

func afficherMenu() {
	println("=== DISTRIBUTEUR ===")
	println("1 - Eau       : 1 €")
	println("2 - Soda      : 2 €")
	println("3 - Café      : 2 €")
	println("4 - Chocolat  : 3 €")
	println("0 - Quitter")
}

func obtenirPrix(choix int) int {
	switch choix {
	case 1:
		return 1
	case 2:
		return 2
	case 3:
		return 2
	case 4:
		return 3
	}
	return 0
}

func afficherBoisson(choix int) {
	switch choix {
	case 1:
		println("Eau")
	case 2:
		println("Soda")
	case 3:
		println("Café")
	case 4:
		println("Chocolat")
	}
}

func main() {
	for {
		afficherMenu()

		var choix int
		fmt.Print("choisir une boisson : ")
		fmt.Scan(&choix)

		if choix == 0 {
			println("Au revoir !")
			return
		}

		prix := obtenirPrix(choix)
		if prix == 0 {
			println("Choix invalide !")
			continue
		}

		afficherBoisson(choix)

		var montant int
		fmt.Print("Montant inséré : ")
		fmt.Scan(&montant)

		if montant < prix {
			println("Montant insuffisant !")
			fmt.Printf("Il manque %d €.\n", prix-montant)
			continue
		}

		println("Merci !")
		fmt.Printf("Votre monnaie : %d €\n", montant-prix)
	}
}
