package controller

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"nhlpool.com/service/go/nhlpool/data"
	"nhlpool.com/service/go/nhlpool/store"
)

func TestAddPlayer(t *testing.T) {
	store.Clean()
	request := data.AddPlayerRequest{Name: "name", Email: "email", Admin: true, Password: "password"}
	reply := AddPlayer(request)
	assert.NotNil(t, reply, "Should not be nil")
	assert.Equal(t, reply.Player.Name, "name", "Invalid name")
	assert.Equal(t, reply.Player.Email, "email", "Invalid email")
	assert.Equal(t, reply.Player.Admin, true, "Invalid admin")
	assert.Equal(t, reply.Player.Password, "", "Password should be empty")
}

func TestGetPlayers(t *testing.T) {
	store.Clean()
	request := data.AddPlayerRequest{Name: "name", Email: "email", Admin: true, Password: "password"}
	AddPlayer(request)
	reply := GetPlayers()
	assert.NotNil(t, reply, "Should not be nil")
	assert.Equal(t, len(reply.Players), 1, "Should be only one player")

}

func TestImportPlayer(t *testing.T) {
	store.Clean()
	adminAddRequest := data.AddPlayerRequest{Name: "admin", Email: "email", Admin: true, Password: "password"}
	adminAddReply := AddPlayer(adminAddRequest)
	loginRequest := data.LoginRequest{Password: "password"}
	loginReply := Login(adminAddReply.Player.ID, loginRequest)

	player := data.NewPlayer("id", "name", "email", true, nil, "password")
	request := data.ImportPlayerRequest{}
	request.SessionID = loginReply.SessionID
	request.Player = *player
	reply := ImportPlayer(request)
	assert.NotNil(t, reply, "Should not be nil")
	assert.Equal(t, reply.Player.ID, "id", "Invalid id")
	assert.Equal(t, reply.Player.Name, "name", "Invalid name")
	assert.Equal(t, reply.Player.Email, "email", "Invalid email")
	assert.Equal(t, reply.Player.Admin, true, "Invalid admin")
	assert.Equal(t, reply.Player.Password, "password", "Invalid pasword")
}

func TestImportPlayerNotAdmin(t *testing.T) {
	store.Clean()
	adminAddRequest := data.AddPlayerRequest{Name: "admin", Email: "email", Admin: false, Password: "password"}
	adminAddReply := AddPlayer(adminAddRequest)
	loginRequest := data.LoginRequest{Password: "password"}
	loginReply := Login(adminAddReply.Player.ID, loginRequest)

	player := data.NewPlayer("id", "name", "email", true, nil, "password")
	request := data.ImportPlayerRequest{}
	request.SessionID = loginReply.SessionID
	request.Player = *player
	reply := ImportPlayer(request)
	assert.NotNil(t, reply, "Should not be nil")
	assert.Equal(t, reply.Result.Code, data.ACCESSDENIED, "Should be access denied")
}
