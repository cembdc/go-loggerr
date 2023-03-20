package main

import "fmt"

var genericMap map[string]chan Logger

func init() {
	genericMap = make(map[string]chan Logger)
	genericMap["InfoLogger"] = make(chan Logger)
	genericMap["ErrorLogger"] = make(chan Logger)

	go Start(genericMap["InfoLogger"])
	go Start(genericMap["ErrorLogger"])
}

func main() {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			genericMap["InfoLogger"] <- &InfoLog{Message: fmt.Sprintf("Divided by 2, num: %d", i)}
		} else {
			genericMap["ErrorLogger"] <- &ErrorLog{Message: fmt.Sprintf("Can not divided by 2, num: %d", i), Error: fmt.Errorf("Odd number")}
		}
	}
}
