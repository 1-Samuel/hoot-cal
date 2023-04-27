package match

import (
	"fmt"
	"github.com/1-samuel/hoot-cal/owl"
	ics "github.com/arran4/golang-ical"
	"sort"
	"time"
)

type Usecase struct {
	repo owl.Repository
}

func (u Usecase) FindAll() ([]owl.Match, error) {
	matches, err := u.repo.Get()

	if err != nil {
		return nil, err
	}

	sort.Slice(matches, func(i, j int) bool {
		return matches[i].Start.After(matches[j].End)
	})
	return matches, nil
}

func (u Usecase) FindAllCal() (*ics.Calendar, error) {
	matches, err := u.repo.Get()

	if err != nil {
		return nil, err
	}

	cal := ics.NewCalendar()
	cal.SetMethod(ics.MethodRequest)

	for _, match := range matches {
		if time.Now().Before(match.End) {
			event := cal.AddEvent(fmt.Sprintf("%d@owl", match.ID))
			event.SetStartAt(match.Start)
			event.SetEndAt(match.End)
			event.SetSummary(match.Teams[0].AbbreviatedName + " - " + match.Teams[1].AbbreviatedName)
			event.SetDescription(match.Teams[0].Name + " - " + match.Teams[1].Name + "\n" + match.Event)
			event.SetURL(match.Link)
		}
	}

	return cal, nil
}
