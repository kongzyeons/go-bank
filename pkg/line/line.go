package line

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/kongzyeons/go-bank/config"
)

type LineAPI interface {
	SendMessage(message string) error
}

type lineAPI struct {
	AccToken string
	UserID   string
}

func NewLineAPI() LineAPI {
	cfg := config.InitConfig()

	return &lineAPI{
		AccToken: cfg.LineAccToken,
		UserID:   cfg.LineUserIDToken,
	}
}

func (self *lineAPI) SendMessage(messageJson string) error {
	url := "https://api.line.me/v2/bot/message/push"
	method := "POST"

	pushReq := PushRequest{
		To: self.UserID,
		Messages: []Message{
			{Type: "text", Text: messageJson},
		},
	}
	// Marshal the struct to JSON
	jsonData, err := json.Marshal(pushReq)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonData))

	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", self.AccToken))

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}
	log.Println(string(res.Status), string(body))
	return nil
}
