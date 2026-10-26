# OpenShort
An open-source, self-hosted URL shortener built with Go. it's lightweight enough to run on a tiny Linux machine, so you can own your short links instead of paying for a monthly subscription service.

## Feature

- Landing page with a URL field and empty-input validation
- Result page that shows the original URL and a copy button
- Go web server that serves the static site
- Random 6-character code generator
- In-memory store for code -> URL, protected by a mutex
- Short URL generation
- Redirect URL To original Url with short URL
- Copy feedback
- Landing URL validation
- Saved shortener link
- Link saved in local storage

## Tech stack
- HTML, CSS, JavaScript, and Go

## Preview 

![OpenShort Preview](static/docs/Dashboard.png)
![OpenShort Preview](static/docs/Ready-Page.png)
![OpenShort Preview](static/docs/404-Page.png)

## Requirements
```
Go 1.27.1
```
## Installation

### Clone repository
```
git clone https://github.com/Zarrr-Dev/OpenShort.git
```

### Enter the folder
```
cd OpenShort
```
### Run the server
```
go run main.go
```
### Open the application
```
http://localhost:8080
```
## Structure folder
```
OPENSHORT/
├── static/
|   ├── assets/
|   |   ├── logo.png
|   |
|   ├──docs/
|   |  ├──Dashboard.png
|   |  ├──Ready-Page.png
|   |  ├──404-Page.png
|   |  
|   ├── 404.html
|   ├── riwayat.html
|   ├── index.css
|   ├── index.html
|   ├── index.js
|   ├── ready.html
|   ├── ready.js
|
├── .gitignore
├── go.mod
├── main.go
├── README.md
```

## Demo link
```
https://openshort-production.up.railway.app/
```
## 🚧 Project Status
MVP: Shortening and Redirect work