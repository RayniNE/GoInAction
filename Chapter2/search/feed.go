package search

import (
	"encoding/json"
	"os"
)

const dataFile = "data/datajson"

type Feed struct {
	Name string `json:"site"`
	URI  string `json:"link"`
	Tyoe string `json:"type"`
}

func RetrieveFeeds() ([]*Feed, error) {
	// We open the file.
	file, err := os.Open(dataFile)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var feeds []*Feed
	err = json.NewDecoder(file).Decode(&feeds)

	return feeds, err

}
