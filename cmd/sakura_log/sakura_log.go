package main

import (
	"flag"
	"fmt"

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

	flag.Parse()

	fmt.Println("Your scenery:", "onlyLOG")
}

func main() {
	discord.DefineScenery("onlyLOG")
	discord.MakeDiscordSession(Email, Passw)
}
