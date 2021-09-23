package models

import (
	"strings"
	"time"
)

type Joke struct {
	Id         string    `json:"id"`
	Value      string    `json:"value"`
	Url        string    `json:"url"`
	IconUrl    string    `json:"icon_url"`
	Categories []string  `json:"categories"`
	CreatedAt  ChuckDate `json:"created_at"`
	UpdatedAt  ChuckDate `json:"updated_at"`
}

type ChuckDate struct {
	time.Time
}

func (c *ChuckDate) UnmarshalJSON(b []byte) (err error) {
	layout := "2006-01-02 15:04:05"

	s := strings.Trim(string(b), `"`) //remove quotes
	if s == "null" {
		return
	}

	c.Time, err = time.Parse(layout, s)
	return
}
