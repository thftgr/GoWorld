package commandHandler

import (
	"discordMusicBot/youtubeHandler"
	"fmt"
	"github.com/bwmarrin/discordgo"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"time"
)

func play(s *discordgo.Session, m *discordgo.MessageCreate, c []string) {
	//선택지 추가
	//url := `https://www.youtube.com/watch?v=D2Ekp3m3IQk&list=PLFsi9YnqVto7_Ryh3-WiTBQNB6ZV62ljd`
	if !youtubeHandler.YoutubeRegex.MatchString(c[0]) {
		fmt.Println("not youtube url")
		return
	}
	v := youtubeHandler.VideoRegix.FindAllStringSubmatch(c[0], -1)
	if !(v != nil && len(v) > 0 && len(v[0]) > 1) {
		fmt.Println("not found video url")
		return
	}
	//
	//client := youtube.Client{}
	//video, err := client.GetVideo(key)
	//if err != nil {
	//	panic(err)
	//}
	//
	//return video.Formats.WithAudioChannels() // only get videos with audio
	embed := &discordgo.MessageEmbed{
		Author:      &discordgo.MessageEmbedAuthor{},
		Color:       0x00ff00, // Green
		Description: "This is a discordgo embed",
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "I am a field",
				Value:  "I am a value",
				Inline: true,
			},
			{
				Name:   "I am a second field",
				Value:  "I am a value",
				Inline: true,
			},
		},
		Image: &discordgo.MessageEmbedImage{
			URL: "https://cdn.discordapp.com/avatars/119249192806776836/cc32c5c3ee602e1fe252f9f595f9010e.jpg?size=2048",
		},
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: "https://cdn.discordapp.com/avatars/119249192806776836/cc32c5c3ee602e1fe252f9f595f9010e.jpg?size=2048",
		},
		Timestamp: time.Now().Format(time.RFC3339), // Discord wants ISO8601; RFC3339 is an extension of ISO8601 and should be completely compatible.
		Title:     "I am an Embed",
	}

	s.ChannelMessageSendEmbed(m.ChannelID, embed)

	//s.ChannelMessageSendReply(m.ChannelID, msg, &discordgo.MessageReference{
	//	MessageID: m.ID,
	//	ChannelID: m.ChannelID,
	//	GuildID:   m.GuildID,
	//})

}
func mp4ToOpus() {

	_ = ffmpeg.Input("./tmp/in1.mp4", ffmpeg.KwArgs{"ss": 1}).
		Output("./tmp/out1.mp4", ffmpeg.KwArgs{"t": 1}).OverWriteOutput().Run()
}