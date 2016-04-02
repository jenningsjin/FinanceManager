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

/**
 * @brief Initializes a session manager
 */
func (sm *SessionManager) Init() {
	sm.CookieMap = make(map[string]string)
}

/**
 * @brief Generates a new session id
 * @details guaranteed to be unique
 * The session id is stored in the session manager
 * and can only be deleted with DeleteSession
 */
func (sm* SessionManager) GenerateNewSessionId(username string) string {
	cookie := sm.generateRandomCookie(CookieLength)

	for {
		if _, ok := sm.CookieMap[cookie]; ok {
			cookie = sm.generateRandomCookie(CookieLength)
		} else {
			break
		}
	}

	sm.CookieMap[cookie] = username
	return cookie
}

/**
 * @brief Checks is a seesion id exists
 */
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

/**
 * @brief "Private" functions used by GenerateNewSessionId
 * @details generates a random cookie
 */
func (sm *SessionManager) generateRandomCookie(length int) string {
	rand.Seed(time.Now().Unix())
	
	cookie := make([]byte, length)
	for i := 0; i < length; i++ {
		cookie[i] = CookieChars[rand.Intn(len(CookieChars))]
	}
	return string(cookie)
}






