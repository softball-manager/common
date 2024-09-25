package player

type Player struct {
	PK        string   `json:"pk" dynamodbav:"pk"`
	Name      string   `json:"name" dynamodbav:"sk"`
	Positions []string `json:"positions" dynamodbav:"positions,omitempty"`
	Stats     []Stats  `json:"stats" dynamodbav:"stats,omitempty"`
}

type Stats struct {
	Year            string
	GameStats       GameStats
	CalculatedStats CalculatedStats
}

type GameStats struct {
	PlateApperances int `json:"plateApperances"`
	Hits            int `json:"hits"`
	Singles         int `json:"singles"`
	Doubles         int `json:"doubles"`
	Triples         int `json:"triples"`
	HomeRuns        int `json:"homeRuns"`
	Walks           int `json:"walks"`
	StrikeOuts      int `json:"strikeOuts"`
	RunsBattedIn    int `json:"rbis"`
}

type CalculatedStats struct {
	BattingAverage     float32 `json:"battingAverage"`
	OnBasePlusSlugging float32 `json:"onBasePlusSlugging"`
	OnBasePercentage   float32 `json:"onBasePercentage"`
	Slugging           float32 `json:"slugging"`
}
