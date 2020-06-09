package main

import (
	"encoding/json"
	"fmt"
	"log"
)

var (
	// in memory representation of entire json data state
	/*
		{
			"name": {
				"Joe": 3,
				"Evan": 1,
			},
			"address": {
				"street": {
					"crooked lane": 1,
					"some street": 2,
				},
				"city": {
					"SF": 2,
					"NY": 1,
				}
			},
			"qualifications": {
				"0": {
					"BS": 1,
					"BTech": 1,
				},
				"1": {
					"MS": 1,
				}
			}
		}
	*/
	JSONData map[string]interface{}
)

func main() {
	JSONData = make(map[string]interface{}, 0)
	testJSONData := [][]byte{
		[]byte(`{"name" : "Joe", "address" : {"street" : "montgomery st", "number": 101, "city": "new york", "state": "ny"}}`),
		[]byte(`{"name" : "Evan", "address" : {"street" : "Santa Theresa st", "number": 201, "city": "sfo", "state": "ca"}}`),
		[]byte(`{"name" : "Joe", "qualifications" : ["BS", "MS"]}`),
	}

	for _, d := range testJSONData {
		var data map[string]interface{}
		if err := json.Unmarshal(d, &data); err != nil {
			log.Fatalln("failed to unmarshal json data: ", err.Error())
		}

		walkJSON(data, JSONData)
	}

	fmt.Println(JSONData)

}
