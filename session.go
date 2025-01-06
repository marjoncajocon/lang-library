package session

import (
	"fmt"
	"math/rand/v2"
)

type Kv struct {
	Key   string
	Value string
}

type Session struct {
	Id   string
	Data []Kv
}

// where are using Ram to access the session
var SessionStorage = []*Session{}

func GetSession(key string) *Session {
	for _, x := range SessionStorage {
		if x.Id == key {
			return x
		}
	}

	return nil
}

func DeleteSession(key string) {
	for index, x := range SessionStorage {
		if x.Id == key {
			SessionStorage = append(SessionStorage[:index], SessionStorage[index+1:]...)
		}
	}

}

func GetAllSession() []*Session {
	return SessionStorage
}

func (s *Session) New(key string) string {
	s.Id = ""
	if key == "" {
		// create a niew key
		chr := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()_+=|[]{}<>,./:;" // 87

		tchar := ""
		for i := 0; i < 16; i++ {
			index := rand.IntN(87 - 0)
			tchar += string(chr[index])
		}

		n := ""
		for i := 0; i < 16; i++ {
			index := rand.IntN(87 - 0)
			n += string(chr[index])
		}

		s.Id = fmt.Sprintf("%s-%s", n, tchar)

	} else {

		for _, x := range SessionStorage {
			if s.Id == x.Id {
				s.Id = key
				break
			}
		}

	}

	return s.Id
}

func (s *Session) GetKey() string {

	return s.Id
}

func (s *Session) Set(key string, value string) *Session {
	data := Kv{Key: key, Value: value}

	s.Data = append(s.Data, data)
	return s
}

func (s *Session) Get(key string) string {
	for _, x := range s.Data {
		if x.Key == key {
			return x.Value
		}
	}
	return ""
}

func (s *Session) Store() {
	// call the global variable to store the new session
	SessionStorage = append(SessionStorage, s)
}
