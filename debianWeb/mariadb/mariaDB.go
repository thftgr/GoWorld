package mariadb

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"io"
	"thftgr.com/GoWorld/debianWeb/src"
	"time"
)

var Maria *sql.DB

type Log struct {
	ip           string
	method       string
	status       int
	uri          string
	latencyHuman string
	time         string
	bytesIn      uint64
	bytesOut     uint64
	captcha      int
}

func Connect() {

	db, err := sql.Open("mysql", src.Config.SQL.Id+":"+src.Config.SQL.Passwd+"@tcp("+src.Config.SQL.Url+")/")
	if Maria = db; db != nil {
		db.SetMaxOpenConns(10)
	}
	if err != nil {
		panic(err)
	}
	if err := Maria.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("connected RDBMS")

}

func GetBeatmapScores(Id, Mode, Page int, Relax bool) (i map[string]interface{}, err error) {
	i = make(map[string]interface{})
	var query string
	if Relax {
		query = QueryBeatmapScoresRX
	} else {
		query = QueryBeatmapScores
	}

	if Page > 0 {
		Page *= 50
	}

	rows, err := Maria.Query(query, Id, Mode, Page, 50, Id)
	if err != nil {
		return
	}
	defer rows.Close()

	var scores []score
	for rows.Next() {

		var s score
		err = rows.Scan(&s.Scoreid, &s.Userid, &s.Username, &s.Country, &s.Score, &s.Rank, &s.MaxCombo, &s.Mods,
			&s.Count300, &s.Count100, &s.Count50, &s.Countmiss, &s.Time, &s.PlayMode, &s.Accuracy, &s.Pp, &s.Fc)
		if err != nil {
			return
		}
		scores = append(scores, s)
	}
	i["Scores"] = scores
	switch Mode {
	//difficulty_std difficulty_taiko difficulty_ctb difficulty_mania
	case 0:
		query = fmt.Sprintf(QueryBeatmapScoresBeatmap, "difficulty_std")
	case 1:
		query = fmt.Sprintf(QueryBeatmapScoresBeatmap, "difficulty_taiko")
	case 2:
		query = fmt.Sprintf(QueryBeatmapScoresBeatmap, "difficulty_ctb")
	case 3:
		query = fmt.Sprintf(QueryBeatmapScoresBeatmap, "difficulty_mania")
	}

	rows, err = Maria.Query(query, Id, Id, Id, Id)
	if err != nil {
		return
	}
	defer rows.Close()
	var Beatmap beatmap
	if rows.Next() {

		err = rows.Scan(
			&Beatmap.BeatmapId,
			&Beatmap.BeatmapsetId,
			&Beatmap.Artist,
			&Beatmap.Title,
			&Beatmap.Diffname,
			&Beatmap.Difficulty,
			&Beatmap.Mode,
			&Beatmap.BPM,
			&Beatmap.AR,
			&Beatmap.CS,
			&Beatmap.OD,
			&Beatmap.HP,
			&Beatmap.TotalLength,
			&Beatmap.HitLength,
			&Beatmap.Playcount,
			&Beatmap.Passcount,
			&Beatmap.MaxCombo,
			&Beatmap.CircleCount,
			&Beatmap.SpinnerCount,
			&Beatmap.SliderCount,
			&Beatmap.Creator,
			&Beatmap.CreatorID,
		)
		if err != nil {
			return
		}

	} else {
		err = errors.New("map Not Found")
		return
	}
	i["Beatmap"] = Beatmap

	var beatmaps []beatmapDF
	rows, err = Maria.Query(QueryBeatmaps, Beatmap.BeatmapsetId)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var t beatmapDF
		err = rows.Scan(&t.BeatmapId, &t.Diffname, &t.Difficulty, &t.Mode)
		if err != nil {
			return
		}
		beatmaps = append(beatmaps, t)

	}
	i["Beatmaps"] = beatmaps
	return
}

func AliveUsername(username string) (alive bool, err error) {
	raws, err := Maria.Query(QueryUsername, username)
	if err != nil {
		return
	}
	defer raws.Close()
	return !raws.Next(), nil
}

func AliveEmail(email string) (alive bool, err error) {
	raws, err := Maria.Query(QueryEmail, email)
	if err != nil {
		return
	}
	defer raws.Close()
	return !raws.Next(), nil
}

func LoginData(username, password string) (v map[string]interface{}, err error) {
	rows, err := Maria.Query(QueryLoginData, username, username)
	if err != nil {
		return
	}
	defer rows.Close()
	if rows.Next() {
		v = make(map[string]interface{})

		var (
			UserID     int
			Username   string
			Email      string
			Password   string
			MainMode   int
			IsBan      bool
			Permission int
			Support    int
			Status     []byte
		)

		err = rows.Scan(&UserID, &Username, &Email, &Password, &MainMode, &IsBan, &Permission, &Support, &Status)
		if err != nil {
			return
		}
		m := md5.New()
		if _, err = io.WriteString(m, password); err != nil {
			return
		}

		if err = bcrypt.CompareHashAndPassword([]byte(Password), []byte(hex.EncodeToString(m.Sum(nil)))); err != nil {
			return
		}

		v["userID"] = UserID
		v["username"] = Username
		v["email"] = Email
		v["mainMode"] = MainMode
		v["isBan"] = IsBan
		v["permission"] = Permission
		v["support"] = Support
		v["status"] = string(Status)

	}
	return
}

func Logout(data map[string]interface{}) (err error) {
	t, err := time.Parse("2006-01-02T15:04:05-07:00", data["expiration"].(string))
	if err != nil {
		log.Println(err)
		return
	}
	rows, err := Maria.Query(QueryLogout, data["request_id"].(string), t.Format("2006-01-02 15:04:05"))

	if err != nil {
		log.Println(err)
		return
	}

	defer rows.Close()

	return
}

func CheckToken(hash string) (err error) {
	rows, err := Maria.Query(QueryCheckTokenLogout, hash)
	if err != nil {
		return
	}
	if rows.Next() {
		return errors.New("Token invalid")
	}
	return
}

func InsertVerifyCode(userid, vc, dc interface{}) (err error) {
	rows, err := Maria.Query(QueryInsertVerifyCode, userid, vc, dc, vc, dc)
	if err != nil {
		return
	}
	defer rows.Close()
	return
}

func CheckVerifyCode(userid, code string) (err error) {
	rows, err := Maria.Query(QueryVerifyCode, userid, code)
	if err != nil {
		return
	}
	defer rows.Close()
	if rows.Next() {
		return
	}
	return errors.New("FAIL")
}

func CheckDisableCode(userid, code string) (err error) {
	rows, err := Maria.Query(QueryDisableCode, userid, code)
	if err != nil {
		return
	}
	defer rows.Close()
	if rows.Next() {
		return
	}
	return errors.New("FAIL")
}

func SetAccountStatus(userid, status interface{}) (err error) {
	rows, err := Maria.Query(QueryUpdateAccountStatus, status, userid, userid)
	if err != nil {
		log.Println(err)
		return
	}
	defer rows.Close()
	return

}

func InsertAPILog(s ...interface{}) (err error) {
	rows, err := Maria.Query(QueryAPILog, s...)
	if err != nil {
		return
	}
	defer rows.Close()
	return
}

func InsertRegisterUser() (err error) {
	tx, err := Maria.Begin()
	if err != nil {
		return
	}
	defer tx.Rollback()

	res, err := Maria.Exec(QueryRegister01)
	if err != nil {
		return errors.New(err.Error() + "Whoops, an error slipped in. You might have been registered, though. I don't know.")
	}
	uid, _ := res.LastInsertId()
	Maria.Exec(QueryRegister02, uid)
	Maria.Exec(QueryRegister03, uid)
	Maria.Exec(QueryRegister04, uid)
	Maria.Exec(QueryRegister05, uid)

	return tx.Commit()

}
