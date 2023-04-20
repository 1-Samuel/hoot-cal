package hoot_cal

import (
	"encoding/json"
	"errors"
	"net/http"
)

const url = "https://us.api.blizzard.com/owl/v1/owl2"

type Repository interface {
	Get() (*OwlResponse, error)
}

type matchRespoitoryApi struct {
	client http.Client
}

func NewMatchRepositoryApi(client http.Client) Repository {
	return matchRespoitoryApi{client: client}
}

func (r matchRespoitoryApi) Get() (*OwlResponse, error) {
	resp, err := r.client.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close() // TODO log error

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("could not get matches from api")
	}

	response := &OwlResponse{}
	err = json.NewDecoder(resp.Body).Decode(response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
