package mariadb
type beatmap struct {
	BeatmapId    int     `json:"beatmap_id"`
	BeatmapsetId int     `json:"beatmapset_id"`
	Artist       string  `json:"artist"`
	Title        string  `json:"title"`
	Diffname     string  `json:"Diffname"`
	Difficulty   float64 `json:"Difficulty"`
	Mode         int     `json:"Mode"`
	BPM          float64 `json:"BPM"`
	AR           float64 `json:"AR"`
	CS           float64 `json:"CS"`
	OD           float64 `json:"OD"`
	HP           float64 `json:"HP"`
	TotalLength  int     `json:"TotalLength"`
	HitLength    int     `json:"HitLength"`
	Playcount    int     `json:"Playcount"`
	Passcount    int     `json:"Passcount"`
	MaxCombo     int     `json:"MaxCombo"`
	CircleCount  int     `json:"CircleCount"`
	SpinnerCount int     `json:"SpinnerCount"`
	SliderCount  int     `json:"SliderCount"`
	Creator      string  `json:"Creator"`
	CreatorID    int     `json:"CreatorID"`
}
type score struct {
	Scoreid   string `json:"scoreid"`
	Userid    string `json:"userid"`
	Username  string `json:"username"`
	Country   string `json:"country"`
	Score     string `json:"score"`
	Rank      string `json:"rank"`
	MaxCombo  string `json:"max_combo"`
	Mods      string `json:"mods"`
	Count300  string `json:"count300"`
	Count100  string `json:"count100"`
	Count50   string `json:"count50"`
	Countmiss string `json:"countmiss"`
	Time      string `json:"time"`
	PlayMode  string `json:"play_mode"`
	Accuracy  string `json:"accuracy"`
	Pp        string `json:"pp"`
	Fc        string `json:"fc"`
}

type beatmapDF struct {
	BeatmapId  int     `json:"beatmap_id"`
	Diffname   string  `json:"Diffname"`
	Difficulty float64 `json:"Difficulty"`
	Mode       int     `json:"Mode"`
}
