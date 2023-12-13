package fight

import "rpg/intro"

type Player struct {
	Name    string
	Hp      int
	Attack1 int
	Attack2 int
	Attack3 int
	Attack4 int
}

var Player1 = Player{intro.Pseudo, 100, 15, 10, 20, 30}
