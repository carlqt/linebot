package main

import (
	"net/http"
	"strings"

	"github.com/carlqt/linebot/bing"
	"github.com/carlqt/linebot/line"
)

func lineReply(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	text, _ := ctx.Value("text").(string)
	q := getSearchText(text)
	imageURL, err := bing.SearchImage(q)

	replyToken, _ := ctx.Value("replyToken").(string)

	if err != nil {
		line.Send("Sorry, the term "+q+" has no results", replyToken)
	} else {
		line.SendImage(imageURL, replyToken)
	}
	w.WriteHeader(200)
}

func getSearchText(q string) string {
	return strings.Replace(q, "/pic ", "", 1)
}
