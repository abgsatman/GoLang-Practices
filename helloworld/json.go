package main

import (
	"encoding/json"
	"fmt"
)

var jsonStr = []byte(`{
    "id"  : 1,
    "content" : { "foo": 1881, "bar": "lego" }
}`)

type Bar struct {
	Id      int64
	Content json.RawMessage
}

func jsonparser() {
	var bar Bar

	err := json.Unmarshal(jsonStr, &bar)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bar.Content)[0])
}
