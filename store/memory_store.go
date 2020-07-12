package store

import (
	"errors"

	"nhlpool.com/service/go/nhlpool/data"
)

// MemoryStore Is a data store that keep it only in memory
type MemoryStore struct {
	players  map[string]*data.Player
	sessions map[string]data.LoginData
	leagues  map[string]*data.League
}

// NewMemoryStore Create a new memory store
func NewMemoryStore() Store {
	store := &MemoryStore{}
	store.players = make(map[string]*data.Player)
	store.sessions = make(map[string]data.LoginData)
	store.leagues = make(map[string]*data.League)
	return store
}

// Clean Empty the store
func (ms *MemoryStore) Clean() error {
	ms.players = make(map[string]*data.Player)
	ms.sessions = make(map[string]data.LoginData)
	ms.leagues = make(map[string]*data.League)
	return nil
}

// GetPlayers Return a list of all players
func (ms *MemoryStore) GetPlayers() ([]data.Player, error) {
	var players []data.Player
	for _, player := range ms.players {
		players = append(players, *player)
	}
	return players, nil
}

func (ms *MemoryStore) getPlayerByName(name string) *data.Player {
	for _, player := range ms.players {
		if player.Name == name {
			return player
		}
	}
	return nil
}

// GetPlayer Return the player of the specified ID
func (ms *MemoryStore) GetPlayer(id string) *data.Player {
	player, ok := ms.players[id]
	if !ok {
		return ms.getPlayerByName(id)
	}
	return player
}

// AddPlayer Add a new player
func (ms *MemoryStore) AddPlayer(player *data.Player) error {
	_, ok := ms.players[player.ID]
	if ok {
		return errors.New("Player already exist")
	}
	ms.players[player.ID] = player
	return nil
}

// UpdatePlayer Update a player
func (ms *MemoryStore) UpdatePlayer(player *data.Player) error {
	_, ok := ms.players[player.ID]
	if !ok {
		return errors.New("Player do not exist")
	}
	ms.players[player.ID] = player
	return nil
}

// DeletePlayer Delete a player
func (ms *MemoryStore) DeletePlayer(player *data.Player) error {
	_, ok := ms.players[player.ID]
	if !ok {
		return errors.New("Player do not exist")
	}
	delete(ms.players, player.ID)
	return nil
}

// AddSession Add a new session
func (ms *MemoryStore) AddSession(session *data.LoginData) error {
	ms.sessions[session.SessionID] = *session
	return nil
}

// DeleteSession Delete a session
func (ms *MemoryStore) DeleteSession(session *data.LoginData) error {
	delete(ms.sessions, session.SessionID)
	return nil
}

// GetSession Return a session using it ID
func (ms *MemoryStore) GetSession(sessionID string) (*data.LoginData, error) {
	session, ok := ms.sessions[sessionID]
	if !ok {
		return nil, errors.New("Do not exist")
	}
	return &session, nil
}

// GetSessionByPlayer Return a session using it player name
func (ms *MemoryStore) GetSessionByPlayer(player *data.Player) (*data.LoginData, error) {
	for _, session := range ms.sessions {
		if session.Player.ID == player.ID {
			return &session, nil
		}
	}
	return nil, errors.New("Do not exist")
}

// AddLeague Add a new league
func (ms *MemoryStore) AddLeague(league *data.League) error {
	_, ok := ms.leagues[league.ID]
	if ok {
		return errors.New("League already exist")
	}
	ms.leagues[league.ID] = league
	return nil
}

// UpdateLeague Update a league info
func (ms *MemoryStore) UpdateLeague(league *data.League) error {
	_, ok := ms.leagues[league.ID]
	if !ok {
		return errors.New("League do not exist")
	}
	ms.leagues[league.ID] = league
	return nil
}

// DeleteLeague Delete a league
func (ms *MemoryStore) DeleteLeague(league *data.League) error {
	_, ok := ms.leagues[league.ID]
	if !ok {
		return errors.New("League do not exist")
	}
	delete(ms.leagues, league.ID)
	return nil
}

// GetLeague Get a league
func (ms *MemoryStore) GetLeague(leagueID string) (*data.League, error) {
	league, ok := ms.leagues[leagueID]
	if !ok {
		return nil, errors.New("League do not exist")
	}
	return league, nil
}

// GetLeagues Get all leagues
func (ms *MemoryStore) GetLeagues() ([]data.League, error) {
	var leagues []data.League
	for _, league := range ms.leagues {
		leagues = append(leagues, *league)
	}
	return leagues, nil
}
