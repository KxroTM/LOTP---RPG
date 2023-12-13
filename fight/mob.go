package fight

import (
	"math/rand"
)

type Mob struct {
	Name         string
	HealthPoints int
	Attack1      int
	Attack2      int
	Attack3      int
	Attack4      int
}

var Shadow = Mob{"Shadow of the Abyss", 100, 15, 20, 10, 17}
var Gutripper = Mob{"Gutripper of the Void", 120, 20, 25, 30, 35}
var AbyssalBlade = Mob{"Abyssal Blade", 150, 25, 40, 30, 35}
var LunarVanguard = Mob{"Lunar Vanguard", 180, 30, 50, 40, 35}
var Grimfang = Mob{"Grimfang, le Loup des Abysses", 250, 40, 60, 55, 70}

func MobAttack(Monstre Mob) int {
	attacks := []int{Monstre.Attack1, Monstre.Attack2, Monstre.Attack3, Monstre.Attack4}
	randomAttack := attacks[rand.Intn(len(attacks))]
	return randomAttack
}
