package data

// GetSeasonsReply Define the information about a season in the league
type GetSeasonsReply struct {
	Result  Result   `json:"result"`
	Seasons []Season `json:"seasons"`
}
