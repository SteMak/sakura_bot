package time_formats

import (
	"time"
	"fmt"
	"strconv"

	"github.com/bwmarrin/discordgo"
)

func TimeOfMessage(m *discordgo.MessageCreate) string {

	messageTime, _ := m.Timestamp.Parse()
	messageTimeStamp := strconv.FormatInt(messageTime.Unix(), 10)
	return messageTimeStamp
}

func TimeByID(id string) time.Time {

	const offset = 1420070400000

	n, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		fmt.Println("ERROR getting time by message:", err.Error())
	}

	timestamp := n>>22 + offset

	var (
		s = int64(timestamp / 1000)
		ns = int64(timestamp % 1000 * 1000000)
	)

	return time.Unix(s, ns)
}

func EnoughTimeRest(a, b time.Time) bool {
	return a.UnixNano() - b.UnixNano() > 2500000000
} 

func StrTime(t time.Time) string {
	year, month, day := t.Date()
	hours, minutes, seconds := t.Clock()
	nanoseconds := int(t.UnixNano() / 1000000) % 1000

	stryear := makeZeros(strconv.Itoa(year), 4, true)
	strmonth := makeZeros(strconv.Itoa(int(month)), 2, true)
	strday := makeZeros(strconv.Itoa(day), 2, true)
	strhours := makeZeros(strconv.Itoa(hours), 2, true)
	strminutes := makeZeros(strconv.Itoa(minutes), 2, true)
	strseconds := makeZeros(strconv.Itoa(seconds), 2, true)
	strnanoseconds := makeZeros(strconv.Itoa(nanoseconds), 3, false)

	return stryear + "." + strmonth + "." + strday + " " + strhours + ":" + strminutes + ":" + strseconds + "." + strnanoseconds 
}

func makeZeros(str string, length int, revert bool) string {

	for i := len(str); i < length; i++ {
		if revert {
			str = "0" + str
		} else {
			str = str + "0"
		}
	}

	return str
}
