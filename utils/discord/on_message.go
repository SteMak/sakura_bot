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
	code1, code2  [2]string
	sakuraTime    [2]time.Time
	isProcessFree [2]bool = [2]bool{true, true}
	scenery       string
)

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	if channel.RightChannel(m.ChannelID, "onlyLOG") {

		onlyLOG(s, m)
	}
	if agroSakura() && channel.RightChannel(m.ChannelID, scenery) {

		if channel.NameOfChannel(m.ChannelID) == "PUB" {

			onlyTavPub(s, m, 0)
		}
		if channel.NameOfChannel(m.ChannelID) == "TAV" {

			onlyTavPub(s, m, 1)
		}
	}
}

func onlyLOG(s *discordgo.Session, m *discordgo.MessageCreate) {

	matchSakura, _ := regexp.Match(`^\d+ случайных 🌸 появились! Напишите `+"`.pick и код с картинки`"+`, чтобы собрать их\.$`, []byte(m.Content))
	if matchSakura && m.Author.String() == "AniLibria.TV#4439" {

		sakuraTime := timeformats.TimeByID(m.ID)
		strSTime := timeformats.StrTime(sakuraTime)
		code1, code2 := magicKodes(m, timeformats.TimeOfMessage(m))
		currency := strings.Split(m.Content, " ")[0]

		fmt.Println(magiclog.FairyLog("SAKURA", currency, channel.NameOfChannel(m.ChannelID), strSTime, code1+" "+code2))
	}

	matchPick, _ := regexp.Match(`^\.pick \w\w\w\w$`, []byte(m.Content))
	if matchPick {

		pickTime := timeformats.TimeByID(m.ID)
		strPTime := timeformats.StrTime(pickTime)
		sendedCode := strings.Split(m.Content, " ")[1]

		fmt.Println(magiclog.FairyLog("PICK", sendedCode, channel.NameOfChannel(m.ChannelID), strPTime, m.Author.Username))
	}

	if len(m.Embeds) > 0 && m.Author.String() == "AniLibria.TV#4439" {

		matchWin, _ := regexp.Match(`^\**<@!\d+>\** собрал \d+🌸$`, []byte(m.Embeds[0].Description))
		if matchWin {

			winTime := timeformats.TimeByID(m.ID)
			strWTime := timeformats.StrTime(winTime)
			winner, _ := s.User(findWinnerID(m.Embeds[0].Description))

			fmt.Println(magiclog.FairyLog("WINNER", findCountOfSakura(m.Embeds[0].Description), channel.NameOfChannel(m.ChannelID), strWTime, winner.Username))
		}
	}
}

func onlyTavPub(s *discordgo.Session, m *discordgo.MessageCreate, tavpub int) {

	matchSakura, _ := regexp.Match(`^\d+ случайных 🌸 появились! Напишите `+"`.pick и код с картинки`"+`, чтобы собрать их\.$`, []byte(m.Content))
	if matchSakura && m.Author.String() == "AniLibria.TV#4439" {

		sakuraTime[tavpub] = timeformats.TimeByID(m.ID)
		code1[tavpub], code2[tavpub] = magicKodes(m, timeformats.TimeOfMessage(m))

		agroOnSakura(s, m, tavpub)
	}

	matchPick, _ := regexp.Match(`^\.pick \w\w\w\w$`, []byte(m.Content))
	if matchPick {

		pickTime := timeformats.TimeByID(m.ID)
		sendedCode := strings.Split(m.Content, " ")[1]

		agroOnPick(s, m, pickTime, sendedCode, tavpub)
	}
}

// DefineScenery define scenery in local evironment
func DefineScenery(str string) {
	scenery = str
}

func agroSakura() bool {

	return scenery == "SAKURA" || scenery == "onlyPUB" || scenery == "onlyTAV"
}

func agroOnSakura(s *discordgo.Session, m *discordgo.MessageCreate, tavpub int) {

	time.Sleep(2500 * time.Millisecond)
	s.ChannelTyping(m.ChannelID)
	time.Sleep((2000 + time.Duration(rand.Intn(1000))) * time.Millisecond)

	if isProcessFree[tavpub] && code1[tavpub] != "" && code2[tavpub] != "" {

		isProcessFree[tavpub] = false

		sends.SendRandomMessage(s, m, code1[tavpub], code2[tavpub])
		code1[tavpub], code2[tavpub] = "", ""

		isProcessFree[tavpub] = true
	}
}

func agroOnPick(s *discordgo.Session, m *discordgo.MessageCreate, pickTime time.Time, sendedCode string, tavpub int) {

	if m.Author.ID == s.State.User.ID {

		code1[tavpub], code2[tavpub] = "", ""
	}

	if timeformats.EnoughTimeRest(pickTime, sakuraTime[tavpub]) && isProcessFree[tavpub] && code1[tavpub] != "" && code2[tavpub] != "" {

		isProcessFree[tavpub] = false

		sends.SendMessageOnOther(s, m, sendedCode, code1[tavpub], code2[tavpub])
		code1[tavpub], code2[tavpub] = "", ""

		isProcessFree[tavpub] = true
	}
}

func findWinnerID(desc string) string {

	winstr := strings.Split(desc, " ")[0]
	winstr = strings.Split(winstr, "!")[1]
	winstr = strings.Split(winstr, ">")[0]

	return winstr
}

func findCountOfSakura(desc string) string {

	countstr := strings.Split(desc, " ")[2]
	countstr = strings.Split(countstr, "🌸")[0]

	return countstr
}

func magicKodes(m *discordgo.MessageCreate, time string) (string, string) {

	if len(m.Attachments) < 1 {
		return "", ""
	}

	url.GetImageByURL((*m.Attachments[0]).URL, time)
	imagework.ConvertImage(time)
	return imagework.ParseImage(time)
}
