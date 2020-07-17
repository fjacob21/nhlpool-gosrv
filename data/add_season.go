package data

// AddSeasonRequest Is the info for an add season request
type AddSeasonRequest struct {
	Year int `json:"year"`
}

// AddSeasonReply Is the reply to an add season request
type AddSeasonReply struct {
	Result Result `json:"result"`
	Season Season `json:"season"`
}
