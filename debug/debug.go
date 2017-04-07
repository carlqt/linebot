package debug

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
)

func dumpOut(r *http.Response) {
	dump, err := httputil.DumpResponse(r, true)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(dump[:]))
}

func jsonRequest(r *http.Request) {
	var out bytes.Buffer
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Indent(&out, body, "", "  ")
	out.WriteTo(os.Stdout)
}

func jsonResponse(r *http.Response) {
	var out bytes.Buffer
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Indent(&out, body, "", "  ")
	out.WriteTo(os.Stdout)
}
