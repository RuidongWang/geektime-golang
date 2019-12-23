package jsontest

import "testing"

import "encoding/json"

import "fmt"

var jsonStr = `{
	"basic_info" : {
		"name" : "Mike",
		"age" : 30
	},
	"job_info" : {
		"skill" : ["Java", "Go", "C"]
	}
}`

func TestEmbeddedJson(t *testing.T) {
	e := new(Employee)
	err := json.Unmarshal([]byte(jsonStr), e)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(*e)
	if v, err := json.Marshal(e); err == nil {
		fmt.Println(string(v))
		// fmt.Println(string(v.basic_info))
	} else {
		t.Error(err)
	}
}
