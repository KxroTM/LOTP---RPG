package help

import (
	"fmt"
	"os"
	"rpg/tools"
)

func Help() {
	val := ""
	tools.Clear()
	readhelp()
	fmt.Print("\n")
	for i := 0; val != "y"; i++ {
		tools.Clear()
		readhelp()
		fmt.Print("\n")
		fmt.Print("As-tu fini de lire ? Ecris y si oui... ")
		fmt.Scan(&val)
		tools.Clear()
	}
}

func readhelp() {
	file, err := os.ReadFile("help/help.txt")

	if err != nil {
		panic(err)
	}
	fmt.Print(string(file))
}
