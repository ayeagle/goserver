// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"os"

// 	"goserver/middleware"
// 	"goserver/routes/home"
// 	"goserver/routes/profile"
// 	"goserver/routes/search"

// 	"github.com/joho/godotenv"
// )

// func main() {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatalf("Error loading .env file")
// 	}

// 	env := os.Getenv("ENV")
// 	var serverURL string

// 	if env == "PROD" {
// 		serverURL = os.Getenv("PROD")
// 	} else if env == "LOCAL" {
// 		serverURL = os.Getenv("LOCAL")
// 	} else {
// 		log.Fatalf("Could not start server: \n%s\n", "No valid server URL chosen.")
// 	}

// 	mainRouter := createRouters()
// 	wrappedRouter := middleware.CORSMiddleware(mainRouter)

// 	fmt.Printf("Starting server on %s\n", serverURL)
// 	if err := http.ListenAndServe(serverURL, wrappedRouter); err != nil {
// 		log.Fatalf("Could not start server: %s\n", err)
// 	}
// }

// func createRouters() *http.ServeMux {

// 	// primary router
// 	mainRouter := http.NewServeMux()

// 	// subrouters
// 	homeRouter := http.NewServeMux()
// 	profileRouter := http.NewServeMux()
// 	searchRouter := http.NewServeMux()

// 	// register primary subroutes
// 	home.RegisterRoutes(homeRouter)
// 	profile.RegisterRoutes(profileRouter)
// 	search.RegisterRoutes(searchRouter)

// 	// register subrouters with the main router
// 	mainRouter.Handle("/", homeRouter)

// 	mainRouter.Handle("/home/", http.StripPrefix("/home", homeRouter))
// 	mainRouter.Handle("/profile/", http.StripPrefix("/profile", profileRouter))
// 	mainRouter.Handle("/search/", http.StripPrefix("/search", searchRouter))

// 	return mainRouter
// }

package main

import (
	"context"
	"fmt"
	"net/http"

	"goserver/middleware"
	"goserver/routes/home"
	"goserver/routes/profile"
	"goserver/routes/search"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/handlerfunc"
)

var wrappedRouter http.Handler

// func init() {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatalf("Error loading .env file")
// 	}

// 	mainRouter := createRouters()
// 	wrappedRouter = middleware.CORSMiddleware(mainRouter)
// }

func main() {
	lambda.Start(handler)
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Print("----------------------------------------------------\n")
	fmt.Print("----------------------------------------------------\n")
	fmt.Print("----------------------------------------------------\n")

	fmt.Print("at least this is running\n")

	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatalf("Error loading .env file")
	// }

	fmt.Printf("Domain Name: %s\n", request.RequestContext.DomainName)
	fmt.Printf("Stage: %s\n", request.RequestContext.Stage)
	fmt.Printf("Domain Prefix: %s\n", request.RequestContext.DomainPrefix)

	mainRouter := createRouters()
	wrappedRouter =
		middleware.StripAWSDefaultPathing(
			"/default/bennies",
			middleware.AppendTrailingSlashMiddleware(
				middleware.CORSMiddleware(mainRouter)))

	fmt.Print("Wrapped router created\n")

	proxyEvent := handlerfunc.New(wrappedRouter.ServeHTTP)
	// http.Handle("/", middleware.StripAWSDefaultPathing("/default/benniesGoHandler", wrappedRouter))
	fmt.Print("After proxy event\n")
	fmt.Printf("ctx: %s\n", ctx.Err())
	fmt.Printf("Body: %s\n", request.Body)
	fmt.Printf("HTTPM: %s\n", request.HTTPMethod)
	fmt.Printf("Headrers: %s\n", request.Headers)
	fmt.Printf("PathParams: %s\n", request.PathParameters)

	// fmt.Sprintf("Hello %s!", request.Body)
	// return events.APIGatewayProxyResponse{
	// 	StatusCode: 200,
	// 	Body:       "Hello from Lambda",
	// }, nil

	// return wrappedRouter.ServeHTTP(ctx, request)
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

	fmt.Print("Routers finished creation \n")

	return mainRouter
}
