// package routes

// import (
// 	"goserver/routes/home"
// 	"goserver/routes/search"
// 	"log"
// 	"net/http"
// 	"os"
// 	"path/filepath"
// 	"reflect"
// 	"strings"
// )

// // RegisterHandlersForPackage dynamically registers handlers for a given package and base path
// func RegisterHandlersForPackage(packageName string, basePath string) {
// 	files, err := readFilesInDirectory(basePath)
// 	if err != nil {
// 		log.Fatalf("Error reading directory %s: %v", basePath, err)
// 	}

// 	for _, file := range files {
// 		if isRouterFile(file, packageName) {
// 			continue
// 		}

// 		handlerName := deriveHandlerName(file)
// 		routePath := deriveRoutePath(file)

// 		handlerFunc := getHandlerFunc(packageName, handlerName)
// 		if handlerFunc != nil {
// 			registerHandler(routePath, handlerFunc)
// 		}
// 	}
// }

// // readFilesInDirectory returns a list of Go file names in the specified directory
// func readFilesInDirectory(directoryPath string) ([]string, error) {
// 	var files []string
// 	err := filepath.Walk(directoryPath, func(path string, info os.FileInfo, err error) error {
// 		if err != nil {
// 			return err
// 		}
// 		if !info.IsDir() && strings.HasSuffix(info.Name(), ".go") {
// 			files = append(files, info.Name())
// 		}
// 		return nil
// 	})
// 	return files, err
// }

// // isRouterFile checks if a file is the router file for the package
// func isRouterFile(fileName string, packageName string) bool {
// 	return fileName == packageName+"_router.go"
// }

// // deriveHandlerName converts a file name to a handler function name (e.g., search_all.go -> SearchAllHandler)
// func deriveHandlerName(fileName string) string {
// 	name := strings.TrimSuffix(fileName, ".go")
// 	parts := strings.Split(name, "_")
// 	for i := range parts {
// 		parts[i] = strings.Title(parts[i])
// 	}
// 	return strings.Join(parts, "") + "Handler"
// }

// // deriveRoutePath converts a file name to a route path (e.g., search_all.go -> /search/all)
// func deriveRoutePath(fileName string) string {
// 	name := strings.TrimSuffix(fileName, ".go")
// 	return "/" + strings.ReplaceAll(name, "_", "/")
// }

// // getHandlerFunc retrieves a handler function by name using reflection
// func getHandlerFunc(packageName string, handlerName string) func(http.ResponseWriter, *http.Request) {
// 	pkg := reflect.ValueOf(Packages[packageName])
// 	method := pkg.MethodByName(handlerName)
// 	if method.IsValid() {
// 		return method.Interface().(func(http.ResponseWriter, *http.Request))
// 	}
// 	return nil
// }

// // registerHandler registers the handler function with the HTTP server
// func registerHandler(routePath string, handlerFunc func(http.ResponseWriter, *http.Request)) {
// 	http.HandleFunc(routePath, handlerFunc)
// }

// // Packages holds references to all handler functions for each package
// var Packages = map[string]interface{}{
// 	"home": &HomePackage,
// 	// "profile": &ProfilePackage,
// 	"search": &SearchPackage,
// }

// // HomePackage holds handler functions for the home package
// var HomePackage = struct {
// 	HomeHandler func(http.ResponseWriter, *http.Request)
// }{
// 	HomeHandler: home.HomeHandler,
// }

// // // ProfilePackage holds handler functions for the profile package
// // var ProfilePackage = struct {
// // 	ProfileHandler func(http.ResponseWriter, *http.Request)
// // }{
// // 	ProfileHandler: ProfileHandler,
// // }

// // SearchPackage holds handler functions for the search package
// var SearchPackage = struct {
// 	SearchAllHandler func(http.ResponseWriter, *http.Request)
// 	SearchOneHandler func(http.ResponseWriter, *http.Request)
// }{
// 	SearchAllHandler: search.SearchAllHandler,
// 	SearchOneHandler: search.SearchOneHandler,
// }

package routes

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
)

var packages = make(map[string]interface{})

// RegisterHandlersForPackage dynamically registers handlers for a given package and base path
func RegisterHandlersForPackage(packageName string, basePath string, pkg interface{}) {
	files, err := readFilesInDirectory(basePath)
	if err != nil {
		log.Fatalf("Error reading directory %s: %v", basePath, err)
	}

	packages[packageName] = pkg

	for _, file := range files {
		if isRouterFile(file, packageName) {
			continue
		}

		// fmt.Printf("Starting server endpoint for: %s\n", file)

		handlerName := deriveHandlerName(file)
		routePath := deriveRoutePath(file)

		fmt.Printf("Handler name: %s\n", handlerName)

		fmt.Printf("Route path: %s\n", routePath)

		handlerFunc := getHandlerFunc(packageName, handlerName)
		fmt.Print("\nmade it this far 1")

		if handlerFunc != nil {
			registerHandler(routePath, handlerFunc)
		}
	}
}

// readFilesInDirectory returns a list of Go file names in the specified directory
func readFilesInDirectory(directoryPath string) ([]string, error) {
	var files []string
	err := filepath.Walk(directoryPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		fmt.Print("\nmade it this far 2")

		if !info.IsDir() && strings.HasSuffix(info.Name(), ".go") {
			files = append(files, info.Name())
		}
		return nil
	})
	return files, err
}

// isRouterFile checks if a file is the router file for the package
func isRouterFile(fileName string, packageName string) bool {
	fmt.Print("\nmade it this far 3")

	return fileName == packageName+"_router.go"
}

// deriveHandlerName converts a file name to a handler function name (e.g., search_all.go -> SearchAllHandler)
func deriveHandlerName(fileName string) string {
	name := strings.TrimSuffix(fileName, ".go")
	parts := strings.Split(name, "_")
	fmt.Print("\nmade it this far 4")

	for i := range parts {
		parts[i] = strings.Title(parts[i])
	}
	return strings.Join(parts, "") + "Handler"
}

// deriveRoutePath converts a file name to a route path (e.g., search_all.go -> /search/all)
func deriveRoutePath(fileName string) string {
	name := strings.TrimSuffix(fileName, ".go")
	fmt.Print("\nmade it this far 5")

	return "/" + strings.ReplaceAll(name, "_", "/")
}

// getHandlerFunc retrieves a handler function by name using reflection
func getHandlerFunc(packageName string, handlerName string) func(http.ResponseWriter, *http.Request) {
	pkg := reflect.ValueOf(packages[packageName])
	method := pkg.MethodByName(handlerName)
	if method.IsValid() {
		return method.Interface().(func(http.ResponseWriter, *http.Request))
	}
	return nil
}

// registerHandler registers the handler function with the HTTP server
func registerHandler(routePath string, handlerFunc func(http.ResponseWriter, *http.Request)) {
	fmt.Printf("Setting up... \nPath: %s\n Func:\n", routePath)

	// fmt.Println("name:", GetFunctionName(handlerFunc))

	http.HandleFunc(routePath, handlerFunc)
}

func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}
