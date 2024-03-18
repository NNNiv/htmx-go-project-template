package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"

	"gotth/internal/templates"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	fmt.Println("Listening on port http://localhost:3000")
	http.Handle("/", templ.Handler(templates.Index()))
	http.ListenAndServe(":3000", nil)

}
