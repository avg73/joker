package repository

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/avg73/joker/pkg/models"
)

func GetJoke() (*models.Joke, error) {
	resp, err := http.Get("https://api.chucknorris.io/jokes/random")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	joke := new(models.Joke)
	err = json.Unmarshal(body, joke)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return joke, nil
}
