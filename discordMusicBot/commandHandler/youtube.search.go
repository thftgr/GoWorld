package commandHandler

import (
	"discordMusicBot/youtubeHandler"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"strconv"
	"strings"
	"time"
)

func searchYoutube(s *discordgo.Session, m *discordgo.MessageCreate, c []string) {
	log.Println(c)
	embed := &discordgo.MessageEmbed{
		Title:       "youtube Search",
		Author:      &discordgo.MessageEmbedAuthor{},
		Color:       0x00ff00, // Green
		Description: "Description",
		Image: &discordgo.MessageEmbedImage{
			URL: "https://cdn.discordapp.com/avatars/119249192806776836/cc32c5c3ee602e1fe252f9f595f9010e.jpg?size=2048",
		},
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: "https://cdn.discordapp.com/avatars/119249192806776836/cc32c5c3ee602e1fe252f9f595f9010e.jpg?size=2048",
		},
		Timestamp: time.Now().Format(time.RFC3339), // Discord wants ISO8601; RFC3339 is an extension of ISO8601 and should be completely compatible.

	}
	yd := youtubeHandler.Search(strings.Join(c[1:], "+"), "")
	for i, item := range yd.Items {
		embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
			Name:   strconv.Itoa(i) + "." + item.Snippet.Title,
			Value:  item.Snippet.ChannelTitle,
			Inline: false,
		})
	}
	fmt.Println(s.ChannelMessageSendEmbed(m.ChannelID, embed))
}
func bold(s *string) (ss *string) {
	z := "**" + *s + "**"
	return &z
}
