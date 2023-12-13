package fight

import (
	"fmt"
	"rpg/tools"
	"time"
)

var CompteurFight int
var Atq int
var Kill1 string = "false"
var Kill2 string = "false"
var Kill3 string = "false"
var Kill4 string = "false"

func MenuFight(MobHp int) {
	var choice string
	fmt.Println("[1] Attaque")
	fmt.Println("[2] Sac à dos")
	fmt.Println("[3] Fuir")
	fmt.Scan(&choice)
	for choice != "1" && choice != "2" && choice != "3" {
		tools.Clear()
		fmt.Println("[1] Attaque")
		fmt.Println("[2] Sac à dos")
		fmt.Println("[3] Fuir")
		fmt.Scan(&choice)
	}
	if choice == "1" {
		var choix string
		ListAttack()
		for choix != "1" && choix != "2" && choix != "3" && choix != "4" {
			tools.Clear()
			ListAttack()
			fmt.Scan(&choix)
			if choix == "1" {
				Atq = Player1.Attack1
			}
			if choix == "2" {
				Atq = Player1.Attack2
			}
			if choix == "3" {
				Atq = Player1.Attack3
			}
			if choix == "4" {
				Atq = Player1.Attack4
			}
		}
	}
	if choice == "2" {
		// inventory.AfficheConsoInv()
	}
	if choice == "3" {
		tools.Clear()
		fmt.Print("La fuite n'est pas une option envisagable !")
		time.Sleep(2000 * time.Millisecond)
		MenuFight(MobHp)
	}
}

func ListAttack() {
	fmt.Println("[1]				[2]")
	fmt.Println("[3]				[4]")
}

func StartFight(MobHp int, MobUsed Mob) {
	tools.Clear()
	var Round int
	for Player1.Hp > 0 && MobHp > 0 {
		Round++
		tools.Clear()
		fmt.Println("Tu as", Player1.Hp, " HP.")
		fmt.Println(MobUsed.Name, " a", MobHp, " HP.")
		fmt.Print("\n")
		MenuFight(MobHp)
		MobHp = MobHp - Atq
		time.Sleep(1000 * time.Millisecond)
		if MobHp > 0 {
			atq := MobAttack(MobUsed)
			Player1.Hp = Player1.Hp - atq
			fmt.Print("\n")
			fmt.Println(MobUsed.Name, " vient de te faire perdre ", atq, " HP")
			fmt.Print("\n")
			time.Sleep(2000 * time.Millisecond)
		}
	}
	if Player1.Hp <= 0 {
		tools.GameOver()
	}
	if MobHp <= 0 {
		CompteurFight++
		MobUsed.HealthPoints = 0
		if MobUsed.Name == "Shadow of the Abyss" {
			Kill1 = "true"
		}
		if MobUsed.Name == "Gutripper of the Void" {
			Kill2 = "true"
		}
		if MobUsed.Name == "Abyssal Blade" {
			Kill3 = "true"
		}
		if MobUsed.Name == "Lunar Vanguard" {
			Kill4 = "true"
		}
		CinematicFight(MobUsed.Name)
		tools.Clear()
	}
}

func LastFight() {

}

func CinematicFight(mob string) {

}
