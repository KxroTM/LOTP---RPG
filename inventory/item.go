package inventory

type Item struct {
	Name    string
	Type    string
	Hp      int
	Armor   int
	Dammage int
	PosX    int
	PosY    int
}

var x001 = Item{"Épée de la Foudre", "Weapon", 0, 0, 20, 17, 32}
