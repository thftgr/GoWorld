package src

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"strings"
)

func QueryOnly(db *sql.DB, sql string, parm ...interface{}) error {
	sqlString := GetSql(sql)
	_, err := db.Query(sqlString, parm...)
	return err
}
func QueryGetJsonObject(db *sql.DB, sql string, parm ...interface{}) interface{} {
	sqlString := GetSql(sql)

	rows, err := db.Query(sqlString, parm...)
	if err != nil {
		log.Println(err)
		return ""
	}

	columns, err := rows.Columns()
	if err != nil {
		log.Println(err)

		return ""
	}
	count := len(columns)
	tableData := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)

	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		_ = rows.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				if strings.Index(string(b), "[{") == 0 {
					var dat []map[string]interface{}
					_ = json.Unmarshal(b, &dat)
					v = dat
					//fmt.Println(dat)

				} else if strings.Index(string(b), "{") == 0 {
					var dat map[string]interface{}
					_ = json.Unmarshal(b, &dat)
					v = dat
					//fmt.Println(string(b))
					//fmt.Println(dat)

				} else if strings.Index(string(b), "[") == 0 {
					dataJson := string(b)
					var arr []string
					_ = json.Unmarshal([]byte(dataJson), &arr)
					v = arr
					//type add struct {
					//	Key string `:`
					//}

				} else if string(b) == "true" {
					v = true
				} else if string(b) == "false" {
					v = false
					//} else if tmp, err := strconv.Atoi(string(b)); err == nil {
					//	v = tmp
				} else {
					v = string(b)
				}

			} else {
				v = val
			}
			entry[col] = v
		}
		//fmt.Println(entry)
		tableData = append(tableData, entry)
		return entry
	}

	return tableData
}

func QueryGetJsonArray(db *sql.DB, sql string, parm ...interface{}) interface{} {
	sqlString := GetSql(sql)

	rows, err := db.Query(sqlString, parm...)
	if err != nil {
		log.Println(err)
		return ""
	}

	columns, err := rows.Columns()
	if err != nil {
		log.Println(err)

		return ""
	}
	count := len(columns)
	tableData := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)

	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		_ = rows.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				if strings.Index(string(b), "{") == 0 {
					var dat map[string]interface{}
					_ = json.Unmarshal(b, &dat)
					v = dat
				} else if strings.Index(string(b), "[") == 0 {
					dataJson := string(b)
					var arr []string
					_ = json.Unmarshal([]byte(dataJson), &arr)
					v = arr
				} else {
					v = string(b)
				}

			} else {
				v = val
			}
			entry[col] = v
		}
		tableData = append(tableData, entry)
	}
	return tableData
}
func GetSql(path string) (query string) {
	b, err := ioutil.ReadFile("./sql/" + path)
	errCheck(err)
	return string(b)
}

func errCheck(err error) (e bool) {
	e = err != nil
	if e {
		log.Println(err)
	}
	return
}
