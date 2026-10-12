package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"

)

func main() {
    rand.Seed(time.Now().UnixNano())
    fmt.Println(generateCode(6))
    port := flag.String("p", "8080", "port to serve on")
    directory := flag.String("d", ".", "the directory of static file to host")
    flag.Parse()

    http.Handle("/", http.FileServer(http.Dir(*directory)))

    log.Printf("serving %s on HTTP port: %s\n", *directory, *port)
    log.Fatal(http.ListenAndServe(":"+*port, nil))
}

const characters = "abdcefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789"

func generateCode(length int) string {
    result := make([]byte, length)
    
    for i := 0; i < length; i++ {
        result[i] = characters[rand.Intn(len(characters))]
    }
    return string(result)
}

 var store = SaveCode{data: make(map[string]string)}

 type SaveCode struct {
    mu sync.Mutex
    data map[string]string
 }

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
 }