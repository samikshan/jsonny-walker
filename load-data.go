package main

import (
	"fmt"
	"strconv"
)

func walkJSON(data map[string]interface{}, parent map[string]interface{}) map[string]interface{} {
	for k, v := range data { // qualifications,  [BS, MS]
		// fmt.Println(k, v)
		switch v.(type) {
		case []interface{}:
			if _, ok := parent[k]; !ok {
				parent[k] = make(map[string]interface{}, 0)
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
				parent[k].(map[string]int)[vStr] = 0
			}
			parent[k].(map[string]int)[vStr]++
		}
	}

	return parent
}
