package discord

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/SteMak/sakura_bot/utils/magic_log"

	"github.com/bwmarrin/discordgo"
)

func MakeDiscordSession(email, passw string) {

	fmt.Println("1 Autorisation:", email)

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New(email, passw)
	if err != nil {
		fmt.Println("ERROR creating Discord session:", err)
		return
	}

	fmt.Println("2 Discord session created")

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	fmt.Println("3 Registred the messageCreate func")

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("ERROR opening connection:", err)
		return
	}

	fmt.Println("4 Opened a websocket")

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	magic_log.WriteInLog("\n\nNEW SESSION STARTED\n\n", true)
	magic_log.WriteInLog("\n\nNEW SESSION STARTED\n\n", false)

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}
