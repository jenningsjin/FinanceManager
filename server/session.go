package main

import (
	"math/rand"
	"time"
)

const CookieChars = "abcdefghijklmnopqrstuvwxyszABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!#$%&'*+-.^_`|~"
const CookieLength = 512

type SessionManager struct {
	CookieMap map[string]string
}

func (sm *SessionManager) Init() {
	sm.CookieMap = make(map[string]string)
}

func (sm* SessionManager) GenerateNewSessionId(username string) string {
	cookie := sm.GenerateRandomCookie(CookieLength)

	for {
		if _, ok := sm.CookieMap[cookie]; ok {
			cookie = sm.GenerateRandomCookie(CookieLength)
		} else {
			break
		}
	}

	sm.CookieMap[cookie] = username
	return cookie
}

func (sm* SessionManager) SessionExists(cookie string) (string, bool) {
	if username, ok := sm.CookieMap[cookie]; ok {
		return username, true
	} else {
		return "", false
	}
}

func (sm* SessionManager) DeleteSession(cookie string) {
	delete(sm.CookieMap, cookie)
}


// "private", used by GetSessionId
func (sm *SessionManager) GenerateRandomCookie(length int) string {
	rand.Seed(time.Now().Unix())
	
	cookie := make([]byte, length)
	for i := 0; i < length; i++ {
		cookie[i] = CookieChars[rand.Intn(len(CookieChars))]
	}
	return string(cookie)
}






