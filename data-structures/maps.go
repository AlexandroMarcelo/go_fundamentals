package main

import "fmt"

func main() {
	// map[<key type>]<value type>
	m := make(map[string]int)

	m["Alex"] = 23
	m["Juan"] = 33

	fmt.Println("map ->", m)

	for i, v := range m {
		fmt.Println(i, v)
	}

	fmt.Println(m["Alex"])
	fmt.Println(m["Rubén"]) // zero value, because the key does not exist (0)

	// Check if a key exist
	value, exist := m["Alex"]
	// value, exist := m["Raúl"]
	fmt.Println(value, exist)

}
