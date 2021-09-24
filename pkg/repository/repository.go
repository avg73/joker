package repository

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/avg73/joker/pkg/models"
)

func getBody(path string) ([]byte, error) {
	resp, err := http.Get(path)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

// GetJoke gets random Joke (from a given category if it is not an empty string)
func GetJoke(category string) (*models.Joke, error) {
	path := "https://api.chucknorris.io/jokes/random"
	if category != "" {
		path = fmt.Sprintf("%s?category=%s", path, category)
	}

	body, err := getBody(path)
	if err != nil {
		return nil, err
	}

	joke := new(models.Joke)
	err = json.Unmarshal(body, joke)
	if err != nil {
		return nil, err
	}

	return joke, nil
}

func GetCategories() ([]string, error) {
	path := "https://api.chucknorris.io/jokes/categories"
	body, err := getBody(path)
	if err != nil {
		return nil, err
	}

	categories := new([]string)
	err = json.Unmarshal(body, categories)
	if err != nil {
		return nil, err
	}

	return *categories, nil
}
