# cookies-n-cream

# React
* Install [nodejs](https://nodejs.org/en/) 16.15.1
* cd into /app
* run `npm start`

# Golang (Linux)
* Install golang (go version go1.18.3 linux/amd64)
* You will need to set the environment variable for go
  * `sudo vi ~/.profile` or `sudo vi ~/.bashrc`
  * add `export PATH=$PATH:/usr/local/go/bin` to the top and save
  * run `source ~/.profile`
* Install http framework [gin](https://github.com/gin-gonic/gin#installation) using `go get -u github.com/gin-gonic/gin`
* Install `go get github.com/lib/pq`
* cd `/internal`
* If it is first time: `go mod init github.com/cookies`

Note: You may need to run `go mod tidy` if your dependencies are not working

# Golang (MacOS)
* Install golang (go version go1.18.3 linux/amd64). can use `brew install golang` if you have homebrew
* You will need to set the environment variable for go
  * `touch ~/.zshrc`
  * `open ~/.zshrc`
  * add `export GOPATH=/Users/username/go` and `export PATH=$GOPATH/bin:$PATH`. Each should be on a separate line, then save. username needs to be changed to your own
  * run `source ~/.zshrc`
* Install http framework [gin](https://github.com/gin-gonic/gin#installation) using `go install github.com/gin-gonic/gin@latest`
* cd `/internal`
* `go run sample-data.go`

Note: You may need to run `go mod tidy` if your dependencies are not working
