package main

import (
	"encoding/json"
	"testing"
)

func Benchmark_Rule(t *testing.B) {
	jsonStr := `{
        "country": {
            "op": "in",
            "val": ["cn", "us"]
        },
        "version": {
            "op": ">=",
            "val": 88888
        },
        "age": {
            "op": "<=",
            "val": 30
        }
    }`
	ruleNum := 1000
	list := make([]Rule, 0)
	for i := 0; i < ruleNum; i++ {

		var ruleItem Rule
		err := json.Unmarshal([]byte(jsonStr), &ruleItem)
		if err != nil {
			t.Fatal(err)
		}
		list = append(list, ruleItem)
	}

	data := map[string]interface{}{
		"country": "cn",
		"version": 90000,
		"age":     28,
	}

	for i := 0; i < t.N; i++ {
		for j := 0; j < len(list); j++ {
			matchRule(list[j], data)
		}
	}

}
