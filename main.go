package main

import (
	"fmt"
	"net/http"

	"github.com/DebPaine/go-templ/hello"
	"github.com/a-h/templ"
)

func main() {
	component := hello.Hello("Deb")
	// If we want to print the component on terminal
	// component.Render(context.Background(), os.Stdout)
	http.Handle("/hello", templ.Handler(component))
	port := 3000
	http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
	fmt.Printf("Server started, listening on port %v", port)
}
