package src

var ErrorCode = struct {
	ERR1 string
	ERR2 string
	ERR3 string
}{"ERR1_", "ERR2_", "ERR3_"}

type User struct {
	AvatarUrl     *string `json:"avatar_url"`
	CountryCode   *string `json:"country_code"`
	DefaultGroup  *string `json:"default_group"`
	Id            int     `json:"id"`
	IsActive      bool    `json:"is_active"`
	IsBot         bool    `json:"is_bot"`
	IsOnline      bool    `json:"is_online"`
	IsSupporter   bool    `json:"is_supporter"`
	LastVisit     *string `json:"last_visit"`
	PmFriendsOnly bool    `json:"pm_friends_only"`
	ProfileColour *string `json:"profile_colour"`
	Username      *string `json:"username"`
	CoverUrl      *string `json:"cover_url"`
	Discord       *string `json:"discord"`
	HasSupported  bool    `json:"has_supported"`
	Interests     *string `json:"interests"`
	JoinDate      *string `json:"join_date"`
	Kudosu        struct {
		Total     int `json:"total"`
		Available int `json:"available"`
	} `json:"kudosu"`
	Location     *string  `json:"location"`
	MaxBlocks    int      `json:"max_blocks"`
	MaxFriends   int      `json:"max_friends"`
	Occupation   *string  `json:"occupation"`
	PlayMode     *string  `json:"playmode"`
	PlayStyle    []string `json:"playstyle"`
	PostCount    int      `json:"post_count"`
	ProfileOrder []string `json:"profile_order"`
	Skype        *string  `json:"skype"`
	Title        *string  `json:"title"`
	TitleUrl     *string  `json:"title_url"`
	Twitter      *string  `json:"twitter"`
	Website      *string  `json:"website"`
	Country      struct {
		Code *string `json:"code"`
		Name *string `json:"name"`
	} `json:"country"`
	Cover struct {
		CustomUrl string `json:"custom_url"`
		Url       string `json:"url"`
		Id        string `json:"id"`
	} `json:"cover"`

	AccountHistory []struct {
		Id        int     `json:"id"`
		Type      *string `json:"type"`
		Timestamp *string `json:"timestamp"`
		Length    int     `json:"length"`
	} `json:"account_history"`
	ActiveTournamentBanner []struct {
		Id           int     `json:"id"`
		TournamentId int     `json:"tournament_id"`
		Image        *string `json:"image"`
	} `json:"active_tournament_banner"`
	Badges []struct {
		AwardedAt   *string `json:"awarded_at"`
		Description *string `json:"description"`
		ImageUrl    *string `json:"image_url"`
		Url         *string `json:"url"`
	} `json:"badges"`
	BeatmapPlaycountsCount   int `json:"beatmap_playcounts_count"`
	FavouriteBeatmapsetCount int `json:"favourite_beatmapset_count"`
	FollowerCount            int `json:"follower_count"`
	GraveyardBeatmapsetCount int `json:"graveyard_beatmapset_count"`
	Groups                   []struct {
		Id             *string `json:"id"`
		Identifier     *string `json:"identifier"`
		IsProbationary *string `json:"is_probationary"`
		Name           *string `json:"name"`
		ShortName      *string `json:"short_name"`
		Description    *string `json:"description"`
		Colour         *string `json:"colour"`
		Playmodes      *string `json:"playmodes"`
	} `json:"groups"`
	LovedBeatmapsetCount int `json:"loved_beatmapset_count"`

	MonthlyPlaycounts []struct {
		StartDate *string `json:"start_date"`
		Count     int     `json:"count"`
	} `json:"monthly_playcounts"`
	Page struct {
		Html string `json:"html"`
		Raw  string `json:"raw"`
	} `json:"page"`
	PreviousUsernames                []string `json:"previous_usernames"`
	RankedAndApprovedBeatmapsetCount int      `json:"ranked_and_approved_beatmapset_count"`
	ReplaysWatchedCounts             []struct {
		StartDate *string `json:"start_date"`
		Count     int     `json:"count"`
	} `json:"replays_watched_counts"`
	ScoresBestCount   int `json:"scores_best_count"`
	ScoresFirstCount  int `json:"scores_first_count"`
	ScoresRecentCount int `json:"scores_recent_count"`

	Statistics struct {
		Level struct {
			Current  int `json:"current"`
			Progress int `json:"progress"`
		} `json:"level"`

		Pp                     float64 `json:"pp"`
		PpRank                 int     `json:"pp_rank"`
		RankedScore            int64   `json:"ranked_score"`
		HitAccuracy            float64 `json:"hit_accuracy"`
		PlayCount              int     `json:"play_count"`
		PlayTime               int     `json:"play_time"`
		TotalScore             int64   `json:"total_score"`
		TotalHits              int64   `json:"total_hits"`
		MaximumCombo           int     `json:"maximum_combo"`
		ReplaysWatchedByOthers int     `json:"replays_watched_by_others"`
		IsRanked               bool    `json:"is_ranked"`
		GradeCounts            struct {
			Ss  int `json:"ss"`
			Ssh int `json:"ssh"`
			S   int `json:"s"`
			Sh  int `json:"sh"`
			A   int `json:"a"`
		} `json:"grade_counts"`
		Rank struct {
			Global  int `json:"global"`
			Country int `json:"country"`
		} `json:"rank"`
	} `json:"statistics"`
	SupportLevel            int `json:"support_level"`
	UnrankedBeatmapsetCount int `json:"unranked_beatmapset_count"`
	UserAchievements        []struct {
		AchievedAt    string `json:"achieved_at"`
		AchievementId int    `json:"achievement_id"`
	} `json:"user_achievements"`
	RankHistory struct {
		Mode *string `json:"mode"`
		Data []int   `json:"data"`
	} `json:"rankHistory"`
	RankHistory_ struct {
		Mode *string `json:"mode"`
		Data []int   `json:"data"`
	} `json:"rank_history"`
}

type ScoresBest []struct {
	Id         int      `json:"id"`
	BestId     int      `json:"best_id"`
	UserId     int      `json:"user_id"`
	Accuracy   float64  `json:"accuracy"`
	Mods       []string `json:"mods"`
	Score      int      `json:"score"`
	MaxCombo   int      `json:"max_combo"`
	Perfect    bool     `json:"perfect"`
	Statistics struct {
		Count50   int `json:"count_50"`
		Count100  int `json:"count_100"`
		Count300  int `json:"count_300"`
		CountGeki int `json:"count_geki"`
		CountKatu int `json:"count_katu"`
		CountMiss int `json:"count_miss"`
	} `json:"statistics"`
	Pp        float32 `json:"pp"`
	CreatedAt *string `json:"created_at"`
	Mode      *string `json:"mode"`
	ModeInt   int     `json:"mode_int"`
	Replay    bool    `json:"replay"`
	Beatmap   struct {
		DifficultyRating float32 `json:"difficulty_rating"`
		Id               int     `json:"id"`
		Mode             *string `json:"mode"`
		TotalLength      int     `json:"total_length"`
		Version          *string `json:"version"`
		Accuracy         float32 `json:"accuracy"`
		Ar               float32 `json:"ar"`
		BeatmapsetId     int     `json:"beatmapset_id"`
		Bpm              int     `json:"bpm"`
		Convert          bool    `json:"convert"`
		CountCircles     bool    `json:"count_circles"`
		CountSliders     bool    `json:"count_sliders"`
		CountSpinners    bool    `json:"count_spinners"`
		Cs               float32 `json:"cs"`
		DeletedAt        *string `json:"deleted_at"`
		Drain            float32 `json:"drain"`
		HitLength        int     `json:"hit_length"`
		IsScoreable      bool    `json:"is_scoreable"`
		LastUpdated      *string `json:"last_updated"`
		ModeInt          int     `json:"mode_int"`
		Passcount        int     `json:"passcount"`
		Playcount        int     `json:"playcount"`
		Ranked           int     `json:"ranked"`
		Status           *string `json:"status"`
		Url              *string `json:"url"`
	} `json:"Beatmap"`
	Beatmapset struct {
		Artist        *string `json:"artist"`
		ArtistUnicode *string `json:"artist_unicode"`
		Covers        struct {
			Cover       *string `json:"cover"`
			Cover2x     *string `json:"cover@2x"`
			Card        *string `json:"card"`
			Card2x      *string `json:"card@2x"`
			List        *string `json:"list"`
			List2x      *string `json:"list@2x"`
			SlimCover   *string `json:"artist_unicode"`
			SlimCover2x *string `json:"artist_unicode@2x"`
		} `json:"covers"`
		Creator        *string `json:"creator"`
		FavouriteCount int     `json:"favourite_count"`
		Hype           *string `json:"hype"`
		Id             int     `json:"id"`
		PlayCount      int     `json:"play_count"`
		PreviewUrl     *string `json:"preview_url"`
		Source         *string `json:"source"`
		Status         *string `json:"status"`
		Title          *string `json:"title"`
		TitleUnicode   *string `json:"title_unicode"`
		UserId         *string `json:"user_id"`
		Video          *string `json:"video"`
	} `json:"beatmapset"`
	Weight struct {
		Percentage float32 `json:"percentage"`
		Pp         int     `json:"pp"`
	} `json:"weight"`
	User struct {
		AvatarUrl     *string `json:"avatar_url"`
		CountryCode   *string `json:"country_code"`
		DefaultGroup  *string `json:"default_group"`
		Id            int     `json:"id"`
		IsActive      bool    `json:"is_active"`
		IsBot         bool    `json:"is_bot"`
		IsOnline      bool    `json:"is_online"`
		IsSupporter   bool    `json:"is_supporter"`
		LastVisit     *string `json:"last_visit"`
		PmFriendsOnly bool    `json:"pm_friends_only"`
		ProfileColour *string `json:"profile_colour"`
		Username      *string `json:"username"`
	} `json:"user"`
}

type RecentActivity []struct { // 최근 1달 플레이 기록
	Created_at *string `json:"created_at"`
	CreatedAt  *string `json:"createdAt"`
	Id         int     `json:"id"`
	Type       *string `json:"type"`
	ScoreRank  *string `json:"scoreRank"`
	Rank       int     `json:"rank"`
	Mode       *string `json:"mode"`
	beatmap    struct {
		Title *string `json:"title"`
		Url   *string `json:"url"`
	}
	user struct {
		Username *string `json:"username"`
		Url      *string `json:"url"`
	}
}
