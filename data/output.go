package data

import (
	"container/heap"
	"log"
)

type PathData struct {
	Path        string
	OccFraction float64
	TopKLeaves  []*PathLeaf
}

type PathLeaf struct {
	Val      interface{}
	Fraction float64
}

func getKLeavesData(leaves map[string]*Leaf, parentPath string, parentFreq, K int) []*PathLeaf {
	topKLeaves := make([]*PathLeaf, 0)
	topKStack := make([]*Leaf, 0)

	heapLen := Paths[parentPath].Len()
	if heapLen < K {
		K = heapLen
	}
	for i := 0; i < K; i++ {
		popped := heap.Pop(Paths[parentPath]).(*Leaf)
		log.Println("Popped: ", popped.Count, popped.Value)
		topKStack = append(topKStack, popped)
	}

	for len(topKStack) > 0 {
		n := len(topKStack)
		top := topKStack[n-1]
		topKStack = topKStack[0 : n-1]
		heap.Push(Paths[parentPath], top)

		log.Println("Top value: ", top.Value)
		log.Println("Top count: ", top.Count)
		log.Println("Parent freq: ", parentFreq)

		l := PathLeaf{
			Val:      top.Value,
			Fraction: float64(top.Count) / float64(parentFreq),
		}

		log.Printf("Adding leaf element. Value: %v, fraction: %f", l.Val, l.Fraction)

		topKLeaves = append(topKLeaves, &l)
	}

	return topKLeaves
}

func GetPaths(K int, Threshold float64) []*PathData {
	paths := make([]*PathData, 0)

	nObjects := JSONData["nObjects"].(int)
	if nObjects == 0 {
		return paths
	}

	var traverseData func(parent map[string]interface{}, parentPath string, parentFreq int)
	traverseData = func(parent map[string]interface{}, parentPath string, parentFreq int) {
		leavesData := make([]*PathLeaf, 0)
		for k, v := range parent {
			prefix := parentPath + k + "/"
			data := &PathData{
				Path: prefix,
			}

			freq := v.(map[string]interface{})["freq"].(int)
			if occFraction := float64(freq) / float64(nObjects); occFraction < Threshold {
				continue
			} else {
				data.OccFraction = occFraction
			}

			components := v.(map[string]interface{})["components"]

			switch components.(type) {
			case map[string]interface{}:
				traverseData(v.(map[string]interface{})["components"].(map[string]interface{}), prefix, freq)
			case map[string]*Leaf:
				leavesData = getKLeavesData(components.(map[string]*Leaf), prefix, freq, K)
			}

			data.TopKLeaves = leavesData

			paths = append(paths, data)
		}
	}

	// recursively populate paths
	traverseData(JSONData["components"].(map[string]interface{}), "", -1)

	return paths
}
