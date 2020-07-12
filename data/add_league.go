package data

// AddLeagueRequest Is the info for an add league request
type AddLeagueRequest struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Website     string `json:"website"`
}

// AddLeagueReply Is the reply to an add league request
type AddLeagueReply struct {
	Result Result `json:"result"`
	League League `json:"league"`
}
