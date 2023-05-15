package match

import (
	"fmt"
	"github.com/1-samuel/hoot-cal/owl"
	ics "github.com/arran4/golang-ical"
)

type Usecase struct {
	repo owl.Repository
}

func (u Usecase) FindAll() ([]owl.Match, error) {
	matches, err := u.repo.Get()

	if err != nil {
		return nil, err
	}

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
		event := cal.AddEvent(fmt.Sprintf("%d@owl", match.ID))
		event.SetStartAt(match.Start)
		event.SetEndAt(match.End)
		event.SetSummary(match.Teams[0].AbbreviatedName + " - " + match.Teams[1].AbbreviatedName)
		event.SetDescription(match.Teams[0].Name + " - " + match.Teams[1].Name + "\n" + match.Event)
	}

	return cal, nil
}

func (u Usecase) FindActive() ([]owl.ActiveMatch, error) {
	match, err := u.repo.GetActive()

	if err != nil {
		return nil, err
	}

	return match, nil
}
