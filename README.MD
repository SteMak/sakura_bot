# Sakura-bot

It is a bot that check text chat on a Discord server "AniLibria.TV" and make log or even answer on special messages that bring local money. This messages have certain structure: special text and picture with captcha.

## Getting Started

Firstly, let's download all that we need for running bot.

### Prerequisites

You need to have:
- git,
- golang,
- tesseract,
- some cretivity!

### Installing

You should:
- install Golang
```
$ sudo apt install golang-go
```
- install git
```
$ sudo apt install git
```
- install tesseract
```
$ sudo apt install libtesseract-dev
```
- install leptonica
```
$ sudo apt install libleptonica-dev
```
- get Go sources of my repository
```
$ go get "github.com/SteMak/sakura_bot/cmd/app"
```
- clone my repository, 
```
$ git clone https://github.com/SteMak/sakura_bot
```
### Run application
- change directory to sakura_bot
```
$ cd sakura_bot
```
- build application, 
```
$ go build cmd/app/app.go
```
- run it in bash.
```
$ ./app -e email@example.com -p password -s onlyLOG
```

## Built With

* [DiscordGo](https://github.com/bwmarrin/discordgo) - repository which gives Discord API is used.
* [Gosseract](https://github.com/otiai10/gosseract) - Tool for parsing text.
* [Bild](https://github.com/anthonynsimon/bild) - Used to make captcha easily scanned by [Gosseract](https://github.com/otiai10/gosseract).
