package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
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

func main() {
	//https://osu.ppy.sh/beatmaps/packs/2331/raw
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://osu.ppy.sh/beatmaps/packs/2331/raw", nil)

	if err != nil {
		return
	}

	req.Header.Add("Authorization", "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJhdWQiOiI1IiwianRpIjoiYjZlYWNjNGI4NWVmYjFmMWZjOWZlZTQzYzY2ZWQ4MjZjODY4ZjBhNDI1NTU2YjcwNzVjYWY3NDY4MmE4YmQzZjY4NzA5YWRhZThiYTQyNzEiLCJpYXQiOjE2NDM5NTMxMzAuNDk0MzEzLCJuYmYiOjE2NDM5NTMxMzAuNDk0MzE3LCJleHAiOjE2NDQwMzk0ODIuNzE2NDgsInN1YiI6IjgxNDYyMzIiLCJzY29wZXMiOlsiKiJdfQ.JzrieW8JcUDnXA12R1MAIjfkn0mb8Uytb0OUjnQQBuyA_kIOClipeEn5StsXVg-jEGsOZgsem546gE59METRqcEeLb9zPOI0LR2-TreVWw-MZpXqJpLVfXI_5QkuaHbWL2FTvnt_eZoJjgmlpmPcmlIxf9j2QbtfwzIbhy1MytZew9DQix7lgfxzQ3DoBU8CjU6tza43NavJZYP0vuDHkJfGKTWqwaEIXUFRPbvlTH1YkPQ5TenmHkiFiRUAKIhyRqYdErZS6zAAoYpTvinizDU9rNFNLIm0Kb-q4sf87rEyK2lGRaUzYkqeu8khN_FRy3I5u15rRGluNWa0XdW9orCRSCqxCq9EFY-Gbg7_6dpphI8x2e5iVyKMl9WaxcNM3YDmKYFsG3RKWyY9KT4MOdzM602IZGl8ZfkA2XRyy0enTaS2hqnAnoNpt3tFMwbSeL50cZi6pm2qNk9GL_ZwtFsaZYKLC20cx03iFh_nfceeeZD54cfu58ar2hCEEPIE7XTFkU4hPQ1qTRKczzrPK5OC5_Ln0dLsuCEYdFdDfjE45aNGexfDzQHzPw4LkamFFuCbFUU0ubiutjIxoWHf-L5IUJX788FQyuCzLxA9emIqZwUTs77NFSltLF5rdkE3E3Ue_2mUCsgVuDw0bEkqtCaAlluS1fQXBYizvlSnzhs")

	res, err := client.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {

	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	fmt.Println(string(body))
}
