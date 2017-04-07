package line

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

var replyURL = "https://api.line.me/v2/bot/message/reply"

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

type WebhookResponse struct {
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

type ReplyMessage struct {
	Text               string `json:"text"`
	Type               string `json:"type"`
	OriginalContentURL string `json:"originalContentUrl"`
	PreviewImageURL    string `json:"previewImageUrl"`
}

type Reply struct {
	Messages   []ReplyMessage `json:"messages"`
	ReplyToken string         `json:"replyToken"`
}

func lineRequest(verb string, body []byte) *http.Request {
	req, _ := http.NewRequest(verb, replyURL, bytes.NewBuffer(body))
	req.Header.Add("Authorization", "Bearer "+os.Getenv("line_access_token"))
	req.Header.Add("Content-Type", "application/json")

	return req
}

func (r Reply) SendImage(imageURL string, replyToken string) {
	rMsg := newImageReplyMessage(imageURL)

	reply := Reply{
		ReplyToken: replyToken,
		Messages: []ReplyMessage{
			rMsg,
		},
	}

	marshal, err := json.Marshal(reply)
	if err != nil {
		log.Println(err)
	}

	client := http.Client{}
	req := lineRequest("POST", marshal)

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}

	defer resp.Body.Close()
	//dumpOut(resp)
}

func newImageReplyMessage(u string) ReplyMessage {
	return ReplyMessage{
		OriginalContentURL: u,
		PreviewImageURL:    u,
		Type:               "image",
	}
}
