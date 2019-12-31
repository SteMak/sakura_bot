package timeformats

import (
	"fmt"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
)

// TimeOfMessage get timestamp of message
func TimeOfMessage(m *discordgo.MessageCreate) string {

	messageTime, _ := m.Timestamp.Parse()
	messageTimeStamp := strconv.FormatInt(messageTime.Unix(), 10)
	return messageTimeStamp
}

// TimeByID find time with nanoseconds using id
func TimeByID(id string) time.Time {

	const offset = 1420070400000

	n, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		fmt.Println("ERROR getting time by message:", err.Error())
	}

	timestamp := n>>22 + offset

	var (
		s  = int64(timestamp / 1000)
		ns = int64(timestamp % 1000 * 1000000)
	)

	return time.Unix(s, ns)
}

// EnoughTimeRest check if it is normal distance in time between messages
func EnoughTimeRest(a, b time.Time) bool {
	return a.UnixNano()-b.UnixNano() > 3500000000
}

// StrTime convert unix time ti string
func StrTime(t time.Time) string {

	return t.Format("01.02 15:04:05.000")
}
