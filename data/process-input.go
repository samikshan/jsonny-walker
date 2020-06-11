package data

import (
	"container/heap"
	"fmt"
	"log"
	"strconv"
)

var (
	JSONData map[string]interface{}
	Paths    map[string]*LeafHeap
)

func ProcessJSONInput(data, parent map[string]interface{}, parentPath string) map[string]interface{} {
	log.Println(data)
	log.Println(parent)

	for k, v := range data {
		switch v.(type) {
		case []interface{}:
			if _, ok := parent[k]; !ok {
				parent[k] = make(map[string]interface{}, 0)
				parent[k].(map[string]interface{})["freq"] = 0
				parent[k].(map[string]interface{})["components"] = make(map[string]interface{}, 0)
			}

			parent[k].(map[string]interface{})["freq"] = parent[k].(map[string]interface{})["freq"].(int) + 1

			arrMap := make(map[string]interface{}, 0)
			for i, val := range v.([]interface{}) {
				idxStr := strconv.Itoa(i)
				arrMap[idxStr] = val
			}

			values := parent[k].(map[string]interface{})["components"].(map[string]interface{})
			prefix := parentPath + k + "/"
			parent[k].(map[string]interface{})["components"] = ProcessJSONInput(arrMap, values, prefix)

		case map[string]interface{}:
			if _, ok := parent[k]; !ok {
				parent[k] = make(map[string]interface{}, 0)
				parent[k].(map[string]interface{})["freq"] = 0
				parent[k].(map[string]interface{})["components"] = make(map[string]interface{}, 0)
			}

			parent[k].(map[string]interface{})["freq"] = parent[k].(map[string]interface{})["freq"].(int) + 1

			values := parent[k].(map[string]interface{})["components"].(map[string]interface{})
			prefix := parentPath + k + "/"
			parent[k].(map[string]interface{})["components"] = ProcessJSONInput(v.(map[string]interface{}), values, prefix)

		default:
			prefix := parentPath + k + "/"
			if _, ok := Paths[prefix]; !ok {
				Paths[prefix] = new(LeafHeap)
			}

			vStr := fmt.Sprintf("%v", v)
			if _, ok := parent[k]; !ok {
				parent[k] = make(map[string]interface{}, 0)
				parent[k].(map[string]interface{})["freq"] = 0
				parent[k].(map[string]interface{})["components"] = make(map[string]*Leaf, 0)
			}

			parent[k].(map[string]interface{})["freq"] = parent[k].(map[string]interface{})["freq"].(int) + 1
			var leaf *Leaf
			if _, ok := parent[k].(map[string]interface{})["components"].(map[string]*Leaf)[vStr]; !ok {
				// new leaf.. push to min heap
				leaf = &Leaf{
					Count: 1,
					Index: -1,
					Value: vStr,
				}

				heapLen := Paths[prefix].Len()
				leaf.Index = heapLen
				heap.Push(Paths[prefix], leaf)
			} else {
				leaf = parent[k].(map[string]interface{})["components"].(map[string]*Leaf)[vStr]
				leaf.Count++

				fmt.Println(vStr, leaf.Count)

				if leaf.Index == -1 {
					// leaf element not present in heap
					heapLen := Paths[prefix].Len()
					leaf.Index = heapLen
					heap.Push(Paths[prefix], *leaf)
				} else {
					Paths[prefix].Update(leaf)
				}
			}
			parent[k].(map[string]interface{})["components"].(map[string]*Leaf)[vStr] = leaf
		}
	}

	return parent
}
