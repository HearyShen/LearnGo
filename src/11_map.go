package main

import "fmt"

// run with commands:
// go run src\11_map.go

func main() {
	// create a map
	emptyMap := map[string]string{}
	printMap(emptyMap)

	// create and init a map
	city2id := map[string]string{
		"Suzhou":   "0512",
		"Beijing":  "010",
		"Shanghai": "021",
	}
	printMap(city2id)

	// create a map with make(type)
	id2city := make(map[string]string)

	id2city["0512"] = "Suzhou"
	id2city["010"] = "Beijing"
	id2city["021"] = "Shanghai"

	fmt.Printf("Query: 010, Result: %s\n", id2city["010"])
	printMap(id2city)
}

func printMap(srcMap map[string]string) {
	fmt.Printf("len = %d\n", len(srcMap))
	for k, v := range srcMap {
		fmt.Println(k, v, " ")
	}
	fmt.Println()
}
