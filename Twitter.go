package main

import (
	"encoding/json"
	"fmt"
	"time"
)

var input = `
{
    "created_at": "Thu May 31 00:00:01 +0000 2012"
}
`

type Timestamp time.Time

func main() {
	var val map[string]interface{}

	json.Unmarshal([]byte(input), &val)

	t := val["created_at"].(string)

	var timestamp Timestamp

	timestamp.UnmarshalJSON([]byte(t))
}

func (t *Timestamp) UnmarshalJSON(b []byte) error {

	v, _ := time.Parse(time.RubyDate, string(b))

	fmt.Println(v)

	return nil
}
