package Route

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"thftgr.com/GoWorld/debianWeb/mariadb"
	"thftgr.com/GoWorld/debianWeb/middlewareFunc"
)

func BeatmapScores(c echo.Context) (err error) {
	defer func() {
		fmt.Println(err)
	}()
	id, err := strconv.Atoi(c.QueryParam("beatmap_id")) //int
	if err != nil {
		middlewareFunc.RequestIDToJson(http.StatusBadRequest,c,"beatmap_id : int")
		return
	}
	m, err := strconv.Atoi(c.QueryParam("mode"))//int
	if err != nil {
		middlewareFunc.RequestIDToJson(http.StatusBadRequest,c,"mode : int")
		return
	}
	p, err := strconv.Atoi(c.QueryParam("page"))//int
	if err != nil {
		middlewareFunc.RequestIDToJson(http.StatusBadRequest,c,"page : int")
		return
	}
	rx, err := strconv.ParseBool(c.QueryParam("relax"))//bool
	if err != nil {
		middlewareFunc.RequestIDToJson(http.StatusBadRequest,c,"relax : bool - 1, t, T, TRUE, true, True, 0, f, F, FALSE, false, False")
		return
	}

	scores, err := mariadb.GetBeatmapScores(id,m,p,rx)
	if err != nil {
		middlewareFunc.RequestIDToJson(http.StatusNotFound, c, "beatmap not found")
		return
	}
	return c.JSON(200, scores)
}
