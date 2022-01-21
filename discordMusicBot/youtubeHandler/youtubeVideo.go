//package main

package youtubeHandler

import (
	"github.com/kkdai/youtube/v2"
	"regexp"
)

var YoutubeRegex, _ = regexp.Compile(`(?:(?:youtu(?:be)?|)(?:com|be))`)
var VideoRegix, _ = regexp.Compile(`(?:v=)([A-z0-9-_]+?)(?:$|&)`)
var PlayListRegix, _ = regexp.Compile(`(?:list=)([A-z0-9-_]+?)(?:$|&)`) //https://www.youtube.com/watch?v=D2Ekp3m3IQk&list=PLFsi9YnqVto7_Ryh3-WiTBQNB6ZV62ljd

//func main() {
//	url := `https://www.youtube.com/watch?v=D2Ekp3m3IQk&list=PLFsi9YnqVto7_Ryh3-WiTBQNB6ZV62ljd`
//	if !youtubeRegex.MatchString(url) {
//		fmt.Println("not youtube url")
//		return
//	}
//	//var vid,list bool
//	v := videoRegix.FindAllStringSubmatch(url, -1)
//	vid := v != nil && len(v) > 0 && len(v[0]) > 1
//
//	l := playListRegix.FindAllStringSubmatch(url, -1)
//	list := l != nil && len(l) > 0 && len(l[0]) > 1
//	fmt.Println(vid, list)
//	printJson(Audio(v[0][1]))
//	printJson(AudioList(l[0][1]))
//
//	//client := youtube.Client{}
//
//	//pl := playListRegix.FindAllStringSubmatch(url, -1)
//
//}
func Audio(key string) (fl youtube.FormatList) {
	client := youtube.Client{}
	video, err := client.GetVideo(key)
	if err != nil {
		panic(err)
	}

	return video.Formats.WithAudioChannels() // only get videos with audio
	//stream, _, err := client.GetStream(video, &formats[0])
	//if err != nil {
	//	panic(err)
	//}
	//
	//file, err := os.Create("video.mp4")
	//if err != nil {
	//	panic(err)
	//}
	//defer file.Close()
	//_, err = io.Copy(file, stream)
	//if err != nil {
	//	panic(err)
	//}
}

// Example usage for playlists: downloading and checking information.
func AudioList(key string) (l []*youtube.PlaylistEntry) {
	client := youtube.Client{}

	playlist, err := client.GetPlaylist(key)
	if err != nil {
		panic(err)
	}
	//
	///* ----- Enumerating playlist videos ----- */
	//header := fmt.Sprintf("Playlist %s by %s", playlist.Title, playlist.Author)
	//println(header)
	//println(strings.Repeat("=", len(header)) + "\n")
	//
	//for k, v := range playlist.Videos {
	//	fmt.Printf("(%d) %s - '%s'\n", k+1, v.Author, v.Title)
	//}

	/* ----- Downloading the 1st video ----- */
	return playlist.Videos
	//entry := playlist.Videos[0]
	//video, err := client.VideoFromPlaylistEntry(entry)
	//if err != nil {
	//	panic(err)
	//}
	//// Now it's fully loaded.
	//
	//fmt.Printf("Downloading %s by '%s'!\n", video.Title, video.Author)
	//
	//stream, _, err := client.GetStream(video, &video.Formats[0])
	//if err != nil {
	//	panic(err)
	//}
	//
	//file, err := os.Create("video.mp4")
	//
	//if err != nil {
	//	panic(err)
	//}
	//
	//defer file.Close()
	//_, err = io.Copy(file, stream)
	//
	//if err != nil {
	//	panic(err)
	//}
	//
	//println("Downloaded /video.mp4")
}
