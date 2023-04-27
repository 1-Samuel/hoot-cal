package api_client

import (
	"encoding/json"
	"errors"
	"github.com/1-samuel/hoot-cal/owl"
	"net/http"
	"time"
)

const url = "https://us.api.blizzard.com/owl/v1/owl2"

type respoitoryApi struct {
	client http.Client
}

func NewRepositoryApi(client http.Client) owl.Repository {
	return respoitoryApi{client: client}
}

func (r respoitoryApi) Get() ([]owl.Match, error) {
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
	return mapMatches(response), nil
}

func mapMatches(response *Response) []owl.Match {
	matches := make([]owl.Match, len(response.Matches))

	for _, match := range response.Matches {
		teams := make([]int, 0, len(match.Teams))

		for _, value := range match.Teams {
			teams = append(teams, value.ID)
		}

		teamRed := response.Teams[teams[0]]
		teamBlue := response.Teams[teams[1]]

		match := owl.Match{
			ID: match.ID,
			Teams: []owl.Team{
				{
					Name:            teamRed.Name,
					AbbreviatedName: teamRed.Code,
					Icon:            teamRed.Logo,
				},
				{
					Name:            teamBlue.Name,
					AbbreviatedName: teamBlue.Code,
					Icon:            teamBlue.Logo,
				},
			},
			Start: time.Time(match.StartTimestamp),
			End:   time.Time(match.EndTimestamp),
		}

		matches = append(matches, match)
	}
	return matches
}
