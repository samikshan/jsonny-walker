package main

import (
	"log"

	"github.com/samikshan/upgraded-umbrella/jsonny-walker/data"
	"github.com/samikshan/upgraded-umbrella/jsonny-walker/server"
)

func main() {
	// iniialise server
	sv := server.New()

	data.JSONData = make(map[string]interface{}, 0)
	data.JSONData["nObjects"] = 0
	data.JSONData["components"] = make(map[string]interface{}, 0)

	data.Paths = make(map[string]*data.LeafHeap)

	if sv == nil {
		log.Fatal("failed to initialise server")
	}

	// Start server
	log.Fatal(sv.E.Start(":1323"))
}
