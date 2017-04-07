package main

import (
	"net/http"
	"strings"

	"github.com/carlqt/linebot/bing"
	"github.com/carlqt/linebot/line"
)

func lineReply(w http.ResponseWriter, r *http.Request) {
	var reply line.Reply
	ctx := r.Context()

	text, _ := ctx.Value("text").(string)
	q := getSearchText(text)
	imageURL := bing.SearchImage(q)

	replyToken, _ := ctx.Value("replyToken").(string)
	reply.SendImage(imageURL, replyToken)
	w.WriteHeader(200)
}

func getSearchText(q string) string {
	return strings.Replace(q, "/pic ", "", 1)
}
