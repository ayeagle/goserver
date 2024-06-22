ENV=LOCAL go run main.go to start servy

GOOS=linux GOARCH=arm64 go build -o bootstrap main.go

zip bootstrap.zip bootstrap


GOOS=linux GOARCH=arm64 go build -o bootstrap main.go && zip bootstrap.zip bootstrap
