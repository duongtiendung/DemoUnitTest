package main

import "fmt"

func hello(name string) string {
	if name == "" {
		return fmt.Sprintf("What is your name ?")
	} else {
		return fmt.Sprintf("Hello %s", name)
		//return fmt.Sprintf("Hello ")
	}
}
