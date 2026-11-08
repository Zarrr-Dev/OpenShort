package handler

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"OpenShort/internal/shortcode"
	"OpenShort/internal/store"
	)
	
	func ShortenHandler(w http.ResponseWriter, r *http.Request) {
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
    for {
    	code = shortcode.GenerateCode(6)
    	_, ada := store.Default.Search(code)
    	if !ada {
    		break
    	}
    }
    store.Default.Simpan(urlAsli, code)
    fmt.Fprint(w, code)
 }

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
   code := strings.TrimPrefix(r.URL.Path, "/s/")
   urlAsli, ada := store.Default.Search(code)

   if  !ada  {
   isi, err := os.ReadFile("web/404.html")
   if err != nil {
   	http.NotFound(w, r)
   	return
   }
   w.Header().Set("Content-Type", "text/html; charset=utf-8")
   w.WriteHeader(http.StatusNotFound)
   w.Write(isi)
   return
   }
   http.Redirect(w, r, urlAsli, http.StatusFound)
   return
 }