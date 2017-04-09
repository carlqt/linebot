package devexcuse

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"time"
)

type DevExcuses struct {
	DevExcuse []string `json:"devExcuse"`
}

func Excuse() string {
	var excuses DevExcuses

	jsonFile, err := ioutil.ReadFile("./devexcuse/devexcuse.json")

	if err != nil {
		log.Printf("dev excuse: %s", err)
	}

	err = json.Unmarshal(jsonFile, &excuses)
	if err != nil {
		log.Fatal(err)
	}

	length := len(excuses.DevExcuse)
	randomInt := random(0, length)
	return excuses.DevExcuse[randomInt]
}

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}
