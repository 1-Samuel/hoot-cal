package owl

import "time"

type Match struct {
	ID    int       `json:"id"`
	Teams []Team    `json:"teams"`
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
	Event string    `json:"event"`
	Link  string    `json:"link"`
}

type Team struct {
	Name            string `json:"name"`
	AbbreviatedName string `json:"abbreviatedName"`
	Icon            string `json:"icon"`
}
