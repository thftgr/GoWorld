package src

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var Config struct {
	Server struct {
		Port   string `json:"port"`
		Domain string `json:"domain"`
	} `json:"server"`
	Jwt struct {
		Key string `json:"key"`
	} `json:"jwt"`
	ReCaptcha struct {
		SiteKey   string `json:"siteKey"`
		ServerKey string `json:"serverKey"`
	} `json:"reCAPTCHA"`
	SQL struct {
		Id     string `json:"id"`
		Passwd string `json:"passwd"`
		Url    string `json:"url"`
	} `json:"sql"`
	Mail struct {
		Mode   string `json:"mode"`
		Host   string `json:"host"`
		From   string `json:"from"`
		Id     string `json:"id"`
		Passwd string `json:"passwd"`
	} `json:"mail"`
	InvalidUsername []string `json:"invalidUsername"`
}

func LoadSetting() {
	b, err := ioutil.ReadFile("./config.json")
	if err != nil {
		out, err := os.Create("./config.json")
		if err != nil {
			panic(err)
		}
		defer out.Close()
		body, err := json.MarshalIndent(Config, "", "    ")
		if err != nil {
			panic(err)
		}
		// Write the body to file
		if _, err = out.Write(body); err != nil {
			panic(err)
		}
	}
	err = json.Unmarshal(b, &Config)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
}
