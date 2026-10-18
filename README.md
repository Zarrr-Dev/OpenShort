# OpenShort

An open-source, self-hosted URL shortener built with Go and SQLite. its lightweight enough to run on a tiny Linux machine, so you can own your short links instead of paying for a monthly subscription service.

## About project

An open-source URL shortener solution that is lightweight enough to run locally on a tiny Linux machine, and is free of charge with no subscription required.

## Feature
```
-Landing page with a URL field and empty-input validation
- Result page that shows the original URL and a copy button
- Go web server that serves the static site
- Random 6-character code generator
- In-memory store for code -> URL, protected by a mutex
```

## Tech Stack

- HTML, CSS, JavaScript, and Go


## Preview 

![OpenShort Preview](static/docs/dashboard.png)

![OpenShort Preview](static/docs/page2.png)

## Provision
```
Make sure it is installed:
-npm/pnpm
-Golang
```
## How to Install

### Clone Repository
```
git clone https://github.com/Zarrr-Dev/OpenShort.git
```

### Enter the folder
cd OPENSHORT

### Install Dependencies
npm install

### Running go
go run main.go -d static

## Structure Folder
```
OPENSHORT/
├── static/
|   ├── assets/
|   |   ├── logo.png
|   |
|   ├──docs/
|      ├──dashboard.png
|      ├──page2.png
|      
├── 404.html
|
├── detail.html
|
├── index.css
|
├── index.html
|
├── index.js
|
├── ready.html
|
├── ready.js
|
├── go.mod
|
├── main.go
|
├── README.md
```

🚧 Project Status

In Progress — OpenShort is currently under development.