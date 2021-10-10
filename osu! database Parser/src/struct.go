package src

type DataBase struct {
	OsuVersion        int32     `json:"osuVersion"`
	FolderCount       int32     `json:"folderCount"`
	AccountUnlocked   bool      `json:"accountUnlocked"`
	AccountUnlockTime uint64    `json:"accountUnlockTime"`
	Username          string    `json:"userName"`
	BeatmapsCount     int32     `json:"beatmapCount"`
	Beatmaps          []Beatmap `json:"beatmaps"`
	UserPermissions   int32
}
type Beatmap struct {
	ArtistName           string
	ArtistNameUnicode    string
	SongTitle            string
	SongTitleUnicode     string
	CreatorName          string
	Difficulty           string
	AudioFileName        string
	MD5                  string
	OsuFileName          string
	RankedStatus         string
	HitCircles           int16
	Sliders              int16
	Spinners             int16
	LastModificationTime string
	ApproachRate         float32
	CircleSize           float32
	HPDrain              float32
	OverallDifficulty    float32
	SliderVelocity       float64
	StarRatingMania      map[string]float64
	StarRatingCtb        map[string]float64
	StarRatingTaiko      map[string]float64
	StarRatingStandard   map[string]float64
	DrainTime            int32
	TotalTime            int32
	PreviewOffset        int32
	TimingPoint          []struct {
		BPM    float64
		Offset float64
		Bool   bool
	}
	BeatmapID                    int32
	BeatmapSetID                 int32
	ThreadID                     int32
	GradeAchievedInOsu           int
	GradeAchievedInTaiko         int
	GradeAchievedInCTB           int
	GradeAchievedInMania         int
	LocalBeatmapOffset           int16
	StackLeniency                float32
	OsuGameplayMode              int
	SongSource                   string
	SongTags                     string
	OnlineOffset                 int16
	FontUsedForTheTitleOfTheSong string
	UnPlayed                     bool
	LastPlayed                   int64
	IsOsz2                       bool
	BeatmapFolderName            string
	OsuRrepositoryCheckedAt      string
	IgnoreBeatmapSound           bool
	IgnoreBeatmapSkin            bool
	DisableStoryboard            bool
	DisableVideo                 bool
	VisualOverride               bool
	LastModificationTime2        int32
	ManiaScrollSpeed             int
}
