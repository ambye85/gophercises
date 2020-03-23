package cyoa

import (
	"encoding/json"
	"io"
)

type Story map[string]Chapter

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

type Option struct {
	Chapter string `json:"arc"`
	Text    string `json:"text"`
}

func LoadStory(f io.Reader) (Story, error) {
	d := json.NewDecoder(f)
	var story Story
	if err := d.Decode(&story); err != nil {
		return nil, err
	}

	return story, nil
}
