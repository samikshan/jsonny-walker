package main

import (
	"container/heap"
	"log"
)

type leaf struct {
	val      interface{}
	fraction float64
}

func getLeavesData(leaves map[string]*Leaf, parentPath string, parentFreq int) []*leaf {
	topKLeaves := make([]*leaf, 0)
	topKStack := make([]*Leaf, 0)

	for Paths[parentPath].Len() > 0 {
		popped := heap.Pop(Paths[parentPath]).(*Leaf)
		log.Println("Popped: ", popped.Count, popped.Value)
		topKStack = append(topKStack, popped)
	}

	// fmt.Println(topKStack)

	for len(topKStack) > 0 {
		n := len(topKStack)
		top := topKStack[n-1]
		topKStack = topKStack[0 : n-1]
		heap.Push(Paths[parentPath], top)

		log.Println("Top value: ", top.Value)
		log.Println("Top count: ", top.Count)
		log.Println("Parent freq: ", parentFreq)

		l := leaf{
			val:      top.Value,
			fraction: float64(top.Count) / float64(parentFreq),
		}

		log.Printf("Adding leaf element. Value: %v, fraction: %f", l.val, l.fraction)

		topKLeaves = append(topKLeaves, &l)
	}

	return topKLeaves
}

func getPaths() []interface{} {
	paths := make([]interface{}, 0)

	nObjects := JSONData["nObjects"].(int)
	if nObjects == 0 {
		return paths
	}

	var traverseData func(parent map[string]interface{}, parentPath string, parentFreq int)
	traverseData = func(parent map[string]interface{}, parentPath string, parentFreq int) {
		leavesData := make([]*leaf, 0)
		for k, v := range parent {
			// fmt.Println(k, v)
			// switch v.(type) {
			// case map[string]interface{}:
			prefix := parentPath + k + "/"
			data := make([]interface{}, 0)
			data = append(data, prefix)

			freq := v.(map[string]interface{})["freq"].(int)
			if occFraction := float64(freq) / float64(nObjects); occFraction < Threshold {
				continue
			} else {
				data = append(data, occFraction)
			}

			components := v.(map[string]interface{})["components"]

			switch components.(type) {
			case map[string]interface{}:
				// prefix += "/"
				traverseData(v.(map[string]interface{})["components"].(map[string]interface{}), prefix, freq)
			case map[string]*Leaf:
				leavesData = getLeavesData(components.(map[string]*Leaf), prefix, freq)
			}

			data = append(data, leavesData)

			paths = append(paths, data)
		}
	}

	rootComponents := JSONData["components"]
	switch rootComponents.(type) {
	case map[string]interface{}:
		traverseData(JSONData["components"].(map[string]interface{}), "", -1)
	case *Leaf:
		// process leaves
	}

	return paths
}
