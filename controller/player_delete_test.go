package controller

import (
	"testing"

	"nhlpool.com/service/go/nhlpool/store"

	"github.com/stretchr/testify/assert"
	"nhlpool.com/service/go/nhlpool/data"
)

func TestDeletePlayer(t *testing.T) {
	store.Clean()
	request := data.AddPlayerRequest{Name: "name", Email: "email", Admin: false, Password: "password"}
	reply := AddPlayer(request)
	loginReq := data.LoginRequest{Password: "password"}
	loginReply := Login(reply.Player.ID, loginReq)
	deleteReq := data.DeletePlayerRequest{SessionID: loginReply.SessionID}
	deleteReply := DeletePlayer(reply.Player.ID, deleteReq)
	assert.Equal(t, deleteReply.Result.Code, data.SUCCESS, "Should be a success")
	players := GetPlayers()
	assert.Equal(t, len(players.Players), 0, "Should have zero player")
}

func TestDeletePlayerAdmin(t *testing.T) {
	store.Clean()
	requestAdmin := data.AddPlayerRequest{Name: "admin", Email: "email", Admin: true, Password: "password"}
	replyAdmin := AddPlayer(requestAdmin)
	request := data.AddPlayerRequest{Name: "name", Email: "email", Admin: false, Password: "password"}
	reply := AddPlayer(request)
	loginReq := data.LoginRequest{Password: "password"}
	loginReply := Login(replyAdmin.Player.ID, loginReq)
	deleteReq := data.DeletePlayerRequest{SessionID: loginReply.SessionID}
	deleteReply := DeletePlayer(reply.Player.ID, deleteReq)
	assert.Equal(t, deleteReply.Result.Code, data.SUCCESS, "Should be a success")
	players := GetPlayers()
	assert.Equal(t, len(players.Players), 1, "Should have one player")
}

func TestDeletePlayerNoAccess(t *testing.T) {
	store.Clean()
	requestAdmin := data.AddPlayerRequest{Name: "admin", Email: "email", Admin: false, Password: "password"}
	replyAdmin := AddPlayer(requestAdmin)
	request := data.AddPlayerRequest{Name: "name", Email: "email", Admin: false, Password: "password"}
	reply := AddPlayer(request)
	loginReq := data.LoginRequest{Password: "password"}
	loginReply := Login(replyAdmin.Player.ID, loginReq)
	deleteReq := data.DeletePlayerRequest{SessionID: loginReply.SessionID}
	deleteReply := DeletePlayer(reply.Player.ID, deleteReq)
	assert.Equal(t, deleteReply.Result.Code, data.ACCESSDENIED, "Should be access denied")
	players := GetPlayers()
	assert.Equal(t, len(players.Players), 2, "Should have two player")
}

func TestDeleteInvalidPlayer(t *testing.T) {
	store.Clean()
	request := data.AddPlayerRequest{Name: "name", Email: "email", Admin: false, Password: "password"}
	reply := AddPlayer(request)
	loginReq := data.LoginRequest{Password: "password"}
	loginReply := Login(reply.Player.ID, loginReq)
	deleteReq := data.DeletePlayerRequest{SessionID: loginReply.SessionID}
	deleteReply := DeletePlayer("invalidid", deleteReq)
	assert.Equal(t, deleteReply.Result.Code, data.NOTFOUND, "Should be an error")
	players := GetPlayers()
	assert.Equal(t, len(players.Players), 1, "Should have one player")
}

func TestDeleteEmptyPlayer(t *testing.T) {
	store.Clean()
	request := data.AddPlayerRequest{Name: "name", Email: "email", Admin: false, Password: "password"}
	reply := AddPlayer(request)
	loginReq := data.LoginRequest{Password: "password"}
	loginReply := Login(reply.Player.ID, loginReq)
	deleteReq := data.DeletePlayerRequest{SessionID: loginReply.SessionID}
	deleteReply := DeletePlayer("", deleteReq)
	assert.Equal(t, deleteReply.Result.Code, data.NOTFOUND, "Should be an error")
	players := GetPlayers()
	assert.Equal(t, len(players.Players), 1, "Should have one player")
}

func TestDeletePlayerInvalidSession(t *testing.T) {
	store.Clean()
	request := data.AddPlayerRequest{Name: "name", Email: "email", Admin: false, Password: "password"}
	reply := AddPlayer(request)
	deleteReq := data.DeletePlayerRequest{SessionID: "invalidsession"}
	deleteReply := DeletePlayer(reply.Player.ID, deleteReq)
	assert.Equal(t, deleteReply.Result.Code, data.ACCESSDENIED, "Should be an error")
	players := GetPlayers()
	assert.Equal(t, len(players.Players), 1, "Should have one player")
}

func TestDeletePlayerEmptySession(t *testing.T) {
	store.Clean()
	request := data.AddPlayerRequest{Name: "name", Email: "email", Admin: false, Password: "password"}
	reply := AddPlayer(request)
	deleteReq := data.DeletePlayerRequest{SessionID: ""}
	deleteReply := DeletePlayer(reply.Player.ID, deleteReq)
	assert.Equal(t, deleteReply.Result.Code, data.ACCESSDENIED, "Should be an error")
	players := GetPlayers()
	assert.Equal(t, len(players.Players), 1, "Should have one player")
}
