package owl

import (
	"encoding/json"
	"errors"
	"net/http"
)

const url = "https://us.api.blizzard.com/owl/v1/owl2"

type Repository interface {
	Get() (*Response, error)
}

type respoitoryApi struct {
	client http.Client
}

func NewRepositoryApi(client http.Client) Repository {
	return respoitoryApi{client: client}
}

func (r respoitoryApi) Get() (*Response, error) {
	resp, err := r.client.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close() // TODO log error

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("could not get matches from api")
	}

	response := new(Response)
	err = json.NewDecoder(resp.Body).Decode(response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
