package youtubeHandler

import (
	"bytes"
	"discordMusicBot/config"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type YoutubeSearchData struct {
	Query string
	//======
	Kind          string `json:"kind"`
	Etag          string `json:"etag"`
	NextPageToken string `json:"nextPageToken"`
	PrevPageToken string `json:"prevPageToken"`
	RegionCode    string `json:"regionCode"`
	PageInfo      struct {
		TotalResults   int `json:"totalResults"`
		ResultsPerPage int `json:"resultsPerPage"`
	} `json:"pageInfo"`
	Items []struct {
		Kind string `json:"kind"`
		Etag string `json:"etag"`
		Id   struct {
			Kind    string `json:"kind"`
			VideoId string `json:"videoId"`
		} `json:"id"`
		Snippet struct {
			PublishedAt time.Time `json:"publishedAt"`
			ChannelId   string    `json:"channelId"`
			Title       string    `json:"title"`
			Description string    `json:"description"`
			Thumbnails  struct {
				Default struct {
					Url    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"default"`
				Medium struct {
					Url    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"medium"`
				High struct {
					Url    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"high"`
			} `json:"thumbnails"`
			ChannelTitle         string    `json:"channelTitle"`
			LiveBroadcastContent string    `json:"liveBroadcastContent"`
			PublishTime          time.Time `json:"publishTime"`
		} `json:"snippet"`
	} `json:"items"`
}

func (v *YoutubeSearchData) Next() (data *YoutubeSearchData) {
	if v.NextPageToken != "" {
		return Search(v.Query, v.NextPageToken)
	}
	return
}
func (v *YoutubeSearchData) Prev() (data *YoutubeSearchData) {
	if v.PrevPageToken != "" {
		return Search(v.Query, v.PrevPageToken)
	}
	return
}

func Search(query, pageToken string) (data *YoutubeSearchData) {
	urlBuf := bytes.Buffer{}
	urlBuf.WriteString("https://www.googleapis.com/youtube/v3/search?")
	//urlBuf.WriteString("part=snippet")
	//urlBuf.WriteString("part=contentDetails")
	urlBuf.WriteString("part=player")
	//urlBuf.WriteString("part=statistics")
	//urlBuf.WriteString("part=status")

	urlBuf.WriteString("&order=viewCount")
	urlBuf.WriteString("&type=video")
	urlBuf.WriteString("&videoDefinition=high")
	urlBuf.WriteString("&key=" + config.Config.Youtube.Token)
	urlBuf.WriteString("&pageToken=" + pageToken)
	urlBuf.WriteString("&q=" + query)

	url := urlBuf.String()
	res, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return
	}
	if res.StatusCode != http.StatusOK {
		log.Println(res.Status, "url:", url)
		return
	}
	b, err := ioutil.ReadAll(res.Body)
	err = json.Unmarshal(b, &data)
	data.Query = query
	return

}
