package line

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
)

var accessToken = "j8PMTefgEHgStNfN77eH9+UkNFh4P0hGiLqttfp9GUumAn/dMbyEtHBDA6io7A7Qyrwf6xBSQ3Lu8nCBUdB8pA8IhSJ32Ary404fhnnjSu8QjKE6MZD82qzmIQBpcAKVfZojU1TNmcGst4ZUzk/WBQdB04t89/1O/w1cDnyilFU="
var replyURL = "https://api.line.me/v2/bot/message/reply"

type Reply struct {
	Events []Event `json:"events"`
}

type Event struct {
	Type       string  `json:"type"`
	ReplyToken string  `json:"replyToken"`
	Message    Message `json:"message"`
}

type Message struct {
	Type string `json:"type"`
	ID   string `json:"id"`
	Text string `json:"text"`
}

type SendMessage struct {
	Text               string `json:"text"`
	Type               string `json:"type"`
	OriginalContentURL string `json:"originalContentUrl"`
	PreviewImageURL    string `json:"previewImageUrl"`
}

type SendMessages struct {
	Message    []SendMessage `json:"messages"`
	ReplyToken string        `json:"replyToken"`
}

const lineURL = "https://api.line.me/v2/bot/message/reply"

func (r Reply) Send(m string) {
	messages := SendMessages{
		ReplyToken: r.Events[0].ReplyToken,
		Message: []SendMessage{
			SendMessage{
				Text: m,
				Type: "text",
			},
		},
	}

	marshal, err := json.Marshal(messages)

	if err != nil {
		log.Fatal(err)
	}

	client := http.Client{}
	req, _ := http.NewRequest("POST", replyURL, bytes.NewBuffer(marshal))
	req.Header.Add("Authorization", "Bearer "+accessToken)
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	dumpOut(resp)
}

func (r Reply) SendImage() {
	messages := SendMessages{
		ReplyToken: r.Events[0].ReplyToken,
		Message: []SendMessage{
			SendMessage{
				OriginalContentURL: "https://imgflip.com/s/meme/Success-Kid.jpg",
				PreviewImageURL:    "https://connorlenahan.files.wordpress.com/2014/07/success-kid.jpg",
				Type:               "image",
			},
		},
	}

	marshal, err := json.Marshal(messages)

	if err != nil {
		log.Fatal(err)
	}

	client := http.Client{}
	req, _ := http.NewRequest("POST", replyURL, bytes.NewBuffer(marshal))
	req.Header.Add("Authorization", "Bearer "+accessToken)
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	dumpOut(resp)
}

func dumpOut(r *http.Response) {
	dump, err := httputil.DumpResponse(r, true)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(dump[:]))
}

func jsonRequest(r *http.Request) {
	var out bytes.Buffer
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Indent(&out, body, "", "  ")
	out.WriteTo(os.Stdout)
}
