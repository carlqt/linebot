package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/carlqt/linebot/line"
)

func validateRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		headerToken := r.Header.Get("X-Line-Signature")
		if headerToken == "" {
			http.Error(w, "not from line", 404)
			return
		}

		ctx := context.WithValue(r.Context(), "headerToken", headerToken)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// I don't know how to authenticate. Documentation is useless
func validateSignature(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

// Checks if replytoken exists
func replyHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var webhook line.WebhookResponse

		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&webhook)
		if err != nil {
			log.Println(err)
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, "replyToken", webhook.Events[0].ReplyToken)
		ctx = context.WithValue(ctx, "text", webhook.Events[0].Message.Text)

		if willReply(webhook) {
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			return
		}
	})
}

func willReply(w line.WebhookResponse) bool {
	if w.Events[0].ReplyToken != "" && keyWordMatch(w.Events[0].Message.Text) {
		return true
	}

	return false
}

func keyWordMatch(text string) bool {
	rx, err := regexp.Compile(`^\/pic`)
	if err != nil {
		log.Println(err)
	}

	matched := rx.MatchString(strings.ToLower(text))

	return matched
}
