package discord

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"time"

	"github.com/SteMak/sakura_bot/utils/channel"
	"github.com/SteMak/sakura_bot/utils/imagework"
	"github.com/SteMak/sakura_bot/utils/magiclog"
	"github.com/SteMak/sakura_bot/utils/sends"
	"github.com/SteMak/sakura_bot/utils/timeformats"
	"github.com/SteMak/sakura_bot/utils/url"

	"github.com/bwmarrin/discordgo"
)

var (
	code1, code2  string = "", ""
	sakuraTime    time.Time
	isProcessFree bool = true
	scenery       string
)

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	matchSakura, _ := regexp.Match(`^\d+ случайных 🌸 появились! Напишите `+"`.pick и код с картинки`"+`, чтобы собрать их\.$`, []byte(m.Content))
	if matchSakura && channel.RightChannel(m.ChannelID, scenery) && m.Author.String() == "AniLibria.TV#4439" {

		sakuraTime = timeformats.TimeByID(m.ID)
		strSTime := timeformats.StrTime(sakuraTime)
		code1, code2 = magicKodes(m, timeformats.TimeOfMessage(m))

		currency := strings.Split(m.Content, " ")[0]
		fmt.Println(magiclog.FairyLog("SAKURA", currency, channel.NameOfChannel(m.ChannelID), strSTime, code1+" "+code2))

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

		pickTime := timeformats.TimeByID(m.ID)
		strPTime := timeformats.StrTime(pickTime)

		sendedCode := strings.Split(m.Content, " ")[1]

		fmt.Println(magiclog.FairyLog("PICK", sendedCode, channel.NameOfChannel(m.ChannelID), strPTime, m.Author.Username))

		if agroSakura() {

			if m.Author.ID == s.State.User.ID {

				code1, code2 = "", ""
			}
			if timeformats.EnoughTimeRest(pickTime, sakuraTime) && isProcessFree && code1 != "" && code2 != "" {

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

			winTime := timeformats.TimeByID(m.ID)
			strWTime := timeformats.StrTime(winTime)

			winner, _ := s.User(findWinnerID(m.Embeds[0].Description))

			fmt.Println(magiclog.FairyLog("WINNER", "", channel.NameOfChannel(m.ChannelID), strWTime, winner.Username))
		}
	}
}

// DefineScenery define scenery in local evironment
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

	if len(m.Attachments) < 1 {
		return "", ""
	}

	url.GetImageByURL((*m.Attachments[0]).URL, time)
	imagework.ConvertImage(time)
	return imagework.ParseImage(time)
}
