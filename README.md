ENV=LOCAL go run main.go to start servy

GOOS=linux GOARCH=arm64 go build -o bootstrap

zip boostrap.zip bootstrap
