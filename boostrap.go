package main

import (
	"context"
	"log"
	"net/http"

	"goserver/middleware"
	"goserver/routes/home"
	"goserver/routes/profile"
	"goserver/routes/search"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/handlerfunc"
	"github.com/joho/godotenv"
)

var wrappedRouter http.Handler

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	mainRouter := createRouters()
	wrappedRouter = middleware.CORSMiddleware(mainRouter)
}

func main() {
	lambda.Start(handler)
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	proxyEvent := handlerfunc.New(wrappedRouter.ServeHTTP)
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "Hello from Lambda",
	}, nil
	return proxyEvent.ProxyWithContext(ctx, request)
}

func createRouters() *http.ServeMux {
	// primary router
	mainRouter := http.NewServeMux()

	// subrouters
	homeRouter := http.NewServeMux()
	profileRouter := http.NewServeMux()
	searchRouter := http.NewServeMux()

	// register primary subroutes
	home.RegisterRoutes(homeRouter)
	profile.RegisterRoutes(profileRouter)
	search.RegisterRoutes(searchRouter)

	// register subrouters with the main router
	mainRouter.Handle("/", homeRouter)
	mainRouter.Handle("/home/", http.StripPrefix("/home", homeRouter))
	mainRouter.Handle("/profile/", http.StripPrefix("/profile", profileRouter))
	mainRouter.Handle("/search/", http.StripPrefix("/search", searchRouter))

	return mainRouter
}
