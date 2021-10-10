package src

import (
	"os/exec"
	"strconv"
)

func NewFfmpeg() {
	run := exec.Command("./discord-osu-bot/ffmpeg/ffmpeg",
		"-i", "pipe:0",
		"-f", "s16le",
		"-ab", strconv.Itoa(fm.Bitrate),
		"-ar", fm.AudioSampleRate,
		"-ac", strconv.Itoa(fm.AudioChannels),
		"-async", "1",
		"pipe:1")
}
