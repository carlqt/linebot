package debug

import (
	"bytes"
	"encoding/json"
	"io"
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

func JsonDump(r io.ReadCloser) {
	var out bytes.Buffer
	body, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}

	json.Indent(&out, body, "", "  ")
	out.WriteTo(os.Stdout)
}
