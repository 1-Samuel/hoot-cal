package hoot_cal

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/arran4/golang-ical"
	"golang.org/x/oauth2/clientcredentials"
	"log"
	"net/http"
	"os"
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

type OwlResponse struct {
	Players  map[int]Player     `json:"players"`
	Teams    map[int]Team       `json:"teams"`
	Matches  map[int]Match      `json:"matches"`
	Segments map[string]Segment `json:"segments"`
}

func main() {
	client := configureApiClient()

	mux := http.NewServeMux()
	mux.HandleFunc("/owl.ics", func(w http.ResponseWriter, request *http.Request) {
		w.Header().Set("Content-type", "text/calendar")
		w.Header().Set("charset", "utf-8")
		w.Header().Set("Content-Disposition", "inline")
		w.Header().Set("filename", "owl.ics")

		resp, err := client.Get("https://us.api.blizzard.com/owl/v1/owl2")

		if err != nil {
			writeError(http.StatusInternalServerError, "error requesting matches", w, err)
		}

		defer resp.Body.Close()

		if resp.StatusCode == http.StatusOK {
			response := &OwlResponse{}
			json.NewDecoder(resp.Body).Decode(response)
			cal := ics.NewCalendar()
			cal.SetMethod(ics.MethodRequest)
			for _, match := range response.Matches {
				if time.Now().Before(time.Time(match.EndTimestamp)) {
					teams := make([]int, 0, len(match.Teams))

					for _, value := range match.Teams {
						teams = append(teams, value.ID)
					}

					teamRed := response.Teams[teams[0]]
					teamBlue := response.Teams[teams[1]]

					event := cal.AddEvent(fmt.Sprintf("%d@owl", match.ID))
					event.SetStartAt(time.Time(match.StartTimestamp))
					event.SetEndAt(time.Time(match.EndTimestamp))
					event.SetSummary(teamRed.Name + " - " + teamBlue.Name)
				}
			}

			writeSuccess(cal.Serialize(), w)
		} else {
			println("api responsed with status code: " + strconv.Itoa(resp.StatusCode))
		}
	})

	http.ListenAndServe(":8080", mux)

}

func writeError(status int, message string, w http.ResponseWriter, err error) {
	log.Print("ERROR: ", err.Error())
	w.WriteHeader(status)
	w.Write([]byte(message))
}

func writeSuccess(message string, w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(message))
}

func configureApiClient() *http.Client {
	clientID := os.Getenv("BLIZZARD_CLIENT_ID")
	clientSecret := os.Getenv("BLIZZARD_CLIENT_SECRET")

	if clientID == "" || clientSecret == "" {
		panic("env not set")
	}

	config := clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenURL:     "https://oauth.battle.net/token",
	}

	client := config.Client(context.TODO())
	return client
}
