package session

import (
	"crypto/rand"
	"fmt"
)

func GenerateId() string {
	b := make([]byte, 16)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

type sessionData struct {
	Username string
}

type Session struct {
	Data map[string]*sessionData
}

func NewSession() *Session {
	obj := new(Session)
	obj.Data = make(map[string]*sessionData)
	return obj
}

func (obj *Session) Init(username string) string {
	sessionId := GenerateId()
	data := &sessionData{username}
	obj.Data[sessionId] = data
	return sessionId
}
func (obj *Session) GetUser(sessionId string) string {
	data := obj.Data[sessionId]
	if data == nil {
		return ""
	}
	return data.Username
}