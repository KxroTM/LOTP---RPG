package main

import (
	"fmt"
	"rpg/carte"
	"rpg/event"
	"rpg/save"
	"rpg/tools"
)

func main() {
	tools.Clear()
	save.Menu()
	e2 := false
	tools.Clear()
	for e2 == false {
		carte.AfficherCarte()
		e2 = true
	}
	e3 := false
	var input string
	for e3 == false {
		fmt.Scan(&input)
		event.EventMap(input)
		event.Event()
	}

}
