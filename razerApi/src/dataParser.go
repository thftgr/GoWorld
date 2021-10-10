package src

func ParseIntModeSQLFile(data int) (sqlFile string) {
	//+4
	switch data {
	case 0:
		sqlFile = "users.osu.sql"
	case 1:
		sqlFile = "users.taiko.sql"
	case 2:
		sqlFile = "users.fruits.sql"
	case 3:
		sqlFile = "users.mania.sql"
	case 4:
		sqlFile = "users.osu.relax.sql"
	case 5:
		sqlFile = "users.taiko.relax.sql"
	case 6:
		sqlFile = "users.fruits.relax.sql"
	case 7:
		sqlFile = "users.mania.relax.sql"

	default:
		sqlFile = "users.osu.sql"
	}
	return
}

func ParseStringModeSQLFile(data string) (sqlFile string) {
	//.relax
	switch data {
	case "osu":
		sqlFile = "users.osu.sql"
	case "fruits":
		sqlFile = "users.fruits.sql"
	case "taiko":
		sqlFile = "users.taiko.sql"
	case "mania":
		sqlFile = "users.mania.sql"
	case "osu.relax":
		sqlFile = "users.osu.relax.sql"
	case "fruits.relax":
		sqlFile = "users.fruits.relax.sql"
	case "taiko.relax":
		sqlFile = "users.taiko.relax.sql"
	case "mania.relax":
		sqlFile = "users.mania.relax.sql"
	default:
		sqlFile = "users.osu.sql"
	}
	return
}
