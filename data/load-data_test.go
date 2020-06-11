package data

import (
	"encoding/json"
	"reflect"
	"testing"
)

var (
	testJSONData = [][]byte{
		[]byte(`{"name" : "Joe", "address" : {"street" : "montgomery st", "number": 101, "city": "new york", "state": "ny"}}`),
		// []byte(`{“name” : “Evan”, “address” : {“street” : “Santa Theresa st”, “number”: 201, “city”: “sfo”, “state”: “ca”}}`),
	}
)

func TestWalkJSON(t *testing.T) {
	for _, v := range testJSONData {
		var data map[string]interface{}
		if err := json.Unmarshal(v, &data); err != nil {
			t.Error(err.Error())
		}
		got := make(map[string]interface{}, 0)
		want := map[string]interface{}{
			"name": map[string]interface{}{
				"Joe": 1,
			},
			"address": map[string]interface{}{
				"street": map[string]interface{}{
					"montgomery st": 1,
				},
				"number": map[string]interface{}{
					"101": 1,
				},
				"city": map[string]interface{}{
					"new york": 1,
				},
				"state": map[string]interface{}{
					"ny": 1,
				},
			},
		}

		walkJSON(data, got)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got: %v; Want: %v", got, want)
		}
	}
}
