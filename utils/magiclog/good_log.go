package magiclog

import (
	"fmt"
	"os"
	"strings"
)

// FairyLog make different types of log that used for monitoring sakura
func FairyLog(q, w, e, r, t string) string {

	q1 := q
	w1 := w
	e1 := e
	r1 := r
	t1 := t

	if q == "SAKURA" {
		q1 = "\x1b[96m" + q
		t1 = "\x1b[38;5;227m" + t
	} else if q == "PICK" {
		q1 = "\x1b[93m" + q
	} else if q == "WINNER" {
		q1 = "\x1b[92m" + q
		t1 = "\x1b[38;5;227m" + t
	}

	r1 = strings.Split(r1, " ")[1]

	q1 = makeLength(q1, 12)
	w1 = makeLength(w1, 5)
	e1 = makeLength(e1, 4)
	r1 = makeLength(r1, 13)

	q = makeLength(q, 7)
	w = makeLength(w, 5)
	e = makeLength(e, 4)
	r = makeLength(r, 24)

	log := q + w + e + r + t
	colouredLog := "\x1b[1m" + q1 + "\x1b[0m" + "\x1b[38;5;198m" + w1 + "\x1b[38;5;202m" + e1 + "\x1b[38;5;212m" + r1 + "\x1b[38;5;217m" + t1

	WriteInLog(log+"\n", false)
	WriteInLog(colouredLog+"\n", true)

	return colouredLog
}

func makeLength(str string, length int) string {

	for i := len(str); i < length; i++ {
		str = str + " "
	}

	return str
}

// WriteInLog write text of log in files
func WriteInLog(text string, coloured bool) {

	var (
		filepath string
	)
	if coloured {
		filepath = "logs/colouredLog.log"
	} else {
		filepath = "logs/log.log"
	}

	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println("ERROR opening file", err.Error())
		return
	}

	defer file.Close()
	file.WriteString(text)

	return
}
