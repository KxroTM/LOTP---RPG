package tools

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

func Clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func ExitGame() {
	Clear()
	for i := 0; i < 26; i++ {
		fmt.Print("\n")
	}
	for i := 0; i < 66; i++ {
		fmt.Print(" ")
	}
	fmt.Print("Au revoir et à bientôt...")
	for i := 0; i < 26; i++ {
		fmt.Print("\n")
	}
	os.Exit(0)
}

func GameText() {
	fmt.Print("                           ▄▄▄█████▓ ██░ ██ ▓█████     ▒█████   ██▀███  ▓█████▄ ▓█████  ██▀███      ▒█████    █████▒")
	fmt.Print("\n")
	fmt.Print("                           ▓  ██▒ ▓▒▓██░ ██▒▓█   ▀    ▒██▒  ██▒▓██ ▒ ██▒▒██▀ ██▌▓█   ▀ ▓██ ▒ ██▒   ▒██▒  ██▒▓██   ▒")
	fmt.Print("\n")
	fmt.Print("                           ▒ ▓██░ ▒░▒██▀▀██░▒███      ▒██░  ██▒▓██ ░▄█ ▒░██   █▌▒███   ▓██ ░▄█ ▒   ▒██░  ██▒▒████ ░")
	fmt.Print("\n")
	fmt.Print("                           ░ ▓██▓ ░ ░▓█ ░██ ▒▓█  ▄    ▒██   ██░▒██▀▀█▄  ░▓█▄   ▌▒▓█  ▄ ▒██▀▀█▄     ▒██   ██░░▓█▒  ░")
	fmt.Print("\n")
	fmt.Print("                             ▒██▒ ░ ░▓█▒░██▓░▒████▒   ░ ████▓▒░░██▓ ▒██▒░▒████▓ ░▒████▒░██▓ ▒██▒   ░ ████▓▒░░▒█░")
	fmt.Print("\n")
	fmt.Print("                             ▒ ░░    ▒ ░░▒░▒░░ ▒░ ░   ░ ▒░▒░▒░ ░ ▒▓ ░▒▓░ ▒▒▓  ▒ ░░ ▒░ ░░ ▒▓ ░▒▓░   ░ ▒░▒░▒░  ▒ ░")
	fmt.Print("\n")
	fmt.Print("                               ░     ▒ ░▒░ ░ ░ ░  ░     ░ ▒ ▒░   ░▒ ░ ▒░ ░ ▒  ▒  ░ ░  ░  ░▒ ░ ▒░     ░ ▒ ▒░  ░")
	fmt.Print("\n")
	fmt.Print("                             ░       ░  ░░ ░   ░      ░ ░ ░ ▒    ░░   ░  ░ ░  ░    ░     ░░   ░    ░ ░ ░ ▒   ░ ░")
	fmt.Print("\n")
	fmt.Print("                                     ░  ░  ░   ░  ░       ░ ░     ░        ░       ░  ░   ░            ░ ░")
	fmt.Print("\n")
	fmt.Print("▄▄▄█████▓ ██░ ██ ▓█████      █████▒▓█████  ██▓     ██▓ ███▄    █ ▓█████  ░  ██▓███  ▓█████ ▄▄▄       ███▄    █  █    ██ ▄▄▄█████▓  ██████")
	fmt.Print("\n")
	fmt.Print("▓  ██▒ ▓▒▓██░ ██▒▓█   ▀    ▓██   ▒ ▓█   ▀ ▓██▒    ▓██▒ ██ ▀█   █ ▓█   ▀    ▓██░  ██▒▓█   ▀▒████▄     ██ ▀█   █  ██  ▓██▒▓  ██▒ ▓▒▒██    ▒")
	fmt.Print("\n")
	fmt.Print("▒ ▓██░ ▒░▒██▀▀██░▒███      ▒████ ░ ▒███   ▒██░    ▒██▒▓██  ▀█ ██▒▒███      ▓██░ ██▓▒▒███  ▒██  ▀█▄  ▓██  ▀█ ██▒▓██  ▒██░▒ ▓██░ ▒░░ ▓██▄")
	fmt.Print("\n")
	fmt.Print("░ ▓██▓ ░ ░▓█ ░██ ▒▓█  ▄    ░▓█▒  ░ ▒▓█  ▄ ▒██░    ░██░▓██▒  ▐▌██▒▒▓█  ▄    ▒██▄█▓▒ ▒▒▓█  ▄░██▄▄▄▄██ ▓██▒  ▐▌██▒▓▓█  ░██░░ ▓██▓ ░   ▒   ██▒")
	fmt.Print("\n")
	fmt.Print("  ▒██▒ ░ ░▓█▒░██▓░▒████▒   ░▒█░    ░▒████▒░██████▒░██░▒██░   ▓██░░▒████▒   ▒██▒ ░  ░░▒████▒▓█   ▓██▒▒██░   ▓██░▒▒█████▓   ▒██▒ ░ ▒██████▒▒")
	fmt.Print("\n")
	fmt.Print("  ▒ ░░    ▒ ░░▒░▒░░ ▒░ ░    ▒ ░    ░░ ▒░ ░░ ▒░▓  ░░▓  ░ ▒░   ▒ ▒ ░░ ▒░ ░   ▒▓▒░ ░  ░░░ ▒░ ░▒▒   ▓▒█░░ ▒░   ▒ ▒ ░▒▓▒ ▒ ▒   ▒ ░░   ▒ ▒▓▒ ▒ ░")
	fmt.Print("\n")
	fmt.Print("    ░     ▒ ░▒░ ░ ░ ░  ░    ░       ░ ░  ░░ ░ ▒  ░ ▒ ░░ ░░   ░ ▒░ ░ ░  ░   ░▒ ░      ░ ░  ░ ▒   ▒▒ ░░ ░░   ░ ▒░░░▒░ ░ ░     ░    ░ ░▒  ░ ░")
	fmt.Print("\n")
	fmt.Print("  ░       ░  ░░ ░   ░       ░ ░       ░     ░ ░    ▒ ░   ░   ░ ░    ░      ░░          ░    ░   ▒      ░   ░ ░  ░░░ ░ ░   ░      ░  ░  ░")
	fmt.Print("\n")
	fmt.Print("          ░  ░  ░   ░  ░              ░  ░    ░  ░ ░           ░    ░  ░               ░  ░     ░  ░         ░    ░                    ░")
	fmt.Print("\n")
}

func DesignIntro() {
	fmt.Print("\n")
	fmt.Print("                          ░░▒█▓░░░░░                            ░░░░░░▒▓▓█░░")
	fmt.Print("\n")
	fmt.Print("                          ░░▒▓███▒▒░░░                          ░░░░░▓████░░")
	fmt.Print("\n")
	fmt.Print("                            ░░░▒▒▓█▒░░░                         ░▒░▒███▓▓░░")
	fmt.Print("\n")
	fmt.Print("                            ░██▒░░▒▓██░░░                 ░░ ░░░▓█▓▒▒▒███▓░")
	fmt.Print("\n")
	fmt.Print("                            ░███▒░░▒▓▓▒██░░░░░░ ░░░░▒░ ░░░░▓░░██▓▒▒▒▓████░")
	fmt.Print("\n")
	fmt.Print("                            ░██▒░▒▓▒▒▒▓▓▓▓███▒░ ░▒▒░▒░░▓████▓▓▓▓▒▒▓▓▓████░")
	fmt.Print("\n")
	fmt.Print("                           ░░▓█▓▓▓▓█▓▓▓█████▒▓▓▓▓███▓▓▓▓▓▓▓▓▓█▓▒▓▓█████▓░░")
	fmt.Print("\n")
	fmt.Print("                           ░▒█▓▓▓▓████▓▓▓▒▒▒▓▓▓▓█▒█▓▓█▓▒▒▒▓▓▓▓▓▓█████████░")
	fmt.Print("\n")
	fmt.Print("                            ░▒▓███▓████▓▒▒▒▒▓▓▒▓█░▒▒▒▓█▓▓▓▓▓▒▓▓████████▒▒░")
	fmt.Print("\n")
	fmt.Print("                         ░▒▒░▒▒█████▓▓▓▒▒▓▒▒██▓▓▒░░░░▓▓▓██▓▓█▓▓▓▓███▓██▓▒▒▒░░")
	fmt.Print("\n")
	fmt.Print("                               ░░▓▓▓▒▒▓▓▓▒▓▓▓▓▓▓▒░   ▓▓██▓▓▓▓███████▓░")
	fmt.Print("\n")
	fmt.Print("                            ░░░░░▒▒▒▒▒▓▓█████▓▓▒░░ ░░░▓▓█████████████▓░")
	fmt.Print("\n")
	fmt.Print("                            ░░░▒▒▒▒▓▓▓██▓▓█▒███▓░░░░░▒███▒▓█▒██████████")
	fmt.Print("\n")
	fmt.Print("                            ░▒▓██▓█▓▓▓▒▓▓▒█▒▒▓██▓░░░▒▒█▓▒▓▓█▒█▓▓███████▓▒░")
	fmt.Print("\n")
	fmt.Print("                          ░▒▒▓▓▓█▓█▓▓▒▒▒▓▒░▒██▓▓▒▒▒▒▒▓▓██▓▒░▒▓▒▓▓███▓▓███░░")
	fmt.Print("\n")
	fmt.Print("                       ░░░▒▓███▓▓▓█▓▓▓▒▒▒░░░░░░░▓▓▒▓▓▓░░░▒▒▓▓▓████████████▒░░")
	fmt.Print("\n")
	fmt.Print("                      ░░░░░░░░▒▓███████▓▓░░░░░░░░▒█▓▓░░░░▒▒▒▓██████████░░░░░░")
	fmt.Print("\n")
	fmt.Print("                             ░░░░▒▓▓▓▓▓▓▓░░░░░░▒▓▓███▒░▒▒▒▒▓████████▓░░░░░░ ░")
	fmt.Print("\n")
	fmt.Print("                          ░░░░░░░▒▒█▓▓▓▓▓▓▓▓▓▓▓▒▒▒▒▒▓▓▓▓█████████▒░░░░░░░░")
	fmt.Print("\n")
	fmt.Print("                          ░░░░░░▒▒▒▒▓▓▓▓▓▓▓▓████████████████████████▓░░░░░")
	fmt.Print("\n")
	fmt.Print("                            ░▒▒▒▒▒▒▒▒▒▓▓▓█████████████████████████████▓")
	fmt.Print("\n")
	fmt.Print("                           ░░░▒▒░░░░░▒░░░░▒▒▒▒▓▓▓▓▓▓▓██████▓██▓████████▓▒░")
	fmt.Print("\n")
	fmt.Print("                           ░░▒░░░░░░░░░░░░░░▒▒▒▒▓▓▓▓▓▓▓▓▓▓▓▒▓█▓▓▓▓▓████▓▓░")
	fmt.Print("\n")
	fmt.Print("                         ░░░░▒░░░░░░  ░░   ░░░░░░░░░░▒▓▓▓▓▒▒▒▓▒▓▓▓▓▓▓███▓░░")
	fmt.Print("\n")
	fmt.Print("                         ░▒▒▒▒░░░░░░░       ░░░░░░░░░░░▒▒▒▒▒▒▓▒▓▓▓▓▓▓█████░")
	fmt.Print("\n")
	fmt.Print("                          ░░░░░▒░  ░░░░░░░ ░░░  ░░░░░░ ░░░░▒▒▓▓▓▓▓▓▓▓██░░░")
	fmt.Print("\n")
	fmt.Print("                           ░░▒▒▓▒▒▒░░░░░░░ ░░░  ░░░░░░ ░░░▒▒▒▓▓▓▓▓██████▓░")
	fmt.Print("\n")
	fmt.Print("                           ░░░▒▓▓▒░░░░░░░░      ░░░░░░   ░▒▓▒▓▓▓▓▓████░░░░")
	fmt.Print("\n")
	fmt.Print("                           ░░░░▒▒▒▓░▓▓▒░░░░░░░░░░░░░▒░░░░░▒▓▓▓████████▓▓▒░")
	fmt.Print("\n")
	fmt.Print("                         ░░░░░░░░▒██▓▓▒▒▒░░░░░▒▒░░░▒▒░▒▒▒▓▓▓███████████░░░")
	fmt.Print("\n")
	fmt.Print("                         ░░░░▒▒░░▒▓██▓▒▒▒░░░▒░░▒░░▒▒▒▓▒▓▓▓▓▓████████████▓░░░")
	fmt.Print("\n")
	fmt.Print("                        ░▒▒▒▒▒▒▒▒▒▒▓████▓▒▒▒▓▒▒▓▓▒▓▓▒▒▒▓▓▓▓███████████████░░")
	fmt.Print("\n")
	fmt.Print("         ░░░            ░▒▒▒▒▒▒▒▒▒▒▒▓█▒█████▓▓▓▓▓▓▓▓▓▓▓▒▒▓▓▓▓▒████████████░")
	fmt.Print("\n")
	fmt.Print("    ░░░░░░░░░░░         ░░▒▒▒▓▓▒▒▒▒▓▓█▓▒▓▓▓▓███▓███▓████████▓▒███████████▒░")
	fmt.Print("\n")
	fmt.Print("   ░░░░▒▓▓███▓▓▓░░░░░    ░░▒▓█▓▓▒▒▒▓▓▓█▓▒▓▓▓▓██▓████████████▓████████████░")
	fmt.Print("\n")
	fmt.Print("  ░▒▒░▒▓▓████████▓▓▒░    ░▒▒▓█▓█▓▒▒▒▓▓█▓▒▒▒▒▓▓█████████████▓▓▓█████████")
	fmt.Print("\n")
	fmt.Print("  ░▒▒▓▓▓███████████▓░    ░▒▒▓█▓██▓▒▒▓▓█▓▒▒▒▒▒▓███████████▓▓▓▓▓████████████▓░")
	fmt.Print("\n")
	fmt.Print(" ░▒▓▓▓▓▓▒██████████▓▒▒     ░▒▓▓▓██▓▓▓██▓▓▒▒▒▓▒▓▓▓███████▓▓▒▓▓███████▓███▓░")
	fmt.Print("\n")
	fmt.Print("░▒▓▓▓▓▒▓▓▓█████████▓░░   ░░░░▓▓████▓▓███▓▒▒▒▒▒▓▓████████▓▒▒▓▓██████████░░░")
	fmt.Print("\n")
	fmt.Print("░░░▓▓▓▓▓▓▓▓██░░███░░░    ░░░░░░████████▓▓▒▒▒▒▒▓▓▓██████▓▓▒▒▓▓███████▓░░")
	fmt.Print("\n")
	fmt.Print("  ░▒▒▓▓▓▓▓█▓▒▒▒▒▒▒   ░░░ ░░░▒▒▒░░░▓██████▓▓▒▒▒▒▓▓██████▓▒▒▓▓██████▒░░")
	fmt.Print("\n")
	fmt.Print(" ░░████▓▓██████░░░░░ ░░░░░▓█▓▒▓▓▓▓███████▓▒▒▒▒▒▓██████▓▒▒▒▒▓███▒░░░")
	fmt.Print("\n")
	fmt.Print("  ░▒▒▓▓██▓▓▓▓███▓█▓▓▓▓▓▓▓█▓▓██████████▓███▒▒▒▒▒▓▓█████▒▒▒▓▓█████▒░░░")
	fmt.Print("\n")
	fmt.Print("  ░░░░░▓▓▓██▓▓▓▓▓▓█████████████████▓▒░░▒▓█▒▒░░▒▓▓█████▒▒▒▒▓███▓▓▓██▓░")
	fmt.Print("\n")
	fmt.Print("         ░░░▓████████████████████▓▓░▒▒░▒▒▓██▓░░▒▓▓███▒░▒▒▓▓█▓▓▓▓▓▓▓██░")
	fmt.Print("\n")
	fmt.Print("      ░░░  ░░▒▒▒▒▒▒▒▒▒▒▒▒▒▓▓█████▓▓▒▒▓▓▓▒▓█▓▒▒▒▒▒▓███▒▒░░▓▓█▓▓█▓▓▓▓██░░░░░")
	fmt.Print("\n")
	fmt.Print("░░░░░░░░▒▒░░░▓▓▓▓▓▓▓▒▓▓▓▓▒░▒░▒▒▒▒▒██████████████▓███████████████▓▓▓▓▓▓▓░▒▒▒▒▒")
	fmt.Print("\n")
}

func ReadFile(Fichierf string, NumeroLigneSouhaiteF int) int {
	fichier, err := os.Open("Savetxt/" + Fichierf)
	if err != nil {
		fmt.Print("Fichier introuvable !")
		fmt.Scan(&fichier)
	}
	defer fichier.Close()
	scanner := bufio.NewScanner(fichier)
	numeroLigneSouhaite := NumeroLigneSouhaiteF
	ligneActuelle := 0
	var valeur int
	for scanner.Scan() {
		ligneActuelle++
		if ligneActuelle == numeroLigneSouhaite {
			valeur, err = strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("Erreur lors de la conversion en entier:", err)
			}
			break
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Erreur lors de la lecture du fichier:", err)
	}
	return valeur
}

func ReadFileS(Fichierf string, NumeroLigneSouhaiteF int) string {
	fichier, err := os.Open("Savetxt/" + Fichierf)
	if err != nil {
		fmt.Print("Fichier introuvable !")
		fmt.Scan(&fichier)
	}
	defer fichier.Close()
	scanner := bufio.NewScanner(fichier)
	numeroLigneSouhaite := NumeroLigneSouhaiteF
	ligneActuelle := 0
	var valeur string
	for scanner.Scan() {
		ligneActuelle++
		if ligneActuelle == numeroLigneSouhaite {
			valeur = scanner.Text()
			break
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Erreur lors de la lecture du fichier:", err)
	}
	return valeur
}

func GameOver() {
	Clear()
	fmt.Print("GAME OVER")
	os.Exit(0)
}
