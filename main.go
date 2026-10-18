package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

func main() {
    http.HandleFunc("/api/shorten", shortenHandler)
    http.HandleFunc("/s/", redirectHandler)

    rand.Seed(time.Now().UnixNano())
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

 func shortenHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
       http.Error(w, "Hanya Post yang diizinkan", http.StatusMethodNotAllowed)
       return
    }
       urlAsli := r.FormValue("url")

   if urlAsli == "" {
      http.Error(w, "URL Tidak Boleh Kosong", http.StatusBadRequest)
      return
   }
      parsedURL, err := url.ParseRequestURI(urlAsli)
   if err != nil || parsedURL.Host == "" ||
   (parsedURL.Scheme != "http" && parsedURL.Scheme != "https") {
      http.Error(w, "URL tidak valid", http.StatusBadRequest)
      return
   }

   var code string
    code = generateCode(6)
    store.Simpan(urlAsli, code)
    fmt.Fprint(w, code)
 }

 func redirectHandler(w http.ResponseWriter, r *http.Request) {
   code := strings.TrimPrefix(r.URL.Path, "/s/")
   urlAsli, ada := store.Search(code)

   if  !ada  {
   http.NotFound(w, r)
   return
   }
      http.Redirect(w, r, urlAsli, http.StatusFound)
      return
 }