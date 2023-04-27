package api_client

import (
	"strconv"
	"time"
)

type unixTime time.Time

func (t unixTime) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(time.Time(t).Unix(), 10)), nil
}

func (t *unixTime) UnmarshalJSON(s []byte) (err error) {
	q, err := strconv.ParseInt(string(s), 10, 64)

	if err != nil {
		return err
	}
	*(*time.Time)(t) = time.Unix(q/1000, 0)
	return
}

type Player struct {
	Number        int    `json:"number"`
	PreferredSlot int    `json:"preferredSlot"`
	GivenName     string `json:"givenName"`
	Teams         []struct {
		ID            int   `json:"id"`
		EarliestMatch int64 `json:"earliestMatch,omitempty"`
	} `json:"teams"`
	Name         string   `json:"name"`
	FamilyName   string   `json:"familyName"`
	Competitions []string `json:"competitions"`
	Role         string   `json:"role"`
	ID           int      `json:"id"`
	HeadshotURL  string   `json:"headshotUrl"`
	AlternateIds []struct {
		Competitions []string `json:"competitions"`
		ID           int      `json:"id"`
	} `json:"alternateIds"`
	CurrentTeam int `json:"currentTeam"`
}

type Team struct {
	ID           int      `json:"id"`
	Competitions []string `json:"competitions"`
	Name         string   `json:"name"`
	Roster       []int    `json:"roster"`
	Code         string   `json:"code"`
	AlternateIds []struct {
		Competitions []string `json:"competitions"`
		ID           int      `json:"id"`
	} `json:"alternateIds"`
	Logo           string `json:"logo"`
	Icon           string `json:"icon"`
	PrimaryColor   string `json:"primaryColor"`
	SecondaryColor string `json:"secondaryColor"`
}

type Match struct {
	CompetitionID        string   `json:"competitionId"`
	Conclusion           string   `json:"conclusion"`
	EndTimestamp         unixTime `json:"endTimestamp"`
	ID                   int      `json:"id"`
	LocalTimeZone        string   `json:"localTimeZone"`
	LocalScheduledDate   string   `json:"localScheduledDate"`
	SeasonID             string   `json:"seasonId"`
	StartTimestamp       unixTime `json:"startTimestamp"`
	ActualStartTimestamp int64    `json:"actualStartTimestamp"`
	ActualEndTimestamp   int64    `json:"actualEndTimestamp"`
	State                string   `json:"state"`
	Teams                map[int]struct {
		ID    int `json:"id"`
		Score int `json:"score"`
	} `json:"teams"`
	Winner  int      `json:"winner"`
	Players []string `json:"players"`
	Games   []string `json:"games"`
}

type Segment struct {
	Id              string `json:"id"`
	Name            string `json:"name"`
	CompetitionId   string `json:"competitionId"`
	SeasonId        string `json:"seasonId"`
	FirstMatchStart int64  `json:"firstMatchStart"`
	LastMatchStart  int64  `json:"lastMatchStart"`
}

type Response struct {
	//Players  map[int]Player     `json:"players"`
	Teams   map[int]Team  `json:"teams"`
	Matches map[int]Match `json:"matches"`
	//Segments map[string]Segment `json:"segments"`
}
