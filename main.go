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
var chiMux *chi.Mux

// Taken from https://github.com/mytrile/nocache
var cacheHeaders = map[string]string{
	"Cache-Control":   "public, max-age=0, s-maxage=31536000",
	"X-Accel-Expires": "31536000",
}

// var etagHeaders = []string{
// 	"ETag",
// 	"If-Modified-Since",
// 	"If-Match",
// 	"If-None-Match",
// 	"If-Range",
// 	"If-Unmodified-Since",
// }

func init() {
	flag.BoolVar(&localFlag, "local", false, "run local server")
	flag.StringVar(&portFlag, "port", "3333", "port to listen to")
	dirFlag := flag.String("dir", ".", "the directory to serve")
	flag.Parse()

	chiMux = chi.NewRouter()
	chiMux.Use(middleware.Logger)
	chiMux.Use(middleware.GetHead)
	chiMux.Use(CaheHeadersMiddleware)

	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, *dirFlag))
	FileServer(chiMux, "/", filesDir)

	httpLambda = lambdachi.New(chiMux)
}

func CaheHeadersMiddleware(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {

		// // Delete any ETag headers that may have been set
		// for _, v := range etagHeaders {
		// 	if r.Header.Get(v) != "" {
		// 		r.Header.Del(v)
		// 	}
		// }

		// Set our NoCache headers
		for k, v := range cacheHeaders {
			w.Header().Set(k, v)
		}

		h.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
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
func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}
