package Route

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"thftgr.com/GoWorld/debianWeb/mariadb"
)

type parmLeaderBoard struct {
	Mode  int    `form:"mode"`
	Page  int    `form:"page"`
	Relax string `form:"relax"`
	Type  string `form:"type"`
}
type respdLeaderBoard struct {
	Id        int     `json:"id"`
	Username  string  `json:"username"`
	Country   string  `json:"country"`
	Pp        int     `json:"pp"`
	Score     int     `json:"score"`
	Accuracy  float64 `json:"accuracy"`
	Playcount int     `json:"playcount"`
	Level     int     `json:"level"`
}

const leaderBoardQuery = "select %s from Ainu.%s where id in (select id from Ainu.users where ban_datetime = 0) order by %s desc %s;"

func LeaderBoard(c echo.Context) (err error) {
	//{
	// "id": 8545,
	// "username": "rvyr",
	// "country": "US",
	// "pp": 11000,
	// "score": 340696169,
	// "accuracy": 96.734375,
	// "playcount": 177,
	// "level": 42
	//}

	var (
		db = &mariadb.Maria
		q  parmLeaderBoard

		limit   = " limit 50 "
		sel     = "id,username,country,pp_std,ranked_score_std,avg_accuracy_std,playcount_std,level_std"
		table   = "users_stats"
		orderBy = "pp_std"
	)

	err = c.Bind(&q)
	if err != nil {
		c.NoContent(http.StatusBadRequest)
		return err
	}
	//rx?

	if parseBool, _ := strconv.ParseBool(q.Relax); parseBool {
		table = "rx_stats"
	}

	//limit
	if q.Page > 0 {
		limit = fmt.Sprintf(" limit %d,50 ", q.Page*50)
	}

	//mode
	if q.Mode != 0 {
		switch q.Mode {
		case 1:
			orderBy = "pp_taiko"
			sel = "id,username,country,pp_taiko,ranked_score_taiko,avg_accuracy_taiko,playcount_taiko,level_taiko"
		case 2:
			orderBy = "pp_ctb"
			sel = "id,username,country,pp_ctb,ranked_score_ctb,avg_accuracy_ctb,playcount_ctb,level_ctb"
		case 3:
			orderBy = "pp_mania"
			sel = "id,username,country,pp_mania,ranked_score_mania,avg_accuracy_mania,playcount_mania,level_mania"
		}
	}

	rows, err := (*db).Query(fmt.Sprintf(leaderBoardQuery, sel, table, orderBy, limit))
	if err != nil {
		c.NoContent(http.StatusInternalServerError)
		return err
	}
	defer rows.Close()
	var d []respdLeaderBoard
	for rows.Next() {
		var dd respdLeaderBoard
		rows.Scan(&dd.Id, &dd.Username, &dd.Country, &dd.Pp, &dd.Score, &dd.Accuracy, &dd.Playcount, &dd.Level)
		d = append(d, dd)
	}

	return c.JSON(http.StatusOK, d)
}
