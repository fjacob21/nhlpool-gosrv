package data

// AddTeamConference Define the information about a team conference
type AddTeamConference struct {
	ID     string `json:"id"`
	League League `json:"league"`
	Name   string `json:"name"`
}

// AddTeamVenue Define the information about a team venue
type AddTeamVenue struct {
	ID       string `json:"id"`
	League   League `json:"league"`
	City     string `json:"city"`
	Name     string `json:"name"`
	Timezone string `json:"timezone"`
	Address  string `json:"address"`
}

// AddTeamRequest Is the info for an add team request
type AddTeamRequest struct {
	ID           string            `json:"id"`
	Abbreviation string            `json:"abbreviation"`
	Name         string            `json:"name"`
	Fullname     string            `json:"fullname"`
	City         string            `json:"city"`
	Active       bool              `json:"active"`
	CreationYear string            `json:"creation_year"`
	Website      string            `json:"website"`
	Venue        AddTeamVenue      `json:"venue"`
	Conference   AddTeamConference `json:"conference"`
}

// AddTeamReply Is the reply to an add team request
type AddTeamReply struct {
	Result Result `json:"result"`
	Team   Team   `json:"team"`
}
