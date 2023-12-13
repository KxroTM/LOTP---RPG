package save

import (
	"fmt"
	"os"
	"rpg/carte"
	"rpg/fight"
	"rpg/intro"
	"rpg/tools"
	"strconv"
	"time"
)

func Menu() {
	var ValMenu string
	tools.Clear()
	ApparanceMenu()
	fmt.Print("\n")
	fmt.Scan(&ValMenu)
	for ValMenu != "1" && ValMenu != "2" && ValMenu != "3" {
		tools.Clear()
		ApparanceMenu()
		fmt.Print("\n")
		fmt.Print("Touche inconnue !")
		fmt.Print("\n")
		fmt.Scan(&ValMenu)
	}
	if ValMenu == "1" {
		tools.Clear()
		intro.Intro()
		carte.AfficherCarte()
	}
	if ValMenu == "2" {
		cheminDossier := "Savetxt/"
		dossier, err := os.Open(cheminDossier)
		if err != nil {
			fmt.Println("Erreur lors de l'ouverture du dossier :", err)
		}
		defer dossier.Close()
		fichiers, err := dossier.Readdir(0)
		if err != nil {
			fmt.Println("Erreur lors de la liste des fichiers du dossier :", err)
		}
		if len(fichiers) == 0 {
			fmt.Println("Le dossier est vide.")
			fmt.Print("Vous allez être renvoyé au menu principal...")
			time.Sleep(2000 * time.Millisecond)
			Menu()
		}
		StartSave()
		for carte.PosX == 0 && carte.PosY == 0 {
			tools.Clear()
			ApparanceMenu()
			fmt.Print("\n")
			fmt.Println("Sauvegarde introuvable !")
			StartSave()
		}
	}
	if ValMenu == "3" {
		tools.ExitGame()
	}
}

func ApparanceMenu() {
	tools.GameText()
	fmt.Print("\n")
	fmt.Print("\n")
	fmt.Print("\n")
	fmt.Print("[1] - Nouvelle partie")
	fmt.Print("\n")
	fmt.Print("[2] - Charger une sauvegarde")
	fmt.Print("\n")
	fmt.Print("[3] - Quitter le jeu")
	fmt.Print("\n")
}

func Pause() {
	var rep string
	tools.Clear()
	fmt.Println("			[Pause]")
	fmt.Println("[r] Reprendre")
	fmt.Println("[s] Sauvegarder")
	fmt.Print("[x] Quitter le jeu")
	fmt.Print("\n")
	fmt.Scan(&rep)
	for rep != "r" && rep != "s" && rep != "x" {
		tools.Clear()
		fmt.Println("			[Pause]")
		fmt.Println("[r] Reprendre")
		fmt.Println("[s] Sauvegarder")
		fmt.Print("[x] Quitter le jeu")
		fmt.Print("\n")
		fmt.Scan(&rep)
	}
	if rep == "r" {
		tools.Clear()
		carte.AfficherCarte()
	}
	if rep == "x" {
		tools.ExitGame()
	}
	if rep == "s" {
		tools.Clear()
		var Save string
		fmt.Print("\n")
		fmt.Print("Nom de la sauvegarde :")
		fmt.Print("\n")
		fmt.Scan(&Save)
		nomDuFichier := "Savetxt/" + Save + ".txt"
		_, err := os.Stat(nomDuFichier)
		var reponse string
		if err == nil {
			Alreadyc := Save
			tools.Clear()
			fmt.Println("Cette sauvegarde existe déjà !")
			fmt.Print("Veux-tu écraser la sauvegarde ?")
			fmt.Print("\n")
			fmt.Scan(&reponse)
			for reponse != "y" && reponse != "n" {
				tools.Clear()
				fmt.Println("Cette sauvegarde existe déjà !")
				fmt.Print("Veux-tu écraser la sauvegarde ?")
				fmt.Print("\n")
				fmt.Scan(&reponse)
			}
			if reponse == "y" {
				FichierSauvegarde, err := os.Create(nomDuFichier)
				if err != nil {
					fmt.Println("Erreur lors de la création du fichier :", err)
					return
				}
				defer FichierSauvegarde.Close()
				contenu := strconv.Itoa(carte.PosX) + "\n" + strconv.Itoa(carte.PosY) + "\n" + strconv.Itoa(fight.CompteurFight) + "\n" + fight.Kill1 + "\n" + fight.Kill2 + "\n" + fight.Kill3 + "\n" + fight.Kill4
				_, err = FichierSauvegarde.WriteString(contenu)
				if err != nil {
					fmt.Println("Erreur lors de l'écriture dans le fichier :", err)
					return
				}
			} else if reponse == "n" {
				for Alreadyc == Save {
					tools.Clear()
					fmt.Print("\n")
					fmt.Print("Nom de la sauvegarde :")
					fmt.Print("\n")
					fmt.Scan(&Save)
					nomDuFichier := "Savetxt/" + Save + ".txt"
					FichierSauvegarde, err := os.Create(nomDuFichier)
					if err != nil {
						fmt.Println("Erreur lors de la création du fichier :", err)
						return
					}
					defer FichierSauvegarde.Close()
					contenu := strconv.Itoa(carte.PosX) + "\n" + strconv.Itoa(carte.PosY) + "\n" + strconv.Itoa(fight.CompteurFight) + "\n" + fight.Kill1 + "\n" + fight.Kill2 + "\n" + fight.Kill3 + "\n" + fight.Kill4
					_, err = FichierSauvegarde.WriteString(contenu)
					if err != nil {
						fmt.Println("Erreur lors de l'écriture dans le fichier :", err)
						return
					}

				}
			}
		} else if os.IsNotExist(err) {
			FichierSauvegarde, err := os.Create(nomDuFichier)
			if err != nil {
				fmt.Println("Erreur lors de la création du fichier :", err)
				return
			}
			defer FichierSauvegarde.Close()
			contenu := strconv.Itoa(carte.PosX) + "\n" + strconv.Itoa(carte.PosY) + "\n" + strconv.Itoa(fight.CompteurFight) + "\n" + fight.Kill1 + "\n" + fight.Kill2 + "\n" + fight.Kill3 + "\n" + fight.Kill4
			_, err = FichierSauvegarde.WriteString(contenu)
			if err != nil {
				fmt.Println("Erreur lors de l'écriture dans le fichier :", err)
				return
			}
		}
		tools.Clear()
		fmt.Println("\n")
		fmt.Println("Sauvegarde en cours ...")
		fmt.Println("5%")
		time.Sleep(1000 * time.Millisecond)
		tools.Clear()
		fmt.Println("\n")
		fmt.Println("Sauvegarde en cours ...")
		fmt.Println("11%")
		time.Sleep(1000 * time.Millisecond)
		tools.Clear()
		fmt.Println("\n")
		fmt.Println("Sauvegarde en cours ...")
		fmt.Println("17%")
		time.Sleep(1000 * time.Millisecond)
		tools.Clear()
		fmt.Println("\n")
		fmt.Println("Sauvegarde en cours ...")
		fmt.Println("26%")
		time.Sleep(1000 * time.Millisecond)
		tools.Clear()
		fmt.Println("\n")
		fmt.Println("Sauvegarde en cours ...")
		fmt.Println("33%")
		time.Sleep(1000 * time.Millisecond)
		tools.Clear()
		fmt.Println("\n")
		fmt.Println("Sauvegarde en cours ...")
		fmt.Println("49%")
		time.Sleep(850 * time.Millisecond)
		tools.Clear()
		fmt.Println("\n")
		fmt.Println("Sauvegarde en cours ...")
		fmt.Println("60%")
		time.Sleep(800 * time.Millisecond)
		tools.Clear()
		fmt.Println("\n")
		fmt.Println("Sauvegarde en cours ...")
		fmt.Println("73%")
		time.Sleep(750 * time.Millisecond)
		tools.Clear()
		fmt.Println("\n")
		fmt.Println("Sauvegarde en cours ...")
		fmt.Println("88%")
		time.Sleep(750 * time.Millisecond)
		tools.Clear()
		fmt.Println("\n")
		fmt.Println("Sauvegarde en cours ...")
		fmt.Println("97%")
		time.Sleep(600 * time.Millisecond)
		tools.Clear()
		fmt.Println("\n")
		fmt.Println("Sauvegarde en cours ...")
		fmt.Println("100%")
		time.Sleep(500 * time.Millisecond)
		tools.Clear()
		fmt.Println("\n")
		fmt.Println("Sauvegarde terminé !")
		time.Sleep(1000 * time.Millisecond)
		Pause()
	}
}

func StartSave() {

	var Choix string
	fmt.Print("\n")
	fmt.Print("Quelle sauvegarde voulez-vous charger ? (Nom exact sans le '.txt')")
	fmt.Print("\n")
	Dossier := "Savetxt/"
	dir, err := os.Open(Dossier)
	if err != nil {
		fmt.Println(err)
	}
	defer dir.Close()
	fichiers, err := dir.Readdir(0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(Dossier)
	for _, fichier := range fichiers {
		fmt.Println(fichier.Name())
	}
	fmt.Print("\n")
	fmt.Scan(&Choix)
	carte.PosX = tools.ReadFile(Choix+".txt", 1)
	carte.PosY = tools.ReadFile(Choix+".txt", 2)
	fight.CompteurFight = tools.ReadFile(Choix+".txt", 3)
	fight.Kill1 = tools.ReadFileS(Choix+".txt", 4)
	fight.Kill2 = tools.ReadFileS(Choix+".txt", 5)
	fight.Kill3 = tools.ReadFileS(Choix+".txt", 6)
	fight.Kill4 = tools.ReadFileS(Choix+".txt", 7)

}
