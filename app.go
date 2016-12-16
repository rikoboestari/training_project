package main

import (
	"log"
	"net/http"

	_ "github.com/lib/pq" // Dependency to connect postgres

	"github.com/julienschmidt/httprouter"
	"github.com/rikoboestari/training_project/talk_training"
)

func init() {
	talk_training.Init()
}

func main() {
	log.Printf("App starting ...")
	router := httprouter.New()

	router.GET("/v1/talks", talk_training.ReadTalks)
	router.POST("/v1/talks", talk_training.WriteTalks)

	log.Printf("App listen on 3000")
	log.Fatal(http.ListenAndServe(":3000", router))

	talk_training.Clean()
}
