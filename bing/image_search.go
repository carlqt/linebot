package bing

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"os"
	"time"
)

// get all values
// filter only images that has jpeg for encodingFormat
// randomly select between 0 to the count of images
// get thumbnail image
// parse to remove \

type SearchValue struct {
	Name         string `json:"name"`
	ThumbnailURL string `json:"thumbnailUrl"`
}

type SearchResponse struct {
	Values []SearchValue `json:"value"`
}

var key1 = "b3dc1ad1afea4cb6b89bcb5e2b97ea5a"
var key2 = "d39700bc690a47c1ba487e35e09bfcd9"

var searchURL = "https://api.cognitive.microsoft.com/bing/v5.0/images/search"

func init() {
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
	req.Header.Add("Ocp-Apim-Subscription-Key", key1)
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

func dumpOut(r *http.Response) {
	dump, err := httputil.DumpResponse(r, true)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(dump[:]))
}

func jsonR(r *http.Response) {
	var out bytes.Buffer
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Indent(&out, body, "", "  ")
	out.WriteTo(os.Stdout)
}
