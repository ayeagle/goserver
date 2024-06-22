ENV=LOCAL go run main.go to start servy

GOOS=linux GOARCH=arm64 go build -o bootstrap main.go

zip bootstrap.zip bootstrap


GOOS=linux GOARCH=arm64 go build -o bootstrap main.go && zip bootstrap.zip bootstrap

INVOKE FUNCTION
https://iyzb5li6u2.execute-api.us-east-1.amazonaws.com/default/

WITH ROUTE
https://iyzb5li6u2.execute-api.us-east-1.amazonaws.com/default/benniesGoHandler

