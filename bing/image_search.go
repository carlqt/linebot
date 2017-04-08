package bing

import (
	"encoding/json"
	"errors"
	"log"
	"math/rand"
	"net"
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

func SearchImage(term string) (string, error) {
	req := bingImageRequestBuilder("GET", term)
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	if img := randomImage(resp); img == "" {
		return "", errors.New("Sorry, " + term + " has no results")
	} else {
		return img, nil
	}
}

func bingImageRequestBuilder(verb, term string) *http.Request {
	req, _ := http.NewRequest(verb, searchURL, nil)
	qs := req.URL.Query()
	qs.Add("q", term)
	qs.Add("size", "Small")
	qs.Add("count", "35")
	qs.Add("X-Search-ClientIP", ipAddr())
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

	if randomInt < 0 {
		return ""
	} else {
		return values[randomInt].ThumbnailURL
	}
}

func random(min, max int) int {
	if max == 0 {
		log.Println("0 search results")
		return -1
	}

	log.Printf("%d images", max)
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

func ipAddr() string {
	var ipAddr string
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Fatal(err)
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ipAddr = ipnet.IP.String()
			}
		}
	}

	return ipAddr
}
