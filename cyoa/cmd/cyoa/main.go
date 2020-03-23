package main

import (
	"flag"
	"fmt"
	"github.com/ambye85/gophercises/cyoa"
	"html/template"
	"net/http"
	"os"
)

func main() {
	host := flag.String("host", "127.0.0.1", "the host address to bind the cyoa web application on")
	port := flag.Int("port", 8080, "the port the cyoa web application will listen on")
	filename := flag.String("file", "gopher.json", "the json file with the cyoa story")
	flag.Parse()
	fmt.Printf("using the story in %s\n", *filename)

	f, err := os.Open(*filename)
	if err != nil {
		fmt.Printf("could not load %s", *filename)
	}

	story, err := cyoa.LoadStory(f)
	if err != nil {
		exit(err)
	}

	fmt.Printf("starting the cyoa web application on %s:%d", *host, *port)
	tpl := template.Must(template.New("").Parse("Hello, World!"))
	err = http.ListenAndServe(fmt.Sprintf("%s:%d", *host, *port), cyoa.CreateHandler(
		story,
		cyoa.WithTemplate(tpl),
	))
	if err != nil {
		exit(err)
	}
}

func exit(err error) {
	fmt.Println("error:", err)
	os.Exit(1)
}
