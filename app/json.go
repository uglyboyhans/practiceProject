package main

import (
	"encoding/json"
	"fmt"
)

type Event struct {
	StartTime int
	EventName string
}

type EventSlice struct {
	Events []Event
}

func testJson() {
	var e Event
	str := `{"StartTime":123456,"EventName":"accOn","Addr":"beijing"}`
	json.Unmarshal([]byte(str),&e)
	fmt.Print(e)
}
