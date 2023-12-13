package inventory

import (
	"fmt"
	"rpg/tools"
)

var Inventaire []string

func AfficheInv() { //affiche l'inventaire
	var val string
	for val != "y" {
		tools.Clear()
		fmt.Print(Inventaire)
		fmt.Print("\n")
		fmt.Print("Pour fermer l'inventaire appuyez sur y... ")
		fmt.Scan(&val)
		tools.Clear()
	}
}

func AddInv(item Item) { //ajoute un item à l'inventaire
	Inventaire = append(Inventaire, item.Name)
	fmt.Print(item.Name, " vient d'être ajouté à votre inventaire")
}

// func AfficheConsoInv() {
// 	var Invconso []string
// 	for i := 0; i < len(Inventaire); i++ {
// 		for range Inventaire {
// 			if Inventaire[i].Type == "consommable" {
// 				Invconso = append(Invconso, Inventaire[i])
// 			}
// 		}
// 	}
// 	fmt.Print(Invconso)
// }
