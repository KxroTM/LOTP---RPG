package carte

import (
	"fmt"
	"rpg/fight"
	"rpg/tools"
	"time"
)

var PosX = 50
var PosY = 44
var TailleX = 101
var TailleY = 46

func MapInit(tailleX, tailleY, posX, posY int, caractereHaut, caractereBas, caractereGauche, caractereDroit, coinHautGauche, coinHautDroit, coinBasGauche, coinBasDroit, caracterePersonnage rune) {
	carte := make([][]rune, tailleY)
	for i := range carte {
		carte[i] = make([]rune, tailleX)
		for j := range carte[i] {
			if i == 0 {
				if j == 0 {
					carte[i][j] = coinHautGauche
				} else if j == tailleX-1 {
					carte[i][j] = coinHautDroit
				} else {
					carte[i][j] = caractereHaut
				}
			} else if i == tailleY-1 {
				if j == 0 {
					carte[i][j] = coinBasGauche
				} else if j == tailleX-1 {
					carte[i][j] = coinBasDroit
				} else {
					carte[i][j] = caractereBas
				}
			} else if j == 0 {
				carte[i][j] = caractereGauche
			} else if j == tailleX-1 {
				carte[i][j] = caractereDroit
			} else {
				carte[i][j] = ' '
			}
		}
	}
	carte[44][40] = 'I'
	carte[posY][posX] = caracterePersonnage
	carte[1][50] = 'M'
	var pos1 = fight.Kill1
	var pos2 = fight.Kill2
	var pos3 = fight.Kill3
	var pos4 = fight.Kill4

	if pos1 == "false" {
		carte[34][25] = 'X'
	}
	if pos2 == "false" {
		carte[34][75] = 'X'
	}
	if pos3 == "false" {
		carte[12][15] = 'X'
	}
	if pos4 == "false" {
		carte[12][85] = 'X'
	}
	for i := range carte {
		fmt.Println(string(carte[i]))
	}
}

func AfficherCarte() {
	caractereHaut := '─'
	caractereBas := '─'
	caractereGauche := '│'
	caractereDroit := '│'
	coinHautGauche := '┌'
	coinHautDroit := '┐'
	coinBasGauche := '└'
	coinBasDroit := '┘'
	caracterePersonnage := 'O'
	MapInit(TailleX, TailleY, PosX, PosY, caractereHaut, caractereBas, caractereGauche, caractereDroit, coinHautGauche, coinHautDroit, coinBasGauche, coinBasDroit, caracterePersonnage)
	fmt.Print("[h] Pour afficher le menu d'aide")
	fmt.Print("\n")
}

func Mouvement(input string) {
	limitex := TailleX - 1
	limitey := TailleY - 1
	if (PosX >= limitex-1 && input == "d") || (PosX <= 1 && input == "q") || (PosY >= limitey-1 && input == "s") || (PosY <= 1 && input == "z") {
		fmt.Print("\n")
		fmt.Print("Vous avez atteint la bordure de la map !")
		time.Sleep(1500 * time.Millisecond)
		tools.Clear()
		AfficherCarte()

	} else {
		if input == "z" {
			PosY = PosY - 1
			tools.Clear()
			AfficherCarte()
		}
		if input == "s" {
			PosY = PosY + 1
			tools.Clear()
			AfficherCarte()
		}
		if input == "q" {
			PosX = PosX - 1
			tools.Clear()
			AfficherCarte()
		}
		if input == "d" {
			PosX = PosX + 1
			tools.Clear()
			AfficherCarte()
		}
	}
}
