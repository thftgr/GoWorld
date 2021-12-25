package src

import (
	"fmt"
	"io/ioutil"
)

//type File struct {
//	Thumbnail string `json:"thumbnail"`
//	Title     string `json:"title"`
//	Url       string `json:"url"`
//	Type      string `json:"type"`
//}

//폴더 타입
//1. 일반 폴더
//2. 하위폴더가 없는 폴더
//func main() {
//	defer func() {
//		err, _ := recover().(error)
//		if err != nil {
//			return
//		}
//	}()
//	//ReadDir("./data")
//
//	b, _ := json.MarshalIndent(RootDir, "", "    ")
//	fmt.Println(string(b))
//	fmt.Println(RootDir.GetKeys())
//	//fmt.Println("==============================")
//	//b, _ = json.MarshalIndent(files, "", "    ")
//	//fmt.Println(string(b))
//
//}

var RootDir Directory

func init() {
	RootDir = ReadDir("./data", "/")
	//b, _ := json.MarshalIndent(RootDir, "", "    ")
	//fmt.Println(string(b))
}

func ReadDir(path, name string) (dir Directory) {
	dir.Name = name
	defer func() {
		err, _ := recover().(error)
		if err != nil {
			fmt.Println(err)
		}
	}()

	if path[len(path)-1:] == "/" {
		path = path[:len(path)-1]
	}
	fl, _ := ioutil.ReadDir(path)
	dir.Dirs = map[string]Directory{}
	for _, file := range fl {
		fileName := file.Name()

		if file.IsDir() {

			dir.Dirs[fileName] = ReadDir(path+"/"+fileName, fileName)
		} else {
			dir.Files = append(dir.Files, fileName)
		}
	}
	return
}

type Directory struct {
	Name  string
	Files []string
	Dirs  map[string]Directory
}

func (v *Directory) GetKeys() (keys []string) {
	keys = make([]string, 0, len(v.Dirs))
	for k := range v.Dirs {
		keys = append(keys, k)
	}
	return
}
