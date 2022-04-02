package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/api2/call", func(w http.ResponseWriter, r *http.Request) {
		resp, err := http.Get("http://api1:8080")
		if err != nil {
			log.Fatalln(err)
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		//Convert the body to type string
		sb := string(body)
		w.Write([]byte(fmt.Sprintf("Hello from API2, We tried reaching API1\r\nYour request URL: %s\r\nAPI1 Responded with: %s", r.RequestURI, sb)))
	})
	r.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf("Hello from API2\r\nYour request URL: %s", r.RequestURI)))
	})
	fmt.Println("API2 is now running and listening to port 8080")

	http.ListenAndServe(":8080", r)
}
