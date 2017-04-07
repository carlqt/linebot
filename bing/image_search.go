package bing

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
)

// get all values
// filter only images that has jpeg for encodingFormat
// randomly select between 0 to the count of images
// get thumbnail image
// parse to remove \

var key1 = "b3dc1ad1afea4cb6b89bcb5e2b97ea5a"
var key2 = "d39700bc690a47c1ba487e35e09bfcd9"

var searchURL = "https://api.cognitive.microsoft.com/bing/v5.0/images/search"

func SearchImage(term string) {
	req := bingImageRequest("GET", term)
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}

	jsonR(resp)
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
