package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jeremylock93/coya"
)

func main() {
	port := flag.Int("port", 3000, "port to the COYA")
	filename := flag.String("file", "gopher.json", "The JSON story file")
	flag.Parse()
	fmt.Printf("Using the story in %s.\n", *filename)

	f, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}

	story, err := coya.JsonStory(f)

	if err != nil {
		panic(err)
	}

	h := coya.NewHandler(story)
	fmt.Printf("Starting on port: %d \n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))

}
