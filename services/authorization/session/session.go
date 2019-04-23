package session

import (
	"github.com/google/uuid"
)

type sessionData struct {
	Username string
}

//Session structure for saving user session
type Session struct {
	Data map[uuid.UUID]*sessionData
}

//NewSession creates new session
func NewSession() *Session {
	obj := new(Session)
	obj.Data = make(map[uuid.UUID]*sessionData)
	return obj
}

//Init method  initialize the session
func (obj *Session) Init(username string) uuid.UUID {
	sessionID := uuid.New()
	data := &sessionData{username}
	obj.Data[sessionID] = data
	return sessionID
}

//GetUser method returns authorized username
func (obj *Session) GetUser(sessionID uuid.UUID) string {
	data := obj.Data[sessionID]
	if data == nil {
		return ""
	}
	return data.Username
}
