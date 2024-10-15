package team

type Team struct {
	PK       string              `json:"pk" dynamodbav:"pk"`
	SK       string              `json:"sk" dynamodbav:"sk"`
	TeamName string              `json:"teamName" dynamodbav:"teamName"`
	Players  []string            `json:"players" dynamodbav:"players"`
	Seasons  map[string][]string `json:"seasons" dynamodbav:"seasons"`
}
