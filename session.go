// / revision for the session storage and we will using maps for the sessions
package lib

import (
	"fmt"
	"math/rand/v2"
	"sync"
)

var mutex sync.Mutex // im using this to avoid datarace
// where using mutex, since we are using session ram as a temporary storage session
// this is a secure session but, it may not good for large user or about millions user,,
// in the future development for this session we will create a sqlite3 as the storage of the session

type Session struct {
	Id   string
	Data map[string]string
}

// where are using Ram to access the session
//var SessionStorage = []*Session{}

var SessionStorage = make(map[string]Session)

func GetSession(key string) *Session {
	mutex.Lock()
	defer mutex.Unlock()
	// for _, x := range SessionStorage {
	// 	if x.Id == key {
	// 		return x
	// 	}
	// }
	if _, ok := SessionStorage[key]; ok {
		ses := SessionStorage[key]
		return &ses
	}

	return nil
}

func DeleteSession(key string) {
	mutex.Lock()
	defer mutex.Unlock()
	// for index, x := range SessionStorage {
	// 	if x.Id == key {
	// 		SessionStorage = append(SessionStorage[:index], SessionStorage[index+1:]...)
	// 		break
	// 	}
	// }
	delete(SessionStorage, key)
}

func GetAllSession() map[string]Session {
	mutex.Lock()
	defer mutex.Unlock()
	return SessionStorage
}

func (s *Session) New(key string) string {
	s.Data = make(map[string]string) // initialized the data

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

		// for _, x := range SessionStorage {
		// 	if s.Id == x.Id {
		// 		s.Id = key
		// 		break
		// 	}
		// }
		mutex.Lock()
		defer mutex.Unlock()
		if _, ok := SessionStorage[key]; ok {
			s.Id = key
		}

	}

	return s.Id
}

func (s *Session) GetKey() string {
	mutex.Lock()
	defer mutex.Unlock()
	return s.Id
}

func (s *Session) Set(key string, value string) {
	mutex.Lock()
	defer mutex.Unlock()
	// data := Kv{Key: key, Value: value}
	// s.Data = append(s.Data, data)

	s.Data[key] = value /// set value
}

func (s *Session) UnSet(key string) {
	mutex.Lock()
	defer mutex.Unlock()

	// for index, item := range s.Data {
	// 	if item.Key == key {
	// 		s.Data = append(s.Data[:index], s.Data[index+1:]...)
	// 		break
	// 	}
	// }

	delete(s.Data, key)

}

func (s *Session) Get(key string) string {
	mutex.Lock()
	defer mutex.Unlock()

	// for _, x := range s.Data {
	// 	if x.Key == key {
	// 		return x.Value
	// 	}
	// }
	// return ""

	return s.Data[key]
}

func (s *Session) Store() {
	// call the global variable to store the new session
	mutex.Lock()
	defer mutex.Unlock()

	//SessionStorage = append(SessionStorage, s)

	SessionStorage[s.Id] = Session{Id: s.Id, Data: s.Data}
}
