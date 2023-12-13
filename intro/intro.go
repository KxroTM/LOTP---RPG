package intro

import (
	"fmt"
	"rpg/tools"
	"time"
)

var Pseudo string

func Intro() { //Commence l'intro du jeu
	tools.Clear()
	tools.DesignIntro()
	fmt.Print("\n")
	fmt.Print("Bonjour nouveau joueur, nouvelle joueuse et bienvenue à Windhelm.")
	fmt.Print("\n")
	fmt.Print("Avant de commencer j'aimerai d'abord connaître ton nom : ")
	fmt.Scan(&Pseudo)
	tools.Clear()

	tools.DesignIntro()
	fmt.Print("\n")
	fmt.Print("Très bien, je vois, c'est donc toi ", Pseudo, " celui dont tout le monde parle...")
	fmt.Print("\n")
	fmt.Print("Notre monde vivait en paix grâce aux cacahuètes sacrées qui étaient gardées en sécurité dans un temple légendaire...")
	time.Sleep(10000 * time.Millisecond)
	tools.Clear()

	tools.DesignIntro()
	fmt.Print("\n")
	fmt.Print("Cependant, un jour, Grimfang, le Loup des Abysses, a réussi à dérober les cacahuètes sacrées, plongeant le monde dans le chaos.")
	fmt.Print("\n")
	fmt.Print("Et vous êtes le seul à être doté de pouvoirs de cacahuète, je vous en supplie aider nous à retrouver les cacahuètes et rétablir la paix.")
	time.Sleep(11000 * time.Millisecond)
	tools.Clear()

	tools.DesignIntro()
	fmt.Print("\n")
	fmt.Print("...")
	time.Sleep(2000 * time.Millisecond)
	tools.Clear()

	tools.DesignIntro()
	fmt.Print("\n")
	fmt.Print("Voici déjà deux objets qui t'aideront pour ta première quête.")
	time.Sleep(4000 * time.Millisecond)
	tools.Clear()

	tools.DesignIntro()
	fmt.Print("\n")
	fmt.Print("Voici tes items :")
	fmt.Print("\n")
	fmt.Print("Vous avez recu : (arme),(potion)")
	time.Sleep(4500 * time.Millisecond)
	tools.Clear()

	tools.DesignIntro()
	fmt.Print("\n")
	fmt.Print("Chaque Boss que tu affronteras te donnera un objet. Lorsque tes points de vie seront bas, tu pourras les utiliser pour tuer le boss.")
	fmt.Print("\n")
	fmt.Print("Bonne chance à toi, jeune aventurier !")
	time.Sleep(8000 * time.Millisecond)
	tools.Clear()
}
