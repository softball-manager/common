package game

type Game struct {
	PK             string      `json:"pk" dynamodbav:"pk"`
	SK             string      `json:"sk" dynamodbav:"sk"`
	Date           string      `json:"date" dynamodbav:"date"`
	TeamRuns       int         `json:"teamRuns" dynamodbav:"teamRuns"`
	OppositionRuns int         `json:"oppositionRuns" dynamodbav:"oppositionRuns"`
	GameOver       bool        `json:"gameOver" dynamodbav:"gameOver"`
	Win            bool        `json:"win" dynamodbav:"win"`
	Lineup         []Positions `json:"lineup" dynamodbav:"lineup"`
}

type Positions struct {
	Player    string   `json:"player" dynamodbav:"player"`
	Positions []string `json:"positions" dynamodbav:"positions"`
}
