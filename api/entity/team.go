package entity


type Team struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Country string `json:"country"`
	Type string `json:"type"`
	LeagueID int `json:"league_id"`
}

type Leagues struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Country string `json:"country"`
}

type TeamStats struct {
	ID   int    `json:"id"`
	TeamID int `json:"team_id"`
	Year int `json:"year"`
	Attack int `json:"attack"`
	Defence int `json:"defence"`
	Midfield int `json:"midfield"`
}