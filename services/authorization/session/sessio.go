package session

import (
	"crypto/rand"
	"fmt"
)

//GenerateID generates id for the session
func GenerateID() string {
	b := make([]byte, 16)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

type sessionData struct {
	Username string
}

//Session structure for saving user session
type Session struct {
	Data map[string]*sessionData
}

//NewSession creates new session
func NewSession() *Session {
	obj := new(Session)
	obj.Data = make(map[string]*sessionData)
	return obj
}

//Init method  initialize the session
func (obj *Session) Init(username string) string {
	sessionID := GenerateID()
	data := &sessionData{username}
	obj.Data[sessionID] = data
	return sessionID
}

//GetUser method returns authorized username
func (obj *Session) GetUser(sessionID string) string {
	data := obj.Data[sessionID]
	if data == nil {
		return ""
	}
	return data.Username
}
