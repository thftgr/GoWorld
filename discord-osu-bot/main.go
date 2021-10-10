package main

import (
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"io/ioutil"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"thftgr.com/discordOsuBot/discord"
)

var Setting struct {
	Discord struct {
		Token string `json:"token"`
	} `json:"discord"`
}

func init() {
	b, err := ioutil.ReadFile("./config.json")
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("Fail to load user config")
		panic(err)
	}
	err = json.Unmarshal(b, &Setting)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("Fail to load user config")
		panic(err)
	}

}

func main() {
	dg, err := discordgo.New("Bot " + Setting.Discord.Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}
	dg.AddHandler(DiscordMessageEventListener)
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	_ = dg.Close()
}

func DiscordMessageEventListener(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	str := strings.Split(m.Content, " ")
	//m.Content
	if len(str) < 1 {
		return
	}
	switch strings.ToLower(str[0]) {
	case "--song":
		fallthrough
	case "--stop":
		if len(str) < 2 {
			_, _ = s.ChannelMessageSend(m.ChannelID, "use `--song [URL]`")
			return
		}
		g, _ := s.State.Guild(m.GuildID)
		for _, gvc := range g.VoiceStates {
			if m.Author.ID == gvc.UserID {
				discord.StreamHandler(s, str, gvc.GuildID, gvc.ChannelID)
				//TODO 이벤트 핸들러로 전달
				break
			}
		}

	case "rmn":
		g, _ := s.Guild(m.GuildID)
		d := m.Author.ID == g.OwnerID
		if d {
			_, err := s.ChannelMessageSend(m.ChannelID, "is server owner")
			if err != nil {
				fmt.Println(err)
			}
		} else {
			_, err := s.ChannelMessageSend(m.ChannelID, "is not server owner")
			if err != nil {
				fmt.Println(err)
			}
		}
	case "owo":
		_, _ = s.ChannelMessageSend(m.ChannelID, "Golang!")
	}

}
