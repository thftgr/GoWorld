package src

type Beatmap struct {
	Accuracy      float32   `json:"accuracy"`
	Ar            float32   `json:"ar"`
	BeatMapSetId  int       `json:"beatmapset_id"`
	Bpm           float32   `json:"bpm"`
	Convert       bool      `json:"convert"`
	CountCircles  int       `json:"count_circles"`
	CountSliders  int       `json:"count_sliders"`
	CountSpinners int       `json:"count_spinners"`
	Cs            float32   `json:"cs"`
	DeletedAt     Timestamp `json:"deleted_at"`
	Drain         float32   `json:"drain"`
	HitLength     int       `json:"hit_length"`
	IsScoreable   bool      `json:"is_scoreable"`
	LastUpdated   Timestamp `json:"last_updated"`
	ModeInt       int       `json:"mode_int"`
	PassCount     int       `json:"passcount"`
	PlayCount     int       `json:"playcount"`
	Ranked        int       `json:"ranked"`
	Status        string    `json:"status"`
	Url           string    `json:"url"`
}
type BeatmapCompact struct {
	DifficultyRating float32 `json:"difficulty_rating"`
	Id               int     `json:"id"`
	Mode             string  `json:"mode"`
	TotalLength      int     `json:"total_length"`
	Version          string  `json:"version"`
}

type Beatmapset struct {
	AvailabilityDownloadDisabled bool      `json:"availability.download_disabled"`
	AvailabilityMoreInformation  string    `json:"availability.more_information"`
	Bpm                          float32   `json:"bpm"`
	CanBeHyped                   bool      `json:"can_be_hyped"`
	Creator                      string    `json:"creator"`
	DiscussionEnabled            bool      `json:"discussion_enabled"`
	DiscussionLocked             bool      `json:"discussion_locked"`
	HypeCurrent                  int       `json:"hype.current"`
	HypeRequired                 int       `json:"hype.required"`
	IsScoreable                  bool      `json:"is_scoreable"`
	LastUpdated                  Timestamp `json:"last_updated"`
	LegacyThreadUrl              string    `json:"legacy_thread_url"`
	Nominationscurrent           int       `json:"nominations.current"`
	Nominationsrequired          int       `json:"nominations.required"`
	Ranked                       int       `json:"ranked"`
	RankedDate                   Timestamp `json:"ranked_date"`
	Source                       string    `json:"source"`
	Storyboard                   bool      `json:"storyboard"`
	SubmittedDate                Timestamp `json:"submitted_date"`
	Tags                         string    `json:"tags"`
}
type BeatmapsetCompact struct {
	Artist         string `json:"artist"`
	ArtistUnicode  string `json:"artist_unicode"`
	Covers         Covers `json:"covers"`
	Creator        string `json:"creator"`
	FavouriteCount int    `json:"favourite_count"`
	Id             int    `json:"id"`
	PlayCount      int    `json:"play_count"`
	PreviewUrl     string `json:"preview_url"`
	Source         string `json:"source"`
	Status         string `json:"status"`
	Title          string `json:"title"`
	TitleUnicode   string `json:"title_unicode"`
	UserId         int    `json:"user_id"`
	Video          bool   `json:"video"`
}

type Covers struct {
	Cover      string `json:"cover"`
	Cover2     string `json:"cover@2x"`
	Card       string `json:"card"`
	Card2      string `json:"card@2x"`
	List       string `json:"list"`
	List2      string `json:"list@2x"`
	SlimCover  string `json:"slimcover"`
	SlimCover2 string `json:"slimcover@2x"`
}

type ChatChannel struct {
	ChannelId      int           `json:"channel_id"`
	Name           string        `json:"name"`
	Description    string        `json:"description"`
	Icon           string        `json:"icon"`
	Type           string        `json:"type"`
	FirstMessageId int           `json:"first_message_id"`
	LastReadId     int           `json:"last_read_id"`
	LastMessageId  int           `json:"last_message_id"`
	RecentMessages []ChatMessage `json:"recent_messages"`
	Moderated      bool          `json:"moderated"`
	Users          []int         `json:"users"`
}

var ChannelTypes = struct {
	PUBLIC      string
	PRIVATE     string
	MULTIPLAYER string
	SPECTATOR   string
	TEMPORARY   string
	PM          string
	GROUP       string
}{"PUBLIC", "PRIVATE", "MULTIPLAYER", "SPECTATOR", "TEMPORARY", "PM", "GROUP"}

type ChatMessage struct {
	MessageId int         `json:"message_id"`
	SenderId  int         `json:"sender_id"`
	ChannelId int         `json:"channel_id"`
	Timestamp string      `json:"timestamp"`
	Content   string      `json:"content"`
	IsAction  bool        `json:"is_action"`
	Sender    UserCompact `json:"sender"`
}
type Comment struct {
	CommentableId   int    `json:"commentable_id"`
	CommentableType string `json:"commentable_type"`
	CreatedAt       string `json:"created_at"`
	DeletedAt       string `json:"deleted_at"`
	EditedAt        string `json:"edited_at"`
	EditedById      int    `json:"edited_by_id"`
	Id              int    `json:"id"`
	LegacyName      string `json:"legacy_name"`
	Message         string `json:"message"`
	MessageHtml     string `json:"message_html"`
	ParentId        int    `json:"parent_id"`
	Pinned          bool   `json:"pinned"`
	RepliesCount    int    `json:"replies_count"`
	UpdatedAt       string `json:"updated_at"`
	UserId          int    `json:"user_id"`
	VotesCount      int    `json:"votes_count"`
}
type CommentBundle struct {
	CommentableMeta  []CommentableMeta `json:"commentable_meta"`
	Comments         []Comment         `json:"comments"`
	HasMore          bool              `json:"has_more"`
	HasMoreId        int               `json:"has_more_id"`
	IncludedComments []Comment         `json:"included_comments"`
	PinnedComments   []Comment         `json:"pinned_comments"`
	Sort             string            `json:"sort"`
	TopLevelCount    int               `json:"top_level_count"`
	Total            int               `json:"total"`
	UserFollow       bool              `json:"user_follow"`
	UserVotes        []int             `json:"user_votes"`
	Users            []UserCompact     `json:"users"`
}
type CommentableMeta struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Type  string `json:"type"`
	Url   string `json:"url"`
}

type Timestamp string

var GameMode = struct {
	fruits string
	mania  string
	osu    string
	taiko  string
}{"fruits", "mania", "osu", "taiko"}

type Event struct {
	achievement struct {
		CreatedAt        Timestamp `json:"created_at"`
		Id               string    `json:"id"`
		Type             string    `json:"type"`
		Username         string    `json:"username"`
		Url              string    `json:"url"`
		PreviousUsername string    `json:"previousUsername"`
	}
	BeatMapPlayCount struct {
		CreatedAt Timestamp `json:"created_at"`
		Id        string    `json:"id"`
		Type      string    `json:"type"`

		Count   int `json:"count"`
		Beatmap struct {
			Title string `json:"title"`
			Url   string `json:"url"`
		} `json:"beatmap"`
	}
	BeatMapSetApprove struct {
		CreatedAt  Timestamp `json:"created_at"`
		Id         string    `json:"id"`
		Type       string    `json:"type"`
		approval   string    //ranked, approved, qualified, loved.
		BeatmapSet struct {
			Title string `json:"title"`
			Url   string `json:"url"` // /b/num
		} `json:"beatmapset"`
		User struct {
			Username         string `json:"username"`
			Url              string `json:"url"`
			PreviousUsername string `json:"previousUsername"`
		} `json:"user"` // Beatmapset owner.
	}
	BeatMapSetDelete struct {
		CreatedAt  Timestamp `json:"created_at"`
		Id         string    `json:"id"`
		Type       string    `json:"type"`
		BeatmapSet struct {
			Title string `json:"title"`
			Url   string `json:"url"`
		} `json:"beatmapset"`
	}
	BeatMapSetRevive struct {
		CreatedAt  Timestamp `json:"created_at"`
		Id         string    `json:"id"`
		Type       string    `json:"type"`
		BeatmapSet struct {
			Title string `json:"title"`
			Url   string `json:"url"` // /b/num
		} `json:"beatmapset"`
		User struct {
			Username         string `json:"username"`
			Url              string `json:"url"`
			PreviousUsername string `json:"previousUsername"`
		} `json:"user"` // Beatmapset owner.
	}
	BeatMapSetUpdate struct {
		CreatedAt  Timestamp `json:"created_at"`
		Id         string    `json:"id"`
		Type       string    `json:"type"`
		BeatmapSet struct {
			Title string `json:"title"`
			Url   string `json:"url"` // /b/num
		} `json:"beatmapset"`
		User struct {
			Username         string `json:"username"`
			Url              string `json:"url"`
			PreviousUsername string `json:"previousUsername"`
		} `json:"user"` // Beatmapset owner.
	}
	BeatMapSetUpload struct {
		CreatedAt  Timestamp `json:"created_at"`
		Id         string    `json:"id"`
		Type       string    `json:"type"`
		BeatmapSet struct {
			Title string `json:"title"`
			Url   string `json:"url"` // /b/num
		} `json:"beatmapset"`
		User struct {
			Username         string `json:"username"`
			Url              string `json:"url"`
			PreviousUsername string `json:"previousUsername"`
		} `json:"user"` // Beatmapset owner.
	}
	rank struct {
		CreatedAt Timestamp `json:"created_at"`
		Id        string    `json:"id"`
		Type      string    `json:"type"`
		ScoreRank string    `json:"scoreRank"` // S SH ...
		Rank      int       `json:"rank"`
		Mode      string    `json:"mode"`
		Beatmap   struct {
			Title string `json:"title"`
			Url   string `json:"url"` // /b/num
		} `json:"beatmap"`
		User struct {
			Username         string `json:"username"`
			Url              string `json:"url"`
			PreviousUsername string `json:"previousUsername"`
		} `json:"user"` // Beatmapset owner.
	}
	rankLost struct {
		CreatedAt Timestamp `json:"created_at"`
		Id        string    `json:"id"`
		Type      string    `json:"type"`
		Mode      string    `json:"mode"`
		Beatmap   struct {
			Title string `json:"title"`
			Url   string `json:"url"` // /b/num
		} `json:"beatmap"`
		User struct {
			Username         string `json:"username"`
			Url              string `json:"url"`
			PreviousUsername string `json:"previousUsername"`
		} `json:"user"` // Beatmapset owner.
	}
	userSupportAgain struct {
		CreatedAt Timestamp `json:"created_at"`
		Id        string    `json:"id"`
		Type      string    `json:"type"`
		User      struct {
			Username         string `json:"username"`
			Url              string `json:"url"`
			PreviousUsername string `json:"previousUsername"`
		} `json:"user"` // Beatmapset owner.
	}
	userSupportFirst struct {
		CreatedAt Timestamp `json:"created_at"`
		Id        string    `json:"id"`
		Type      string    `json:"type"`
		User      struct {
			Username         string `json:"username"`
			Url              string `json:"url"`
			PreviousUsername string `json:"previousUsername"`
		} `json:"user"` // Beatmapset owner.
	}
	userSupportGift struct {
		CreatedAt Timestamp `json:"created_at"`
		Id        string    `json:"id"`
		Type      string    `json:"type"`
		User      struct {
			Username         string `json:"username"`
			Url              string `json:"url"`
			PreviousUsername string `json:"previousUsername"`
		} `json:"user"` // Beatmapset owner.
	}
	usernameChange struct {
		CreatedAt Timestamp `json:"created_at"`
		Id        string    `json:"id"`
		Type      string    `json:"type"`
		User      struct {
			Username         string `json:"username"`
			Url              string `json:"url"`
			PreviousUsername string `json:"previousUsername"`
		} `json:"user"` // Beatmapset owner.
	}
}
type Group struct {
	Id             int    `json:"id"`
	Identifier     string `json:"identifier"`
	IsProbationary string `json:"is_probationary"`
	HasPlayModes   bool   `json:"has_playmodes"`
	Name           string `json:"name"`
	ShortName      string `json:"short_name"`
	Description    string `json:"description"`
	Colour         string `json:"colour"`
}
type UserCompact struct {
	AvatarUrl     string    `json:"avatar_url"`      //url of user's avatar
	CountryCode   string    `json:"country_code"`    //two-letter code representing user's country
	DefaultGroup  string    `json:"default_group"`   //Identifier of the default Group the user belongs to.
	Id            int       `json:"id"`              //unique identifier for user
	IsActive      bool      `json:"is_active"`       //has this account been active in the last x months?
	IsBot         bool      `json:"is_bot"`          //is this a bot account?
	IsOnline      bool      `json:"is_online"`       //is the user currently online? (either on lazer or the new website)
	IsSupporter   bool      `json:"is_supporter"`    //does this user have supporter?
	LastVisit     Timestamp `json:"last_visit"`      //last access time. null if the user hides online presence
	PmFriendsOnly bool      `json:"pm_friends_only"` //whether or not the user allows PM from other than friends
	ProfileColour string    `json:"profile_colour"`  //colour of username/profile highlight, hex code (e.g. #333333)
	Username      string    `json:"username"`        //user's display name
}

type MostPlayed []struct {
	BeatmapId  int               `json:"beatmap_id"`
	Count      int               `json:"count"`
	Beatmap    BeatmapCompact    `json:"beatmap"`
	BeatmapSet BeatmapsetCompact `json:"beatmapset"`
}

type Statistics struct {
	Count50   int `json:"count_50"`
	Count100  int `json:"count_100"`
	Count300  int `json:"count_300"`
	CountGeki int `json:"count_geki"`
	CountKatu int `json:"count_katu"`
	CountMiss int `json:"count_miss"`
}

type BestPlay []struct {
	Id         int        `json:"id"`
	BestId     int        `json:"best_id"`
	UserId     int        `json:"user_id"`
	Accuracy   float64    `json:"accuracy"`
	Mods       []string   `json:"mods"`
	Score      int        `json:"score"`
	MaxCombo   int        `json:"max_combo"`
	Perfect    bool       `json:"perfect"`
	Statistics Statistics `json:"statistics"`
	Pp         float64    `json:"pp"`
	Rank       string     `json:"rank"`
	CreatedAt  Timestamp  `json:"created_at"`
	Mode       string     `json:"mode"`
	ModeInt    int        `json:"mode_int"`
	Replay     bool       `json:"replay"`

	Beatmap    Beatmap           `json:"beatmap"`
	BeatmapSet BeatmapsetCompact `json:"beatmapset"`
	Weight     struct {
		Percentage float64 `json:"percentage"`
		Pp         float64 `json:"pp"`
	} `json:"weight"`
	User UserCompact `json:"user"`
}

//1
