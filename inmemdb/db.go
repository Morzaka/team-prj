package inmemdb

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"team-project/logger"
)

var (
	db     = map[string]interface{}{}
	dbLock sync.Mutex
	//redis  map[string]interface{}
	R                Redis
	lock             sync.Mutex
	initialisedRedis bool = false
)

type Redis struct {
	DB map[string]interface{}
}

func initialiseRedis(redis *Redis) {
	(*redis).DB = make(map[string]interface{}, 100)
	initialisedRedis = true
}

//LPush pushes empty interface value into Redis
func LPush(redis *Redis, key string, value interface{}) {
	if initialisedRedis == false {
		initialiseRedis(redis)
	}
	lock.Lock()
	defer lock.Unlock()
	(*redis).DB[key] = value

}

//LGet gets empty interface value from Redis
func LGet(redis *Redis, key string, value interface{}) interface{} {

	lock.Lock()
	defer lock.Unlock()

	value, ok := (*redis).DB[key]
	if !ok {
		logger.Logger.Errorf("Key %q not found", key)
		return nil
	}

	return value
}

//LRemove removes key from Redis
func LRange(redis *Redis, key string) {
	lock.Lock()
	defer lock.Unlock()
	_, ok := (*redis).DB[key]
	if !ok {
		logger.Logger.Errorf("Key %q not found", key)
		return
	}
	if ok {
		delete((*redis).DB, key)
	}
}

//Entry is a map entry,fits responses and requests
type Entry struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}

//sendResponse sends Response
func sendResponse(entry *Entry, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	if err := enc.Encode(entry); err != nil {
		log.Printf("error encoding %+v-%s", entry, err)
	}
}

//dbPostHandler Post Key Value by http
func dbPostHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	dec := json.NewDecoder(r.Body)
	entry := &Entry{}
	if err := dec.Decode(entry); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	dbLock.Lock()
	defer dbLock.Unlock()
	db[entry.Key] = entry.Value
	sendResponse(entry, w)
}

//dbGetHandler GetValue by Key by http
func dbGetHandler(w http.ResponseWriter, r *http.Request) {

	key := r.URL.Path[4:] //trim /db
	dbLock.Lock()
	defer dbLock.Unlock()
	value, ok := db[key]
	if !ok {
		http.Error(w, fmt.Sprintf("Key %q not found", key), http.StatusNotFound)
		return
	}
	entry := &Entry{
		Key:   key,
		Value: value,
	}
	sendResponse(entry, w)
}

//Serve function to serve database in http mode
//func serve() {
//	defer recover()
//
//	http.HandleFunc("/db", dbPostHandler)
//	http.HandleFunc("/db/", dbGetHandler)
//	if err := http.ListenAndServe(":3310", nil); err != nil {
//		log.Fatal(err)
//	}
//}
