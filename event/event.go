package event

import (
	"fmt"
	"rpg/carte"
	"rpg/fight"
	"rpg/help"
	"rpg/inventory"
	"rpg/save"
	"rpg/tools"
	"time"
)

func EventMap(input string) {
	if input == "i" {
		inventory.AfficheInv()
		carte.AfficherCarte()
	} else if input == "z" || input == "q" || input == "s" || input == "d" {
		carte.Mouvement(input)
	} else if input == "h" {
		help.Help()
		carte.AfficherCarte()
	} else if input == "p" {
		save.Pause()
	} else {
		fmt.Print("Touche inconnue !")
	}
}

func Event() {
	if carte.PosX == BigBoss.PosMobX && carte.PosY == BigBoss.PosMobY {
		tools.Clear()
		var enter string
		fmt.Println("Veux tu entrer ?")
		fmt.Println("[y] Oui ")
		fmt.Println("[n] Non ")
		fmt.Scan(&enter)
		for enter != "y" && enter != "n" {
			tools.Clear()
			fmt.Println("Veux tu entrer ?")
			fmt.Println("[y] Oui ")
			fmt.Println("[n] Non ")
			fmt.Scan(&enter)
		}
		if enter == "y" {
			if EventBossFinal() == true {
				fight.LastFight()
			} else {
				tools.Clear()
				carte.PosY = carte.PosY + 1
				carte.AfficherCarte()
				fmt.Println("Tu ne peux rentrer, tu n'as pas assez de clé.")
				fmt.Print("Il te manque ", 4-fight.CompteurFight, " clés...")
				fmt.Print("\n")
			}
		}
		if enter == "n" {
			carte.PosY = carte.PosY + 1
			tools.Clear()
			carte.AfficherCarte()
		}
	}
	if EventPos1() == true {
		fight.StartFight(fight.Shadow.HealthPoints, fight.Shadow)
		carte.AfficherCarte()
	}
	if EventPos2() == true {

		fight.StartFight(fight.Gutripper.HealthPoints, fight.Gutripper)
		carte.AfficherCarte()
	}
	if EventPos3() == true {
		fight.StartFight(fight.AbyssalBlade.HealthPoints, fight.AbyssalBlade)
		carte.AfficherCarte()
	}
	if EventPos4() == true {
		fight.StartFight(fight.LunarVanguard.HealthPoints, fight.LunarVanguard)
		carte.AfficherCarte()
	}
	if EventPos5() == true {
		if fight.CompteurFight == 1 {
			tools.Clear()
			fmt.Println("Vous venez de vous faire Heal !")
			fmt.Println("Vous aviez ", fight.Player1.Hp, " Hp")
			fight.Player1.Hp = 100
			fmt.Print("Vous avez maintenant ", fight.Player1.Hp, " HP.")
			time.Sleep(3000 * time.Millisecond)
			tools.Clear()
			carte.AfficherCarte()
		}
		if fight.CompteurFight == 2 {
			tools.Clear()
			fmt.Println("Vous venez de vous faire Heal !")
			fmt.Println("Vous aviez ", fight.Player1.Hp, " Hp")
			fight.Player1.Hp = 120
			fmt.Print("Vous avez maintenant ", fight.Player1.Hp, " HP.")
			time.Sleep(3000 * time.Millisecond)
			tools.Clear()
			carte.AfficherCarte()
		}
		if fight.CompteurFight == 3 {
			tools.Clear()
			fmt.Println("Vous venez de vous faire Heal !")
			fmt.Println("Vous aviez ", fight.Player1.Hp, " Hp")
			fight.Player1.Hp = 140
			fmt.Print("Vous avez maintenant ", fight.Player1.Hp, " HP.")
			time.Sleep(3000 * time.Millisecond)
			tools.Clear()
			carte.AfficherCarte()
		}
		if fight.CompteurFight == 4 {
			tools.Clear()
			fmt.Println("Vous venez de vous faire Heal !")
			fmt.Println("Vous aviez ", fight.Player1.Hp, " Hp")
			fight.Player1.Hp = 160
			fmt.Print("Vous avez maintenant ", fight.Player1.Hp, " HP.")
			time.Sleep(3000 * time.Millisecond)
			tools.Clear()
			carte.AfficherCarte()
		}
	}
}

func EventPos1() bool {
	if carte.PosX == Boss1.PosMobX && carte.PosY == Boss1.PosMobY && fight.Kill1 == "false" {
		return true
	}
	return false
}

func EventPos2() bool {
	if carte.PosX == Boss2.PosMobX && carte.PosY == Boss2.PosMobY && fight.Kill2 == "false" {
		if fight.Kill1 == "false" {
			tools.Clear()
			fmt.Print("Tu dois d'abord aller battre le premier Boss !")
			time.Sleep(3000 * time.Millisecond)
			tools.Clear()
			carte.PosY = carte.PosY + 1
			carte.AfficherCarte()
			return false
		}
		return true
	}
	return false
}

func EventPos3() bool {
	if carte.PosX == Boss3.PosMobX && carte.PosY == Boss3.PosMobY && fight.Kill3 == "false" {
		if fight.Kill2 == "false" {
			tools.Clear()
			fmt.Print("Tu dois d'abord aller battre le second Boss !")
			time.Sleep(3000 * time.Millisecond)
			tools.Clear()
			carte.PosY = carte.PosY + 1
			carte.AfficherCarte()
			return false
		}
		return true
	}
	return false
}

func EventPos4() bool {
	if carte.PosX == Boss4.PosMobX && carte.PosY == Boss4.PosMobY && fight.Kill4 == "false" {
		if fight.Kill3 == "false" {
			tools.Clear()
			fmt.Print("Tu dois d'abord aller battre le troisième Boss !")
			time.Sleep(3000 * time.Millisecond)
			tools.Clear()
			carte.PosY = carte.PosY + 1
			carte.AfficherCarte()
			return false
		}
		return true
	}
	return false
}

func EventPos5() bool {
	if carte.PosX == Pnj.PosMobX && carte.PosY == Pnj.PosMobY {
		return true
	}
	return false
}

func EventBossFinal() bool {
	if fight.CompteurFight == 4 {
		return true
	}
	return false
}
