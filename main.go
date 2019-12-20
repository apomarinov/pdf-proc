package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	viewer()

	// get line starting Y points for every image
	return
	m := map[int][]int{}
	for i := 1; i < 185; i++ {
		m[i] = getLineStartingPointsForImage(fmt.Sprintf("words_4_ordered/%d.jpg", i))
	}

	b, _ := json.Marshal(m)
	fmt.Println(string(b))
}
