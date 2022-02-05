package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strconv"
)

var (
	regPackDownloadLink, _ = regexp.Compile(`(?:<a href=")(.+?)(?:"(\s+?|.+?)class="beatmap-pack-download__link">)`)
	regMapSetId, _         = regexp.Compile(`(?:<a href=")(?:https://osu[.]ppy[.]sh/beatmapsets/)([0-9]+?)(?:"(\s+?|.+?)class="beatmap-pack-items__link">)`)
	regPackName, _         = regexp.Compile(`(?:<div class="beatmap-pack__name">)(.*)(?:</div>)`)
	regPackId, _           = regexp.Compile(`(?:https://osu[.]ppy[.]sh/beatmaps/packs/)([0-9]+?)(?:")`)
	regPackDate, _         = regexp.Compile(`(?:<span class="beatmap-pack__date">)(.*)(?:</span>)`)
	regCreator, _          = regexp.Compile(`(?:<span class="beatmap-pack__author beatmap-pack__author--bold">)(.*)(?:</span>)`)
	regLastPage, _         = regexp.Compile(`(?:<a class="pagination-v2__link" href=".+?">)([0-9]+?)(?:</a>)`)
)
var (
	packName    []string
	packId      []string
	packDate    []string
	packCreator []string
)
var packType = []string{
	"standard",
	"chart",
	"theme",
	"artist",
}

func main() {
	//for i := 0; i < len(packType); i++ {
	page := 0
	bodyString := fetch(packType[0], strconv.Itoa(page))
	m := regLastPage.FindAllStringSubmatch(bodyString, -1)
	for _, sm := range m {
		if len(sm) > 1 {
			t1, _ := strconv.Atoi(sm[1])
			if page < t1 {
				page = t1
			}
		}
	}
	parseBody(bodyString)
	//for j := 1; j <= page; j++ {
	//	time.Sleep(time.Millisecond * 500)
	//	bodyString = fetch(packType[i], strconv.Itoa(j))
	//	parseBody(bodyString)
	//}
	//}
	fmt.Println(len(packId), len(packDate), len(packCreator), len(packName))
	fmt.Printf("| %5s | %10s | %30s | %s \n", "id", "packDate", "packCreator", "packName")
	fmt.Printf("| %5s | %10s | %30s | \n", "-----", "----------", "------------------------------")
	for j := 0; j < len(packId); j++ {
		fmt.Printf("| %5s | %10s | %30s | %s \n", packId[j], packDate[j], packCreator[j], packName[j])
	}

}
func parseBody(bodyString string) {
	m := regPackName.FindAllStringSubmatch(bodyString, -1)
	for _, sm := range m {
		if len(sm) > 1 {
			packName = append(packName, sm[1])
		} else {
			packName = append(packName, "")
		}
	}
	m = regPackId.FindAllStringSubmatch(bodyString, -1)
	for _, sm := range m {
		if len(sm) > 1 {
			packId = append(packId, sm[1])
		} else {
			packId = append(packId, "")
		}
	}
	m = regPackDate.FindAllStringSubmatch(bodyString, -1)
	for _, sm := range m {
		if len(sm) > 1 {
			packDate = append(packDate, sm[1])
		} else {
			packDate = append(packDate, "")
		}
	}
	m = regCreator.FindAllStringSubmatch(bodyString, -1)
	for _, sm := range m {
		if len(sm) > 1 {
			packCreator = append(packCreator, sm[1])
		} else {
			packCreator = append(packCreator, "")
		}
	}
}
func fetch(Type, page string) string {
	res, err := http.Get(fmt.Sprintf("https://osu.ppy.sh/beatmaps/packs?type=%s&page=%s", Type, page))
	if err != nil || res.StatusCode != http.StatusOK {
		if res != nil {
			panic(res.Status)
		}
		panic(err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	return string(body)
}
