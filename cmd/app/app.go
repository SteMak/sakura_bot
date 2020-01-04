package main

import (
	"flag"
	"fmt"
	"time"
	"math/rand"

	"github.com/SteMak/sakura_bot/utils/discord"
)

// Variables used for command line parameters
var (
	Email   string
	Passw   string
	Scenery string
)

func init() {

	flag.StringVar(&Email, "e", "", "Email")
	flag.StringVar(&Passw, "p", "", "Password")
	flag.StringVar(&Scenery, "s", "", "Scenery of work (onlyLOG/onlyPUB/onlyTAV/SAKURA)")

	flag.Parse()

	if Scenery != "onlyPUB" && Scenery != "onlyTAV" && Scenery != "SAKURA" {
		Scenery = "onlyLOG"
	}

	fmt.Println("Your scenery:", Scenery)

	rand.Seed(time.Now().Unix())
}

func main() {
	discord.DefineScenery(Scenery)
	discord.MakeDiscordSession(Email, Passw)
}
