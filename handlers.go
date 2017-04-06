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

var secret = "2d7f970ca104a9252d8069e01ab525dd"

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

func validateSignature(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//token := r.Context().Value("headerToken").(string)
		//_, err := base64.StdEncoding.DecodeString(token)
		//if err != nil {
		//	log.Fatal(err)
		//}

		next.ServeHTTP(w, r)
	})
}

func replyHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var reply line.Reply
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&reply)
		if err != nil {
			log.Println(err)
		}

		if reply.Events[0].ReplyToken == "" {
			return
		}

		ctx := context.WithValue(r.Context(), "replyStruct", reply)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func lineReply(w http.ResponseWriter, r *http.Request) {
	// reply only if the text part has @gif and replyToken is present
	var reply line.Reply
	reply, _ = r.Context().Value("replyStruct").(line.Reply)

	if !keyWordMatch(reply.Events[0].Message.Text) {
		return
	}

	//fmt.Println(reply.Events[0].Message.Text)
	//reply.Send("pek")
	reply.SendImage()
	w.WriteHeader(200)
}

func keyWordMatch(text string) bool {
	rx, err := regexp.Compile(`^\/pic`)
	if err != nil {
		log.Println(err)
	}

	matched := rx.MatchString(strings.ToLower(text))

	return matched
}
