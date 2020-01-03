# Sakura-bot

It is a bot that check text chat on a Discord server "AniLibria.TV" and make log or even answer on special messages that bring local money. This messages have certain structure: special text and picture with a captcha.

## Run by Docker

All that you need is run this command
```
sudo docker run -it --rm -e MAIL=email@example.com -e PASS=password -e SCEN=onlyLOG chestem/sakura_bot
```
`email@example.com` is your e-mail adress of account of Discord

`password` is password of your Discord account

`onlyLOG` is a scenery there are four types of scenary:
- `onlyLOG` (make log only)
- `onlyPUB` (pick sakura only in pub-channel)
- `onlyTAV` (pick sakura only in taverna-channel)
- `SAKURA` (pick sakura anywhere)

## Run by Source

Firstly, let's download all that we need for running bot.

### Prerequisites

You need to have:
- `golang` (1.11 or later),
- `git`,
- `tesseract`,
- `some creativity!`

### Installing

You should:
- install `Golang`
```
$ sudo snap install go --classic
```
- install `git`
```
$ sudo apt install git
```
- install `tesseract`
```
$ sudo apt install libtesseract-dev
$ sudo apt install tesseract-ocr
```
- clone `sakura_bot` repository, 
```
$ git clone https://github.com/SteMak/sakura_bot
```

### Run application

- change directory to `sakura_bot`
```
$ cd sakura_bot
```
- build an application, 
```
$ go build cmd/app/app.go
```
- run it in `bash`.
```
$ ./app -e email@example.com -p password -s onlyLOG
```

### How to run?

In `cmd/` there are some types of assembly:
- `app/app.go` - has flags `-e email` of account on Dicord, `-p password` of account on Dicord, `-s scenary` there are four types of scenary:
  - `onlyLOG` (make log only)
  - `onlyPUB` (pick sakura only in pub-channel)
  - `onlyTAV` (pick sakura only in taverna-channel)
  - `SAKURA` (pick sakura anywhere)
- `sakura_pick/sakura_pick.go` - has flags `-e email` of account on Dicord, `-p password` of account on Dicord, `-s scenary` there are three types of scenary:
  - `onlyPUB` (pick sakura only in pub-channel)
  - `onlyTAV` (pick sakura only in taverna-channel)
  - `SAKURA` (pick sakura anywhere)
- `sakura_log/sakura_log.go` - has flags `-e email` of account on Dicord, `-p password` of account on Dicord (always make only log)

### Unknown bags

If it is a problem with `leptonica headers` try to install `leptonica`
```
$ sudo apt install libleptonica-dev
```

If you have special symbols in e-mail/password print e-mail/password in `""` like
```
$ ./app -e "e*ma?il@ex$am#p%le.com" -p "pa&|ss.w*o$rd" -s onlyLOG
```

## Built With

* [DiscordGo](https://github.com/bwmarrin/discordgo) - repository which gives Discord API is used.
* [Gosseract](https://github.com/otiai10/gosseract) - Tool for parsing text.
* [Bild](https://github.com/anthonynsimon/bild) - Used to make captcha easily scanned by [Gosseract](https://github.com/otiai10/gosseract).

## Badges

[![Go Report Card](https://goreportcard.com/badge/github.com/SteMak/sakura_bot)](https://goreportcard.com/report/github.com/SteMak/sakura_bot)