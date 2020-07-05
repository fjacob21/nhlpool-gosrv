package controller

import (
	"testing"

	"nhlpool.com/service/go/nhlpool/store"

	"github.com/stretchr/testify/assert"
	"nhlpool.com/service/go/nhlpool/data"
)

func TestEditPlayer(t *testing.T) {
	store.Clean()
	request := data.AddPlayerRequest{Name: "name", Email: "email", Admin: false, Password: "password"}
	reply := AddPlayer(request)
	players := GetPlayers()
	assert.Equal(t, len(players.Players), 1, "Should have one player")
	assert.Equal(t, players.Players[0].Name, "name", "Invalid name")
	assert.Equal(t, players.Players[0].Email, "email", "Invalid email")
	loginReq := data.LoginRequest{Password: "password"}
	loginReply := Login(reply.Player.ID, loginReq)
	editReq := data.EditPlayerRequest{SessionID: loginReply.SessionID, Name: "name2", Email: "email2"}
	editReply := EditPlayer(reply.Player.ID, editReq)
	assert.Equal(t, editReply.Result.Code, data.SUCCESS, "Should be a success")
	players = GetPlayers()
	assert.Equal(t, len(players.Players), 1, "Should have one player")
	assert.Equal(t, players.Players[0].Name, "name2", "Invalid name")
	assert.Equal(t, players.Players[0].Email, "email2", "Invalid email")
}

func TestEditPlayerAdmin(t *testing.T) {
	store.Clean()
	requestAdmin := data.AddPlayerRequest{Name: "admin", Email: "email", Admin: true, Password: "password"}
	replyAdmin := AddPlayer(requestAdmin)
	request := data.AddPlayerRequest{Name: "name", Email: "email", Admin: false, Password: "password"}
	reply := AddPlayer(request)
	loginReq := data.LoginRequest{Password: "password"}
	loginReply := Login(replyAdmin.Player.ID, loginReq)
	editReq := data.EditPlayerRequest{SessionID: loginReply.SessionID, Name: "name2", Email: "email2"}
	editReply := EditPlayer(reply.Player.ID, editReq)
	assert.Equal(t, editReply.Result.Code, data.SUCCESS, "Should be a success")
	players := GetPlayers()
	assert.Equal(t, len(players.Players), 2, "Should have two player")
	player := GetPlayer(reply.Player.ID)
	assert.Equal(t, player.Player.Name, "name2", "Invalid name")
	assert.Equal(t, player.Player.Email, "email2", "Invalid email")
}

func TestEditPlayerNoAccess(t *testing.T) {
	store.Clean()
	requestAdmin := data.AddPlayerRequest{Name: "admin", Email: "email", Admin: false, Password: "password"}
	replyAdmin := AddPlayer(requestAdmin)
	request := data.AddPlayerRequest{Name: "name", Email: "email", Admin: false, Password: "password"}
	reply := AddPlayer(request)
	loginReq := data.LoginRequest{Password: "password"}
	loginReply := Login(replyAdmin.Player.ID, loginReq)
	editReq := data.EditPlayerRequest{SessionID: loginReply.SessionID, Name: "name2", Email: "email2"}
	editReply := EditPlayer(reply.Player.ID, editReq)
	assert.Equal(t, editReply.Result.Code, data.ACCESSDENIED, "Should be access denied")
	players := GetPlayers()
	assert.Equal(t, len(players.Players), 2, "Should have two player")
	player := GetPlayer(reply.Player.ID)
	assert.Equal(t, player.Player.Name, "name", "Invalid name")
	assert.Equal(t, player.Player.Email, "email", "Invalid email")
}

func TestEditPlayerInvalidPlayer(t *testing.T) {
	store.Clean()
	request := data.AddPlayerRequest{Name: "name", Email: "email", Admin: false, Password: "password"}
	reply := AddPlayer(request)
	loginReq := data.LoginRequest{Password: "password"}
	loginReply := Login(reply.Player.ID, loginReq)
	editReq := data.EditPlayerRequest{SessionID: loginReply.SessionID, Name: "name2", Email: "email2"}
	editReply := EditPlayer("invalidid", editReq)
	assert.Equal(t, editReply.Result.Code, data.NOTFOUND, "Should be an error")
	players := GetPlayers()
	assert.Equal(t, len(players.Players), 1, "Should have one player")
	player := GetPlayer(reply.Player.ID)
	assert.Equal(t, player.Player.Name, "name", "Invalid name")
	assert.Equal(t, player.Player.Email, "email", "Invalid email")
}

func TestEditPlayerEmptyPlayer(t *testing.T) {
	store.Clean()
	request := data.AddPlayerRequest{Name: "name", Email: "email", Admin: false, Password: "password"}
	reply := AddPlayer(request)
	loginReq := data.LoginRequest{Password: "password"}
	loginReply := Login(reply.Player.ID, loginReq)
	editReq := data.EditPlayerRequest{SessionID: loginReply.SessionID, Name: "name2", Email: "email2"}
	editReply := EditPlayer("", editReq)
	assert.Equal(t, editReply.Result.Code, data.NOTFOUND, "Should be an error")
	players := GetPlayers()
	assert.Equal(t, len(players.Players), 1, "Should have one player")
	player := GetPlayer(reply.Player.ID)
	assert.Equal(t, player.Player.Name, "name", "Invalid name")
	assert.Equal(t, player.Player.Email, "email", "Invalid email")
}

func TestEditPlayerInvalidSession(t *testing.T) {
	store.Clean()
	request := data.AddPlayerRequest{Name: "name", Email: "email", Admin: false, Password: "password"}
	reply := AddPlayer(request)
	editReq := data.EditPlayerRequest{SessionID: "invalidsession", Name: "name2", Email: "email2"}
	editReply := EditPlayer(reply.Player.ID, editReq)
	assert.Equal(t, editReply.Result.Code, data.ACCESSDENIED, "Should be an error")
	players := GetPlayers()
	assert.Equal(t, len(players.Players), 1, "Should have one player")
	player := GetPlayer(reply.Player.ID)
	assert.Equal(t, player.Player.Name, "name", "Invalid name")
	assert.Equal(t, player.Player.Email, "email", "Invalid email")
}

func TestEditPlayerEmptySession(t *testing.T) {
	store.Clean()
	request := data.AddPlayerRequest{Name: "name", Email: "email", Admin: false, Password: "password"}
	reply := AddPlayer(request)
	editReq := data.EditPlayerRequest{SessionID: "", Name: "name2", Email: "email2"}
	editReply := EditPlayer(reply.Player.ID, editReq)
	assert.Equal(t, editReply.Result.Code, data.ACCESSDENIED, "Should be an error")
	players := GetPlayers()
	assert.Equal(t, len(players.Players), 1, "Should have one player")
	player := GetPlayer(reply.Player.ID)
	assert.Equal(t, player.Player.Name, "name", "Invalid name")
	assert.Equal(t, player.Player.Email, "email", "Invalid email")
}
