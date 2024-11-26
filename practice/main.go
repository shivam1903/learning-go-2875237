package main

import (
	"fmt"
	"sort"
)

func main() {
	states := make(map[string]string)

	fmt.Println(states)
	states["WA"] = "Washington"
	states["NY"] = "New York"
	states["CA"] = "California"

	fmt.Println(states)

	california := states["CA"]
	fmt.Println(california)

	delete(states, "NY")
	states["OR"] = "Oregon"
	fmt.Println(states)

	for k, v := range states {
		fmt.Printf("Key is %v and Value is %v\n", k, v)
	}

	// extracting keys as a slice

	keys := make([]string, len(states))
	i := 0
	for k := range states{
		keys[i] = k
		i++
	}
	fmt.Println(keys)
	sort.Strings(keys)
	fmt.Println(keys)

	for i := range keys{
		fmt.Println(states[keys[i]])
	}
}
