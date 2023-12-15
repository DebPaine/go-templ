package main

import (
	"fmt"
	"net/http"

	"github.com/DebPaine/go-templ/hello"
	"github.com/a-h/templ"
)

type Person struct {
	name   string
	salary int
}

func main() {
	person := Person{name: "Deb", salary: 5000}

	helloComponent := hello.Hello(person.name)
	greetingComponent := hello.Greeting(person.name)
	salaryComponent := hello.Salary(person.name, person.salary)
	// If we want to print the component on terminal
	// component.Render(context.Background(), os.Stdout)

	port := 3000
	http.Handle("/hello", templ.Handler(helloComponent))
	http.Handle("/greeting", templ.Handler(greetingComponent))
	http.Handle("/salary", templ.Handler(salaryComponent))
	http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
	fmt.Printf("Server started, listening on port %v", port)
}
