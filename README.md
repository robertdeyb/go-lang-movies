# Installation
Download GO: https://golang.org/dl/
Get Extension in VSCODE for Go
After installed, hit CTR + Shift + P and search Go: Install and Update -> Select all

# Setup
Create a folder for the app
go mod init backend

# For running and building Go
go run cmd/api main.go 
Or
go run ./cmd/api

# For building Go
go build sample.go

# For HTTP Router
go get -u github.com/julienschmidt/httprouter

# For JWT Token
go get github.com/pascaldekloe/jwt
