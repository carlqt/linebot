package main

import (
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/carlqt/linebot/bing"
	"github.com/carlqt/linebot/devexcuse"
	"github.com/carlqt/linebot/line"
)

func lineReply(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	text, _ := ctx.Value("text").(string)
	replyToken, _ := ctx.Value("replyToken").(string)

	switch command(text) {
	case "/pic":
		replyImage(text, replyToken)
	case "/devexcuse":
		line.Send(devexcuse.Excuse(), replyToken)
	}
	w.WriteHeader(200)
}

func replyImage(text, replyToken string) {
	q := getSearchText(text)
	imageURL, err := bing.SearchImage(q)

	if err != nil {
		line.Send("Sorry, the term "+q+" has no results", replyToken)
	} else {
		line.SendImage(imageURL, replyToken)
	}
}

func getSearchText(q string) string {
	return strings.Replace(q, "/pic ", "", 1)
}

func command(text string) string {
	rx, err := regexp.Compile(`^\/pic|^\/devexcuse`)
	if err != nil {
		log.Println(err)
	}

	return rx.FindString(text)
}
