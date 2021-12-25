package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func main() {
	var list = map[string][]fs.FileInfo{}
	path := "D:/tmp/oszz"

	fi, err := ioutil.ReadDir(path)
	if err != nil {
		log.Println(err)
		return
	}
	c, err := regexp.Compile("^([0-9]+)(\\s)")
	if err != nil {
		return
	}
	for _, f := range fi {
		list[strings.TrimSpace(c.FindString(f.Name()))] = append(list[strings.TrimSpace(c.FindString(f.Name()))], f)
	}
	for k, i := range list {
		if len(i) > 1 {
			fmt.Println(k)
			for a := 0; a < len(i); a++ {

				err := os.Rename(path+"/"+i[a].Name(), path+"/dup/"+i[a].Name())
				if err != nil {
					log.Println(err)
				}

			}
		}
	}
	//fmt.Println(os.Mkdir("./az",0775))
	//fmt.Println(os.Rename("./a","./az/z"))

}
