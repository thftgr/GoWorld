package commandHandler

import (
	"github.com/bwmarrin/discordgo"
	"strings"
)

func IsCreated(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot || len(m.Content) < 2 || m.Content[:2] != ";;" {
		return
	}
	args := parseCommand(m.Content)
	if len(args) < 2 {
		return
	}
	switch args[0] {
	case "p", "play":
		play(s, m, args[1:])
	case "s", "stop":

	case "search":
		searchYoutube(s, m, args)

	}

}

func parseCommand(s string) (ss []string) {
	t1 := strings.Split(s[2:], " ")
	for _, i2 := range t1 {
		if i2 != "" {
			ss = append(ss, i2)
		}
	}
	return
}
