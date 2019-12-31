package sends

import (
	"fmt"
	"math/rand"
	
	"github.com/bwmarrin/discordgo"
)

func SendMessageOnOther(s *discordgo.Session, m *discordgo.MessageCreate, alienCode, code1, code2 string) {

	if alienCode == code1 {
		send(s, m, code2)
	} else if alienCode == code2 {
		send(s, m, code1)
	} else {
		SendRandomMessage(s, m, code1, code2)
	}
}

func SendRandomMessage(s *discordgo.Session, m *discordgo.MessageCreate, code1, code2 string) {

	if rand.Intn(2) % 2 == 1 {
		send(s, m, code1)
	} else {
		send(s, m, code2)
	}
}

func send(s *discordgo.Session, m *discordgo.MessageCreate, code string) {

	_, err := s.ChannelMessageSend(m.ChannelID, ".pick " + code)

	if err != nil {
		fmt.Println("ERROR sending message:", err.Error())
	}
}
