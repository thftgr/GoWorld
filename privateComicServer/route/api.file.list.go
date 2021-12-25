package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/thftgr/GoWorld/privateComicServer/src"
	"strings"
)

func GetFileList(c *fiber.Ctx) (err error) {
	defer func() {
		err, _ = recover().(error)
		if err != nil {
			return
		}
	}()
	dirs := strings.Split(c.Query("path", "/"), "/")
	var dirs2 []string
	for _, dir := range dirs {
		if dir != "" {
			dirs2 = append(dirs2, dir)
		}
	}
	var body = map[string]interface{}{}

	var td src.Directory
	td = src.RootDir
	for i := 0; i < len(dirs2); i++ {
		td = td.Dirs[dirs2[i]]
	}

	var (
		dd      []string
		gallery []gal
	)
	base := strings.Join(dirs2, "/")

	for _, s := range td.GetKeys() {
		if len(td.Dirs[s].Files) > 0 {
			gallery = append(gallery, gal{
				Name:      td.Dirs[s].Name,
				Thumbnail: base + "/" + td.Dirs[s].Name + "/" + td.Dirs[s].Files[0],
			})
		}
		if len(td.Dirs[s].Dirs) > 0 {
			dd = append(dd, s)
		}

	}

	body["path"] = base
	body["gallery"] = gallery
	body["dirs"] = dd

	_ = c.JSON(body)

	return
}

type gal struct {
	Name      string
	Thumbnail string
}
