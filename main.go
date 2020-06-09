package main

import (
	"fmt"
	"strconv"
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

// data = {“name” : “Joe”, “address” : {“street” : “montgomery st”, “number”: 101, “city”: “new york”, “state”: “ny”}}
// parent = {"name": {"Joe": 3, "Evan": 1}, ...}
func walkJSON(data map[string]interface{}, parent map[string]interface{}) map[string]interface{} {
	for k, v := range data { // qualifications,  [BS, MS]
		// fmt.Println(k, v)
		switch v.(type) {
		case []interface{}:
			if _, ok := parent[k]; !ok {
				parent[k] = make([]interface{}, 0)
			}

			arrMap := make(map[string]interface{}, 0)
			for i, val := range v.([]interface{}) {
				idxStr := strconv.Itoa(i)
				arrMap[idxStr] = val
			}

			values := parent[k].(map[string]interface{})
			parent[k] = walkJSON(arrMap, values)
		case map[string]interface{}:
			if _, ok := parent[k]; !ok { // parent[address] => !ok
				parent[k] = make(map[string]interface{}, 0) // parent[address] = {}
			}

			values := parent[k].(map[string]interface{})
			parent[k] = walkJSON(v.(map[string]interface{}), values)
		default:
			vStr := fmt.Sprintf("%v", v)
			if _, ok := parent[k]; !ok {
				parent[k] = make(map[string]int, 0)
			}
			if _, ok := parent[k].(map[string]int)[vStr]; !ok {
				parent[k].(map[string]int)[vStr]++
			}
		}
	}

	return parent
}

func main() {

}
