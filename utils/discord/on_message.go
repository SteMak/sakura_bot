package discord

import (
	"time"
	"fmt"
	"math/rand"
	"regexp"
	"strings"

	"sakurabot/utils/magic_log"
	"sakurabot/utils/sends"
	"sakurabot/utils/channel"
	"sakurabot/utils/time_formats"
	"sakurabot/utils/image_work"
	"sakurabot/utils/get_by_URL"

	"github.com/bwmarrin/discordgo"
)

var (
	code1, code2 string = "", ""
	sakuraTime time.Time
	isProcessFree bool = true
	scenery string
)

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	matchSakura, _ := regexp.Match(`^\d+ случайных 🌸 появились! Напишите ` + "`.pick и код с картинки`" + `, чтобы собрать их\.$`, []byte(m.Content))
	if matchSakura && channel.RightChannel(m.ChannelID, scenery) && m.Author.String() == "AniLibria.TV#4439" {

		sakuraTime = time_formats.TimeByID(m.ID)
		strSTime := time_formats.StrTime(sakuraTime)
		code1, code2 = magicKodes(m, time_formats.TimeOfMessage(m))

		currency := strings.Split(m.Content, " ")[0]
		fmt.Println(magic_log.FairyLog("SAKURA", currency, channel.ChannelName(m.ChannelID), strSTime, code1 + " " + code2))

		if agroSakura() {

			time.Sleep(2500 * time.Millisecond)
			s.ChannelTyping(m.ChannelID)
			time.Sleep((2000 + time.Duration(rand.Intn(1000))) * time.Millisecond)

			if isProcessFree && code1 != "" && code2 != "" {

				isProcessFree = false

				sends.SendRandomMessage(s, m, code1, code2)
				code1, code2 = "", ""

				isProcessFree = true
			}
		}
	}

	matchPick, _ := regexp.Match(`^\.pick \w\w\w\w$`, []byte(m.Content))
	if matchPick && channel.RightChannel(m.ChannelID, scenery) {

		pickTime := time_formats.TimeByID(m.ID)
		strPTime := time_formats.StrTime(pickTime)

		sendedCode := strings.Split(m.Content, " ")[1]
		
		fmt.Println(magic_log.FairyLog("PICK", sendedCode, channel.ChannelName(m.ChannelID), strPTime, m.Author.Username))

		if agroSakura() {

			if m.Author.ID == s.State.User.ID {

				code1, code2 = "", ""
			}
			if time_formats.EnoughTimeRest(pickTime, sakuraTime) && isProcessFree && code1 != "" && code2 != "" {

				isProcessFree = false

				sends.SendMessageOnOther(s, m, sendedCode, code1, code2)
				code1, code2 = "", ""

				isProcessFree = true
			}
		}
	}
	
	if len(m.Embeds) > 0 && m.Author.String() == "AniLibria.TV#4439" {

		matchWin, _ := regexp.Match(`^\**<@!\d+>\** собрал \d+🌸$`, []byte(m.Embeds[0].Description))
		if matchWin && channel.RightChannel(m.ChannelID, scenery) && m.Embeds[0].Type == "rich" {

			winTime := time_formats.TimeByID(m.ID)
			strWTime := time_formats.StrTime(winTime)

			winner, _ := s.User(findWinnerID(m.Embeds[0].Description))

			fmt.Println(magic_log.FairyLog("WINNER", "", channel.ChannelName(m.ChannelID), strWTime, winner.Username))
		}
	}
}

func DefineScenery(str string) {
	scenery = str
}

func agroSakura() bool {
	
	return scenery == "SAKURA" || scenery == "onlyPUB" || scenery == "onlyTAV"
}

func findWinnerID(desc string) string {
	winstr := strings.Split(desc, " ")[0]
	winstr = strings.Split(winstr, "!")[1]
	winstr = strings.Split(winstr, ">")[0]

	return winstr
}

func magicKodes(m *discordgo.MessageCreate, time string) (string, string) {

	get_by_URL.GetImageByURL((*m.Attachments[0]).URL, time)
	image_work.ConvertImage(time)
	return image_work.ParseImage(time)
}
