package match

import (
	"fmt"
	"github.com/1-samuel/hoot-cal/owl"
	ics "github.com/arran4/golang-ical"
	"golang.org/x/exp/maps"
	"sort"
	"time"
)

type Usecase struct {
	repo owl.Repository
}

func (u Usecase) FindAll() ([]owl.Match, error) {
	response, err := u.repo.Get()

	if err != nil {
		return nil, err
	}

	matches := maps.Values(response.Matches)
	sort.Slice(matches, func(i, j int) bool {
		return time.Time(matches[i].StartTimestamp).After(time.Time(matches[j].StartTimestamp))
	})
	return matches, nil
}

func (u Usecase) FindAllCal() (*ics.Calendar, error) {
	response, err := u.repo.Get()

	if err != nil {
		return nil, err
	}

	cal := ics.NewCalendar()
	cal.SetMethod(ics.MethodRequest)

	for _, match := range response.Matches {
		teams := make([]int, 0, len(match.Teams))

		for _, value := range match.Teams {
			teams = append(teams, value.ID)
		}

		teamRed := response.Teams[teams[0]]
		teamBlue := response.Teams[teams[1]]

		if time.Now().Before(time.Time(match.EndTimestamp)) {
			event := cal.AddEvent(fmt.Sprintf("%d@owl", match.ID))
			event.SetStartAt(time.Time(match.StartTimestamp))
			event.SetEndAt(time.Time(match.EndTimestamp))
			event.SetSummary(teamRed.Name + " - " + teamBlue.Name)
		}
	}

	return cal, nil
}
