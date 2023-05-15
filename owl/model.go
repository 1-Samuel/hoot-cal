package owl

import "time"

type Match struct {
	ID     int       `json:"id"`
	Teams  []Team    `json:"teams"`
	Start  time.Time `json:"start"`
	End    time.Time `json:"end"`
	Event  string    `json:"event"`
	Status string    `json:"status"`
	Encore bool      `json:"encore"`
}

type Team struct {
	Name            string `json:"name"`
	AbbreviatedName string `json:"abbreviatedName"`
	Icon            string `json:"icon"`
	Score           int    `json:"score"`
	Color           string `json:"color"`
}

type TeamColored struct {
	Team  `json:",inline" bson:",inline"`
	Color string `json:"color"`
}

type ActiveMatch struct {
	UID         string        `json:"uid"`
	Teams       []TeamColored `json:"teams"`
	Status      string        `json:"status"`
	TimeToMatch int           `json:"timeToMatch"`
	LinkToMatch string        `json:"linkToMatch"`
	IsEncore    bool          `json:"isEncore"`
	MatchDate   time.Time     `json:"matchDate"`
}
