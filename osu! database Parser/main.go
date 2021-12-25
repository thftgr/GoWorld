package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"thftgr.com/osuDBParser/src"
)

var dataBase = src.DataBase{}

func GetBytes(key interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(key)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func main() {

	b, err := ioutil.ReadFile("./osu!.db")
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	buffer := bytes.Buffer{}
	buffer.Write(b)
	fmt.Println(buffer.Len())
	//addr := buffer.Len()
	dataBase.OsuVersion = src.Int(&buffer)
	dataBase.FolderCount = src.Int(&buffer)
	dataBase.AccountUnlocked = src.Bool(&buffer)
	dataBase.AccountUnlockTime = src.DateTime(&buffer)
	dataBase.Username = src.String(&buffer)
	dataBase.BeatmapsCount = src.Int(&buffer)

	for i := 0; i < int(dataBase.BeatmapsCount); i++ {
		var beatMap = src.Beatmap{}

		beatMap.ArtistName = src.String(&buffer)
		beatMap.ArtistNameUnicode = src.String(&buffer)
		beatMap.SongTitle = src.String(&buffer)
		beatMap.SongTitleUnicode = src.String(&buffer)
		beatMap.CreatorName = src.String(&buffer)
		beatMap.Difficulty = src.String(&buffer)
		beatMap.AudioFileName = src.String(&buffer)
		beatMap.MD5 = src.String(&buffer)
		beatMap.OsuFileName = src.String(&buffer)
		beatMap.RankedStatus = src.RankedStatus(src.ByteInt(&buffer))
		beatMap.HitCircles = src.Short(&buffer)
		beatMap.Sliders = src.Short(&buffer)
		beatMap.Spinners = src.Short(&buffer)

		beatMap.LastModificationTime = src.WindowsTick(&buffer)

		beatMap.ApproachRate = src.Single(&buffer)
		beatMap.CircleSize = src.Single(&buffer)
		beatMap.HPDrain = src.Single(&buffer)
		beatMap.OverallDifficulty = src.Single(&buffer)
		beatMap.SliderVelocity = src.Double(&buffer)

		var difficulties []map[string]float64
		for i := 0; i < 4; i++ {
			len := src.Int(&buffer)
			diffs := map[string]float64{}
			for i := 0; i < int(len); i++ {
				src.ByteInt(&buffer)
				mode := src.Int(&buffer)
				src.ByteInt(&buffer)
				diff := src.Double(&buffer)
				diffs[src.ModsParser(mode)] = diff
			}
			difficulties = append(difficulties, diffs)
		}
		switch len(difficulties) {
		case 4:
			beatMap.StarRatingMania = difficulties[3]
			fallthrough
		case 3:
			beatMap.StarRatingCtb = difficulties[2]
			fallthrough
		case 2:
			beatMap.StarRatingTaiko = difficulties[1]
			fallthrough
		case 1:
			beatMap.StarRatingStandard = difficulties[0]
		}

		beatMap.DrainTime = src.Int(&buffer)
		beatMap.TotalTime = src.Int(&buffer)
		beatMap.PreviewOffset = src.Int(&buffer)

		type tp struct {
			BPM    float64
			Offset float64
			Bool   bool
		}

		timingPointsLength := src.Int(&buffer)

		for i := 0; i < int(timingPointsLength); i++ {
			beatMap.TimingPoint = append(beatMap.TimingPoint, tp{
				BPM:    src.Double(&buffer),
				Offset: src.Double(&buffer),
				Bool:   src.Bool(&buffer),
			})
		}

		beatMap.BeatmapID = src.Int(&buffer)
		beatMap.BeatmapSetID = src.Int(&buffer)
		beatMap.ThreadID = src.Int(&buffer)
		beatMap.GradeAchievedInOsu = src.ByteInt(&buffer)
		beatMap.GradeAchievedInTaiko = src.ByteInt(&buffer)
		beatMap.GradeAchievedInCTB = src.ByteInt(&buffer)
		beatMap.GradeAchievedInMania = src.ByteInt(&buffer)
		beatMap.LocalBeatmapOffset = src.Short(&buffer)
		beatMap.StackLeniency = src.Single(&buffer)
		beatMap.OsuGameplayMode = src.ByteInt(&buffer)
		beatMap.SongSource = src.String(&buffer)
		beatMap.SongTags = src.String(&buffer)
		beatMap.OnlineOffset = src.Short(&buffer)
		beatMap.FontUsedForTheTitleOfTheSong = src.String(&buffer)
		beatMap.UnPlayed = src.Bool(&buffer)
		beatMap.LastPlayed = src.Long(&buffer)
		beatMap.IsOsz2 = src.Bool(&buffer)
		beatMap.BeatmapFolderName = src.String(&buffer)
		beatMap.OsuRrepositoryCheckedAt = src.WindowsTick(&buffer)
		beatMap.IgnoreBeatmapSound = src.Bool(&buffer)
		beatMap.IgnoreBeatmapSkin = src.Bool(&buffer)
		beatMap.DisableStoryboard = src.Bool(&buffer)
		beatMap.DisableVideo = src.Bool(&buffer)
		beatMap.VisualOverride = src.Bool(&buffer)
		beatMap.LastModificationTime2 = src.Int(&buffer)
		beatMap.ManiaScrollSpeed = src.ByteInt(&buffer)

		dataBase.Beatmaps = append(dataBase.Beatmaps, beatMap)
	}

	dataBase.UserPermissions = src.Int(&buffer)

	//fmt.Println(dataBase)
	bb, _ := json.Marshal(dataBase)
	//
	//fmt.Println(json.Marshal(dataBase))
	//bt, _ := GetBytes(dataBase)
	err = ioutil.WriteFile("./db.json", bb, 0644)
	if err != nil {
		log.Println(err)
	}
	//fmt.Println(dataBase)

	//for i := 0; i < int(dataBase.BeatmapsCount); i++ {
	//	bz, err := json.Marshal(dataBase.Beatmaps[i])
	//	fmt.Println(string(bz))
	//	if err != nil {
	//		fmt.Println(i,err,dataBase.Beatmaps[i])
	//		return
	//	}
	//}

	//fmt.Println("User permissions", src.Int(&buffer))

}
