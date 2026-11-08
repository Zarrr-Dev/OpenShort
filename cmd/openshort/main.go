package main

import (
	"flag"
	"log"
	"net/http"
   "os"
   "OpenShort/internal/handler"
   "OpenShort/internal/store"
)

func main() {
		http.HandleFunc("/api/shorten", handler.ShortenHandler)
		http.HandleFunc("/s/", handler.RedirectHandler)
		
    port := os.Getenv("PORT")
    if port == "" {
      port = "8080"
    }
    directory := flag.String("d", "web", "the directory of web file to host")
    flag.Parse()
    http.Handle("/", http.FileServer(http.Dir(*directory)))
    
    err := store.Load()
    if err != nil && os.IsNotExist(err) == false {
    	log.Println(err)
    }

    log.Printf("serving %s on HTTP port: %s\n", *directory, port)
    log.Fatal(http.ListenAndServe(":"+port, nil))
    }