package models

import (
	"fmt"
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

const layout = "2006-01-02 15:04:05"

func (c *ChuckDate) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), `"`) //remove quotes
	if s == "null" {
		return
	}

	c.Time, err = time.Parse(layout, s)
	return
}

func (c ChuckDate) MarshalJSON() ([]byte, error) {
	if c.Time.IsZero() {
		return nil, nil
	}
	return []byte(fmt.Sprintf(`"%s"`, c.Time.Format(layout))), nil
}
