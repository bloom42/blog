//
// FileServer
// ===========
// This example demonstrates how to serve static files from your filesystem.
//
//
// Boot the server:
// ----------------
// $ go run main.go
//
// Client requests:
// ----------------
// $ curl http://localhost:3333/files/
// <pre>
// <a href="notes.txt">notes.txt</a>
// </pre>
//
// $ curl http://localhost:3333/files/notes.txt
// Notessszzz
//
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	lambdachi "github.com/awslabs/aws-lambda-go-api-proxy/chi"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

var httpLambda *lambdachi.ChiLambda

var localFlag bool
var portFlag string
var dirFlag string
var chiMux *chi.Mux

var cacheHeaders = map[string]string{
	"Cache-Control":   "public, max-age=0, s-maxage=31536000",
	"X-Accel-Expires": "31536000",
}

func init() {
	flag.BoolVar(&localFlag, "local", false, "run local server")
	flag.StringVar(&portFlag, "port", "3333", "port to listen to")
	flag.StringVar(&dirFlag, "dir", ".", "the directory to serve")
	flag.Parse()

	chiMux = chi.NewRouter()
	chiMux.Use(middleware.Logger)
	chiMux.Use(middleware.GetHead)
	chiMux.Use(CaheHeadersMiddleware)
	chiMux.NotFound(NotFoundHandler)

	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, dirFlag))
	FileServer(chiMux, "/", filesDir)

	httpLambda = lambdachi.New(chiMux)
}

func CaheHeadersMiddleware(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		// Set our cache headers
		for k, v := range cacheHeaders {
			w.Header().Set(k, v)
		}

		h.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

// HandlerNotFound simply returns an error indicating that the route does not exist
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Not found handler")
	w.WriteHeader(http.StatusNotFound)
	http.ServeFile(w, r, filepath.Join(dirFlag, "404.html"))
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// If no name is provided in the HTTP request body, throw an error
	return httpLambda.ProxyWithContext(ctx, req)
}

func main() {
	if localFlag {
		log.Println("Starting server", fmt.Sprintf("port=%s", portFlag))
		http.ListenAndServe(":"+portFlag, chiMux)
	} else {
		lambda.Start(Handler)
	}
}

// FileServer conveniently sets up a http.FileServer handler to serve
// static files from a http.FileSystem.
func FileServer(router *chi.Mux, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		router.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	router.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		if _, err := os.Stat(fmt.Sprintf("%s", root) + r.RequestURI); os.IsNotExist(err) {
			router.NotFoundHandler().ServeHTTP(w, r)
		} else {
			fs.ServeHTTP(w, r)
		}
		// fs.ServeHTTP(w, r)
	})
}
