package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/vikashvverma/greeter/config"
	"github.com/vikashvverma/greeter/job"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}
func main() {
	c := config.ReadConfig("./config.json")
	if c == nil {
		log.Fatal("Could not read config file!")
	}
	g := job.NewGreeter(c)
	s := job.NewScheduler(c.Time, g)
	gocron, err := s.Schedule()
	if err != nil {
		log.Fatalf("ListenAndServe: %s", err)
	}
	<-gocron.Start()

	// No need of a server, a server is just to check the app status easily on cloud
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	http.HandleFunc("/", handler)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
