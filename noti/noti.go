// Package noti, shortform for notification will be primarily used to
// send notifications through different channels, eg. telegram, sms,
// windows toast etc.

// We are reading from the yaml file for the configurations for the
// parameters required for http request and other usecases.

package noti

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"gopkg.in/yaml.v2"
)

// Comfig is a struct created to read the secrets and password
// from a yaml file
type Config struct {
	Linaak_home struct {
		Token   string `yaml:"token"`
		Chat_id string `yaml:"chat_id"`
	}
}

// Jres is a struct created to handle the response from the
// telegram api after sending notifications.
type Jres struct {
	Ok bool
}

// Telereq does two things,
// // 1. unmarshalls the token and secret required for the telegram requests
// // 2. sends the telegram requests for sending a simple text message
// The input is just the text to be sent and the return value is a bool
// which is unmarshalled from the http response json and passed as the return
// value.
func Telereq(text string) bool {

	// The below snippet just unmarshalls the token from the yaml

	f, err := os.Open("noti.yaml")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var cfg Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	fmt.Println(cfg)
	if err != nil {
		log.Fatal(err)
	}

	// The request for telegram notification starts here

	urlTele := "https://api.telegram.org/bot" + cfg.Linaak_home.Token + "/"
	path := "sendMessage?"
	prm1 := "chat_id=" + cfg.Linaak_home.Chat_id
	prm2 := "&text=" + text
	final_url := urlTele + path + prm1 + prm2
	fmt.Println(final_url)
	res, err := http.Get(final_url)
	if err != nil {
		log.Fatal(err)
	}
	b, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	var m Jres
	err = json.Unmarshal(b, &m)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(m)
	return m.Ok

}
