package main

import "fmt"

func main() {
	cats := map[string]string {
		"cat1": "Uchi",
		"cat2": "Ofy",
	}
	//cats["cat1"] = "Chiiii"
	//delete(cats, "cat1")
	printMap(cats)
}

func printMap(c map[string]string) {
	for key, value := range c {
		fmt.Printf("%-v is %-v \n", key, value)
	}
}