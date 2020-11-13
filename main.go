package main

import (
	"fmt"

	"caching/cache"
)

// main function of the program
func main() {
	var capacity int
	if _, err := fmt.Scan(&capacity); err!= nil {
		panic(err)
	}
	var cachePolicy string
	if _, err := fmt.Scan(&cachePolicy); err!= nil {
		panic(err)
	}

	// Create object based on input
	c := cache.NewCache(cachePolicy, capacity)
	c.Set("A", "A")
	c.Set("B", "B")
	c.Set("C", "C")
	c.Set("D", "D")
	fmt.Println(c.Get("A").(string))
}
