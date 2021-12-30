//package main
package src

import (
	"encoding/json"
	"fmt"
	"github.com/pterm/pterm"
	"io/ioutil"
	"regexp"
	"strings"
)

var RootDir DirectoryTree

func init() {
	RootDir = ReadDir("./data")
	//b, _ := json.MarshalIndent(RootDir, "", "    ")
	//fmt.Println(string(b))
}

var video, _ = regexp.Compile(`([.](mp4|webm|mov|avi))$`)
var image, _ = regexp.Compile(`([.](png|jpg|jpeg|webp|pdf))$`)

// ReadDir
// example:
// path: "./dir1/dir2"
// name: "dir3"
func ReadDir(path string) (dir DirectoryTree) {
	fmt.Println("scanning path: ", path)

	defer func() {
		err, _ := recover().(error)
		if err != nil {
			pterm.Error.Println(err)
		}
	}()

	dir.Path = path
	dir.Dir = map[string]DirectoryTree{}
	dirs, _ := ioutil.ReadDir(path)

	for _, f := range dirs {
		if !f.IsDir() {
			dir.Files = append(dir.Files, f.Name())
			continue
		}
		dir.Dirs = append(dir.Dirs, f.Name())
		dir.Dir[f.Name()] = ReadDir(path + "/" + f.Name())
	}
	return
}

type DirectoryTree struct {
	Path  string                   `json:"path"`  // 전체 경로
	Files []string                 `json:"files"` // 파일명
	Dirs  []string                 `json:"dirs"`  // 파일명
	Dir   map[string]DirectoryTree `json:"dir"`   // 폴더 구조
}

type Directory struct {
	Path      string    `json:"path"`      // 전체 경로
	Files     []string  `json:"files"`     // 파일명
	Dirs      []string  `json:"dirs"`      // 파일명
	Galleries []Gallery `json:"galleries"` // 파일명
}
type Gallery struct {
	Thumb       string `json:"thumb"`
	GalleryName string `json:"galleryName"`
}

func (v *DirectoryTree) GetInfo(path string) (dir Directory) {
	a := strings.Split(path, "/")
	var patches []string
	for _, p := range a {
		if p != "" {
			patches = append(patches, p)
		}
	}

	//요청 디렉토리까지 진입
	tmp := RootDir
	for _, p := range patches {
		tmp = tmp.Dir[p]
	}

	var t Directory
	//TODO 파일 처리
	t.Path = strings.ReplaceAll(tmp.Path, "./data", "")
	if t.Path == "" {
		t.Path = "/"
	}
	t.Files = append(t.Files, tmp.Files...)
	for _, s := range tmp.Dirs {
		tt := tmp.Dir[s]
		img := 0
		vid := 0
		for _, file := range tt.Files {
			if image.Match([]byte(file)) {
				img++
			}
			if video.Match([]byte(file)) {
				vid++
			}
		}

		if img > 0 {
			t.Galleries = append(t.Galleries, Gallery{
				Thumb:       s + "/" + tt.Files[0],
				GalleryName: s,
			})
		}

		if len(tt.Dirs) > 0 {
			t.Dirs = append(t.Dirs, s)
		}

	}

	return t
}

func printJsonFormat(v interface{}) {
	data, _ := json.MarshalIndent(v, "", "    ")
	fmt.Println(string(data))
}

//func main() {
//	printJsonFormat(RootDir)
//	printJsonFormat(RootDir.GetInfo("/"))
//}

//
