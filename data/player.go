package data

import (
	"crypto/sha256"
	"fmt"
	"time"
)

// Player Define the information about a player in the pool
type Player struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Admin     bool       `json:"admin"`
	LastLogin *time.Time `json:"last_login"`
	Password  string     `json:"password"`
}

var salt = "superhero"

// UserHash Generate hash for user ID
func UserHash(name string) string {
	hash := sha256.Sum256([]byte(salt + name))
	return fmt.Sprintf("%x", hash)
}

// PasswordHash Generate hash for password
func PasswordHash(id string, password string) string {
	hash := sha256.Sum256([]byte(salt + id + password))
	return fmt.Sprintf("%x", hash)
}

// CreatePlayer Create a new player and generate id and password
func CreatePlayer(name string, email string, admin bool, password string) *Player {
	player := &Player{Name: name}
	player.ID = UserHash(player.Name)
	player.Email = email
	player.Admin = admin
	player.LastLogin = nil
	player.Password = PasswordHash(player.ID, password)
	return player
}

// NewPlayer Return a player object using all info
func NewPlayer(id string, name string, email string, admin bool, lastLogin *time.Time, password string) *Player {
	player := &Player{Name: name}
	player.ID = id
	player.Email = email
	player.Admin = admin
	player.LastLogin = lastLogin
	player.Password = password
	return player
}

// PasswordOK Test password
func (p *Player) PasswordOK(password string) bool {
	passHash := PasswordHash(p.ID, password)
	//fmt.Printf("psw:%v test:%v\n", p.Password, passHash)
	return passHash == p.Password
}

// ChangePassword Test password
func (p *Player) ChangePassword(password string) {
	oldpsw := p.Password
	p.Password = PasswordHash(p.ID, password)
	fmt.Printf("Nes psw:%v old:%v\n", p.Password, oldpsw)
}
