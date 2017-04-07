package bing

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type SearchValue struct {
	Name         string `json:"name"`
	ThumbnailURL string `json:"thumbnailUrl"`
}

type SearchResponse struct {
	Values []SearchValue `json:"value"`
}

var searchURL = "https://api.cognitive.microsoft.com/bing/v5.0/images/search"

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func SearchImage(term string) string {
	req := bingImageRequest("GET", term)
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	return randomImage(resp)
}

func bingImageRequest(verb, term string) *http.Request {
	req, _ := http.NewRequest(verb, searchURL, nil)
	qs := req.URL.Query()
	qs.Add("q", term)
	qs.Add("size", "Small")
	qs.Add("count", "35")
	req.Header.Add("Ocp-Apim-Subscription-Key", os.Getenv("bing_key1"))
	req.Header.Add("User-Agent", "LineBot/1.0")
	req.URL.RawQuery = qs.Encode()
	return req
}

func randomImage(resp *http.Response) string {
	var searchResp SearchResponse

	err := json.NewDecoder(resp.Body).Decode(&searchResp)
	if err != nil {
		log.Println(err)
	}

	values := searchResp.Values
	randomInt := random(0, len(values))
	return values[randomInt].ThumbnailURL
}

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}
