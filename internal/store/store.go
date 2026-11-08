package store

import (
	"encoding/json"
	"os"
	"log"
	"sync"
	)
	
	const dataFile = "data/data.json"
	
	type SaveCode struct {
    mu sync.Mutex
    data map[string]string
 }
 
 var Default = SaveCode{data: make(map[string]string)}


 func (s *SaveCode) Search(code string) (string, bool) {
    s.mu.Lock()
    defer s.mu.Unlock()
    urlAsli, ada := s.data[code]
    return urlAsli, ada
}

 func (s *SaveCode) Simpan(urlAsli, code string) {
    s.mu.Lock()
    defer s.mu.Unlock()
    s.data[code] = urlAsli
    b, err := json.Marshal(s.data)
    if err != nil {
    	log.Println(err)
    }
    err = os.WriteFile(dataFile, b, 0644)
    if err != nil {
    	log.Println(err)
    }
 }
 
 func Load() error {
 	b, err := os.ReadFile(dataFile)
 	
 	if err != nil {
 		return err
 	}
 	Default.mu.Lock()
 	defer Default.mu.Unlock()
 	return json.Unmarshal(b, &Default.data)
 }