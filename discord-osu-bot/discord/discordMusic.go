package discord

import (
	"encoding/binary"
	"fmt"
	"github.com/bwmarrin/dgvoice"
	"github.com/bwmarrin/discordgo"
	"github.com/kkdai/youtube/v2"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

//var voiceChannelList = map[string]*discordgo.VoiceConnection{}

func StreamHandler(s *discordgo.Session, msg []string, gid, cID string) {
	client := youtube.Client{}
	end := make(chan bool)

	video, err := client.GetVideo(msg[1])
	if err != nil {
		panic(err)
	}
	fm := youtube.Format{}
	for _, format := range video.Formats {
		if !strings.Contains(format.MimeType, "audio") {
			continue
		}
		if fm.Bitrate < format.Bitrate {
			fm = format
		}
	}
	v, _ := s.ChannelVoiceJoin(gid, cID, false, true)
	//voiceChannelList[gid] = v
	defer v.Disconnect()
	defer v.Close()
	//defer delete(voiceChannelList, gid)

	run := exec.Command("./ffmpeg/ffmpeg",
		"-i", "pipe:0",
		"-f", "s16le",
		"-ab", strconv.Itoa(fm.Bitrate),
		"-ar", fm.AudioSampleRate,
		"-ac", strconv.Itoa(fm.AudioChannels),
		"-async", "1",
		"pipe:1")

	ffmpegin, err := run.StdinPipe()
	if err != nil {
		fmt.Println("StdoutPipe Error", err)
		return
	}

	ffmpegout, err := run.StdoutPipe()
	if err != nil {
		fmt.Println("StdoutPipe Error", err)
		return
	}

	if err := run.Start(); err != nil {
		fmt.Println("RunStart Error", err)
		return
	}

	defer run.Process.Kill()
	defer ffmpegout.Close()

	go func() {
		buf := make([]byte, 1920)
		for {
			res, _, err := client.GetStream(video, &fm)
			if err != nil {
				log.Println(err)
				time.Sleep(time.Millisecond * 10)
				continue
			}
			for {
				n, err := res.Read(buf)
				ffmpegin.Write(buf[:n])
				if err != nil {
					break
				}
			}
			if fm.InitRange != nil {
				break
			}
		}
		end <- true

	}()

	send := make(chan []int16, 2)
	defer close(send)

	go func() {
		dgvoice.SendPCM(v, send)
		end <- true
		fmt.Println("dgvoice.SendPCM(v, send)")
	}()
	for {
		// read data from ffmpeg stdout
		audiobuf := make([]int16, 1920)

		if binary.Read(ffmpegout, binary.LittleEndian, &audiobuf) != nil {
			log.Println(err)
			break
		}

		select {
		case send <- audiobuf:
		case <-end:
			break
		}
	}

}
