package main

import (
	. "github.com/matthewojenkins/programservice/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", HandleHomePage)
	http.HandleFunc("/status", HandleStatusPage)
	http.HandleFunc("/program", HandleProgramRequest)

	http.HandleFunc("/set", HandleSetRequest)
	http.HandleFunc("/workout", HandleWorkoutRequest)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
