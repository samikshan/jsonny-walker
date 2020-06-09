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
			"nObjects": 4,
			"components": {
				"name": {
					"freq": 4,
					"components": {
						"Joe": 3,
						"Evan": 1,
					},
				},
				"address": {
					"freq": 3,
					"components": {
						"street": {
							"freq": 3,
							"components": {
								"crooked lane": 1,
								"some street": 2,
							}
						},
						"city": {
							"freq": 3,
							"components": {
								"SF": 2,
								"NY": 1,
							},
						},
					},
				},
				"qualifications": {
					"freq": 1,
					"components": {
						"0": {
							components: {
								"BS": 1,
								"BTech": 1,
							},
						},
						"1": {
							"components": {
								"MS": 1,
							}
						}
					}
				}
			}
		}
	*/
	JSONData map[string]interface{}
)

func main() {
	JSONData = make(map[string]interface{}, 0)
	JSONData["nObjects"] = 0
	JSONData["components"] = make(map[string]interface{}, 0)

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

		JSONData["nObjects"] = JSONData["nObjects"].(int) + 1
		JSONData["components"] = walkJSON(data, JSONData["components"].(map[string]interface{}))
	}

	fmt.Println(JSONData)

}
