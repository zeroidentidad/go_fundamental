package main

import "configformats/json_"

func main() {
	if err := MarshalAll(); err != nil {
		panic(err)
	}

	if err := UnmarshalAll(); err != nil {
		panic(err)
	}

	if err := json_.OtherJSONExamples(); err != nil {
		panic(err)
	}
}
